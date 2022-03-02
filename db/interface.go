package db

import (
	"github.com/culionbear/qtool/db/iterator"
	"github.com/culionbear/qtool/db/hash"
	"github.com/culionbear/qtool/template"
	"github.com/culionbear/qtool/qerror"
)

type Table interface {
	//Set value in table when value is not found before
	Set([]byte, template.Node) qerror.Error
	//Set value in table
	SetX([]byte, template.Node)
	//Get value in table with name
	Get([]byte) (template.Node, qerror.Error)
	//Del value in table with key list
	Del(...[]byte) int
	//Regexp string to get value in table
	Regexp([]byte) ([]template.Node, qerror.Error)
	//Get Iterator with key
	Iterator([]byte) (iterator.Node, qerror.Error)
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

func Set(key []byte, value template.Node) qerror.Error {
	return Manager.Set(key, value)
}

func SetX(key []byte, value template.Node) {
	Manager.SetX(key, value)
}

func Get(key []byte) (template.Node, qerror.Error) {
	return Manager.Get(key)
}

func Del(key... []byte) int {
	return Manager.Del(key...)
}

func Regexp(str []byte) ([]template.Node, qerror.Error) {
	return Manager.Regexp(str)
}

func Iterator(key []byte) (iterator.Node, qerror.Error) {
	return Manager.Iterator(key)
}

func Rename(dst, src []byte) qerror.Error {
	return Manager.Rename(dst, src)
}

func Cover(dst, src []byte) {
	Manager.Cover(dst, src)
}

func Exist(key []byte) bool {
	return Manager.Exist(key)
}