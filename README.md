# abbreviate

Shorten your strings using common abbreviations.

[![build status](https://travis-ci.org/dnnrly/abbreviate.svg?branch=master)](https://travis-ci.org/dnnrly/abbreviate)
[![codecov](https://codecov.io/gh/dnnrly/abbreviate/branch/master/graph/badge.svg)](https://codecov.io/gh/dnnrly/abbreviate)
[![godoc](https://godoc.org/github.com/dnnrly/abbreviate?status.svg)](http://godoc.org/github.com/dnnrly/abbreviate)
[![report card](https://goreportcard.com/badge/github.com/dnnrly/abbreviate)](https://goreportcard.com/report/github.com/dnnrly/abbreviate)

## Motivation

This tool comes out of a frustration of the name of resources (in my specific
case, AWS stack names) being too long. Wouldn't it be nice if we could have a
tool that would be able to suggest shorter alternatives if your original name
is too long.

## Installation

```bash
go get github.com/dnnrly/abbreviate
make build
```

## Usage

```
This tool will attempt to shorten the string provided using common abbreviations
specified by language and 'set'. Word boundaries will be detected using title case
and non-letters.

Hosted on Github - https://github.com/dnnrly/abbreviate

If you spot a bug, feel free to raise an issue or fix it and make a pull
request. We're really interested to see more abbreviations added or corrected.

Usage:
  abbreviate [action] [flags]
  abbreviate [command]

Available Commands:
  camel       Abbreviate a string and convert it to camel case
  help        Help about any command
  original    Abbreviate the string using the original word boundary separators
  pascal      Abbreviate a string and convert it to pascal case
  print       Print abbreviations in this set
  snake       Abbreviate a string and convert it to snake case

Flags:
  -c, --custom string     Custom abbreviation set
  -h, --help              help for abbreviate
  -l, --language string   Language to select (default "en-us")
      --list              List all abbreviate sets by language
  -m, --max int           Maximum length of string, keep on abbreviating while the string is longer than this limit
  -n, --newline           Add newline to the end of the string (default true)
  -s, --set string        Abbreviation set (default "common")

Use "abbreviate [command] --help" for more information about a command.
```

Examples:
```
$ abbreviate original strategy-limited
stg-ltd

$ abbreviate original strategy-limited --max 11
strategy-ltd

$ abbreviate camel --max 99 strategy-limited
strategyLimited
```

## Code of Conduct
This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Contributing
Pull requests are welcome. See the [contributing guide](CONTRIBUTING.md) for more details.

Please make sure to update tests as appropriate.

## License
[Apache 2](https://choosealicense.com/licenses/apache-2.0/)
