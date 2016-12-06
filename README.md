README
======

Export fields from a SOLR index to tabular form, fast.

* https://cwiki.apache.org/confluence/display/solr/Pagination+of+Results

Especially section: *Fetching A Large Number of Sorted Results: Cursors*

> As an alternative to increasing the "start" parameter to request subsequent
> pages of sorted results, Solr supports using a "Cursor" to scan through
> results.  Cursors in Solr are a logical concept, that doesn't involve
> caching any state information on the server.  Instead the sort values of the
> last document returned to the client are used to compute a "mark"
> representing a logical point in the ordered space of sort values.

Installation
------------

Via debian or rpm [package](https://github.com/ubleipzig/solrdump/releases).

Or via go tool:

```
$ go get github.com/ubleipzig/solrdump/...
```

Usage
-----

```shell
solrdump -h
Usage of solrdump:
  -fl string
        field or fields to export, separate multiple values by comma
  -q string
        SOLR query (default "*:*")
  -rows int
        number of rows returned per request (default 1000)
  -server string
        SOLR server, host post and collection (default "http://localhost:8983/solr/example")
  -sort string
        sort order (only unique fields allowed) (default "id asc")
  -verbose
        show progress
  -version
        show version and exit
```

Export id and title field for all documents:

```shell
$ solrdump -server https://localhost:8983/solr/biblio -q '*:*' -fl id,title
{"id":"0000001864","title":"Veröffentlichungen des Museums für Völkerkunde zu Leipzig"}
{"id":"0000002001","title":"Festschrift zur Feier des 500jährigen Bestehens der ... /"}
...
```

Export documents matching a query and postprocess with jq:

```
$ solrdump -server https://localhost:8983/solr/biblio -q '"title:"topic model"' -fl id,title | \
  jq -r .title | \
  head -10

A generic approach to topic models and its application to virtual communities /
Topic models for image retrieval on large scale databases
On the use of language models and topic models in the web new algorithms for filtering, ...
Integration von Topic Models und Netzwerkanalyse bei der Bestimmung des Kundenwertes
Time dynamic topic models /
...
```
