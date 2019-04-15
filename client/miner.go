package client

import (
	"PunkCoin/common"
	"PunkCoin/core"
	"PunkCoin/p2p"
	"errors"
	"fmt"
)

//如果是想作为矿工来运行 难度值就会变高 如果只是作为一名用户想要进行交易 则难度值会降低
//不过即便是用户运行着难度值低的程序 如果一旦计算出来的结果符合主块 也能获得奖励

type miner struct {

	address common.Address

	//是想作为用户还是矿工
	isminer bool

	//todo:这里要改成bits位数 用来减少内存占用
	powtargetbits int

	//检验机制里面已经会有主链和dag结构了
	MainChain *core.Mainchain
	//TxPool *core.TxPool
	Dag *core.Dag


	//检验机制
	Check *core.Check

	balance int
}

var m miner


//矿工初始化难度要求
func init(){

	//初始化主块链
	m.MainChain = core.NewMainChain()

	m.isminer = true
	m.powtargetbits = common.TargetForMBBits

	//初始化校验工具
	m.Check = core.NewCheck()
}

func Newminer(isminer bool,address common.Address) *miner{
	m.isminer = isminer
	m.address = address
	m.balance = 0
	return &m
}



//发送交易 参数：对方地址，转账金额
func (m *miner) SendTx(address *common.Address,amount int){

	//先检查是否有大于金额的余额
	if m.GetBalance() > amount{

		//用户的余额足够创建这笔交易的话 将所有可用输入找出 并创建输出
		txInputs,txOutputs := m.createTx(address,amount)

		//pow创建一个区块
		b,err := m.mineBlock(txInputs,txOutputs)
		fmt.Println(b)
		if err !=nil{
			fmt.Errorf(err.Error())
		}else {
			p2p.SendToNet(b) //将区块发送到网络中
		}
	}

}

//todo：这里还有需要添加的内容
func (m *miner) GetBalance() int {
	return m.balance
}

//切换用户模式和矿工模式 就是难度的差别
func (m *miner) ChangeMode(){
	if m.isminer{
		m.isminer = false
		m.powtargetbits = common.TargetForTxBits
	}else{
		m.isminer = true
		m.powtargetbits = common.TargetForMBBits
	}

}



//todo:这里还有需要添加的内容
//在自己所有可支付区块里面找到金额大于amount的所有可用块 构建输出 若所有输入大于amount则要构建一笔转回给自己的输出
//生成Txinput和Txoutput
func (m *miner) createTx(address *common.Address,amount int) ([]core.TxInput,[]core.TxOutput){
	output1 := core.TxOutput{}
	output2 := core.TxOutput{}
	outputs := []core.TxOutput{output1,output2}
	input1 := core.TxInput{}
	input2 := core.TxInput{}
	inputs := []core.TxInput{input1,input2}
	return inputs,outputs
}

//挖矿的时候如果不做交易 也可以挖空块 不过输出要有一笔是给自己的奖励
//挖出一个区块的过程 如果挑选的两个验证块不满足则重新挑选
func (m *miner) mineBlock(inputs []core.TxInput,outputs []core.TxOutput) (*core.Block,error) {

	var b *core.Block

	//获取当前区块链中最新的主块 //这里最后要改成获取到的是直接哈希值 而不是一个区块
	//如果验证交易失败的话 就重新进行一次
work:
	{
		mb := m.MainChain.Getlatest()
		fmt.Println("get latestMb ...",mb)
		//任意获取两笔其他的交易区块 //直接获取到他们的哈希值 而不是区块实例
		b1 := m.Dag.GetTransaction()
		b2 := m.Dag.GetTransaction()

		//检验两笔交易是否合理 如果验证了两笔交易没问题的话 创建区块
		if (m.Check.CheckoutTx(b1) && m.Check.CheckoutTx(b2)) {
			b = core.CreateBlock(&m.address,mb, b1, b2, inputs, outputs, m.powtargetbits)
			if b == nil{
				goto work
			}

		} else { //如果验证交易失败 说明那两笔交易区块不符合要求 不链入账本中

			goto work
		}

	}

	if b != nil {
		return b, nil
	} else {
		return nil, fmt.Errorf("mistake")
	}
}

//矿工接收一个区块 blocktype是1则为主块 否则为普通块
func (m *miner) ReceiveBlock() error{
	b := p2p.ReceiveFromNet()
	isok,blocktype := m.CheckBlock(b)
	if isok{
		if blocktype{
			m.MainChain.Add(b)
		}
		m.Dag.Add(b)
		return nil
	}
	return errors.New("this block is illegal")

}

func (m *miner) CheckBlock(b *core.Block) (bool,bool){

	//接收到一个区块 检验是否区块结构合理 确定区块类型
	isok,blocktype,err := m.Check.CheckoutBlock(b)
	if err != nil{
		return false,false
	}else{
		return isok,blocktype
	}

}



