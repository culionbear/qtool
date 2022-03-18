package db

import (
	"github.com/culionbear/qtool/db/hash"
	"github.com/culionbear/qtool/persistence"
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type Table interface {
	//Set value in table when value is not found before
	Set([]byte, template.Node) qerror.Error
	//Set value in table
	SetX([]byte, template.Node)
	//Update value in table
	Update([]byte, template.Node) qerror.Error
	//Get value in table with name
	Get([]byte) template.Node
	//Gets value list in table with name list
	Gets(...[]byte) []template.Node
	//Del node in table with key
	Del([]byte) qerror.Error
	//Dels node int table with key list
	Dels(...[]byte) int
	//Regexp string to get value in table
	Regexp([]byte) [][]byte
	//Get Iterator with key
	Iterator([]byte) hash.Node
	//Range iterators
	Iterators([]byte, func(hash.Node) bool)
	//Rename src to dst
	Rename([]byte, []byte) qerror.Error
	//Cover src to dst, if dst is not found then rename src to dst
	Cover([]byte, []byte) qerror.Error
	//Exist key
	Exist([]byte) bool
}

var Manager Table

func InitDB(c persistence.Config) (err error) {
	Manager, err = hash.New(c)
	return err
}
