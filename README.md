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

* `kx`: if `x` comes from an array of a json field, `kx` is the name of the field. e.g. for `{"k": [1,2,3]}`, `kx` would be `k`.

* `now`: `time.Now()` at the start of the program, provided to avoid calling `time.Now()` repeatedly in large files.

## examples

- Filter lines: print only even numbers parsed from a file

```bash
gedi -i testdata/lines.txt 'atoi(x) % 2 == 0'
```

- Map CSV rows: output the second column of every CSV row (map mode)

```bash
gedi -t csv -i testdata/HSI.csv -m map 'x[1]'
```

- Filter JSON Lines: select objects where `status == "OK"`

```bash
gedi -t jsonl -i testdata/jsonl.jsonl 'x.status == "OK"'
```

- Reduce (aggregate): count records (shorthand `-r` sets reduce mode)

```bash
gedi -i testdata/lines.txt -r 'acc + 1'
```

- Skip header lines: skip the first line of an input before processing

```bash
gedi -s 1 -t csv -i testdata/noheader.csv 'atoi(x[0]) > 100'
```

- SSV with max fields: read at most 3 fields per record (useful for space-separated files)

```bash
gedi -t ssv -n 3 -f data.ssv 'x[2] == "foo"'
```

Time helpers

The library provides helpers to parse and compare times. For example, to find log lines whose timestamp (first 20 chars) is within the last 24 hours:

```bash
gedi -f testdata/dates.log 'x[0:20] | localtime() | within("-24h")'
```

Flags and modes

- `-t`, `--type`: file type. One of `line`, `csv`, `ssv`, `jsonl`, `json`, or `auto` (default: `auto`, guesses from filename).
- `-m`, `--mode`: mode of operation. `auto` (infer), `filter` (default), `map`, `reduce`.
- `-r`, `--reduce`: shortcut to set reduce mode.
- `-s`, `--skip`: skip N lines from start of input.
- `-n`, `--max`: for `ssv` reader, maximum number of fields per record.

Common variables and helpers available in expressions

- `ix`: current record index (0-based).
- `x`: current record (type depends on file type).
- `kx`: when iterating over a JSON field that is an array, `kx` is the key name for the items.
- `now`: program start time (useful for relative comparisons).
- Duration constants: `ms`, `sec`, `min`, `hour`, `day`, `week`, `month`, `year`.

Useful functions

- `localtime(string)`, `utctime(string)`, `tztime(string, tz)`: parse a timestamp string assuming a local/UTC/specified timezone.
- `unixtime(int64)`: convert unix timestamps (ms/s/us) to time.Time.
- `within(t, duration)`: check whether `t` is within a duration relative to `now` (e.g. `"-24h"`).
- `after(t1, t2)`, `before(t1, t2)`: compare times.
- `gt`, `lt`, `ge`, `le`: general comparison helpers for numbers and times.
- `grep(input, regex [, group])`: return regex matches from `input`; if `regex` has groups, `group` selects which group to return.

See the examples above for common usage patterns. The tool is designed for fast, streaming transforms and filters on large files.

Useful regex patterns

The expression environment includes a set of pre-defined regular expression constants you can use directly in expressions (they are injected into the `env` map in `expr.go`). The names and typical uses are:

- `reDate`: common date patterns
- `reTime`: common time patterns
- `reLink`: URL / link pattern
- `reEmail`: email address pattern
- `reIPv4`: IPv4 address pattern
- `reIPv6`: IPv6 address pattern
- `reIP`: generic IP address pattern (IPv4 or IPv6)
- `reNotKnownPort`: pattern for ports that are not known/expected
- `reMD5Hex`: 32-character hex MD5 digest
- `reSHA1Hex`: 40-character hex SHA-1 digest
- `reSHA256Hex`: 64-character hex SHA-256 digest
- `reGUID`: GUID/UUID pattern
- `reMACAddress`: MAC address pattern
- `reGitRepo`: Git repository URL or path pattern

You can use these constants inside expressions, for example:

```
x | grep(reEmail)
```
