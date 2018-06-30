package goecho_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/umaumax/goecho"
)

func callBuiltinEcho(args []string) (out string, err error) {
	echoCmd := "echo " + `"` + strings.Join(args, " ") + `"`
	raw, err := exec.Command("zsh", "-c", echoCmd).Output()
	out = string(raw)
	return
}

func TestSimple(t *testing.T) {
	list := []string{
		"nanoha",
		`fate\chayate`,
		`fate \c hayate`,
	}
	for i := 0; i < 26; i++ {
		list = append(list, `\`+string('a'+i), `\`+string('A'+i))
	}
	for i := 0; i < 0x100; i++ {
		list = append(list, `\x`+fmt.Sprintf("%02x", i))
	}

	for _, v := range list {
		args := []string{v}
		got := goecho.Echo(false, false, args)
		want, err := callBuiltinEcho(args)
		if err != nil {
			t.Fatalf("builtin echo err:%s", err)
		}
		if got != want {
			t.Fatalf("test [%v], want [%v(%v)], but [%v(%v)]:", args, want, []byte(want), got, []byte(got))
		}
	}
}
