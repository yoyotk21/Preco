package client

import (
	"log"
	"net/rpc"
	"preco/common"
	"preco/dpfc"
	"preco/dpfc/field"
	"preco/prg"
)

type Client struct {
	Port1  string
	Port2 string
	Addr1 string
	Addr2 string
}

func listToField(li []uint64) []field.FP {
	fieldList := make([]field.FP, 0)

	for _, val := range li {
		fieldList = append(fieldList, field.FP(val))
	}

	return fieldList
}

// assumes both lists are of the same length
// might throw errors if not
func addFieldLists(li1, li2 []field.FP) []field.FP {
	res := make([]field.FP, 0)

	for i := range li1 {
		res = append(res, field.Add(li1[i], li2[i]))
	}

	return res
}

func (client *Client) Write(rowInd int, rowToAdd []uint64) {
	// connecting to both servers
	conn1, err := rpc.Dial("tcp", client.Addr1+":"+client.Port1)
	if err != nil {
		log.Fatal(err)
	}
	defer conn1.Close()

	conn2, err := rpc.Dial("tcp", client.Addr2+":"+client.Port2)
	if err != nil {
		log.Fatal(err)
	}
	defer conn2.Close()

	s := field.RandomFieldElement()

	clientDPF := dpfc.ClientDPFInitialize()

	arr := make([]uint64, common.Length)
	arr[rowInd] = uint64(s)

	keyA, keyB := clientDPF.GenDPFKeys(uint64(rowInd), 64) // cannot set index to s
	miniServer := dpfc.ServerDPFInitialize(clientDPF.PrfKey)

	sA := miniServer.BatchEval(keyA, arr)
	sB := miniServer.BatchEval(keyB, arr)

	m := addFieldLists(listToField(prg.Prg(sA[rowInd])), listToField(prg.Prg(sB[rowInd])))
	m = addFieldLists(m, listToField(rowToAdd))
	kABit, kBBit := clientDPF.GenDPFKeys(uint64(rowInd), 64)

	conn1.Call("Database.Write", common.Call{InitKey: clientDPF.PrfKey, Send: keyA, Bit: kABit, M: m}, nil)
	conn2.Call("Database.Write", common.Call{InitKey: clientDPF.PrfKey, Send: keyB, Bit: kBBit, M: m}, nil)
}
