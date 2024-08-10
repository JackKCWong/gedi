# gedi

A simple streaming editor like sed and the like, but written in Go and [expr](https://expr-lang.org/).


# usage

Each record (depending on the file type, it can be a line(`string`) / row(`[]string`) / json element(`map[string]any`)) is read into var `x` and can be referenced in the expression. By default it operates in `filter` mode, which simply prints the record if the expression evaluates to true.

```
gedi -f examples/lines.txt 'atoi(x) % 2 == 0'
```

Additional vars / functions:

* `ix`: the current record number.

* `x`: the current record. When `filetype` is `line`, it's the current line as `string`, when `filetype` is `csv`, it's the current row as `[]string`.

* `now`: `time.Now()` at the start of the program, provided to avoid calling `time.Now()` repeatedly in large files.

* `localtime(string)`: guess time from a given string as if it's a local time. e.g. assuming local is HTK, `"2023-01-01 00:00:00 WARN foobar" | localtime()` gives `2023-01-01 00:00:00 UTC+08:00`

* `utctime(string)`: guess time from a given string as if it's a UTC time. e.g. `"2023-01-01 00:00:00 WARN foobar" | utctime()` gives `2023-01-01 00:00:00 UTC+0000`

* `tztime(string, string)`: guess time from a given string as if it's a given timezone time. e.g. `"2023-01-01 00:00:00 WARN foobar" | tztime("UTC+8")` gives `2023-01-01 00:00:00 UTC+08:00`

* `within(time.Time, string)`: checks if a given time is within a given duration comapred to `now`. e.g. find log lines that are within the last 24 hours: `x[0:20] | localtime() | within("-24h")`


# TODOs

[x] supports line by line files

[x] supports csv

[] supports jsonl

[] supports jsonarray
