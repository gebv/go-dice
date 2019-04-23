package stddatabase

import "text/template"

var tplSqlDriverValue = template.Must(template.New("database/driver#value").Parse(`
var _ driver.Valuer = (*{{.StructName}})(nil)

// Value returns a driver Value.
func (p {{.StructName}}) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(p); err != nil {
		return nil, errors.Wrap(err, "Failed encode {{.StructName}}.")
	}
	return buf.Bytes(), nil
}`))

var tplSqlDriverScan = template.Must(template.New("database/driver#scan").Parse(`
func (p *{{.StructName}}) Scan(in interface{}) error {
	switch v := in.(type) {
	{{ if .EmptyIsError }}
	case nil:
		return errors.New("{{.StructName}}: got empty value")
	{{ else }}
	case nil:
		return nil
	{{ end }}
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
