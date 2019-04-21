package core

import (
	"PunkCoin/common"
	"bytes"
	"crypto/sha256"
	"math/big"
)

func PrepareData(miner *common.Address,Mainblock *common.BlockHash,BlockOne,BlockTwo *common.BlockHash,TxInputs []TxInput,SendTo []TxOutput,targetbits int,nonce *big.Int,number *big.Int) []byte {
	txinHash := [][]byte{}
	for _,tx := range(TxInputs){
		txinHash = append(txinHash, tx.ToHash())
	}
	txin := bytes.Join(txinHash,[]byte{})

	txoutHash := [][]byte{}
	for _,txout := range(SendTo){
		if txout.OutputAddress == nil {
			continue
		}
		txoutHash = append(txoutHash,txout.ToHash())
	}
	txout := bytes.Join(txoutHash,[]byte{})

	data := bytes.Join(
		[][]byte{
			[]byte((*miner)[:]),
			[]byte((*Mainblock)[:]),
			[]byte((*BlockOne)[:]),
			[]byte((*BlockTwo)[:]),

			common.IntToHex(int64(targetbits)),
			common.IntToHex(nonce.Int64()),
			common.IntToHex(number.Int64()),
		},[]byte{},
	)

	data = bytes.Join([][]byte{data,txin,txout},[]byte{})

	return data
}
//将区块进行hash
func HashBlock(b *Block) []byte {
	var hash [32]byte
	data := PrepareData(b.MinerAddress,b.Mainblock,b.BlockOne,b.BlockTwo,b.TxInputs,b.SendTo,b.targetbits,b.nonce,b.Number)
	hash = sha256.Sum256(data)//计算出哈希
	return hash[:]
}
