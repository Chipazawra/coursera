package curseraBasics

import (
	"bufio"
	"fmt"
	"io"
)

func IvokeDefer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend FIRST:", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend SECOND:", err)
			panic("second panic")
		}
	}()
	fmt.Println("Some userful work")
	panic("something bad happend")
	//return
}

func Uniq(in io.Reader, out io.Writer) error {

	input := bufio.NewScanner(in)
	var prev string

	for input.Scan() {
		txt := input.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt
		fmt.Fprintln(out, txt)
	}

	return nil

}
