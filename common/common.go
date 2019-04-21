package common

import (
	"math"
)

type BlockHash [32]byte
type Address [32]byte

const MaxNonce = math.MaxInt64


const TargetForTxBits = 2
const TargetForMBBits = 9


const CoinBase = 25

//最大距离十个主块 超过十个主块 属于年老区块 如果想要进入主链只能重新发送
const MaxDistance = 10