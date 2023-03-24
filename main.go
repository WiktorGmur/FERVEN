// Started on 23.03.2023
// FERVEN - Username reconnaissance tool

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/happierall/l"
)

func banner() {
	banner, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(banner))
}

func help() {
	// Displays possible arguments
	fmt.Println("Possible switches:")
}

func main() {
	banner()

	arg1 := os.Args[1]
	if arg1 == "-h" {
		help()
	}
	l.Log("STARTING WITH THE PROVIDED FLAG:", arg1)
}
