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
	f, args := doFormat(format, a...)
	return fmt.Sprintf(f, args...)
}

func doFormat(fmt string, a ...interface{}) (string, []interface{}) {
	var args []interface{}

	for i, v := range a {
		val, ok := v.(string)
		if ok {
			format(&fmt, i, &val)
		}
		args = append(args, val)
	}
	return fmt, args
}

func format(fmt *string, index int, val *string) {
	fmts := strings.Split(*fmt, "%")
	fpos := index + 1
	f := fmts[fpos]
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

	if width >= 0 {
		*val = substrWithWidth(*val, width)
	}

	if padding <= 0 && width <= 0 {
		return
	}

	if padding > 0 {
		padding = padding - countWideChars(*val)
		if padding < 0 {
			padding = 0
		}
		fmts[fpos] = strconv.Itoa(padding) + f[spos:]
		if padAndLen[0][0] == 0x30 {
			fmts[fpos] = "0" + fmts[fpos]
		}
	}

	if minus {
		fmts[fpos] = "-" + fmts[fpos]
	}

	*fmt = strings.Join(fmts, "%")

	return
}

func charWidth(c rune) int {
	if len(string(c)) == 1 {
		return 1
	}
	return 2
}

func substrWithWidth(s string, width int) string {
	var (
		ret string
		n   int
	)
	for _, c := range s {
		n += charWidth(c)
		if n > width {
			break
		}
		ret += string(c)
	}
	return ret
}

func countWideChars(s string) int {
	n := 0
	for _, c := range s {
		n += charWidth(c)
	}
	return n - len([]rune(s))
}
