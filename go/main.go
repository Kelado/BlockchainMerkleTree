package main

import (
	"bytes"
	"fmt"
	"log"

	mt "github.com/txaty/go-merkletree"
	"golang.org/x/crypto/sha3"
)

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

var VALID_ROOT = "e065e9ea2d2d058750cf07789fe451fd53e8f69c799dca21177b6ecc2f31bea0"

var VALID_PROOF = [][]byte{
	{68, 209, 36, 96, 224, 169, 169, 83, 180, 78, 2, 19, 248, 112, 40, 205, 67, 146, 185, 112, 64, 235, 254, 26, 206, 10, 153, 18, 244, 7, 194, 159},
	{236, 62, 161, 74, 23, 222, 11, 226, 31, 208, 59, 208, 175, 102, 35, 114, 202, 64, 195, 60, 62, 145, 163, 21, 42, 26, 217, 36, 89, 181, 242, 31},
	{173, 94, 172, 253, 212, 12, 110, 255, 167, 234, 141, 69, 73, 71, 244, 156, 250, 211, 211, 81, 39, 130, 242, 250, 99, 114, 192, 201, 34, 215, 29, 24},
	{102, 225, 53, 254, 128, 191, 77, 149, 85, 32, 250, 1, 107, 164, 142, 18, 75, 36, 218, 210, 196, 199, 138, 56, 88, 93, 250, 116, 120, 191, 250, 171},
}

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

func MyPassFunc(data []byte) ([]byte, error) {
	return data, nil
}

func main() {

	blocks := generateDataBlocks()

	c := mt.Config{
		HashFunc:           MyHashFunc,
		Mode:               mt.ModeProofGenAndTreeBuild,
		SortSiblingPairs:   true,
		DisableLeafHashing: false,
	}

	tree, _ := mt.New(&c, blocks)

	mr := tree.Root
	hexMr := fmt.Sprintf("%x", mr)
	fmt.Println(hexMr)
	fmt.Printf("valid: %v\n", hexMr == VALID_ROOT)

	leaf := &DataBlock{
		data: WALLET_ADDRESS,
	}

	proof, err := tree.Proof(leaf)
	if err != nil {
		log.Fatalf("error getting proof: %v", err)
	}

	fmt.Println("\nproof:")
	byteArr := [][]byte{}
	for _, sibling := range proof.Siblings {
		bytes := []byte(sibling)
		fmt.Println(bytes)
		byteArr = append(byteArr, bytes)
	}

	fmt.Printf("valid: %v\n", compareByteByteArrays(byteArr, VALID_PROOF))

	fmt.Println(proof)
}

func compareByteByteArrays(a [][]byte, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !bytes.Equal(a[i], b[i]) {
			return false
		}
	}

	return true
}
