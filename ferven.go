// Started on 23.03.2023
// FERVEN - Username reconnaissance tool

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/happierall/l"
)

func banner() {
	banner, err := os.ReadFile("banner.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(banner))
}

func showHelp() {
	fmt.Println("FERVEN is an OSINT tool which looks for the provided username on various services.")
	fmt.Println("It queries the username and returns the output if the server returned a valid response.")
}

func checkUserExists(username, service string) bool {
	var url string

	switch service {
	case "youtube":
		url = fmt.Sprintf("https://www.youtube.com/%s", username)
	case "reddit":
		url = fmt.Sprintf("https://www.reddit.com/user/%s/about.json", username)
	case "twitter":
		url = fmt.Sprintf("https://api.twitter.com/1.1/users/show.json?screen_name=%s", username)
	case "instagram":
		url = fmt.Sprintf("https://www.instagram.com/%s/?__a=1", username)
	case "facebook":
		url = fmt.Sprintf("https://www.facebook.com/%s", username)
	case "linkedin":
		url = fmt.Sprintf("https://www.linkedin.com/in/%s", username)
	case "github":
		url = fmt.Sprintf("https://api.github.com/users/%s", username)
	case "gitlab":
		url = fmt.Sprintf("https://gitlab.com/api/v4/users?username=%s", username)
	case "bitbucket":
		url = fmt.Sprintf("https://api.bitbucket.org/2.0/users/%s", username)
	case "soundcloud":
		url = fmt.Sprintf("https://soundcloud.com/%s", username)
	case "pinterest":
		url = fmt.Sprintf("https://pinterest.com/%s/", username)
	case "twitch":
		url = fmt.Sprintf("https://twitch.tv/%s", username)
	case "vimeo":
		url = fmt.Sprintf("https://vimeo.com/%s", username)
	case "medium":
		url = fmt.Sprintf("https://medium.com/@%s", username)
	case "flickr":
		url = fmt.Sprintf("https://flickr.com/people/%s/", username)
	case "dribbble":
		url = fmt.Sprintf("https://dribbble.com/%s", username)
	case "hackerrank":
		url = fmt.Sprintf("https://hackerrank.com/%s", username)
	case "codecademy":
		url = fmt.Sprintf("https://codecademy.com/%s", username)
	case "codepen":
		url = fmt.Sprintf("https://codepen.com/%s", username)

	}
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
	return resp.StatusCode == 200
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
		fmt.Println("FOUND: ")
		services := []string{
			"youtube",
			"reddit",
			"twitter",
			"instagram",
			"facebook",
			"linkedin",
			"github",
			"gitlab",
			"bitbucket",
			"soundcloud",
			"pinterest",
			"twitch",
			"vimeo",
			"medium",
			"flickr",
			"dribble",
			"hackerrank",
			"codecademy",
			"codepen",
		}
		for _, service := range services {
			if checkUserExists(username, service) {
				switch service {
				case "youtube":
					l.Printf("Youtube: https://www.youtube.com/%s", username)
				case "reddit":
					l.Printf("Reddit: https://www.reddit.com/user/%s/about.json", username)
				case "twitter":
					l.Printf("Twitter: https://api.twitter.com/1.1/users/show.json?screen_name=%s", username)
				case "instagram":
					l.Printf("Instagram: https://www.instagram.com/%s/?__a=1", username)
				case "facebook":
					l.Printf("Facebook: https://www.facebook.com/%s", username)
				case "linkedin":
					l.Printf("LinkedIn: https://www.linkedin.com/in/%s", username)
				case "github":
					l.Printf("Github: https://api.github.com/users/%s", username)
				case "gitlab":
					l.Printf("GitLab: https://gitlab.com/api/v4/users?username=%s", username)
				case "bitbucket":
					l.Printf("Bitbucket: https://api.bitbucket.org/2.0/users/%s", username)
				case "soundcloud":
					l.Printf("SoundCloud: https://soundcloud.com/%s", username)
				case "pinterest":
					l.Printf("Pinterest: https://pinterest.com/%s", username)
				case "twitch":
					l.Printf("Twitch: https://www.twitch.tv/%s", username)
				case "vimeo":
					l.Printf("Vimeo: https://vimeo.com/%s", username)
				case "medium":
					l.Printf("Medium: https://medium.com/@%s", username)
				case "flickr":
					l.Printf("Flickr: https://flickr.com/people/%s", username)
				case "behance":
					l.Printf("Behance: https://behance.com/%s", username)
				case "dribble":
					l.Printf("Dribble: https://dribble.com/%s", username)
				case "hackerrank":
					l.Printf("Hackerrank: https://hackerrank.com/%s", username)
				case "codecademy":
					l.Printf("Codecademy: https://codecademy.com/%s", username)
				case "codepen":
					l.Printf("Codepen: https://codepen.io/%s", username)
				}
			}
		}
	}
}
