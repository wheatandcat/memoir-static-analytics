# github.com/wheatandcat/memoir-static-analytics/checkcustomerror


[![pkg.go.dev][gopkg-badge]][gopkg]

## Install

```bash
$ go install github.com/wheatandcat/memoir-static-analytics/checkcustomerror/cmd/checkcustomerror@v0.0.6
```

## How to use

```bash
$ go vet -vettool=$(which checkcustomerror) ./...
```

## Version Up

```bash
$ git checkout main
$ git pull --ff-only origin main
$ git tag -a v1.0.0 -m 'リリース内容'
$ git push origin v1.0.0
```
