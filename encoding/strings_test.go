package encoding

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Encoding_Strings_UnmarshalJSON_failure(t *testing.T) {
	type tst struct {
		inp string
	}

	var testCases = []tst{
		// Case 001, whitespace
		{
			inp: `" "`,
		},
		// Case 002, string
		{
			inp: `"some-string"`,
		},
		// Case 003, mixed array
		{
			inp: `["Yes",1]`,
		},
		// Case 004, number
		{
			inp: `123`,
		},
		// Case 005, object
		{
			inp: `{}`,
		},
		// Case 006, invalid array
		{
			inp: `"["`,
		},
		// Case 007, invalid JSON string
		{
			inp: `''`,
		},
		// Case 008, no string array
		{
			inp: `[Yes,No]`,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var out Strings
			{
				err = json.Unmarshal([]byte(tc.inp), &out)
				if err == nil {
					t.Fatalf("expected json.Unmarshal to fail")
				}
			}
		})
	}
}

func Test_Encoding_Strings_UnmarshalJSON_success(t *testing.T) {
	type tst struct {
		inp string
		out Strings
	}

	var testCases = []tst{
		// Case 000
		{
			inp: `null`,
			out: nil,
		},
		// Case 001
		{
			inp: `[]`,
			out: Strings{},
		},
		// Case 002
		{
			inp: `["Yes","No"]`,
			out: Strings{"Yes", "No"},
		},
		// Case 003
		{
			inp: `[""]`,
			out: Strings{""},
		},
		// Case 004
		{
			inp: `"[\"foo\",\"bar\"]"`,
			out: Strings{"foo", "bar"},
		},
		// Case 005
		{
			inp: `"[]"`,
			out: Strings{},
		},
		// Case 006
		{
			inp: `""`,
			out: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var out Strings
			{
				err = json.Unmarshal([]byte(tc.inp), &out)
				if err != nil {
					t.Fatal(err)
				}
			}

			if dif := cmp.Diff(tc.out, out); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
