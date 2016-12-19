![logo](logo.png)

cvlin (síblɪŋ : しぶりん) is CSV Lint tool. CsV LINt.

## How To Use

* Write rule file in toml format

```toml
id    = "A0[0-9]"
name  = ".*"
point = "^\\d+$"
```

* Prepare CSV file you want to lint

```csv
# id,name,point
A01,Shibuya Rin,100
A02,,200
```

* Check if your CSV is valid (i.e. your CSV satisfies each regexp or not)

```shell
# specify rule file with -r or --rule option
$ cvlin -r rule.toml subject.csv
Valid.
```

* If your CSV is invalid (see ID of 2nd row):

```csv
# id,name,point
A01,Shibuya Rin,100
B02,,200
```

* you get error message

```
$ cvlin -r rule.toml invalid.csv 
Invalid. ( line: 1, column: 0, value: B02, rules: A0[0-9] )
```

## Install

Just download, unzip and place it in $PATH directory. For example, on Linux:

```shell
VERSION=v1.0.0; sudo wget https://github.com/megane42/cvlin/releases/download/$VERSION/cvlin-linux-amd64 -O /usr/local/bin/cvlin
sudo chmod +x /usr/local/bin/cvlin
```

## Rule File

* Rule file is written in toml format.
    * The left-hand side is the name of column. This is just for documentation.
    * The right-hand side is the regexp which the correspond column should satisfy.

* Regexp is read as "**string literal**", not as "raw string literal". So you must escape all special chars.
    * good: `\\d`
    * bad: `\d`

### Embedded Default Rule

* For further portability, you can embed default rule file into the binary.
* Embedded default rule is used when you invoke cvlin without -r option.
* How to:
  * Prepare go build environment
  * `go get -u github.com/jteeuwen/go-bindata/...`
  * `go get github.com/megane42/cvlin/`
  * Write your rule in `default_rules.toml`
  * `cd $GOPATH/src/github.com/megane42/cvlin`
  * `go-bindata --pkg cvlin -o cvlin/bindata.go default_rules.toml`

## For Developer

### Test
* `go test github.com/megane42/cvlin/...`

### Contribution

1. Fork ([https://github.com/megane42/cvlin/fork](https://github.com/megane42/cvlin/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Create a new Pull Request

### Release (just for me)
1. `go install`
1. `sh ./build.sh`
1. `git tag $(cvlin -v)`
1. `ghr $(cvlin -v) ./bin`

## Author

* megane42
    * https://github.com/megane42
    * https://twitter.com/tsdnm

## License

* MIT
