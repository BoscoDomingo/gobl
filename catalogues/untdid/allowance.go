package untdid

import (
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/i18n"
	"github.com/invopop/gobl/pkg/here"
)

const (
	// ExtKeyAllowance is used to identify the UNTDID 5189 allownce codes
	// used in discounts.
	ExtKeyAllowance cbc.Key = "untdid-allowance"
)

var extAllowance = &cbc.Definition{
	Key: ExtKeyAllowance,
	Name: i18n.String{
		i18n.EN: "UNTDID 5189 Allowance",
	},
	Desc: i18n.String{
		i18n.EN: here.Doc(`
			UNTDID 5189 code used to describe the allowance type. This list is based on the
			[EN16931 code list](https://ec.europa.eu/digital-building-blocks/sites/display/DIGITAL/Registry+of+supporting+artefacts+to+implement+EN16931#RegistryofsupportingartefactstoimplementEN16931-Codelists)
			values table which focusses on invoices and payments.
		`),
	},
	Values: []*cbc.Definition{
		{
			Code: "41",
			Name: i18n.NewString("Bonus for works ahead of schedule"),
		},
		{
			Code: "42",
			Name: i18n.NewString("Other bonus"),
		},
		{
			Code: "60",
			Name: i18n.NewString("Manufacturer’s consumer discount"),
		},
		{
			Code: "62",
			Name: i18n.NewString("Due to military status"),
		},
		{
			Code: "63",
			Name: i18n.NewString("Due to work accident"),
		},
		{
			Code: "64",
			Name: i18n.NewString("Special agreement"),
		},
		{
			Code: "65",
			Name: i18n.NewString("Production error discount"),
		},
		{
			Code: "66",
			Name: i18n.NewString("New outlet discount"),
		},
		{
			Code: "67",
			Name: i18n.NewString("Sample discount"),
		},
		{
			Code: "68",
			Name: i18n.NewString("End-of-range discount"),
		},
		{
			Code: "70",
			Name: i18n.NewString("Incoterm discount"),
		},
		{
			Code: "71",
			Name: i18n.NewString("Point of sales threshold allowance"),
		},
		{
			Code: "88",
			Name: i18n.NewString("Material surcharge/deduction"),
		},
		{
			Code: "95",
			Name: i18n.NewString("Discount"),
		},
		{
			Code: "100",
			Name: i18n.NewString("Special rebate"),
		},
		{
			Code: "102",
			Name: i18n.NewString("Fixed long term"),
		},
		{
			Code: "103",
			Name: i18n.NewString("Temporary"),
		},
		{
			Code: "104",
			Name: i18n.NewString("Standard"),
		},
		{
			Code: "105",
			Name: i18n.NewString("Yearly turnover"),
		},
	},
}
