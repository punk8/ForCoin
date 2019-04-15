package core

import (
	"PunkCoin/common"
	"crypto/sha256"
	"math/big"
)

type Mainchain struct {

	ChainBlocks [](*Block)
	TargetforTx *big.Int
	TargetforMb *big.Int
}


type Dag struct {
	Dag map[*common.BlockHash]*Block //存放普通区块
}




//var targetforTx = big.NewInt(math.MaxInt64)
//var targetforMb = big.NewInt(math.MaxInt64)
var mc = &Mainchain{}
var dag = &Dag{}


//返回一个主块链实例
func NewMainChain() *Mainchain {
	return mc
}
func NewDag() *Dag{
	return dag
}

//初始化一个主块链
func init(){

	//genesis:=
	genesis := mc.CreateGenesisMainBlock()
	//fmt.Println(genesis.Hash)
	mc.Settarget(2,5)
	mc.ChainBlocks =  [](*Block){}
	mc.Add(genesis)
	dag.Dag = make(map[*common.BlockHash]*Block)
	dag.Add(genesis)

}

//全局的难度值 会有一个线程通过计算区块增长速度 来修改难度值 这里暂时设置为 前导零两个
func (mc *Mainchain) Settarget(n1,n2 uint) {

	mc.TargetforTx = big.NewInt(1)
	mc.TargetforMb = big.NewInt(1)

	//TargetforTx > TragetForMb 所以小于targetforTx的比较容易
	mc.TargetforTx.Rsh(mc.TargetforTx, uint(256-n1))
	mc.TargetforMb.Rsh(mc.TargetforMb, uint(256-n2))
}



func (mc *Mainchain) Add(block *Block){
	mc.ChainBlocks = append(mc.ChainBlocks, block)
}


func (mc *Mainchain) Getlatest() *common.BlockHash {
	return mc.ChainBlocks[len(mc.ChainBlocks)-1].Hash
}



func (mc *Mainchain) CreateGenesisMainBlock() *Block {

	//todo：
	//这里填写初始块输出的两个地址
	//tout1 := TxOutput{}
	//tout2 := TxOutput{}


	//sendto := []TxOutput{tout1,tout2}

	minerName := "punk"
	address := common.Address(sha256.Sum256([]byte(minerName)))

	mainblock := "MainBlock"
	MB := common.BlockHash(sha256.Sum256([]byte(mainblock)))

	blockOne := "blockOne"
	B1 := common.BlockHash(sha256.Sum256([]byte(blockOne)))

	blockTwo := "blockTwo"
	B2 := common.BlockHash(sha256.Sum256([]byte(blockTwo)))
	TxInputs := []TxInput{}
	SendTo := []TxOutput{}

	amount := 10

	targetbits := 3

	genesishash := [32]byte{2 ,246 ,70 ,22 ,170 ,146 ,154 ,243 ,166 ,14 ,199 ,197 ,155 ,12 ,234 ,103 ,143 ,170, 13, 200, 78, 52, 115, 152, 131, 47, 138, 100, 235 ,204,9,54}

	hash := common.BlockHash(genesishash)

	//hash,nonce := Pow([32]byte{0},[32]byte{0}, [32]byte{0}, [32]byte{0}, nil, sendto, 0, mc.TargetforMb)
	genesis := NewBlock(&address,&MB,&B1,&B2,TxInputs,SendTo,amount,targetbits,big.NewInt(1),&hash)
	//
	//mc.Add(genesis)
	//return genesis
	return genesis
}



func (d *Dag) Add(b *Block){

	d.Dag[b.Hash] = b

}

func (d *Dag) FindBlockByHash(hash *common.BlockHash) *Block {
	return d.Dag[hash]
}

//todo:这里应该有个算法来选择dag中的哪两笔交易应该被选中
func (d *Dag) GetTransaction() (*common.BlockHash){
	return nil
}