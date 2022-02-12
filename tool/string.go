package tool

import "bytes"

func Joint(str ...string) string {
	var writer bytes.Buffer
	for _, v := range str {
		writer.WriteString(v)
	}
	return writer.String()
}