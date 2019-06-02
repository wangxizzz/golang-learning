package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json" // 变量名小写，表示包内可见

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns. 哪怕函
	//数意外崩溃终止，也能保证关键字 defer 安排调用的函数会被执行
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
