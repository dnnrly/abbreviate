# abbreviate

Shorten your strings using common abbreviations.

## Installation


```bash
go get github.com/dnnrly/abbreviate/abbreviate
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
