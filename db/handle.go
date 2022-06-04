package db

import (
	"github.com/culionbear/qtool/db/hash"
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

//Set value in table when value is not found before
func Set(key []byte, value template.Node) *qerror.Error {
	return m.Set(key, value)
}

//Set value in table
func SetX(key []byte, value template.Node) {
	m.SetX(key, value)
}

//Update value in table
func Update(key []byte, value template.Node) *qerror.Error {
	return m.Update(key, value)
}

//Get value in table with name
func Get(key []byte) template.Node {
	return m.Get(key)
}

//Gets value list in table with name list
func Gets(keys ...[]byte) []template.Node {
	return m.Gets(keys...)
}

//Del node in table with key
func Del(key []byte) *qerror.Error {
	return m.Del(key)
}

//Dels node int table with key list
func Dels(keys ...[]byte) int {
	return m.Dels(keys...)
}

//Regexp string to get value in table
func Regexp(rex []byte) [][]byte {
	return m.Regexp(rex)
}

//Range iterators
func Iterators(callBack func(hash.Node) bool) {
	m.Iterators(callBack)
}

//Rename src to dst
func Rename(dst, src []byte) *qerror.Error {
	return m.Rename(dst, src)
}

//Cover src to dst, if dst is not found then rename src to dst
func Cover(dst, src []byte) *qerror.Error {
	return m.Cover(dst, src)
}

//Exist key
func Exist(key []byte) bool {
	return m.Exist(key)
}

//Size
func Size() int {
	return m.Size()
}
