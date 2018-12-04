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

func formatLoop(format string, a ...interface{}) (string, []interface{}) {

	var args []interface{}

	for i, v := range a {

		val, ok := v.(string)

		if !ok {
			args = append(args, v)
			continue
		}

		newformat, newval := calc(format, i, val)

		if newformat != format {
			format = newformat
		}
		args = append(args, newval)

	}

	return format, args
}

func calc(format string, index int, val string) (string, string) {

	var newval string

	list := strings.Split(format, "%")
	f := list[index+1]
	spos := strings.Index(f, "s")

	// case "%s"
	if spos <= 0 {
		return format, val
	}

	minus := false
	if f[0] == 0x2d {
		minus = true
	}

	// case "%-s"
	if minus && spos == 1 {
		return format, val
	}

	var padAndLen []string

	if minus {
		padAndLen = strings.Split(f[1:spos], ".")
	} else {
		padAndLen = strings.Split(f[0:spos], ".")
	}

	pad := -1
	length := -1

	pad, err := strconv.Atoi(padAndLen[0])
	if err != nil {
		// case "%.2s"
		pad = -1
	}

	if len(padAndLen) > 1 {
		length, err = strconv.Atoi(padAndLen[1])
		if err != nil {
			// case "%.s"
			length = -1
		}
	}

	if length >= 0 {
		n := 0
		for _, c := range val {
			if len(string(c)) == 1 {
				n += 1
			} else {
				n += 2
			}
			if n > length {
				break
			}
			newval += string(c)
		}
	} else {
		newval = val
	}

	if pad <= 0 && length <= 0 {
		return format, newval
	}

	if pad > 0 {
		pad = pad - zenkakuCnt(val)
		if pad <= 0 {
			pad = 1
		}
		if padAndLen[0][0] == 0x30 {
			list[index+1] = "0" + strconv.Itoa(pad) + "s" + f[spos+1:]
		} else {
			list[index+1] = strconv.Itoa(pad) + "s" + f[spos+1:]
		}
	}

	if minus && list[index+1][0] != 0x2d {
		list[index+1] = "-" + list[index+1]
	}

	return strings.Join(list, "%"), newval
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
