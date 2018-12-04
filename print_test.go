package wfmt

import (
	"testing"
)

func TestSprintf(t *testing.T) {

	var fmtTests = []struct {
		fmt string
		val string
		out string
	}{
		// package fmt tests
		{"%s", "abc", "abc"},
		{"%5s", "abc", "  abc"},
		{"%-5s", "abc", "abc  "},
		{"%05s", "abc", "00abc"},
		{"%5s", "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"},
		{"%.5s", "abcdefghijklmnopqrstuvwxyz", "abcde"},
		{"%.5s", "日本語日本語", "日本"},
		{"%.10s", "日本語日本語", "日本語日本"},

		{"%s", "abc", "abc"},
		{"%.s", "abc", ""},
		{"%0.s", "abc", ""},
		{"%.0s", "abc", ""},
		{"%0.0s", "abc", ""},
		{"%1.0s", "abc", " "},
		{"%0.1s", "abc", "a"},
		{"%1.1s", "abc", "a"},
		{"%2.1s", "abc", " a"},
		{"%1.2s", "abc", "ab"},
		{"%-s", "abc", "abc"},
		{"%-.s", "abc", ""},
		{"%-0.s", "abc", ""},
		{"%-.0s", "abc", ""},
		{"%-0.0s", "abc", ""},
		{"%-1.0s", "abc", " "},
		{"%-0.1s", "abc", "a"},
		{"%-1.1s", "abc", "a"},
		{"%-2.1s", "abc", "a "},
		{"%-1.2s", "abc", "ab"},
		{"%s", "あいう", "あいう"},
		{"%1.s", "あいう", "あいう"},
		{"%.1s", "あいう", ""},
		{"%1.1s", "あいう", " "},
		{"%5s", "あいう", "あいう"},
		{"%7s", "あいう", " あいう"},
		{"%7s", "あいう", " あいう"},
		{"%-7s", "あいう", "あいう "},
	}

	for _, tt := range fmtTests {
		t.Run(tt.fmt, func(t *testing.T) {
			s := Sprintf(tt.fmt, tt.val)
			if s != tt.out {
				t.Errorf("got %#v, want %#v", s, tt.out)
			}
		})
	}
}
