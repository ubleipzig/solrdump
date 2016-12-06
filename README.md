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

Via debian or rpm [package](https://github.com/miku/solrtab/releases).

Or via go tool:

```
$ go get github.com/miku/solrtab/...
```

Usage
-----

```shell
$ solrtab -server https://localhost:8983/solr/biblio -q '*:*' -fl id,title
{"id":"0000001864","title":"Veröffentlichungen des Museums für Völkerkunde zu Leipzig"}
{"id":"0000002001","title":"Festschrift zur Feier des 500jährigen Bestehens der ... /"}
...
```

Postprocess with JSON tools.

```
$ solrtab -server https://localhost:8983/solr/biblio -q '"title:"topic model"' -fl id,title | \
  jq -r .title | \
  head -10

A generic approach to topic models and its application to virtual communities /
Topic models for image retrieval on large scale databases
On the use of language models and topic models in the web new algorithms for filtering, ...
Integration von Topic Models und Netzwerkanalyse bei der Bestimmung des Kundenwertes
Time dynamic topic models /
...
```
