package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

func main(){
	im := big.NewInt(math.MaxInt64)
	fmt.Printf("%b\n",im)

	an := big.NewInt(math.MinInt64)
	an.Rsh(im,4)

	fmt.Printf("%b\n", an)
	//fmt.Printf("%x\n",im>>4)

	target := big.NewInt(1);//初始化目标整数
	fmt.Printf("%b\n",target)
	target.Lsh(target,uint(256-2)) //数据转换，小于这个数的都是符合条件的哈希值,二进制移位232位，16进制以为58位
	fmt.Printf("%b\n",target)

	target2 := big.NewInt(1);//初始化目标整数
	fmt.Printf("%b\n",target2)
	target2.Lsh(target2,uint(256-5)) //数据转换，小于这个数的都是符合条件的哈希值,二进制移位232位，16进制以为58位
	fmt.Printf("%b\n",target2)

	fmt.Println(target.Cmp(target2))



}


type ProofOfWork struct {
	block *Block //区块
	target *big.Int ////big是go语言里的精确的大数的标准库，存储计算哈希对比的特定整数
}

//创建一个工作量证明的挖矿对象
func NewProofOfWork(block *Block)*ProofOfWork{
	target := big.NewInt(1);//初始化目标整数
	fmt.Printf("%x\n",target)
	target.Lsh(target,uint(256-targetBits)) //数据转换，小于这个数的都是符合条件的哈希值,二进制移位232位，16进制以为58位
	fmt.Printf("%x\n",target)
	pow:=&ProofOfWork{block,target} //创建一个对象
	return pow
}
//准备数据，进行挖矿计算
func (pow * ProofOfWork) prepareData(nonce int)[]byte{
	data := bytes.Join(
		[][]byte{
			pow.block.PreBlockHash,//上一块哈希
			pow.block.Data,//当前数据
			IntToHex(pow.block.Timestamp),//十六进制
			IntToHex(int64(targetBits)),//位数，十六进制
			IntToHex(int64(nonce)),//保存工作量的证明
		},[]byte{},
	)
	return data
}
//挖矿执行
func (pow * ProofOfWork) Run() (int,[]byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce :=0
	fmt.Printf("当前挖矿计算的数据%s",pow.block.Data)
	for nonce<maxNonce{
		data :=pow.prepareData(nonce)//准备好的数据
		hash = sha256.Sum256(data)//计算出哈希
		fmt.Printf("\r%x",hash)//打印显示哈希

		hashInt.SetBytes(hash[:])//获取要对比的数据
		//10000000000000000000000000000000000000000000000000000000000
		//000000ce7bdc24fb9b3ebc36939c21b3053013acc87dc890e673036e575c2a00
		if hashInt.Cmp(pow.target)==-1{ //挖矿的校验
			break
		}else {
			nonce++
		}
	}
	fmt.Println("\n\n")
	return nonce,hash[:]//nonce 解题的答案，hash当前的哈希
}
//校验挖矿是不是真的成功
func (pow * ProofOfWork) Validate()bool{
	var hashInt big.Int
	data :=pow.prepareData(pow.block.Nonce)//准备好的数据
	hash := sha256.Sum256(data)//计算出哈希
	hashInt.SetBytes(hash[:])//获取对比的数据

	isValid :=(hashInt.Cmp(pow.target))==-1//校验数据
	return isValid

}
