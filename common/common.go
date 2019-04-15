package common

import (
	"math"
)

type BlockHash [32]byte
type Address [32]byte

const MaxNonce = math.MaxInt64


const TargetForTxBits = 2
const TargetForMBBits = 9