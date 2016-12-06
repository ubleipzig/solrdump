README
======

Export fields from a SOLR index to tabular form, fast.

```
$ solrtab -server https://localhost:8983/solr/biblio -q '*:*' -fl id,title
{"id":"0000001864","title":"Veröffentlichungen des Museums für Völkerkunde zu Leipzig"}
{"id":"0000002001","title":"Festschrift zur Feier des 500jährigen Bestehens der ... /"}
...

$ solrtab -server https://localhost:8983/solr/biblio -q '"title:"topic model"' -fl id,title | \
  jq -r .title | head -10
A generic approach to topic models and its application to virtual communities /
Topic models for image retrieval on large scale databases
On the use of language models and topic models in the web new algorithms for filtering, ...
Integration von Topic Models und Netzwerkanalyse bei der Bestimmung des Kundenwertes
Time dynamic topic models /
...
```
