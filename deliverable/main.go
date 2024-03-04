package main

import (
	"fmt"

	// I used a different package that implements merkletree
	// that also implements Proof API
	mt "github.com/txaty/go-merkletree"
	"golang.org/x/crypto/sha3"
)

// Solution 1
// You need to use ALLOWLIST length equals to the power of 2
// e.g. 2,4,8,16 ...

// Solution 2
// You can have as many as you want
var ALLOWLIST = [][]byte{
	[]byte("sei1vwap3mwdddyrcenn4f3fytd3rvxp7tp08eskp8"),
	[]byte("sei1vtpy9gu4xtaddzxqxalj3z02zt53hzdcstpy77"),
	[]byte("sei15ggkgs9cth485n299cftf3v939m85ludgms9pu"),
	[]byte("sei1g9zyz302ln5nzxyca6fqsre9qnrp7khwywtqf7"),
	[]byte("sei1msqmz3euz4hwdvjn22fckfry6tvwk264u2fg0t"),
	[]byte("sei10wx2t6cfcz8qaq77a2dx7fc74pw3ds3f5ssv20"),
	[]byte("sei1s772ntuxfz2w02jwu8z9lnqcmqzfs7uhgpeda7"),
	[]byte("sei13kawxavfda2w8f8sh2xcr92u6yxjdgjce0qvgs"),
	[]byte("sei1rdd6v2ufrf5h0qm3edc69l6g6xsz6l07z3wtd5"),
	[]byte("sei1x2r5r5ap7gnyxqwc3dpp8603aqyh23athy2m9d"),
}

var WALLET_ADDRESS = []byte("sei1vwap3mwdddyrcenn4f3fytd3rvxp7tp08eskp8")

type DataBlock struct {
	data []byte
}

func (t *DataBlock) Serialize() ([]byte, error) {
	return t.data, nil
}

// Generate all data blocks from ALLOWLIST
func generateDataBlocks() (blocks []mt.DataBlock) {
	for i := range ALLOWLIST {
		block := &DataBlock{
			data: ALLOWLIST[i],
		}
		blocks = append(blocks, block)
	}
	return
}

func MyHashFunc(data []byte) ([]byte, error) {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return hash, nil
}

func main() {
	blocks := generateDataBlocks()

	c := mt.Config{
		HashFunc:           MyHashFunc,
		Mode:               mt.ModeTreeBuild,
		SortSiblingPairs:   true,
		DisableLeafHashing: false,
	}

	tree, _ := mt.New(&c, blocks)

	mr := tree.Root
	hexMr := fmt.Sprintf("Root: %x", mr)
	fmt.Println(hexMr)

	leaf := &DataBlock{
		data: WALLET_ADDRESS,
	}

	fmt.Println(tree.Proof(leaf))
}
