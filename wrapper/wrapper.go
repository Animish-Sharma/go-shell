package wrapper

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	m "go-shell/function"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = Cmd(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Cmd(input string) error {
	input = strings.TrimSuffix(input, "\r\n")
	input = strings.TrimSuffix(input, "\r\n")
	args := strings.Split(input, " ")
	// git, browser, app-launcher, background tasks, ssh
	switch args[0] {
	case "cd":
		return m.Cd(args)
	case "mkdir":
		m.MkDir(args)
	case "rmdir":
		m.Rmdir(args)
	case "touch":
		m.Touch(args)
	case "remove":
		m.Remove(args)
	case "pwd":
		m.Pwd()
	case "edit":
		m.TextEditor(args)
	case "ls":
		m.Ls()
	case "cp":
		m.Copy(args)
	case "mv":
		m.MoveFile(args)
	case "cat":
		m.Cat(args)
	case "whoami":
		m.WhoAmI()
	case "clear":
		m.Clear()
	case "exit":
		os.Exit(0)
	default:
		fmt.Print("")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
