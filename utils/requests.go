package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// TODO: ADD HTTP Error handling
func GetRequest(url string, token string) []byte {
	r, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("accept", "application/json")
	r.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}

	resp, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return b
}

// TODO: ADD HTTP Error handling
func PostRequest(url string, body []byte, token string) []byte {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("accept", "application/json")
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return b

}
