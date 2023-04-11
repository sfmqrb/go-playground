package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type HashFunc func(data []byte) uint32

type ConsistentHash struct {
	hashFunc   HashFunc       // hash function
	replicaCnt int            // number of replicas per node
	keys       []int          // sorted list of keys
	hashMap    map[int]string // hash to node mapping
}

func NewConsistentHash(replicaCnt int, hashFunc HashFunc) *ConsistentHash {
	return &ConsistentHash{
		hashFunc:   hashFunc,
		replicaCnt: replicaCnt,
		keys:       []int{},
		hashMap:    make(map[int]string),
	}
}

func (ch *ConsistentHash) AddNode(node string) {
	for i := 0; i < ch.replicaCnt; i++ {
		key := int(ch.hashFunc([]byte(node + strconv.Itoa(i))))
		ch.keys = append(ch.keys, key)
		ch.hashMap[key] = node
	}
	sort.Ints(ch.keys)
}


func (ch *ConsistentHash) RemoveNode(node string) {
	for i := 0; i < ch.replicaCnt; i++ {
		key := int(ch.hashFunc([]byte(node + strconv.Itoa(i))))
		delete(ch.hashMap, key)
		ch.keys = remove(ch.keys, key)
	}
}


func (ch *ConsistentHash) GetNode(key string) string {
	hash := int(ch.hashFunc([]byte(key)))
	idx := sort.Search(len(ch.keys), func(i int) bool {
		return ch.keys[i] >= hash
	})
	if idx == len(ch.keys) {
		idx = 0
	}
	return ch.hashMap[ch.keys[idx]]
}

func remove(slice []int, item int) []int {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
// PrintState prints the current state of the hash ring.
func (ch *ConsistentHash) PrintState() {
	fmt.Println("Consistent Hash Ring State:")
	fmt.Println("Replica Count per Node:", ch.replicaCnt)
	fmt.Println("Virtual Nodes:")
	for key, value := range ch.hashMap {
		fmt.Println("  Virtual Node Hash:", key, " => Node:", value)
	}
	fmt.Println()
}

func main() {
	// Example usage
	hashFunc := func(data []byte) uint32 {
		return crc32.ChecksumIEEE(data)
	}
	ch := NewConsistentHash(3, hashFunc)
	ch.AddNode("Node1")
	ch.AddNode("Node2")
	ch.AddNode("Node3")

	ch.PrintState()

	fmt.Println(ch.GetNode("Key1")) // Output: Node2
	ch.RemoveNode("Node2")
	ch.PrintState()
	fmt.Println(ch.GetNode("Key1")) // Output: Node3
}

