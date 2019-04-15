package core

type Account struct {
	MainChain *Mainchain
	Dag *Dag
}

var acc = &Account{}

//账户余额初始化
func init(){

	//初始化主块链
	acc.MainChain = NewMainChain()

	acc.Dag = NewDag()
}



func GetBalance(txins []TxInput) int {

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


//去掉第一笔交易
func Calculate(txouts[]TxOutput) int{

	out := 0

	for i:=1;i<len(txouts);i++{
		txout := txouts[i]
		out += txout.Amount

	}
	return out
}


