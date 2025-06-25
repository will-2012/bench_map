package main

import (
	"fmt"
	"time"

	"bench_map/newmap"
	"bench_map/newmap/hasher"

	crand "crypto/rand"

	"github.com/ethereum/go-ethereum/common"
)

var init_in_slice = make([]common.Address, 1000000)
var init_notin_slice = make([]common.Address, 100000)

func main() {
	fmt.Println("hello world")
	//run_example()

	setup()

	bench_map()

	bench_new_map()
}

func run_example() {
	// Initial bucket size/array length
	bucketSize := 4

	// %load at which the bucket is resized/doubled and rehashed
	loadFactor := 80

	mapper := newmap.NewHashTable(bucketSize, loadFactor, hasher.Djb2)

	// Set few values
	mapper.Set("fruits", []string{"orange", "apple"})
	mapper.Set("age", 12)
	mapper.Set("is_adult", false)

	//mapper.Display()
	fmt.Println()

	mapper.Set("activity", newmap.KeyVal{"sport", "swimming"})
	mapper.Set("flavour", "spicy")
	mapper.Set("flavour", "spicy")

	mapper.Display()
	fmt.Println()

	mapper.Remove("activity")

	mapper.Set("activity", "football")
	mapper.Set("activiyt", "volley ball")

	is_adult := mapper.Get("is_adult")
	fmt.Println(is_adult)
}

func generateRandomAddress() common.Address {
	addrBytes := make([]byte, 20)
	crand.Read(addrBytes)
	return common.BytesToAddress(addrBytes)
}

func setup() {
	for i := 0; i < 1000000; i++ {
		init_in_slice = append(init_in_slice, generateRandomAddress())
	}
	for i := 0; i < 100000; i++ {
		init_notin_slice = append(init_notin_slice, generateRandomAddress())
	}
}

func bench_map() {
	map_ := make(map[common.Address]bool, 1000000)
	start := time.Now()
	for _, addr := range init_in_slice {
		map_[addr] = true
	}
	for _, addr := range init_notin_slice {
		if _, ok := map_[addr]; ok {
			//fmt.Println("found in map")
			_ = ok
		}
	}
	for _, addr := range init_in_slice {
		if _, ok := map_[addr]; ok {
			//fmt.Println("found in map")
			_ = ok
		}
	}
	elapsed := time.Since(start)
	fmt.Println("map time: ", elapsed)
}

func bench_new_map() {
	map_ := newmap.NewHashTable(1000000, 80, hasher.AddrHash)
	start := time.Now()
	for _, addr := range init_in_slice {
		map_.Set(addr.String(), true)
	}
	for _, addr := range init_notin_slice {
		v := map_.Get(addr.String())
		_ = v
	}
	for _, addr := range init_in_slice {
		v := map_.Get(addr.String())
		_ = v
	}
	elapsed := time.Since(start)
	fmt.Println("newmap time: ", elapsed)
}
