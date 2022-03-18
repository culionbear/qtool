package persistence

import (
	"encoding/json"
	"io"
	"os"
)

//Config is persistence manager config file
type Config struct {
	AofPath  string `json:"aof_path,omitempty"`
	AofTimer int    `json:"aof_timer,omitempty"`
}

//NewConfig with file path
func NewConfig(path string, c *Config) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	buf, err := io.ReadAll(fp)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, c)
}
