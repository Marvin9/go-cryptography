package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type option struct {
	title string
	path  string
}

var options = []option{
	option{
		title: "Shift cipher",
		path:  "./shift/shift_cipher.go",
	},
	option{
		title: "Mono alphabetic cipher",
		path:  "./mono-alphabetic/mono_alphabetic_cipher.go",
	},
	option{
		title: "Playfair cipher",
		path:  "./playfair/playfair_cipher.go",
	},
}

func clear() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	case "windows":
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	var slt int
	for ind, opt := range options {
		fmt.Printf("%v. %v\n", ind+1, opt.title)
	}
	fmt.Printf("Select: ")
	fmt.Scan(&slt)

	if slt > 0 && slt <= len(options) {
		clear()
		cmd := exec.Command("go", "run", options[slt-1].path)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		panic("Invalid option choice")
	}
}
