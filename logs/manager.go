package logs

import (
	"log"
	"os"
)

const (
	errPrefix  = "[ERROR]"
	infoPrefix = "[INFO]"
)

type Manager struct {
	logger *log.Logger
}

func New() *Manager {
	return &Manager{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (m *Manager) Redirect(path string) error {
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	m.logger.SetOutput(fp)
	return nil
}

func (m *Manager) PrintInfo(v ...any) {
	m.logger.Println(v...)
}

func (m *Manager) PrintError(v ...any) {
	m.logger.Println(Sprint(Red, v...))
}
