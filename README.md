# dice

[![Release](https://img.shields.io/github/release/gebv/go-dice.svg)](https://github.com/gebv/go-dice/releases/latest)
[![CircleCI](https://circleci.com/gh/gebv/go-dice/tree/master.svg?style=svg)](https://circleci.com/gh/gebv/go-dice/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gebv/go-dice)](https://goreportcard.com/report/github.com/gebv/go-dice)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-lint.svg)](https://golangci.com)

Code generator

<img align="right" alt="Dice gopher logo" title="Dice gopher logo" src="./logo/with color dice/go2.svg" width="266">

# Features list

* [x] database/sql/driver Scan & Value sqldriver
* [ ] transfrom <from> <-> <to>
* [ ] stringer (basic type, enum, struct)
* [ ] helper builder for struct with default values, with init map, chan
* [ ] документация DTO объекта от его полей с комментариями к полям (doc, inline)
* [ ] формирование DTO + фильтрация->валидация

# Notes and magic, how's without her?

* used `github.com/BurntSushi/toml` for encode toml because not overrides origin values if key is omitted in the file (it is not public possibility?). It is for to preserve default values for config
* нет возможности записать опции аннотации в inline (в силу ограничения формата toml). Другие варианты? Форматы?

# Logo

Attached *.ai, *.pdf, *.svg version logo. Two versions

1. with color dice

<a href=""><img alt="Color dice gopher logo" title="Color dice gopher logo" src="./logo/with color dice/go2.svg" width="150"></a>

2. with white dice

<a href=""><img alt="White dice gopher logo" title="White dice gopher logo" src="./logo/with white dice/go1.svg" width="150"></a>


