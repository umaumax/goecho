package goecho

import (
	"regexp"
	"strconv"
)

func EscapeBackslash(v string) (ret string) {
	runes := []rune(v)
	n := len(runes)
	for i := 0; i < n; i++ {
		r := runes[i]
		if r == '\\' {
			if i+1 < n {
				i++
				r = runes[i]
				switch r {
				case '\\':
				case 'a':
					r = rune(0x07)
				case 'b':
					r = rune(0x08)
				case 'c':
					return
				case 'e':
					r = rune(0x1b)
				case 'f':
					r = rune(0x0c)
				case 'n':
					r = rune(0x0a)
				case 'r':
					r = rune(0x0d)
				case 't':
					r = rune(0x09)
				case 'u', 'U':
					r = rune(0x00)
				case 'v':
					r = rune(0x0b)
				case 'x':
					r = rune(0x00)
					i++
					numStr := regexp.MustCompile(`^[0-9A-Fa-f][0-9A-Fa-f]`).FindString(string(runes[i:]))
					v, _ := strconv.ParseInt(numStr, 16, 64)
					if numStr != "" {
						i += len(numStr)
						r = rune(v)
					}
					i--
				case '0':
					r = rune(0x00)
					i++
					numStr := regexp.MustCompile(`^[0-7][0-7]?[0-7]?`).FindString(string(runes[i:]))
					if numStr != "" {
						v, _ := strconv.ParseInt(numStr, 8, 64)
						i += len(numStr)
						r = rune(v)
					}
					i--
				default:
					r = '\\'
					i--
				}
			}
		}
		ret = string(append([]byte(ret), byte(r)))
	}
	return
}
