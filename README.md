# abbreviate

Shorten your strings using common abbreviations.

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/W7W414S4U)

[![codecov](https://codecov.io/gh/dnnrly/abbreviate/branch/master/graph/badge.svg)](https://codecov.io/gh/dnnrly/abbreviate)
[![godoc](https://godoc.org/github.com/dnnrly/abbreviate?status.svg)](http://godoc.org/github.com/dnnrly/abbreviate)
[![report card](https://goreportcard.com/badge/github.com/dnnrly/abbreviate)](https://goreportcard.com/report/github.com/dnnrly/abbreviate)

![GitHub watchers](https://img.shields.io/github/watchers/dnnrly/abbreviate?style=social)
![GitHub stars](https://img.shields.io/github/stars/dnnrly/abbreviate?style=social)
[![Twitter URL](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fdnnrly%2Fabbreviate)](https://twitter.com/intent/tweet?url=https://github.com/dnnrly/abbreviate)


<button class="button-save large"><a href="https://tidelift.com/subscription/pkg/go-github-com-dnnrly-abbreviate?utm_source=go-github-com-dnnrly-abbreviate&utm_medium=referral&utm_campaign=enterprise">Supported by Tidelift</a></button>

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
  kebab       Abbreviate a string and convert it to kebab case
  separated   Abbreviate a string and convert it with separator passed between words and abbreviations

Flags:
  -c, --custom string     Custom abbreviation set
      --from-front        Shorten from the front
  -h, --help              help for abbreviate
  -l, --language string   Language to select (default "en-us")
      --list              List all abbreviate sets by language
  -m, --max int           Maximum length of string, keep on abbreviating while the string is longer than this limit
  -n, --newline           Add newline to the end of the string (default true)
  -s, --set string        Abbreviation set (default "common")
  -r, --strategy          Set a strategy to use if no match is found for input string(allowed value "removeVowel")

Use "abbreviate [command] --help" for more information about a command.
```

Examples:
```
$ abbreviate original strategy-limited
stg-ltd

$ abbreviate original --max 11 strategy-limited
strategy-ltd

$ abbreviate original --max 11 --from-front strategy-limited
stg-limited

$ abbreviate camel --max 99 strategy-limited
strategyLimited

$ abbreviate kebab StrategyLimited
stg-ltd

$ abbreviate separated StrategyLimited --separator +
stg+ltd

$ abbreviate separated StrategyLimited
stgltd

$ abbreviate original --strategy removeVowel WhatToDo
WhtTD

```

## Code of Conduct
This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Contributing
Pull requests are welcome. See the [contributing guide](CONTRIBUTING.md) for more details.

Please make sure to update tests as appropriate.

## github.com/dnnrly/abbreviate for enterprise

Available as part of the Tidelift Subscription

The maintainers of github.com/dnnrly/abbreviate and thousands of other packages are working with Tidelift to deliver commercial support and maintenance for the open source dependencies you use to build your applications. Save time, reduce risk, and improve code health, while paying the maintainers of the exact dependencies you use. [Learn more.](https://tidelift.com/subscription/pkg/go-github-com-dnnrly-abbreviate?utm_source=go-github-com-dnnrly-abbreviate&utm_medium=referral&utm_campaign=enterprise&utm_term=repo)

## License
[Apache 2](https://choosealicense.com/licenses/apache-2.0/)
