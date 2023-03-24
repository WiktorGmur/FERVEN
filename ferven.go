// Started on 23.03.2023
// FERVEN - Username reconnaissance tool

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func showHelp() {
	// Displays possible arguments
	whatIs :=
		`FERVEN is an OSINT tool which looks for the provided username on various services.
It queries the username and returns the output if the server returned a valid response.`
	fmt.Println(whatIs)
	fmt.Println(" ")
}

func checkUserExistsYoutube(username string) bool {
	url := fmt.Sprintf("https://www.youtube.com/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func checkUserExistsReddit(username string) bool {
	url := fmt.Sprintf("https://www.reddit.com/user/%s/about.json", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func main() {
	banner()
	var username string

	if len(os.Args) > 1 {
		username = os.Args[1]

	}
	if username == "" {
		l.Error("[ERROR] You must provide a username.")
		os.Exit(1)
	}

	if username == "-h" {
		showHelp()
		os.Exit(0)
	} else {
		// this will execute if "-h" is not present
		l.Log("STARTING WITH THE PROVIDED USERNAME:", username)
		l.Print("FOUND: ")
		if checkUserExistsYoutube(username) {
			l.Printf("Youtube: https://www.youtube.com/%s", username)
		}
		if checkUserExistsReddit(username) {
			l.Printf("Reddit: https://www.reddit.com/user/%s/about.json", username)
		}
	}
}
