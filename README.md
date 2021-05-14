# unfmtjson

Format multiline json to one line.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gonejack/unfmtjson)
![Build](https://github.com/gonejack/unfmtjson/actions/workflows/go.yml/badge.svg)
[![GitHub license](https://img.shields.io/github/license/gonejack/unfmtjson.svg?color=blue)](LICENSE)

### Install
```shell
> go get github.com/gonejack/unfmtjson
```

### Usage
```shell
> unfmtjson *.json
```

### Before
```json
{
    "a": 1,
    "b": 2,
    "c": 3
}
```

### After
```json
{"a":1,"b":2,"c":3}
```
