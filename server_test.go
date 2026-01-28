// server_test.go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "testing"
)

func TestGistsAPI_PrintURLs(t *testing.T) {
    url := "http://localhost:8080/Shehbab-Kakkar?page=1&per_page=50" // you can adjust per_page

    resp, err := http.Get(url)
    if err != nil {
        t.Fatalf("Failed to GET %s: %v", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status 200, got %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    var gists []map[string]interface{}
    if err := json.Unmarshal(body, &gists); err != nil {
        t.Fatalf("Failed to parse JSON: %v", err)
    }

    if len(gists) == 0 {
        t.Errorf("Expected at least 1 gist, got 0")
    }

    fmt.Println("Gist URLs for user Shehbab-Kakkar:")
    for _, gist := range gists {
        if htmlURL, ok := gist["html_url"].(string); ok {
            fmt.Println(htmlURL)
        }
    }
}
