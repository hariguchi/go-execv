package execv

import (
	"fmt"
	"testing"

	utils "github.com/hariguchi/go-utils"
)

var (
	ePrintf = utils.Eprintf
)

func TestCmd(t *testing.T) {
	fn := utils.FuncName(false)

	// "ls /"
	cmd := New([]string{"ls", "/"})
	if cmd == nil {
		t.Fatalf("%s: failed to instantiate a command\n", fn)
	}
	fmt.Printf("command: %s\n", cmd.cmdStr)
	err := cmd.Run()
	if err != nil {
		t.Fatalf("%s: cmd.Run: %s\n", fn, err)
	}
	fmt.Printf("stdout:\n%s\n", cmd.Stdout())
	fmt.Printf("stderr:\n%s\n", cmd.Stderr())
	fmt.Printf("--------\n")

	// "mkdir /foo": this must fail without sudo
	cmd = New([]string{"mkdir", "/foo"})
	if cmd == nil {
		t.Fatalf("%s: failed to instantiate a command\n", fn)
	}
	fmt.Printf("command: %s\n", cmd.cmdStr)
	err = cmd.Run()
	if err != nil {
		ePrintf("ERROR (expected): %s: cmd.Run: %s\n", fn, err)
	}
	fmt.Printf("stdout:\n%s\n", cmd.Stdout())
	fmt.Printf("stderr:\n%s\n", cmd.Stderr())
}
