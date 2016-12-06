README
======

Export fields from a SOLR index to tabular form, fast.

```
$ solrtab -server https://localhost:8983/solr/biblio -q '*:*' -fl id,title
{"id":"0000001864","title":"Veröffentlichungen des Museums für Völkerkunde zu Leipzig"}
{"id":"0000002001","title":"Festschrift zur Feier des 500jährigen Bestehens der Universität Leipzig : [1409 - 1909] /"}
...
```
