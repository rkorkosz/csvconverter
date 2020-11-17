package csvconverter

import (
	"encoding/json"
	"io"
)

// Convert reads csv file from `io.Reader`
// and writes it to `io.Writer` as json
func ConvertJSON(in io.Reader, out io.Writer) error {
	c := NewConverter(in, json.NewEncoder(out))
	c.Convert()
	return c.Err()
}
