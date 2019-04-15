package core

import (
	"PunkCoin/common"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type PoW struct {

}


//pow调用basetool的哈希函数来进行哈希运算


func prepareData(miner *common.Address,Mainblock *common.BlockHash,BlockOne,BlockTwo *common.BlockHash,TxInputs []TxInput,SendTo []TxOutput,amount int,targetbits int,nonce int) []byte {

	txinHash := [][]byte{}
	for _,tx := range(TxInputs){
		txinHash = append(txinHash, tx.ToHash())
	}
	txin := bytes.Join(txinHash,[]byte{})

	txoutHash := [][]byte{}
	for _,txout := range(SendTo){
		txoutHash = append(txoutHash,txout.ToHash())
	}
	txout := bytes.Join(txoutHash,[]byte{})

	data := bytes.Join(
		[][]byte{
			[]byte((*miner)[:]),
			[]byte((*Mainblock)[:]),
			//[]byte((*BlockOne)[:]),
			//[]byte((*BlockTwo)[:]),
			common.IntToHex(int64(amount)),
			common.IntToHex(int64(targetbits)),
			common.IntToHex(int64(nonce)),
		},[]byte{},
	)

	data = bytes.Join([][]byte{data,txin,txout},[]byte{})

	return data
}

func Pow(miner *common.Address,Mainblock *common.BlockHash,BlockOne,BlockTwo *common.BlockHash,TxInputs []TxInput,SendTo []TxOutput,amount int,targetbits int) (common.BlockHash,*big.Int) {

	//return [32]byte{},nonce


	var hashInt big.Int
	var hash [32]byte
	nonce := 0


	Target := big.NewInt(1)

	//TargetforTx > TragetForMb 所以小于targetforTx的比较容易
	Target.Lsh(Target, uint(256-targetbits))

	fmt.Println("当前难度:\n",Target)


	for nonce < common.MaxNonce{
		data := prepareData(miner,Mainblock,BlockOne,BlockTwo,TxInputs,SendTo,amount,targetbits,nonce)//准备好的数据
		//fmt.Println(data)
		hash = sha256.Sum256(data)//计算出哈希
		fmt.Printf("[mining...] :%x\n",hash)//打印显示哈希

		hashInt.SetBytes(hash[:])//获取要对比的数据
		//10000000000000000000000000000000000000000000000000000000000
		//000000ce7bdc24fb9b3ebc36939c21b3053013acc87dc890e673036e575c2a00
		if hashInt.Cmp(Target)==-1{ //挖矿的校验
			break
		}else {
			nonce+=1
		}
	}
	return hash,big.NewInt(int64(nonce))//nonce 解题的答案，hash当前的哈希
}


