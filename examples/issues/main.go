package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func bearerAuth(ctx context.Context, req *http.Request) error {
	token := "<TOKEN>"
	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}

func main() {
	hc := http.Client{}

	c, err := client.NewClientWithResponses("http://127.0.0.1", client.WithHTTPClient(&hc))
	if err != nil {
		log.Fatal(err)
	}

	params := client.GetApiV4IssuesParams{}
	resp, err := c.GetApiV4IssuesWithResponse(context.TODO(), &params, bearerAuth)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode())
	}

	for _, issue := range *resp.JSON2XX {
		fmt.Printf("Id: %d, Title: %s, Author: %s\n", *issue.Id, *issue.Title, *issue.Author.Name)
	}
}
