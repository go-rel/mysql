package mysql

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQuote_Panic(t *testing.T) {
	quoter := Quote{}
	assert.PanicsWithValue(t, "unsupported value", func() {
		quoter.Value(1)
	})
}

func TestQuote_ID(t *testing.T) {
	quoter := Quote{}

	cases := []struct {
		input string
		want  string
	}{
		{`foo`, "`foo`"},
		{`foo bar baz`, "`foo bar baz`"},
		{"foo`bar", "`foo``bar`"},
		{"foo\x00bar", "`foo`"},
		{"\x00foo", "``"},
	}

	for _, test := range cases {
		assert.Equal(t, test.want, quoter.ID(test.input))
	}
}

func TestQuote_Value(t *testing.T) {
	quoter := Quote{}

	cases := []struct {
		input string
		want  string
	}{
		{"foo\x00bar", "'foo\\0bar'"},
		{"foo\nbar", "'foo\\nbar'"},
		{"foo\rbar", "'foo\\rbar'"},
		{"foo\x1abar", "'foo\\Zbar'"},
		{"foo\"bar", "'foo\\\"bar'"},
		{"foo\\bar", "'foo\\\\bar'"},
		{"foo'bar", "'foo\\'bar'"},
	}

	for _, test := range cases {
		assert.Equal(t, test.want, quoter.Value(test.input))
	}
}

type customType int

func (c customType) Value() (driver.Value, error) {
	return int(c), nil
}

func TestValueConvert_CustomType(t *testing.T) {
	valuer := ValueConvert{}
	v, err := valuer.ConvertValue(customType(1))
	assert.EqualError(t, err, "non-Value type int returned from Value")
	assert.Nil(t, v)
}

func TestValueConvert_DateTime(t *testing.T) {
	valuer := ValueConvert{}
	v, err := valuer.ConvertValue(time.Unix(1633934368, 0).UTC())
	assert.NoError(t, err)
	assert.Equal(t, "2021-10-11 06:39:28", v)
}
