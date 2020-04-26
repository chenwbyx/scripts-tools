package main

import (
	"fmt"
	"log"
	"math/rand"
)

func init() {
	log.Println("skiplist link with max level", maxLevel)
}


type SkipNode struct {
	next []*SkipNode
	length []int
	Value int
}


func newSkipNode(level int) *SkipNode {
	sn := &SkipNode{}
	sn.next = make([]*SkipNode, level)
	sn.length = make([]int, level)
	return sn
}

const (
	p = 0.25
	maxLevel = 24
)

func randomLevel(p float32) int {
	nLevel := 1
	for rand.Int31n(0xffff) < int32(p*0xffff) && nLevel < maxLevel {
		nLevel++
	}
	return nLevel
}

type SkipList struct {
	head      *SkipNode
	tail      *SkipNode
	maxLevel  int        //max level
	curLevel  int        //Current level
	curLength int        //Current length
}

func NewSkipList() *SkipList {
	sl := &SkipList{}
	sl.curLength = 0
	sl.curLevel = 0
	sl.head = newSkipNode(maxLevel)

	return sl
}

func (sl *SkipList) Insert(v int) {
	var update [maxLevel]*SkipNode
	var length [maxLevel]int
	curNode := sl.head
	for level := sl.curLevel-1; level >= 0; level-- {
		for curNode.next[level] != nil && curNode.next[level].Value < v {
			length[level] += curNode.length[level]
			curNode = curNode.next[level]
		}
		update[level] = curNode
	}
	nLevel := randomLevel(p)
	if nLevel > sl.curLevel {
		for i:=sl.curLevel; i < nLevel; i++ {
			update[i] = sl.head
			update[i].length[i] = sl.curLength
			length[i] = 0
		}
		sl.curLevel = nLevel
	}
	node := newSkipNode(nLevel)
	node.Value = v
	steps := 0
	for i:=0; i < nLevel; i++ {
		node.next[i] = update[i].next[i]
		node.length[i] = update[i].length[i]-steps
		update[i].next[i] = node
		update[i].length[i] = steps+1
		steps += length[i]
	}
	for i:= nLevel; i < sl.curLevel; i++ {
		update[i].length[i] += 1
	}
	sl.curLength++
}

func (sl *SkipList)Remove(v int) {
	var update [maxLevel]*SkipNode
	curNode := sl.head
	for level:=sl.curLevel-1; level >= 0; level-- {
		for curNode.next[level] != nil && curNode.next[level].Value < v {
			curNode = curNode.next[level]
		}
		update[level] = curNode
	}
	if curNode.next[0] == nil || curNode.next[0].Value != v {
		return
	}
	curNode = curNode.next[0] // target to remove
	for level:=sl.curLevel-1; level>=0; level-- {
		if update[level].next[level] == curNode {
			update[level].next[level] = curNode.next[level]
			update[level].length[level] += curNode.length[level]-1
		} else {
			update[level].length[level]--
		}
	}
	for sl.curLevel > 0 && sl.head.next[sl.curLevel-1] == nil {
		sl.head.length[sl.curLevel-1] = 0
		sl.curLevel--
	}
	sl.curLength--
}

// get value = v
func (sl *SkipList)GetRank(v int) int {
	i := 0
	curNode := sl.head
	for level:=sl.curLevel-1; level>=0; level-- {
		for curNode.next[level] != nil && curNode.next[level].Value <= v {
			i += curNode.length[level]
			curNode = curNode.next[level]
		}
	}
	if curNode.Value == v {
		return i
	}
	return -1
}

// get position = r
func (sl *SkipList)GetByRank(r int) int {
	i := 0
	curNode := sl.head
	for level:=sl.curLevel-1; level>=0; level-- {
		for curNode.next[level] != nil && curNode.length[level]+i <= r {
			i += curNode.length[level]
			curNode = curNode.next[level]
		}
		if i == r {
			return curNode.Value
		}
	}
	return i
}

// print
func (sl *SkipList)Dprint() {
	curNode := sl.head
	for curNode != nil && curNode.next[0] != nil {
		fmt.Printf("->%#v", curNode.next[0].Value)
		curNode = curNode.next[0]
	}
	fmt.Println("")
}
