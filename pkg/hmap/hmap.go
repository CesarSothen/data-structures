package hmap

import (
	"fmt"
	"hash/maphash"
	"math"
	"sync"
)

// V is an interface to store any data type as the value associated to the key in the HMap
type V interface{}

// K is an inteface to store any data type as the key to the HMap
type K interface{}

// Bucket is the struct that contains each instance of Key and its Values
type Bucket struct {
	key   K
	value V
}

// HMap is the struct that contains the various instances of Items
type HMap struct {
	buckets []*Bucket
	hash    maphash.Hash
	mutex   sync.Mutex
}

// Method to write the key into the hash algorithm
func writeToHash(key K, hash maphash.Hash) {
	sKey := fmt.Sprintf("%s", key)
	hash.WriteString(sKey)
}

// Method to calculate the index of the key
func calcIndex(h int, length int) int {
	index := h & (length - 1)
	return index
}

// New creates a new HMap structure with a size multiple of 16 greater than minSize
func New(minSize int) *HMap {
	length := int(math.Ceil(float64(minSize)/16) * 16)

	return &HMap{
		buckets: make([]*Bucket, length),
	}
}

// Put adds a new pair of Key and Value to the Hash Map
func (hm *HMap) Put(key K, value V) {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	defer hm.hash.Reset()

	writeToHash(key, hm.hash)
	index := calcIndex(int(hm.hash.Sum64()), len(hm.buckets))

	hm.buckets[index] = &Bucket{
		key:   key,
		value: value,
	}
}

// Remove removes a pair of Key and Value form the Hash Map
func (hm *HMap) Remove(key K) {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	defer hm.hash.Reset()

	writeToHash(key, hm.hash)
	index := calcIndex(int(hm.hash.Sum64()), len(hm.buckets))

	hm.buckets[index] = nil
}

// Get returns the value associated to the given Key
func (hm *HMap) Get(key K) V {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	defer hm.hash.Reset()

	writeToHash(key, hm.hash)
	index := calcIndex(int(hm.hash.Sum64()), len(hm.buckets))

	bucket := hm.buckets[index]

	if bucket == nil {
		return nil
	}

	value := bucket.value

	return value
}

// ContainsKey checks if the given Key exists in the Hash Map
func (hm *HMap) ContainsKey(key K) bool {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	defer hm.hash.Reset()

	writeToHash(key, hm.hash)
	index := calcIndex(int(hm.hash.Sum64()), len(hm.buckets))

	return hm.buckets[index] != nil
}
