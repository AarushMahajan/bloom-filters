package main

import (
	"fmt"
	"hash"
	"math/rand"

	"github.com/spaolacci/murmur3"
)

var murHash []hash.Hash32

type BloomFilter struct {
	filter []bool
	size   int32
}

func init() {
	murHash = make([]hash.Hash32, 0)
	for i := 0; i < 3; i++ {
		murHash = append(murHash, murmur3.New32WithSeed(rand.Uint32()))
	}
}

func initBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		filter: make([]bool, size),
		size:   size,
	}
}

func hashMurmur3(data []byte, i int) uint32 {
	hasher := murHash[i]
	hasher.Write(data)
	result := hasher.Sum32()
	hasher.Reset()
	return result
}

func (b *BloomFilter) Add(key string) {
	data := []byte(key)
	for i := 0; i < len(murHash); i++ {
		hashValue := hashMurmur3(data, i) % uint32(b.size)
		fmt.Println("Index: ", i, " hashValue: ", hashValue, " key: ", key)
		b.filter[hashValue] = true
	}

}

func (b *BloomFilter) Exists(key string) bool {
	data := []byte(key)
	for i := 0; i < len(murHash); i++ {
		hashValue := hashMurmur3(data, i) % uint32(b.size)
		fmt.Println("Index:::::::: ", i, " hashValue: ", hashValue, " key: ", key)
		if !b.filter[hashValue] {
			return false
		}
	}

	return true
}

func main() {

	dataSet := []string{"a", "b", "c", "d"}
	bloom := initBloomFilter(1000)

	for _, i := range dataSet {
		bloom.Add(i)
	}

	fmt.Println(bloom.Exists("d"))
	fmt.Println(bloom.Exists("ab"))

}
