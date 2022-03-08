package db

import (
	"github.com/culionbear/qtool/db/hash"
	"github.com/culionbear/qtool/template"
	"github.com/culionbear/qtool/qerror"
)

type Table interface {
	//Set value in table when value is not found before
	Set([]byte, template.Node) qerror.Error
	//Set value in table
	SetX([]byte, template.Node)
	//Update value in table
	Update([]byte, template.Node) qerror.Error
	//Get value in table with name
	Get([]byte) (template.Node, qerror.Error)
	//Gets value list in table with name list
	Gets(...[]byte) ([]template.Node)
	//Del value in table with key list
	Del(...[]byte) int
	//Regexp string to get value in table
	Regexp([]byte) ([]template.Node, qerror.Error)
	//Get Iterator with key
	Iterator([]byte) (hash.Node, qerror.Error)
	//Range iterators
	Iterators([]byte, func(hash.Node) bool) qerror.Error
	//Rename src to dst
	Rename([]byte, []byte) qerror.Error
	//Cover src to dst, if dst is not found then rename src to dst
	Cover([]byte, []byte) qerror.Error
	//key is exist or not
	Exist([]byte) bool
}

var Manager Table

func init() {
	Manager = hash.New()
}