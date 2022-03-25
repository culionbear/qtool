package logs

import "testing"

func TestColor(t *testing.T) {
	t.Log(Sprint(Red, "213"))
}
