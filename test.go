package main

import (
	"PunkCoin/client"
	"PunkCoin/common"
	"PunkCoin/core"
	"crypto/sha256"
	"fmt"
)

func main() {


	minerName := "punk"
	address := common.Address(sha256.Sum256([]byte(minerName)))

	sendto := "sang"
	sendaddress := common.Address(sha256.Sum256([]byte(sendto)))

	mainblock := "MainBlock"
	MB := common.BlockHash(sha256.Sum256([]byte(mainblock)))

	blockOne := "blockOne"
	B1 := common.BlockHash(sha256.Sum256([]byte(blockOne)))

	blockTwo := "blockTwo"
	B2 := common.BlockHash(sha256.Sum256([]byte(blockTwo)))
	TxInputs := []core.TxInput{}
	SendTo := []core.TxOutput{}

	amount := 11

	targetbits := 6

	byte,nonce := core.Pow(&address,&MB,&B1,&B2,TxInputs,SendTo,amount,targetbits)



	fmt.Printf("nonce:%d hash:%x\n",nonce,byte)

	dag := core.NewDag()
	for k,_ := range dag.Dag{
		fmt.Println("当前的主块哈希:\n",*dag.Dag[k].Hash)
	}

	miner := client.Newminer(false,address)

	fmt.Println("\n\nmining...")
	b := miner.SendTx(&sendaddress,5)


	fmt.Println("当前挖出的块的上个主块哈希:\n",*b.Mainblock)

}
//e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855