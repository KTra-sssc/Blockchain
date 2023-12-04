package main

import (
	"crypto/sha256"
)

type MerkleTree struct {
	Root *Node
}
type Node struct {
	Left  *Node
	Right *Node
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
func CreateMerkleTree(data [][]byte) *MerkleTree {
	var nodes []Node

	//padding bằng cách duplicate dữ liệu cuối nếu số tx không chia hết cho 2
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	//tạo các node lá với dữ liệu là các transaction
	for _, tx := range data {
		node := NewNode(nil, nil, tx)
		nodes = append(nodes, *node)
	}

	//ghép các cặp dữ liệu lại với nhau, việc hash 2 node sẽ được thực hiện trong hàm NewNode
	//ghép nối cho đến node cuối cùng là root của merkle tree
	for i := 0; i < len(data)/2; i++ {

		//các mức của merkle tree
		//mức cao nhất chiều cao của cây
		var level []Node

		for j := 0; j < len(nodes); j += 2 {
			node := NewNode(&nodes[j], &nodes[j+1], nil)
			level = append(level, *node)
		}

		//tiếp tục ghép cặp với các nút vừa tạo
		nodes = level
	}

	merkle_tree := MerkleTree{&nodes[0]}

	return &merkle_tree
}

// NewMerkleNode creates a new Merkle tree node
func NewNode(left, right *Node, data []byte) *Node {
	node := Node{}

	//với các nút lá không có nút trái phải, hash tx của nút đó
	//với các nút không phải nút lá, nối giá trị hash của nút trái và phải
	//rồi hash chuỗi vừa nối
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		tx_concat := append(left.Data, right.Data...)
		hash := sha256.Sum256(tx_concat)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}
