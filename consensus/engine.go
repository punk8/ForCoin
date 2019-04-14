package consensus

import (
	"PunkCoin/common"
	"PunkCoin/core"
	"bytes"
	"math/big"
)

type PoW struct {

}


//pow调用basetool的哈希函数来进行哈希运算


func prepareData(miner *common.Address,Mainblock *common.BlockHash,BlockOne,BlockTwo *common.BlockHash,TxInputs []core.TxInput,SendTo []core.TxOutput,amount int,target *big.Int,nonce *big.Int) []byte {
	data := bytes.Join(
		[][]byte{

		},[]byte{},
	)
	return data
}

func Pow(miner *common.Address,Mainblock *common.BlockHash,BlockOne,BlockTwo *common.BlockHash,TxInputs []core.TxInput,SendTo []core.TxOutput,amount int,target *big.Int) (common.BlockHash,*big.Int) {
	nonce := big.NewInt(0)
	return [32]byte{},nonce
}

