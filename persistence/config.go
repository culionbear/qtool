package persistence

//Config is persistence manager config file
type Config struct {
	QdbPath  string `json:"qdb_path,omitempty"`
	QdbTimer int    `json:"qdb_timer,omitempty"`
}
