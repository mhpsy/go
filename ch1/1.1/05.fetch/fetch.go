package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {

		if !(len(url) >= 7 && (strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://"))) {
			url = "https://" + url
		}

		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
			panic(err)
		}

		defer res.Body.Close()

		_, err = io.Copy(os.Stdout, res.Body)

		fmt.Println("fetching URL:", url)
		fmt.Println("fetching status:", res.Status)

		// b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("fetch reading %s: %v\n", url, err)
			return
		}

		// fmt.Printf("%s", b)
	}
}
