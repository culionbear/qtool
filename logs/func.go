package logs

var m = New()

func Redirect(path string) error {
	return m.Redirect(path)
}

func PrintInfo(v ...any) {
	m.PrintInfo(v...)
}

func PrintError(v ...any) {
	m.PrintError(v...)
}