// Code generated by github.com/gebv/go-dice with details DO NOT EDIT.
// File: /Users/gebv/go/src/github.com/gebv/go-dice/examples/sqldriver.go
// Porcess with args: []
// Package name: examples
// Version: ???
package examples

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// dice: START meta information
// File position: "/Users/gebv/go/src/github.com/gebv/go-dice/examples/sqldriver.go:60:1"
// Annotation name: "sqldriver"
// Rule config:
// {"ScanEmptyAsError":true,"ScanPrint":true,"ValuePrint":true}
// dice: END meta information
// dice: START code

func (p *Struct2) Scan(in interface{}) error {
	switch v := in.(type) {

	case nil:
		return errors.New("Struct2: got empty value")

	case []byte:
		buf := bytes.NewBuffer(v)
		err := json.NewDecoder(buf).Decode(p)
		return errors.Wrap(err, "Failed decode Struct2.")
	case string:
		buf := bytes.NewBufferString(v)
		err := json.NewDecoder(buf).Decode(p)
		return errors.Wrap(err, "Failed decode Struct2.")
	default:
		return fmt.Errorf("Struct2: not expected type %T", in)
	}
}

var _ driver.Valuer = (*Struct2)(nil)

// Value returns a driver Value.
func (p Struct2) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(p); err != nil {
		return nil, errors.Wrap(err, "Failed encode Struct2.")
	}
	return buf.Bytes(), nil
}

// dice: END code

// dice: START meta information
// File position: "/Users/gebv/go/src/github.com/gebv/go-dice/examples/sqldriver.go:60:1"
// Annotation name: "sqldriver"
// Rule config:
// {"ScanEmptyAsError":false,"ScanPrint":false,"ValuePrint":false}
// dice: END meta information
// dice: START code

// dice: END code