package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

const saveQueueLength = 1000

type record struct {
	Key, URL string
}

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save chan record
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),
	}
	if err := s.load(filename); err != nil {
		log.Println("Error loading data in URLStore:", err)
	}
	go s.saveLoop(filename)
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

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count()) // generate the short URL
		if s.Set(key, url) {
			s.save <- record{key, url}
			return key
		}
	}
	panic("shouldn't be here in URLStore Put method")
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
			s.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore:", err) // map hasn't been read correctly
	return err
}
