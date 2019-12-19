package encode

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testJSON struct {
	I int     `json:"i"`
	F float64 `json:"f"`
	S string  `json:"s"`
}

func TestNewJSONEncode(t *testing.T) {
	j := NewJSONEncode(nil)
	assert.Nil(t, j)

}

func TestJSONEncode_Name(t *testing.T) {
	assert.Equal(t, NewJSONEncode("").Name(), "json")
}

func TestJSONEncode_Encode(t *testing.T) {
	need := testJSON{
		I: 100,
		F: 3.14,
		S: "test encode json",
	}

	out := bytes.Buffer{}

	data := []interface{}{need, &need}
	for _, v := range data {
		j := NewJSONEncode(v)
		out.Reset()

		j.Encode(&out)

		got := testJSON{}

		err := json.Unmarshal(out.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, got, need)
	}

}
