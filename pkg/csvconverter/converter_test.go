package csvconverter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestConverterWithJSONEncoder(t *testing.T) {
	tests := []struct {
		name, in, out string
	}{
		{
			name: "basic",
			in: `col1,col2,col3
val11,val12,val13
val21,val22,val23`,
			out: `{"col1":"val11","col2":"val12","col3":"val13"}
{"col1":"val21","col2":"val22","col3":"val23"}
`,
		},
	}
	var outBuf bytes.Buffer
	enc := json.NewEncoder(&outBuf)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				outBuf.Reset()
			})
			in := strings.NewReader(test.in)
			c := NewConverter(in, enc)
			c.Convert()
			if err := c.Err(); err != nil {
				t.Fatal(err)
			}
			if test.out != outBuf.String() {
				t.Errorf("got %s, want %s", &outBuf, test.out)
			}
		})
	}
}
