package logs

import "fmt"

const (
	Black = 30 + iota
	Red
	Green
	Yellow
	Blue
	Purple
	DarkGreen
	White
)

func Sprint(color int, v ...any) string {
	str := ""
	for i := 0; i < len(v); i++ {
		if i != 0 {
			str += " "
		}
		str += "%v"
	}
	str = "\x1b[%dm" + str + " \x1b[0m"
	return fmt.Sprintf(str, append([]any{color}, v...)...)
}
