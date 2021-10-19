package solrdump

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"strings"

	"github.com/sethgrid/pester"
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
		Docs     []json.RawMessage `json:"docs"` // dependent on SOLR schema
	} `json:"response"`
	NextCursorMark string `json:"nextCursorMark"`
}

// PrependSchema http, if missing.
func PrependSchema(s string) string {
	if !strings.HasPrefix(s, "http") {
		return fmt.Sprintf("http://%s", s)
	}
	return s
}

// Dumper can run a data extraction from solr.
type Dumper struct {
	Writer                      io.Writer
	Server                      string
	Fields                      string
	Sort                        string
	Query                       string
	NumRows                     int
	Wt                          string
	SkipCertificateVerification bool
	Verbose                     bool
}

func (d *Dumper) Run() error {
	v := url.Values{}
	v.Set("q", d.Query)
	v.Set("sort", d.Sort)
	v.Set("rows", fmt.Sprintf("%d", d.NumRows))
	v.Set("fl", d.Fields)
	v.Set("wt", "json")
	v.Set("cursorMark", "*")
	var total int
	for {
		link := fmt.Sprintf("%s/select?%s", d.Server, v.Encode())
		if d.Verbose {
			log.Println(link)
		}
		resp, err := pester.Get(link)
		if err != nil {
			return fmt.Errorf("http: %s", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("failed to fetch response body for debugging")
			}
			log.Printf("response body (%d): %s", len(b), string(b))
			return fmt.Errorf("status: %v", resp.Status)
		}
		var response Response
		switch d.Wt {
		case "json":
			// invalid character '\r' in string literal
			dec := json.NewDecoder(resp.Body)
			if err := dec.Decode(&response); err != nil {
				return fmt.Errorf("decode: %s", err)
			}
		default:
			return fmt.Errorf("wt=%s not implemented", d.Wt)
		}
		// We do not defer, since we hard-exit on errors anyway.
		if err := resp.Body.Close(); err != nil {
			return err
		}
		for _, doc := range response.Response.Docs {
			fmt.Println(string(doc))
		}
		total += len(response.Response.Docs)
		if d.Verbose {
			log.Printf("fetched %d docs", total)
		}
		if response.NextCursorMark == v.Get("cursorMark") {
			break
		}
		v.Set("cursorMark", response.NextCursorMark)
	}
	if d.Verbose {
		log.Printf("fetched %d docs", total)
	}
	return nil
}
