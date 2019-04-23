package gen

import (
	"io"
	"text/template"
)

var tplDatabaseValue = template.Must(template.New("database scan value").Parse(`
func (p {{.StructName}}) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(p); err != nil {
		return nil, errors.Wrap(err, "Failed encode {{.StructName}}.")
	}

// 1



	return buf.Bytes(), nil
}

func (p *{{.StructName}}) Scan(in interface{}) error {
	switch v := in.(type) {
	case nil:
		return nil
	case []byte:
		buf := bytes.NewBuffer(v)
		err := json.NewDecoder(buf).Decode(p)
		return errors.Wrap(err, "Failed decode {{.StructName}}.")
	case string:
		buf := bytes.NewBufferString(v)
		err := json.NewDecoder(buf).Decode(p)
		return errors.Wrap(err, "Failed decode {{.StructName}}.")
	default:
		return fmt.Errorf("{{.StructName}}: not expected type %T", in)
	}
}`))

func TplDatabaseValueScan(w io.Writer, structName string) error {
	return tplDatabaseValue.Execute(w, map[string]string{
		"StructName": structName,
	})
}
