package persistence

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"strings"
)

//Config is persistence manager config file
type Config struct {
	XMLName		xml.Name	`json:"-" xml:"config"`
	AofPath		string		`json:"aof_path,omitempty" xml:"aof_path,omitempty"`
	AofTimer	int			`json:"aof_timer,omitempty" xml:"aof_timer,omitempty"`
}

//NewConfig with file path
func NewConfig(path string, c *Config) error {
	var f func([]byte, interface{}) error
	if strings.LastIndex(path, ".json") != -1 {
		f = json.Unmarshal
	} else if strings.LastIndex(path, ".xml") != -1 {
		f = xml.Unmarshal
	} else {
		return errors.New("file is illegal")
	}
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	buf, err := io.ReadAll(fp)
	if err != nil {
		return err
	}
	return f(buf, c)
}