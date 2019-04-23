package examples

import (
	"strings"
)

//go:generate dice

const (
	a = iota
	b
	c
)

const (
	const1 string = "string"
	const2 int64  = 2
)

const const3 = 123.123
const const4 = true

var varString1 = "123"
var varString2 string
var varString21 float64 = 123.123
var varString22 = map[string]string{"a": "b"}
var (
	varString3 string
	varString4 string
)

var varFund1 func() error

type typeFund2 func() error

var (
	varFund3 func() error
	varFund4 func() error
)

func fn1() {
	strings.HasPrefix("qwd", "qwd")
}

func fn2() {

}

// Struct1
// dice:stddb.scan
// dice=stddb.value
type Struct1 struct {
	Proxy struct {
		Address string
		Port    string
	}
}

// Struct2
type (
	// @sqldriver(empty_error=true)
	Struct2 struct {
		Address string
		Port    string
	} // comment
	Struct3 struct {
		Address string
		Port    string
	}

	// Struct4 qwd
	// @sqldriver(
	//  value=false
	//  scan=false
	// )
	Struct4 struct {
		Address string
		Port    string
	}
)

func (Struct1) Struct1Fn() error {
	return nil
}

func (Struct4) Struct4Fn() error {
	return nil
}
