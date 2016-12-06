package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

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
		NumFound int             `json:"numFound"`
		Start    int             `json:"start"`
		Docs     json.RawMessage `json:"docs"`
	} `json:"response"`
	NextCursorMark string `json:"nextCursorMark"`
}

func main() {
	server := flag.String("server", "http://localhost:8983/solr/example", "SOLR server, host post and collection")
	// fields := flag.String("f", "", "field or fields to export")
	flag.Parse()

	// https://cwiki.apache.org/confluence/display/solr/Pagination+of+Results
	// cursorMark and start are mutually exclusive parameters

	// when fetching all docs, you might as well use a simple id sort
	// unless you really need the docs to come back in a specific order
	// $params = [ q => $some_query, sort => 'id asc', rows => $r, cursorMark => '*' ]
	// $done = false
	// while (not $done) {
	//   $results = fetch_solr($params)
	//   // do something with $results
	//   if ($params[cursorMark] == $results[nextCursorMark]) {
	//     $done = true
	//   }
	//   $params[cursorMark] = $results[nextCursorMark]
	// }

	// {
	//   "responseHeader": {
	//     "status": 0,
	//     "QTime": 0,
	//     "params": {
	//       "q": "*:*",
	//       "cursorMark": "*",
	//       "sort": "id asc",
	//       "rows": "1000",
	//       "wt": "json"
	//     }
	//   },
	//   "response": {
	//     "numFound": 0,
	//     "start": 0,
	//     "docs": []
	//   },
	//   "nextCursorMark": "*"
	// }

	v := url.Values{}
	v.Set("q", "*:*")
	v.Set("sort", "id asc")
	v.Set("rows", "1000")
	v.Set("wt", "json")
	v.Set("cursorMark", "*")

	for {
		link := fmt.Sprintf("%s/select?%s", *server, v.Encode())

		log.Println(link)

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		dec := json.NewDecoder(resp.Body)
		var response Response
		if err := dec.Decode(&response); err != nil {
			log.Fatal(err)
		}

		if response.NextCursorMark == v.Get("cursorMark") {
			break
		}
		v.Set("cursorMark", response.NextCursorMark)
	}
}
