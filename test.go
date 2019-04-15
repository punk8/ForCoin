package main

import (
	"PunkCoin/client"
	"PunkCoin/common"
	"PunkCoin/core"
	"crypto/sha256"
	"fmt"
)

func main() {

	fmt.Println("hello world")

	minerName := "punk"
	address := common.Address(sha256.Sum256([]byte(minerName)))

	sendto := "sang"
	sendaddress := common.Address(sha256.Sum256([]byte(sendto)))

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



	fmt.Printf("nonce:%d hash:%x\n",nonce,byte)

	dag := core.NewDag()
	for k,_ := range dag.Dag{
		fmt.Println(dag.Dag[k].Hash)
	}
	fmt.Println(dag.Dag)

	miner := client.Newminer(false,address)

	fmt.Println("mining...")
	miner.SendTx(&sendaddress,5)



}
//e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855