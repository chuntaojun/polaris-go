//@Author: springliao
// @Description:
//@Time: 2021/10/18 17:10

package main

import (
	"log"
	"sync"
)

type T struct {
	lock *sync.RWMutex
	arr  []int
}

func main() {
	t := &T{
		lock: &sync.RWMutex{},
	}

	t.arr = append(t.arr, 1)
	log.Printf("arr len : %d", len(t.arr))
}
