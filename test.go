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

	txout1 := core.TxOutput{}

	b := txout1.OutputAddress == nil

	fmt.Println(b)


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
	txout := core.TxOutput{OutputAddress:&address,Amount:10}
	SendTo := []core.TxOutput{txout}

	fmt.Println("hello world")



	targetbits := 4


	fmt.Println("hello world")



	byte,nonce := core.Pow(&address,&MB,&B1,&B2,TxInputs,SendTo,targetbits)



	fmt.Printf("nonce:%d hash:%x\n",nonce,byte)

	dag := core.NewDag()
	for k,_ := range dag.Dag{
		fmt.Println(dag.Dag[k].Hash)
		fmt.Println(k)

	}
	fmt.Println(dag.Dag)

	miner := client.Newminer(false,address)

	fmt.Println("mining...")
	miner.SendTx(&sendaddress,5)

	//minerName := "punk"
	//address := common.Address(sha256.Sum256([]byte(minerName)))
	//txin := []core.TxInput{}
	//txout := core.TxOutput{OutputAddress:&address,Amount:10}
	//txouts := []core.TxOutput{txout}
	//b := core.GetBalance(txin)
	//c := core.Calculate(txouts)
	//fmt.Println(b,c)


}
//e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855