package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Posts is..
type Posts []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getJSON(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("cannot fetch URL %q: %v", url, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http GET status: %s", resp.Status)
	}
	// We could check the resulting content type
	// here if desired.
	error := json.NewDecoder(resp.Body).Decode(result)
	if error != nil {
		return fmt.Errorf("cannot decode JSON: %v", err)
	}
	return nil
}

// JSONParse is..
func JSONParse() {
	fmt.Println("Welcome to JSON parsing..")
	posts := Posts{}
	getJSON("https://jsonplaceholder.typicode.com/posts", &posts)
	for _, post := range posts {
		fmt.Printf("\nuserId : %d\n", post.UserID)
		fmt.Printf("id : %d\n", post.ID)
		fmt.Printf("title : %s\n", post.Title)
		fmt.Printf("body : %s\n", post.Body)
	}

}
