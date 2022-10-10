package common

import (
	"preco/dpfc"
	"preco/dpfc/field"
)

type Call struct {
	InitKey dpfc.PrfKey
	Send    *dpfc.DPFKey
	Bit     *dpfc.DPFKey
	M       []field.FP
}
