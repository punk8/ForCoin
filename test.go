package main

import (
	"PunkCoin/common"
	"PunkCoin/core"
	"crypto/sha256"
	"fmt"
)

func main() {

	fmt.Println("hello world")

	minerName := "punk"
	address := common.Address(sha256.Sum256([]byte(minerName)))

	mainblock := "MainBlock"
	MB := common.BlockHash(sha256.Sum256([]byte(mainblock)))
	fmt.Println("hello world")

	blockOne := "blockOne"
	B1 := common.BlockHash(sha256.Sum256([]byte(blockOne)))

	blockTwo := "blockTwo"
	B2 := common.BlockHash(sha256.Sum256([]byte(blockTwo)))
	TxInputs := []core.TxInput{}
	SendTo := []core.TxOutput{}
	fmt.Println("hello world")

	amount := 11

	targetbits := 4


	fmt.Println("hello world")



	byte,nonce := core.Pow(&address,&MB,&B1,&B2,TxInputs,SendTo,amount,targetbits)

	fmt.Println(nonce,byte)

}
//e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855