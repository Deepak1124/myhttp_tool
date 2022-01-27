package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Limiting default number of goroutines to prevent exhausting local resources.
const defaultLimit = 10

// validateURL validates the string is a valid url or not
func validateURL(val string) bool {
	_, err := url.ParseRequestURI(val)
	if err != nil {
		fmt.Println("Invalid URL: ", val, "with error:-", err)
		return false
	}
	return true
}

// addHttpsToDomains adds https to domains to create a valid url
func addHttpToDomains(domains []string) []string {
	var urlBuilder strings.Builder

	var urls []string

	for _, domain := range domains {
		if strings.HasPrefix(domain, "http://") || strings.HasPrefix(domain, "https://") {
			if validateURL(domain) {
				urls = append(urls, domain)
			}
		} else {
			urlBuilder.WriteString("http://")
			urlBuilder.WriteString(domain)

			if validateURL(urlBuilder.String()) {
				urls = append(urls, urlBuilder.String())
			}
			urlBuilder.Reset()
		}
	}
	return urls
}

// isCommandValid checks if given command is valid or not
func isCommandValid(parallel int, urls []string) bool {
	if parallel <= 0 || len(urls) == 0 {
		return false
	}
	return true
}

// getHashOfUrls process Url's response and return corresponding Hash
func getHashOfUrls(url string) (err error) {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	text := string(body)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	fmt.Println(url + " " + hash)
	return nil
}

// runCommand runs the given command
func runCommand(job <-chan string, result chan bool) {
	for url := range job {
		err := getHashOfUrls(url)
		if err != nil {
			log.Fatalln("Error: ", err)
		}

		//Sending message through channel
		result <- true
	}
}

func main() {
	parallel := flag.Int("parallel", defaultLimit, "To limit the number of parallel requests")

	flag.Parse()
	args := flag.Args()

	if !isCommandValid(*parallel, args) {
		log.Fatalln("Invalid command! Please check sample commands in README file")
	}

	urls := addHttpToDomains(args)

	jobs := make(chan string, len(urls))
	result := make(chan bool, len(urls))

	if *parallel > len(urls) {
		*parallel = len(urls)
	}

	//Running Parallel commands using multiple goroutines
	for p := 0; p < *parallel; p++ {
		go runCommand(jobs, result)
	}

	//Sending through channel
	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	//Recieving through channel
	for i := 0; i < len(urls); i++ {
		<-result
	}
}
