package core

import (
	"PunkCoin/common"
	"bytes"
	"fmt"
	"math/big"
	"sync"
)

type Check struct {
	MainChain *Mainchain
	Dag *Dag
}

var c = &Check{}

//校验工具初始化
func init(){

	//初始化主块链
	c.MainChain = NewMainChain()

	c.Dag = NewDag()
}


func NewCheck() *Check{
	return c
}


//检验区块结构是否没有问题 哈希值是否没问题并且返回 是主块还是普通块 主块为1 普通块为0
//检验区块的结构包括 哈希值是否满足难度 区块的金额是否等于交易的输入 是否等于交易的输出总和 如果是一个主块的话 会有一笔额外的奖励
func (c *Check) CheckoutBlock(block *Block) (isok bool,blocktype bool,err error){
	//添加线程组
	var wg sync.WaitGroup

	//默认为普通块
	blocktype = false
	//输入是否可用
	tiCanUse := true
	//总共发送给人的金额
	spend := 0
	//输出块总金额
	receive := 0

	//如果该块已经被可信主块连接了 那么不用检验了
	isok,blocktype,err = c.hasBeenCheckByTrustedMB(block)
	if isok{
		return isok,blocktype,err
	}
	
	//如果该块离当前最新的主块已经有M个距离了那就不用检测了 直接返回不符合
	isok,blocktype,err = c.CheckDisBetweenMB(block)
	if !isok{
		return isok,blocktype,err
	}


	//区块哈希值没问题 检测输入输出是否是存在的区块 检测输入输出是否已经被使用过了
 	if bytes.Equal(HashBlock(block),block.Hash[:]){

		//确定区块类型 返回true则为主块
		blocktype = c.DetermineBlockType(block)

		//如果是主块 检查下第一笔输出是否合理
		if blocktype{
			c.isCoinBase(block)
		}

 		//检验输入
 		for i:=0;i<len(block.TxInputs);i++{
			txi := block.TxInputs[i]
			if c.isBlockExit(txi.InputAddress){
				//在数据库里面查询这个输入块 看是否已经被消耗 如果这个输入是一笔矿工奖励的话还要检验是否达到可使用要求
				//如果不是一笔矿工交易的话 index=0的这笔交易无效

				//如果输入可以被使用 输入检查 要检查那个输入块的输出特别是第一笔 即矿工奖励的输出
				if c.InputCanUse(txi){
					continue
				}else{
					tiCanUse = false
				}
			}else {
				tiCanUse = false
			}
		}

 		if !tiCanUse{
 			return false,blocktype,fmt.Errorf("input can't be used")
		}else{
			spend += GetBalance(block.TxInputs)

		}

 		//if spend!=total{
 		//	return false,blocktype,fmt.Errorf("input isn't equal to the block amount")
		//}



		//检验输出
		for i:=0;i<len(block.SendTo);i++{
			txo := block.SendTo[i]
			if c.isAddressExist(txo.OutputAddress){
				receive += txo.Amount
			}else{
				return false,blocktype,fmt.Errorf("output %d is not exist",i)
			}
		}

		if spend < receive{
			return false,blocktype,fmt.Errorf("input isn't equal to the output")
		}

		wg.Add(2)
		b1hash := block.BlockOne
		b2hash := block.BlockTwo
		
		b1 := c.Dag.FindBlockByHash(b1hash)
		b2 := c.Dag.FindBlockByHash(b2hash)

		go func(){
			defer wg.Done()
			c.CheckoutBlock(b1)
		}()
		
		go func(){
			defer wg.Done()
			c.CheckoutBlock(b2)
		}()




		//满足条件的时候
		return true,blocktype,nil

	}


	return false,blocktype,fmt.Errorf("block hash is illegal")
}

func (c *Check) CheckoutTx(block *common.BlockHash) bool {
	return true
}


//检验区块是否存在
func (c *Check) isBlockExit(hash *common.BlockHash) bool{
	return false
}

func (c *Check) InputCanUse(input TxInput)bool{
	canUse := false


	inputindex := input.Index //用来判断是矿工奖励还是普通交易

	inputBlockHash := input.InputAddress

	inputBlock := c.Dag.FindBlockByHash(inputBlockHash)

	inputBlockType := c.DetermineBlockType(inputBlock)

	//如果输入区块不是主块 且输入是第一笔 即使用了矿工奖励做输入 则该笔输入是不可使用的
	if !inputBlockType && inputindex == 0{
		canUse = false
		return canUse
	}


	if c.traceBlock(input){
		canUse = true
		return canUse
	}

	return canUse
}


//用于确定区块类型 利用区块哈希值和目标值对比
func (c *Check) DetermineBlockType(block *Block) bool {
	var hashInt big.Int
	hashInt.SetBytes(block.Hash[:])//获取要对比的数据

	if hashInt.Cmp(c.MainChain.TargetforMb) == -1{
		return true
	}
	return false
}

func (c *Check) traceBlock(input TxInput)bool{
	return false
}

func(c *Check) isAddressExist(address *common.Address) bool {
	return false
}

//检查第一笔交易是不是合法的矿工奖励 包括检查输出第一笔交易是不是输出给区块使用者 输出价格是不是coinbase奖励的金额
//todo：后期改进 不一定限制奖励接收者一定要主块挖矿者，金额也应该逐渐减半 最终只有有限货币
func (c *Check) isCoinBase(block *Block) bool{
	coinBase := block.SendTo[0]
	if coinBase.OutputAddress == block.MinerAddress{
		if coinBase.Amount == common.CoinBase{
			return true
		}
		return false
	}
	return false
}



func (c *Check) hasBeenCheckByTrustedMB(block *Block) (isok bool,blocktype bool,err error){
	return false, false, nil
}


//检测当前区块连接的主块离当前主块是不是太远了 如果距离已经大于MaxDistance 那么认为这笔交易已经过时了
func (c *Check) CheckDisBetweenMB(block *Block) (isok,blocktype bool,err error){
	currenthash := c.MainChain.Getlatest()
	currentheight := c.Dag.FindBlockByHash(currenthash).Number

	blockheight := block.Number

	dis := big.NewInt(0)

	distance := dis.Sub(currentheight,blockheight)
	if distance.Int64()>common.MaxDistance{
		return false,false,nil
	}

	blocktype = c.DetermineBlockType(block)

	return true, blocktype, nil
}