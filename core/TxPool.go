package core

////未放入dag中的
//type TxPool struct {
//	Tx []*Block //最好改成先进先出
//}
//
//
////todo:这里还有内容需要补充
//func (tp *TxPool) GetTransaction() common.BlockHash {
//	block := tp.Tx[0]
//	tp.Tx = tp.Tx[1:]
//
//	return block.Hash
//}