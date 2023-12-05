package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

// Blockchain đại diện cho chuỗi khối trong blockchain
type Blockchain struct {
	DB     *bolt.DB
	Blocks []*Block
}

// Serialize chuyển đổi dữ liệu blockchain thành chuỗi bytes
func Serialize(chain []*Block) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(chain)
	return buf.Bytes(), err
}

// Deserialize chuyển đổi dữ liệu từ chuỗi bytes thành blockchain
func Deserialize(data []byte, chain *[]*Block) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	return decoder.Decode(chain)
}

// NewBlockchain tạo một chuỗi block mới với block khởi đầu (genesis block)
func NewBlockchain(db *bolt.DB) *Blockchain {
	var blocks []*Block

	// Thử lấy blockchain hiện tại từ cơ sở dữ liệu
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Blockchain"))
		if b != nil {
			// Unmarshal the blockchain data
			chainData := b.Get([]byte("chain"))
			if chainData != nil {
				err := Deserialize(chainData, &blocks)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error fetching blockchain from BoltDB:", err)
		os.Exit(1)
	}

	// Nếu không có blockchain tồn tại, tạo mới với genesis block
	if len(blocks) == 0 {
		genesisBlock := NewBlock([]*Transaction{}, []byte{})
		blocks = append(blocks, genesisBlock)
	}

	return &Blockchain{DB: db, Blocks: blocks}
}

// Thêm một block mới vào chuỗi khối
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	// Tạo một block mới
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	// Thêm block vào chuỗi khối
	bc.Blocks = append(bc.Blocks, newBlock)

	// Cập nhật dữ liệu blockchain trong BoltDB
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Blockchain"))
		if err != nil {
			return err
		}

		// Chuyển đổi và lưu blockchain đã cập nhật
		chainData, err := Serialize(bc.Blocks)
		if err != nil {
			return err
		}
		return b.Put([]byte("chain"), chainData)
	})

	if err != nil {
		fmt.Println("Loi cap nhat BoltDB:", err)
		os.Exit(1)
	}
}
