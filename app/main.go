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
	builtins := []string{"echo", "exit", "type"}

	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)
		tokens := strings.Split(command, " ")

		if strings.HasPrefix(command, "type ") {
			if slices.Contains(builtins, tokens[1]) {
				fmt.Println(tokens[1] + " is a shell builtin")
			} else if path, err := exec.LookPath(tokens[1]); err == nil {
				fmt.Printf("%s is %s\n", tokens[1], path)
			} else {
				fmt.Println(tokens[1] + ": not found")
			}
			continue
		}

		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		}

		if _, err := exec.LookPath(tokens[0]); err == nil {
			args := tokens[1:]
			program_name := tokens[0]
			fmt.Printf("Program was passed %d args (including program name).\n", len(tokens))
			for i := 0; i >= len(args); i++ {
				fmt.Printf("Arg #%d (program name): %s\n", i, program_name)
			}
			cmd := exec.Command(program_name, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			continue
		}

		fmt.Println(command + ": command not found")
	}
}
