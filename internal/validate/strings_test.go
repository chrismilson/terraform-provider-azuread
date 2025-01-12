// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
)

func TestNoEmptyStrings(t *testing.T) {
	cases := []struct {
		Value    string
		TestName string
		ErrCount int
	}{
		{
			Value:    "!",
			TestName: "Exclamation",
			ErrCount: 0,
		},
		{
			Value:    ".",
			TestName: "Period",
			ErrCount: 0,
		},
		{
			Value:    "-",
			TestName: "Hyphen",
			ErrCount: 0,
		},
		{
			Value:    "_",
			TestName: "Underscore",
			ErrCount: 0,
		},
		{
			Value:    "10.1.0.0/16",
			TestName: "IP",
			ErrCount: 0,
		},
		{
			Value:    "",
			TestName: "Empty",
			ErrCount: 1,
		},
		{
			Value:    " ",
			TestName: "Space",
			ErrCount: 1,
		},
		{
			Value:    "     ",
			TestName: "FiveSpaces",
			ErrCount: 1,
		},
		{
			Value:    "  1",
			TestName: "DoubleSpaceOne",
			ErrCount: 0,
		},
		{
			Value:    "1 ",
			TestName: "OneSpace",
			ErrCount: 0,
		},
		{
			Value:    "\r",
			TestName: "CarriageReturn",
			ErrCount: 1,
		},
		{
			Value:    "\n",
			TestName: "NewLine",
			ErrCount: 1,
		},
		{
			Value:    "\t",
			TestName: "HorizontalTab",
			ErrCount: 1,
		},
		{
			Value:    "\f",
			TestName: "FormFeed",
			ErrCount: 1,
		},
		{
			Value:    "\v",
			TestName: "VerticalTab",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			diags := NoEmptyStrings(tc.Value, cty.Path{})

			if len(diags) != tc.ErrCount {
				t.Fatalf("Expected NoEmptyStrings to have %d not %d errors for %q", tc.ErrCount, len(diags), tc.TestName)
			}
		})
	}
}

func TestStringIsEmailAddress(t *testing.T) {
	cases := []struct {
		Value    string
		TestName string
		ErrCount int
	}{
		{
			Value:    "j.doe@hashicorp.com",
			TestName: "Valid_EmailAddress",
			ErrCount: 0,
		},
		{
			Value:    "j.doehashicorp.com",
			TestName: "Invalid_EmailAddress_NoAtChar",
			ErrCount: 1,
		},
		{
			Value:    "j/doe@ha$hicorp.com",
			TestName: "Invalid_EmailAddress_InvalidChars",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			diags := StringIsEmailAddress(tc.Value, cty.Path{})

			if len(diags) != tc.ErrCount {
				t.Fatalf("Expected StringIsEmailAddress to have %d not %d errors for %q", tc.ErrCount, len(diags), tc.TestName)
			}
		})
	}
}
