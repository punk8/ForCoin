package core

import (
	"PunkCoin/common"
	"bytes"
	"fmt"
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
	//默认为普通块
	blocktype = false

	tiCanUse := true



	//总共发送给人的金额
	spend := 0



	//输出块总金额
	receive := 0



	//区块哈希值没问题 检测输入输出是否是存在的区块 检测输入输出是否已经被使用过了
 	if bytes.Equal(HashBlock(block),block.Hash[:]){

		//确定区块类型
		blocktype = c.DetermineBlockType(block)

		//如果是主块 检查下第一笔输出是否合理
		if blocktype{
			c.isCoinBase(block.SendTo[0])
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
	return false
}

func (c *Check) traceBlock(input TxInput)bool{
	return false
}

func(c *Check) isAddressExist(address *common.Address) bool {
	return false
}

//检查第一笔交易是不是合法的矿工奖励
func (c *Check) isCoinBase(txouts TxOutput) bool{
	return false
}