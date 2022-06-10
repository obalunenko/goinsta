package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Davincible/goinsta/v3"
)

func main() {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		log.Fatal(err)
	}

	defer insta.Logout()

	log.Println("Download random photo")
	var client http.Client
	request, err := http.NewRequest("GET", "https://picsum.photos/800/800", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	err = insta.Account.ChangeProfilePic(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
