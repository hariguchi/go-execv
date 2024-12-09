package execv

/*
Copyright 2024 Yoichi Hariguchi

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
“Software”), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"bytes"
	"fmt"
	"os/exec"

	utils "github.com/hariguchi/go-utils"
)

type Cmd struct {
	cmd    *exec.Cmd
	cmdStr string
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func New(args []string) *Cmd {
	if len(args) <= 0 {
		return nil
	}
	cmd := new(Cmd)
	cmd.cmd = exec.Command(args[0], args[1:]...)
	cmd.cmd.Stdout = &cmd.stdout
	cmd.cmd.Stderr = &cmd.stderr
	for _, s := range args {
		cmd.cmdStr += s + " "
	}
	cmd.cmdStr = cmd.cmdStr[:len(cmd.cmdStr)-1]

	return cmd
}

func (cmd *Cmd) Run() error {
	fn := utils.FuncName(false)

	err := cmd.cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: os.exec.Run(): %w", fn, err)
	}
	return nil
}

func (cmd *Cmd) Stdout() string {
	return cmd.stdout.String()
}

func (cmd *Cmd) Stderr() string {
	return cmd.stderr.String()
}
