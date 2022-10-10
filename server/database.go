package server

import (
	"preco/common"
	"preco/dpfc"
	"preco/prg"
)

type Database struct {
	data [][]uint64
}

func NewDatabase(data [][]uint64) *Database {
	d := new(Database)
	d.data = data
	return d
}

func (d *Database) Write(data common.Call, reply *uint64) error {
	serverDPF := dpfc.ServerDPFInitialize(data.InitKey)
	s := make([]uint64, 0)
	for _, val := range serverDPF.BatchEval(data.Send, d.data) {
		s = append(s, prg.Prg(val))
	}
	return nil
}
