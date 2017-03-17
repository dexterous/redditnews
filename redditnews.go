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
	Author    string ` json:"author"    `
	Score     int    ` json:"score"     `
	URL       string ` json:"url"       `
	Permalink string ` json:"permalink" `
	Title     string ` json:"title"     `
}

type response struct {
	Data struct {
		Children []struct {
			Link Item         ` json:"data"     `
		}                   ` json:"children" `
	}                     ` json:"data"     `
}

// Get fetches the most recent Items posted to the specified subreddit.
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Set("User-Agent", "golang-reddit-gopher-0.1")

  client := &http.Client{}
	resp, err := client.Do(req)
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

	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Link
	}
	return items, nil
}

func (i Item) String() string {
  return fmt.Sprintf( "[%3d] %s (%s)\nPermalink: %s\nAuthor: %s\n", i.Score, i.Title, i.URL, i.Permalink, i.Author)
}
