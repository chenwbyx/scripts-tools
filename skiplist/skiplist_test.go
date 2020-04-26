package main

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList()
	for i := 10; i < 20; i++ {
		sl.Insert(i)
	}
	sl.Dprint()
	fmt.Println(sl.GetByRank(5))
	fmt.Println(sl.GetRank(15))
	sl.Remove(15)
	sl.Dprint()
}
