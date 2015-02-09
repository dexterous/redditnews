// redditnews package implements a basic client for the Reddit API.
package redditnews

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Item describes a RedditNews item.
type Item struct {
	Author string `json:"author"`
	Score  int    `json:"score"`
	URL    string `json:"url"`
	Title  string `json:"title"`
}

type response struct {
	Data1 struct {
		Children []struct {
			Data2 Item `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

// Get fetches the most recent Items posted to the specified subreddit.
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(r.Data1.Children))
	for i, child := range r.Data1.Children {
		items[i] = child.Data2
	}
	return items, nil
}

func (i Item) String() string {
  return fmt.Sprintf( "[%d] %s (%s)\nAuthor: %s", i.Score, i.Title, i.URL, i.Author)
}
