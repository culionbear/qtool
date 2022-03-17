package db

import (
	"github.com/culionbear/qtool/db/hash"
	"github.com/culionbear/qtool/persistence"
	"github.com/culionbear/qtool/template"
)

type Table interface {
	//Set value in table when value is not found before
	Set([]byte, template.Node) error
	//Set value in table
	SetX([]byte, template.Node)
	//Update value in table
	Update([]byte, template.Node) error
	//Get value in table with name
	Get([]byte) template.Node
	//Gets value list in table with name list
	Gets(...[]byte) []template.Node
	//Del node in table with key
	Del([]byte) error
	//Dels node int table with key list
	Dels(...[]byte) int
	//Regexp string to get value in table
	Regexp([]byte) [][]byte
	//Get Iterator with key
	Iterator([]byte) hash.Node
	//Range iterators
	Iterators([]byte, func(hash.Node) bool)
	//Rename src to dst
	Rename([]byte, []byte) error
	//Cover src to dst, if dst is not found then rename src to dst
	Cover([]byte, []byte) error
	//Exist key
	Exist([]byte) bool
}

var Manager Table

func InitDB(c persistence.Config) (err error) {
	Manager, err = hash.New(c)
	return err
}
