package randomcube

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//ThrowCube returns a random result of throwing
func ThrowCubes(amountOfCubes int) int {

	var result int

	for i := 0; i < amountOfCubes; i++ {
		rand.Seed(time.Now().UnixNano())
		result += 1 + rand.Intn(6)
	}

	return result
}

//CountStatistic returns statistic by results of trowing
func CountStatistic(attempts int, amountOfCubes int) {

	//keys slice is used to sort map entries by the keys (for ordered output)
	var keys []int
	results := make(map[int]int)

	for i := 0; i < attempts; i++ {
		res := ThrowCubes(amountOfCubes)
		results[res] = results[res] + 1
	}

	for k := range results {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range keys {
		var percent float64 = float64(results[key]) / float64(attempts)
		fmt.Printf("%v - %.2f%%\n", key, percent)
	}

}
