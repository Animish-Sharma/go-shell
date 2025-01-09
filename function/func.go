package function

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	helper "go-shell/helper"
)

var errPath = errors.New("path required")

func Cd(args []string) error {
	if len(args) < 2 {
		return errPath
	}
	return os.Chdir(args[1])
}

func MkDir(args []string) error {
	if len(args) < 2 {
		return errors.New("Need a Folder Name")
	}
	return os.Mkdir(args[1], os.ModePerm)
}
func Rmdir(args []string) error {
	if len(args) < 2 {
		return errors.New("Need a Folder Name")
	}
	return os.Remove(args[1])
}

func Touch(args []string) error {
	if len(args) < 3 {
		return errors.New("Need a file name")
	}
	if len(args[3]) > 0 {
		newContent := []byte{}
		x := args[1]
		args = args[2:]
		for _, content := range args {
			newContent = append(newContent, []byte(content)...)
			newContent = append(newContent, ' ')

		}
		return os.WriteFile(x, newContent, 0755)
	}
	return os.WriteFile(args[1], nil, 0755)
}

func Remove(args []string) error {
	if len(args) < 2 {
		fmt.Println("Need the both src and dest of files")
		return errors.New("Need a file name")
	}
	return os.Remove(args[1])
}

func Pwd() error {
	str, err := os.Getwd()
	fmt.Println(str)
	return err
}

func Ls() error {
	entries, err := os.ReadDir("./")
	for _, e := range entries {
		fmt.Println(e.Name())
	}
	return err
}

func Copy(args []string) error {
	if len(args) > 2 {
		return CopyFile(args[1], args[2])
	}
	return errors.New("Either the src or the dest is not defined")
}

func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("Non Regular Source File %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = helper.CopyFileContents(src, dst)
	return
}

func TextEditor(args []string) error {
	if len(args) < 2 {
		return errPath
	}
	contents, err := os.ReadFile(args[1])
	if err != nil {
		return errors.New("File Not Specified Properly")
	}
	fmt.Printf("\nContents of the file %s are: \n%s", args[1], contents)
	fmt.Println("\n Enter text below and type 'SAVE' to save the file")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "SAVE" {
			break
		}
		lines = append(lines, line)
	}
	newContent := []byte{}
	for _, content := range lines {
		newContent = append(newContent, []byte(content)...)
		newContent = append(newContent, '\n')

	}
	err = os.WriteFile(args[1], newContent, 0644)
	if err != nil {
		return err
	}
	fmt.Println("File Saved Succesfully")
	return nil
}

func MoveFile(args []string) error {
	if len(args) < 3 {
		return errPath
	}
	ipFile, err := os.Open(args[1])
	if err != nil {
		return err
	}
	defer ipFile.Close()
	opfile, err := os.Create(args[2])
	if err != nil {
		return nil
	}
	defer opfile.Close()
	_, err = io.Copy(opfile, ipFile)
	if err != nil {
		return err
	}
	ipFile.Close()
	err = os.Remove(args[1])
	if err != nil {
		return err
	}
	return nil
}

func Cat(args []string) error {
	if len(args) < 2 {
		return errPath
	}
	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		return errPath
	}
	fmt.Println(string(data))
	return nil

}

func WhoAmI() error {
	fmt.Println(os.Hostname())
	return nil
}

func Clear() error {
	fmt.Print("\033[H\033[2J")
	return nil
}
