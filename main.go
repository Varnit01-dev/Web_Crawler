package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {

	fetcher := urlFetcher{}

	url := "https://golangbot.com/learn-golang-series/page1"

	urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		os.Exit(1)
	}

	fmt.Println("Extracted URLs:")
	for _, u := range urls {
		fmt.Println(u)
	}
}

type Fetcher interface {
	Fetch(url string) ([]string, error)
}

type urlFetcher struct{}

func (ef urlFetcher) Fetch(url string) ([]string, error) {
	// Simulate fetching URLs
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch URL: %s", response.Status)
	}

	urls := make([]string, 0)
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(bodyStr, "\n") {
		if strings.Contains(line, "<a href=") {
			start := strings.Index(line, "\"") + 1
			end := strings.Index(line[start:], "\"") + start
			urls = append(urls, line[start:end])
		}
	}

	return []string{"https://golangbot.com/learn-golang-series/page1", "https://golangbot.com/learn-golang-series/page2"}, nil
}

func Concurrent(url string, fetcher Fetcher, fetched map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	if fetched[url] {
		return
	}
	fetched[url] = true
	urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	for _, u := range urls {
		wg.Add(1)
		go Concurrent(u, fetcher, fetched, wg)
	}
}
