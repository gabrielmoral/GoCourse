package main

import (
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	var wg sync.WaitGroup

	wg.Add(len(sites))

	for _, site := range sites {
		go func(site string) {
			res, err := http.Get(site)
			if err != nil {
			}

			io.WriteString(os.Stdout, res.Status+"\n")
			wg.Done()
		}(site)
	}

	wg.Wait()
}
