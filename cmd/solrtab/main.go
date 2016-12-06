// https://cwiki.apache.org/confluence/display/solr/Pagination+of+Results
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Response is a SOLR response.
type Response struct {
	Header struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
		Params struct {
			Query      string `json:"q"`
			CursorMark string `json:"cursorMark"`
			Sort       string `json:"sort"`
			Rows       string `json:"rows"`
		} `json:"params"`
	} `json:"header"`
	Response struct {
		NumFound int               `json:"numFound"`
		Start    int               `json:"start"`
		Docs     []json.RawMessage `json:"docs"`
	} `json:"response"`
	NextCursorMark string `json:"nextCursorMark"`
}

// Prepends http, if missing.
func PrependSchema(s string) string {
	if !strings.HasPrefix(s, "http") {
		return fmt.Sprintf("http://%s", s)
	}
	return s
}

func main() {
	server := flag.String("server", "http://localhost:8983/solr/example", "SOLR server, host post and collection")
	fields := flag.String("fl", "", "field or fields to export, separate multiple values by comma")
	query := flag.String("q", "*:*", "SOLR query")
	rows := flag.Int("rows", 1000, "number of rows returned per request")
	sort := flag.String("sort", "id asc", "sort order (only unique fields allowed)")
	verbose := flag.Bool("verbose", false, "show progress")

	flag.Parse()

	*server = PrependSchema(*server)

	v := url.Values{}

	v.Set("q", *query)
	v.Set("sort", *sort)
	v.Set("rows", fmt.Sprintf("%d", *rows))
	v.Set("fl", *fields)

	v.Set("wt", "json")
	v.Set("cursorMark", "*")

	var total int

	for {
		link := fmt.Sprintf("%s/select?%s", *server, v.Encode())

		if *verbose {
			log.Println(link)
		}

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}

		dec := json.NewDecoder(resp.Body)
		var response Response
		if err := dec.Decode(&response); err != nil {
			log.Fatal(err)
		}

		// we do not defer, since we hard-exit on errors anyway
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}

		for _, doc := range response.Response.Docs {
			fmt.Println(string(doc))
		}

		total += len(response.Response.Docs)

		if *verbose {
			log.Printf("fetched %d docs", total)
		}

		if response.NextCursorMark == v.Get("cursorMark") {
			break
		}
		v.Set("cursorMark", response.NextCursorMark)
	}
	log.Printf("fetched %d docs", total)
}
