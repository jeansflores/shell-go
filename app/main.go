package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	builtins := []string{"echo", "exit", "type", "pwd"}

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)
		tokens := strings.Split(command, " ")

		if command == "exit" {
			break
		}

		switch tokens[0] {
		case "type":
			if slices.Contains(builtins, tokens[1]) {
				fmt.Println(tokens[1] + " is a shell builtin")
			} else if path, err := exec.LookPath(tokens[1]); err == nil {
				fmt.Printf("%s is %s\n", tokens[1], path)
			} else {
				fmt.Println(tokens[1] + ": not found")
			}
		case "echo":
			fmt.Println(command[5:])
		case "pwd":
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(cwd)
		default:
			if _, err := exec.LookPath(tokens[0]); err == nil {
				cmd := exec.Command(tokens[0], tokens[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			} else {
				fmt.Println(command + ": command not found")
			}
		}
	}
}
