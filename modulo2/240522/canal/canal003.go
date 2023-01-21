package canal

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.spotify.com",
		"https://www.apple.com",
		"https://www.google.com",
		"https://www.microsoft.com",
		"https://globoesporte.globo.com",
	}

	ch := make(chan string)

	for _, url := range urls {
		go checkStatus(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

}

func checkStatus(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("ocorreu um erro enquanto acessava a url: %s", url)
		ch <- fmt.Sprintf("error: %s", err)
		return
	}

	defer resp.Body.Close()

	u := fmt.Sprintf("url: %s - status: %s", url, resp.Status)
	ch <- u
}
