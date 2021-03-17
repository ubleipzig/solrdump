# README

Export documents from a SOLR index as JSON, fast and simply from the command
line.

* https://cwiki.apache.org/confluence/display/solr/Pagination+of+Results

Requesting large number of documents from SOLR can lead to *Deep Paging*
problems:

> When you wish to fetch a very large number of sorted results from Solr to
> feed into an external system, using very large values for the start or rows
> parameters can be very inefficient.

See also: *Fetching A Large Number of Sorted Results: Cursors*

> As an alternative to increasing the "start" parameter to request subsequent
> pages of sorted results, Solr supports using a "Cursor" to scan through
> results. Cursors in Solr are a logical concept, that doesn't involve caching
> any state information on the server. Instead the sort values of the last
> document returned to the client are used to compute a "mark" representing a
> logical point in the ordered space of sort values.

Requirements
------------

SOLR 4.7 or higher, since the cursor mechanism was introduced with SOLR 4.7
([2014-02-25](https://archive.apache.org/dist/lucene/solr/4.7.0/)) &mdash; see
also [efficient deep paging with
cursors](https://solr.pl/en/2014/03/10/solr-4-7-efficient-deep-paging/).

[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active) ![https://goreportcard.com/report/github.com/ubleipzig/solrdump](https://goreportcard.com/badge/github.com/ubleipzig/solrdump)

This project has been developed for [Project finc](https://finc.info) at [Leipzig University Library](https://ub.uni-leipzig.de).

## Installation

Via debian or rpm [package](https://github.com/ubleipzig/solrdump/releases).

Or via go tool:

```shell
$ go get github.com/ubleipzig/solrdump/...
```

## Usage

```shell
$ solrdump -h
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

Export documents matching a query and postprocess with [jq](https://stedolan.github.io/jq/):

```shell
$ solrdump -server https://localhost:8983/solr/biblio -q 'title:"topic model"' -fl id,title | \
  jq -r .title | \
  head -10

A generic approach to topic models and its application to virtual communities /
Topic models for image retrieval on large scale databases
On the use of language models and topic models in the web new algorithms for filtering, ...
Integration von Topic Models und Netzwerkanalyse bei der Bestimmung des Kundenwertes
Time dynamic topic models /
...
```

## Instant search as one-liner

Using solrdump + [jq](https://stedolan.github.io/jq/) + [fzf](https://github.com/junegunn/fzf) (or [peco](https://github.com/peco/peco)).

```shell
$ solrdump -server http://solr.io/solr/biblio -q 'title:"leipzig"' -fl 'id,source_id,title' |\
    jq -rc '[.source_id, .title[:80]] | @tsv' | fzf -e
```

![](images/8e4zf1ryf2gusi3usv329btt8.gif)

[...](https://asciinema.org/a/N8L01waFUixUfO6AIlOfp6RTC?autoplay=1)
