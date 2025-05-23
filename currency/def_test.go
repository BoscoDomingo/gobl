package currency_test

import (
	"testing"

	"github.com/invopop/gobl/currency"
	"github.com/invopop/gobl/num"
	"github.com/stretchr/testify/assert"
)

func TestDefAmount(t *testing.T) {
	// make a test table
	tests := []struct {
		name     string
		currency currency.Code
		amt      num.Amount
		exp      string
	}{
		{
			name:     "zero EUR",
			currency: currency.EUR,
			amt:      num.MakeAmount(0, 2),
			exp:      "€0,00",
		},
		{
			name:     "thousand EUR",
			currency: currency.EUR,
			amt:      num.MakeAmount(123456, 2),
			exp:      "€1.234,56",
		},
		{
			name:     "million EUR",
			currency: currency.EUR,
			amt:      num.MakeAmount(123456789, 2),
			exp:      "€1.234.567,89",
		},
		{
			name:     "zero USD",
			currency: currency.USD,
			amt:      num.MakeAmount(0, 2),
			exp:      "$0.00",
		},
		{
			name:     "thousand USD",
			currency: currency.USD,
			amt:      num.MakeAmount(123456, 2),
			exp:      "$1,234.56",
		},
		{
			name:     "million USD",
			currency: currency.USD,
			amt:      num.MakeAmount(123456789, 2),
			exp:      "$1,234,567.89",
		},
		{
			name:     "with template",
			currency: currency.AED,
			amt:      num.MakeAmount(123456, 2),
			exp:      "1,234.56 د.إ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := test.currency.Def()
			assert.Equal(t, test.exp, d.FormatAmount(test.amt))
		})
	}
}

func TestDefinitions(t *testing.T) {
	list := currency.Definitions()
	assert.NotEmpty(t, list)
	assert.Equal(t, currency.USD, list[0].ISOCode)
}

func TestDefFormat(t *testing.T) {
	t.Run("with USD", func(t *testing.T) {
		d := currency.USD.Def()
		f := d.Formatter(currency.WithDisambiguateSymbol())
		a := num.MakeAmount(123456, 2)
		assert.Equal(t, "US$1,234.56", f.Amount(a))
	})

	t.Run("with EUR", func(t *testing.T) {
		d := currency.EUR.Def()
		f := d.Formatter(currency.WithDisambiguateSymbol())
		a := num.MakeAmount(123456, 2)
		assert.Equal(t, "€1.234,56", f.Amount(a))
	})

	t.Run("with numeral system", func(t *testing.T) {
		d := currency.USD.Def()
		f := d.Formatter(currency.WithNumeralSystem(num.NumeralArabic))
		a := num.MakeAmount(123456, 2)
		assert.Equal(t, "$١,٢٣٤.٥٦", f.Amount(a))
	})
}

func TestDefFormatPercentage(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		d := currency.USD.Def()
		p := num.MakePercentage(100, 3)
		assert.Equal(t, "10.0%", d.FormatPercentage(p))
	})
}

func TestDefZero(t *testing.T) {
	t.Run("with zero USD", func(t *testing.T) {
		d := currency.USD.Def()
		assert.Equal(t, "0.00", d.Zero().String())
	})
	t.Run("with zero CLP", func(t *testing.T) {
		d := currency.CLP.Def()
		assert.Equal(t, "0", d.Zero().String())
	})
	t.Run("with zero BTC", func(t *testing.T) {
		d := currency.BTC.Def()
		assert.Equal(t, "0.00000000", d.Zero().String())
	})
}

func TestDefRescale(t *testing.T) {
	t.Run("with USD", func(t *testing.T) {
		d := currency.USD.Def()
		a := num.MakeAmount(123456, 4)
		assert.Equal(t, "12.35", d.Rescale(a).String())
	})
	t.Run("with CLP", func(t *testing.T) {
		d := currency.CLP.Def()
		a := num.MakeAmount(123456, 2)
		assert.Equal(t, "1235", d.Rescale(a).String())
	})
}

func TestDefRescaleUp(t *testing.T) {
	t.Run("with USD", func(t *testing.T) {
		d := currency.USD.Def()
		a := num.MakeAmount(123, 0)
		assert.Equal(t, "123.00", d.RescaleUp(a).String())
	})
	t.Run("with USD extra precision", func(t *testing.T) {
		d := currency.USD.Def()
		a := num.MakeAmount(123456, 4)
		assert.Equal(t, "12.3456", d.RescaleUp(a).String())
	})
	t.Run("with CLP", func(t *testing.T) {
		d := currency.CLP.Def()
		a := num.MakeAmount(1234, 0)
		assert.Equal(t, "1234", d.RescaleUp(a).String())
	})
}

func TestDefByISONumber(t *testing.T) {
	t.Run("with 978", func(t *testing.T) {
		d := currency.ByISONumber("978")
		assert.Equal(t, currency.EUR, d.ISOCode)
	})
	t.Run("with 152", func(t *testing.T) {
		d := currency.ByISONumber("152")
		assert.Equal(t, currency.CLP, d.ISOCode)
	})
	t.Run("with 0", func(t *testing.T) {
		d := currency.ByISONumber("0")
		assert.Nil(t, d)
	})
}
