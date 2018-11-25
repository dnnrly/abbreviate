# abbreviate

Shorten your strings using common abbreviations.

[![build status](https://travis-ci.org/dnnrly/abbreviate.svg?branch=master)](https://travis-ci.org/dnnrly/abbreviate)
[![codecov](https://codecov.io/gh/dnnrly/abbreviate/branch/master/graph/badge.svg)](https://codecov.io/gh/dnnrly/abbreviate)
[![godoc](https://godoc.org/github.com/dnnrly/abbreviate?status.svg)](http://godoc.org/github.com/dnnrly/abbreviate)
[![report card](https://goreportcard.com/badge/github.com/dnnrly/abbreviate)](https://goreportcard.com/report/github.com/dnnrly/abbreviate)

## Installation


```bash
go get github.com/gobuffalo/packr/v2/packr2
go get github.com/dnnrly/abbreviate
packr2 install github.com/dnnrly/abbreviate
```

## Usage

```
This tool will attempt to shorten the string provided using common abbreviations
specified by language and 'set'.

Word boundaries will detect camel case and non-letter

Usage:
  abbreviate [string] [flags]

Flags:
  -c, --custom string     Custom abbreviation set
  -h, --help              help for abbreviate
  -l, --language string   Language to select (default "en-us")
      --list              List all abbreviate sets by language
  -m, --max int           Maximum length of string, keep on abbreviating while the string is longer than this limit
  -s, --set string        Abbreviation set (default "common")
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache 2](https://choosealicense.com/licenses/apache-2.0/)
