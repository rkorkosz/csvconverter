package csvconverter

import (
	"encoding/xml"
	"io"
)

// Convert reads csv file from `io.Reader`
// and writes it to `io.Writer` as xml
func ConvertXML(in io.Reader, out io.Writer) error {
	c := NewConverter(in, xml.NewEncoder(out))
	c.Convert()
	return c.Err()
}
