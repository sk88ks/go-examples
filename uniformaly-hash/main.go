package main

import (
	"fmt"
	"hash/fnv"
	"math"
	"sort"

	"github.com/awa/liverpool-server/utils"
	"github.com/k0kubun/pp"
)

var m = make([][]string, 200)

func createKeyByHex(str string) int {
	var key uint8
	for i := range str {
		key += uint8(math.Pow(16, float64(i))) * str[i]
	}
	return int(key % 200)
}

func createKey(str string) int {
	var key uint8
	for i := range str {
		key += str[i]
	}
	return int(key % 200)
}

func hash(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

func createKeyByHash(str string) int {
	hashCode := hash(str)
	return int(hashCode % 200)
}

func main() {
	ids := make([]string, 0, 10000000)

	for i := 0; i < 10000000; i++ {
		ids = append(ids, utils.CreateUniqueID())
	}
	for i := range ids {
		//var key uint8
		//for j := range ids[i] {
		//	key += ids[i][j]
		//}

		//key := createKeyByHex(ids[i])
		//key := createKey(ids[i])
		key := createKeyByHash(ids[i])

		//fmt.Println(ids[i], key)
		m[key] = append(m[key], ids[i])
	}

	nums := []int{}
	pp.Println(len(m))
	for i := range m {
		nums = append(nums, len(m[i]))
		fmt.Println(i, len(m[i]))
	}

	sort.Ints(nums)

	pp.Println(nums)

}
