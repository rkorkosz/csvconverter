package csvconverter

import (
	"encoding/csv"
	"io"
	"strings"
)

type encoder interface {
	Encode(v interface{}) error
}

type Converter struct {
	in  io.Reader
	enc encoder
	err error
}

func NewConverter(in io.Reader, enc encoder) *Converter {
	return &Converter{
		in:  in,
		enc: enc,
	}
}

// Convert reads csv file from `io.Reader`
// and encodes it via encoder
func (c *Converter) Convert() {
	r := csv.NewReader(c.in)
	columns, err := r.Read()
	if err != nil {
		c.err = err
	}
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			c.err = err
		}
		out := rowToMap(columns, row)
		err = c.enc.Encode(out)
		if err != nil {
			c.err = err
		}
	}
}

// Err returns convert error
func (c *Converter) Err() error {
	return c.err
}

func rowToMap(columns, row []string) map[string]interface{} {
	out := make(map[string]interface{})
	for i, val := range row {
		out[columns[i]] = strings.Trim(val, `"`)
	}
	return out
}
