package main

import (
	"fmt"
	"net/http"
	"os"
	"bufio"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <URL> <wordlist>")
		return
	}

	url := os.Args[1]
	wordlist := os.Args[2]

	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dir := scanner.Text()
		fullURL := fmt.Sprintf("%s/%s", url, dir)
		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Println("Error requesting:", fullURL)
			continue
		}
		if resp.StatusCode == 200 {
			fmt.Println("Found:", fullURL)
		}
		resp.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist:", err)
	}
}

