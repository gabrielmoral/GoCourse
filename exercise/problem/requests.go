package main

import (
	"context"
	"io"
	"net/http"
	"os"
)

func main() {
	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
		"https://hangouts.google.com1",
	}
	ctx, cancel := context.WithCancel(context.Background())

	channel := make(chan error)

	for _, url := range sites {
		go request(url, ctx, channel)
	}

	for err := range channel {
		if err != nil {
			break
		}
	}

	cancel()
}

func request(url string, ctx context.Context, errorChannel chan<- error) {

	select {
	case <-ctx.Done():
		return
	default:
		res, err := http.Get(url)
		if err != nil {
			errorChannel <- err
		} else {
			errorChannel <- nil
			io.WriteString(os.Stdout, res.Status+"\n")
		}
	}
}
