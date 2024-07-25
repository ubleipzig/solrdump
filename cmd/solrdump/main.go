// https://cwiki.apache.org/confluence/display/solr/Pagination+of+Results
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/ubleipzig/solrdump"
)

// Version of application.
const Version = "0.1.12"

var (
	server                      = flag.String("server", "http://localhost:8983/solr/example", "SOLR server, host post and collection")
	fields                      = flag.String("fl", "", "field or fields to export, separate multiple values by comma")
	query                       = flag.String("q", "*:*", "SOLR query")
	rows                        = flag.Int("rows", 1000, "number of rows returned per request")
	sort                        = flag.String("sort", "id asc", "sort order (only unique fields allowed)")
	wt                          = flag.String("wt", "json", "output format")
	verbose                     = flag.Bool("verbose", false, "show progress")
	showVersion                 = flag.Bool("version", false, "show version and exit")
	skipCertificateVerification = flag.Bool("k", false, "skip certificate verfication")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}
	if *skipCertificateVerification {
		// TODO: Pester may not pick this up.
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	dumper := solrdump.Dumper{
		Writer:                      os.Stdout,
		Server:                      solrdump.PrependSchema(*server),
		Fields:                      *fields,
		Sort:                        *sort,
		Query:                       *query,
		NumRows:                     *rows,
		Wt:                          *wt,
		SkipCertificateVerification: *skipCertificateVerification,
		Verbose:                     *verbose,
	}
	if err := dumper.Run(); err != nil {
		log.Fatal(err)
	}
}
