package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func build(entrypoint string, output string) {
	project_path, _ := os.Getwd()
	in_abspath := filepath.Join(project_path, entrypoint)
	out_abspath := filepath.Join(project_path, output)
	out_dir := filepath.Dir(out_abspath)

	fmt.Printf("input: %v\n", in_abspath)
	fmt.Printf("output: %v\n", out_abspath)
	_, err := os.Stat(out_dir)

	if err != nil {
		fmt.Println("- creating build directory")
	} else {
		fmt.Println("- cleaning up directory")
		err = os.Remove(output)
	}

	if err != nil {
		fmt.Println("! error while making/cleaning build directory")
		return
	}

	cmd := exec.Command("go", "build", "-o", out_abspath, in_abspath)

	fmt.Printf("- running command: %v\n", cmd.Args)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("! Build process failed!")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
	fmt.Println("Build success!")
}

func main() {
	entry := "cmd/qms-cli/main.go"
	output := "build/qms-cli"
	build(entry, output)
}
