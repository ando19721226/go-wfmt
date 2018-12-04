package wfmt

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	buf := []byte(Sprintf(format, a...))
	return w.Write(buf)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Sprintf(format string, a ...interface{}) string {
	f, args := formatLoop(format, a...)
	return fmt.Sprintf(f, args...)
}

func formatLoop(fmt string, a ...interface{}) (string, []interface{}) {

	var args []interface{}

	for i, v := range a {

		val, ok := v.(string)

		if !ok {
			args = append(args, v)
			continue
		}

		format(&fmt, i, &val)

		args = append(args, val)

	}

	return fmt, args
}

func format(fmt *string, index int, val *string) {

	list := strings.Split(*fmt, "%")
	f := list[index+1]
	spos := strings.Index(f, "s")

	// case "%s"
	if spos <= 0 {
		return
	}

	minus := false
	if f[0] == 0x2d {
		minus = true
	}

	// case "%-s"
	if minus && spos == 1 {
		return
	}

	var padAndLen []string

	if minus {
		padAndLen = strings.Split(f[1:spos], ".")
	} else {
		padAndLen = strings.Split(f[0:spos], ".")
	}

	padding := -1
	width := -1

	padding, err := strconv.Atoi(padAndLen[0])
	if err != nil {
		// case "%.2s"
		padding = -1
	}

	if len(padAndLen) > 1 {
		width, err = strconv.Atoi(padAndLen[1])
		if err != nil {
			// case "%.s"
			width = -1
		}
	}

	var newval string

	if width >= 0 {
		n := 0
		for _, c := range *val {
			if len(string(c)) == 1 {
				n += 1
			} else {
				n += 2
			}
			if n > width {
				break
			}
			newval += string(c)
		}
	} else {
		newval = *val
	}

	*val = newval

	if padding <= 0 && width <= 0 {
		return
	}

	if padding > 0 {
		padding = padding - zenkakuCnt(*val)
		if padding <= 0 {
			padding = 1
		}
		list[index+1] = strconv.Itoa(padding) + "s" + f[spos+1:]
		if padAndLen[0][0] == 0x30 {
			list[index+1] = "0" + list[index+1]
		}
	}

	if minus && list[index+1][0] != 0x2d {
		list[index+1] = "-" + list[index+1]
	}

	*fmt = strings.Join(list, "%")

	return
}

func zenkakuCnt(s string) int {
	cnt := 0
	for _, c := range s {
		if len(string(c)) == 1 {
			cnt += 1
		} else {
			cnt += 2
		}
	}
	return cnt - len([]rune(s))
}
