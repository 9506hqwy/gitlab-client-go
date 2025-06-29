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

	params := client.GetApiV4ProjectsIdRepositoryTagsParams{}
	resp, err := c.GetApiV4ProjectsIdRepositoryTagsWithResponse(context.TODO(), "2", &params, bearerAuth)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode())
	}

	for _, tag := range *resp.JSON200 {
		fmt.Printf("Name: %s, Commit: %s\n", *tag.Name, *tag.Commit.Id)
	}
}
