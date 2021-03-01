package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/rpc"
	"os"
	"sync"
)

const saveQueueLength = 1000

type Store interface {
	Put(url, key *string) error
	Get(key, url *string) error
}

type record struct {
	Key, URL string
}

type ProxyStore struct {
	urls   *URLStore
	client *rpc.Client
}

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save chan record
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}
	if filename != "" {
		s.save = make(chan record, saveQueueLength)
		if err := s.load(filename); err != nil {
			log.Println("Error loading data in URLStore:", err)
		}
		go s.saveLoop(filename)
	}
	return s
}

func (s *URLStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore:", err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	// buffered encoding:
	// b := bufio.NewWriter(f)
	// e := gob.NewEncoder(b)
	// defer b.Flush()
	for {
		r := <-s.save // taking record from channel and encoding it
		if err := e.Encode(r); err != nil {
			log.Println("URLStore:", err)
		}
	}
}

func (s *URLStore) Get(key, url *string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if u, ok := s.urls[*key]; ok {
		*url = u
		return nil
	}
	return errors.New("key not found")
}

func (s *URLStore) Set(key, url *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[*key]; present {
		return errors.New("key already exists")
	}
	s.urls[*key] = *url
	return nil
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url, key *string) error {
	*key = genKey(s.Count()) // generate the short URL
	if err := s.Set(key, url); err != nil {
		return errors.New("error saving URL to map")
	}
	if s.save != nil {
		s.save <- record{*key, *url}
	}
	return nil
}

func (s *URLStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("URLStore:", err)
		return err
	}
	defer f.Close()

	// buffered reading:
	// b := bufio.NewReader(f)
	// d := gob.NewDecoder(b)

	d := json.NewDecoder(f)
	for err != io.EOF {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set(&r.Key, &r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore:", err) // map hasn't been read correctly
	return err
}

func NewProxyStore(addr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Println("Error constructing ProxyStore", err)
	}
	return &ProxyStore{urls: NewURLStore(""), client: client}
}

func (s *ProxyStore) Get(key, url *string) error {
	if err := s.urls.Get(key, url); err == nil {
		return nil
	}
	// rpc call to master
	if err := s.client.Call("Store.Get", key, url); err != nil {
		return err
	}
	s.urls.Set(key, url) // update local cache
	return nil
}

func (s *ProxyStore) Put(url, key *string) error {
	// rpc call to master
	if err := s.client.Call("Store.Put", url, key); err != nil {
		return err
	}
	s.urls.Set(key, url) // update local cache
	return nil
}
