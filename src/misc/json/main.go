package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// IssuesURL ...
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult holds search results
type IssuesSearchResult struct {
	// TotalCount ...
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue ...
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User ...
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues searches https://api.github.com/search/issues
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, ""))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed %s", resp.Status)
	}

	var result IssuesSearchResult

	// Use streaming decoder, json.Decoder, which allows several JSON
	// entities to be decoded in sequence from the same stream.
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
