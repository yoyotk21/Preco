package main

/*
#include "main.c"
*/
import "C"
import (
	"preco/dpfc/field"
)

func add(x, y []uint64) []field.FP {
	ans := make([]field.FP, len(x))
	// assuming x and y are of equal length
	for i := range x {
		ans[i] = field.Add(field.FP(x[i]), field.FP(y[i]))
	}

	return ans
}

func main() {
	name := "Joel"
	C.sayHi(C.CString(name))
	println(name)
}

// client := dpfc.ClientDPFInitialize()
// defer client.Free()

// keyA, keyB := client.GenDPFKeys(5, 64)

// server := dpfc.ServerDPFInitialize(client.PrfKey)
// defer server.Free()

// data := []uint64{1, 2, 3, 4, 5}

// data1 := []field.FP{field.FP(1), field.FP(1), field.FP(2), field.FP(3), field.FP(2)}
// data2 := []field.FP{field.FP(0), field.FP(1), field.FP(1), field.FP(1), field.FP(3)}

// ans0 := server.BatchEval(keyA, data)
// ans1 := server.BatchEval(keyB, data)

// for ind := range ans0 {
// 	data1[ind] = field.Add(field.Multiply(field.FP(ans0[ind]), field.FP(5)), data1[ind])
// }

// for ind := range ans1 {
// 	data2[ind] = field.Add(field.Multiply(field.FP(ans1[ind]), field.FP(5)), data2[ind])
// }

// fmt.Println(fieldAdd(data1, data2))
