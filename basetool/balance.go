package basetool

import (
	"PunkCoin/core"
)

type Account struct {
	MainChain *core.Mainchain
	Dag *core.Dag
}

var acc = &Account{}

//账户余额初始化
func init(){

	//初始化主块链
	acc.MainChain = core.NewMainChain()

	acc.Dag = core.NewDag()
}



func GetBalance(txins []core.TxInput) int {

	balance := 0

	for i:=0;i<len(txins);i++{
		tx := txins[i]
		b := acc.Dag.FindBlockByHash(tx.InputAddress)
		//计算这些utxo的输出加起来的余额
		prevtxout := b.SendTo[tx.Index]
		balance += prevtxout.Amount
	}

	return balance
}

func Calculate(txouts[]core.TxOutput) int{

	out := 0

	for i:=0;i<len(txouts);i++{
		txout := txouts[i]
		out += txout.Amount

	}
	return out
}
