package curseraBasics

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestIvokeDefer(t *testing.T) {
	IvokeDefer()
}

var testOk = `1
2
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestUniq(t *testing.T) {

	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	if err != nil {
		t.Errorf("test fo OK Failed")
	}
	if out.String() != testOkResult {
		t.Errorf("test result no match")
	}
}
