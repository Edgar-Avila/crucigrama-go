package wikipedia

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

type OpenSearchOption struct {
	Text string
	Link  string
}

type QueryExtractResponse struct {
	Query struct {
		Pages map[string]struct {
			Extract string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}

func OpenSearch(topic string) ([]OpenSearchOption, error) {
	// Get request and error handling
	topicQuery := url.QueryEscape(topic)
	path := "https://en.wikipedia.org/w/api.php?action=opensearch&format=json&search=" + topicQuery
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch data")
	}

	// Read response body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse response body
	var parsed []interface{}
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		log.Fatal(err)
	}

	// Extract search options and links
	result := make([]OpenSearchOption, 0)
	size := len(parsed[1].([]interface{}))
	texts := parsed[1].([]interface{})
	links := parsed[3].([]interface{})
	for i := 0; i < size; i++ {
		option := OpenSearchOption{
			Text: texts[i].(string),
			Link:  links[i].(string),
		}
		result = append(result, option)
	}

	return result, nil
}

func QueryExtract(title string) (string, error) {
	// Get request and error handling
	titleQuery := url.QueryEscape(title)
	path := "https://en.wikipedia.org/w/api.php?action=query&redirects=5&prop=extracts&exintro&explaintext&format=json&titles=" + titleQuery
	res, err := http.Get(path)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch data, returned status " + res.Status)
	}

	// Read response body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse response body
	var parsed QueryExtractResponse
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		log.Fatal(err)
	}

	// Get extract
	for _, page := range parsed.Query.Pages {
		return page.Extract, nil
	}

	return "", nil
}
