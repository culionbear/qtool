package persistence

//Config is persistence manager config file
type Config struct {
	AofPath  string `json:"aof_path,omitempty"`
	AofTimer int    `json:"aof_timer,omitempty"`
}
