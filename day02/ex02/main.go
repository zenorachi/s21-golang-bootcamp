package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	console := bufio.NewScanner(os.Stdin)
	args := os.Args[1:]

	for console.Scan() {
		input := console.Text()
		if input == "" {
			continue
		}
		cmd := exec.Command(args[0], append(args[1:], input)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil && !strings.Contains(cmd.String(), fmt.Sprintf("/bin/%s", os.Args[1])) {
			log.Println(cmd, "invalid command")
		}
	}

	if err := console.Err(); err != nil {
		log.Fatal(err)
	}
}
