package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// wget скачивает указанную по ссылке страницу в файл html
func wget(fullURL string) {

	pageURL, err := url.Parse(fullURL)
	if err != nil {
		log.Fatal(err)
	}
	path := pageURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1] + ".html"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fullURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Загружена страница %s размером %d\n", fileName, size)
}

func main() {
	scannner := bufio.NewScanner(os.Stdin)

	for scannner.Scan() {
		pageURL := scannner.Text()
		switch pageURL {
		case "":
			continue
		default:
			wget(pageURL)
		}
	}
}
