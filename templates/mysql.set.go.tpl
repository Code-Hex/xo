{{- $type := .Name -}}
{{- $short := (shortname $type "set" "text" "buf" "ok" "src") -}}
{{- $reverseNames := .ReverseConstNames -}}
// {{ $type }} is the '{{ .Set.SetName }}' set type from schema '{{ .Schema  }}'.
type {{ $type }} string

// String returns the string value of the {{ $type }}.
func ({{ $short }} {{ $type }}) String() string {
	return string({{ $short }})
}

// Value satisfies the sql/driver.Valuer interface for {{ $type }}.
func ({{ $short }} {{ $type }}) Value() (driver.Value, error) {
	return {{ $short }}.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for {{ $type }}.
func ({{ $short }} *{{ $type }}) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
	    return errors.New("invalid {{ $type }}")
	}
    *da = {{ $type }}(buf)
	return nil
}
