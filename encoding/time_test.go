package encoding

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Encoding_Time_UnmarshalJSON_failure(t *testing.T) {
	type tst struct {
		inp string
	}

	var testCases = []tst{
		// Case 000, whitespace
		{
			inp: `" "`,
		},
		// Case 001, string
		{
			inp: `"some-string"`,
		},
		// Case 002, mixed array
		{
			inp: `["Yes",1]`,
		},
		// Case 003, empty array
		{
			inp: `[]`,
		},
		// Case 004, object
		{
			inp: `{}`,
		},
		// Case 005, invalid array
		{
			inp: `"["`,
		},
		// Case 006, invalid JSON string
		{
			inp: `''`,
		},
		// Case 007, no string array
		{
			inp: `[Yes,No]`,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var out Time
			{
				err = json.Unmarshal([]byte(tc.inp), &out)
				if err == nil {
					t.Fatalf("expected json.Unmarshal to fail")
				}
			}
		})
	}
}

func Test_Encoding_Time_UnmarshalJSON_success(t *testing.T) {
	type tst struct {
		inp string
		out Time
	}

	var testCases = []tst{
		// Case 000, zero
		{
			inp: `0`,
			out: Time{time.Unix(0, 0).UTC()},
		},
		// Case 001, small number
		{
			inp: `123`,
			out: Time{time.Unix(123, 0).UTC()},
		},
		// Case 002, unix timestamp
		{
			inp: `1765823498`,
			out: Time{time.Unix(1765823498, 0).UTC()},
		},
		// Case 003, null
		{
			inp: `null`,
			out: Time{time.Time{}},
		},
		// Case 004, parsed time string
		{
			inp: `"2025-12-08T18:03:03.779632Z"`,
			out: Time{time.Unix(0, 1765216983779632000).UTC()},
		},
		// Case 005, empty string
		{
			inp: `""`,
			out: Time{time.Time{}},
		},
		// Case 006, another parsed time string
		{
			inp: `"2025-12-14 02:00:00+00"`,
			out: Time{time.Unix(1765677600, 0).UTC()},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var out Time
			{
				err = json.Unmarshal([]byte(tc.inp), &out)
				if err != nil {
					t.Fatal(err)
				}
			}

			if dif := cmp.Diff(tc.out.String(), out.String()); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
