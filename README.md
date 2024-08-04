# gedi

A simple streaming editor like sed and the like, but written in Go and [expr](https://expr-lang.org/).


# usage

Each record (depending on the file type, it can be a line(`string`) / row(`[]string`) / json element(`map[string]any`)) is read into var `x` and can be referenced in the expression. By default it operates in `filter` mode, which simply prints the record if the expression evaluates to true.

```
gedi -f examples/lines.txt 'atoi(x) % 2 == 0'
```


# TODOs

[x] supports line by line files

[] supports csv

[] supports jsonl

[] supports jsonarray
