package main

import (
	"crypto/sha256"
	"time"
)

type Block struct{

	//版本号
	Version int64
	//前区块哈希值
	PrevblockHash[]byte
	//梅克尔根
	MerkelRoot[]byte
	//时间戳
	TimeStamp int64
	//难度值
	Bits int64
	//随机值
	Nonce int64

	//交易信息
	Data []byte

}
func Newblock(data string,prevBlock []byte)*Block{
	var block Block
	block = Block{
	Version :1,
	PrevblockHash:prevblockHash,
	//Hash TODO
	MerkelRoot:[]byte{},
	TimeStamp:time.Now().Unix(),
	Bits:1,
	Nonce:1,
	Data:[]byte(data)}

	block.SetHash()
	return &block
}

func (block *Block)SetHash(){
	//func Sum256(data []byte) [Size]byte

	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevblockHash,
		block.MerkelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data}

	data := bytes.join(tmp,[]byte{})
	hash := sha256.Sum256(data)
	block.Hash =hash[:]
}

func NewGensisBlock(){
	NewBlock(data:"Gensis Block!",[]byte{})
}
