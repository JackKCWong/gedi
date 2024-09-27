package internal

import (
	"strings"
	"testing"
)

func TestSsvReader_Read(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		maxNumOfFields int
		expectedOutput []Record
	}{
		{
			name:           "Basic SSV parsing",
			input:          "1 2 3\n4 5 6\n7 8 9",
			maxNumOfFields: -1,
			expectedOutput: []Record{
				&record{lineno: 1, raw: "1 2 3", parsed: map[string]any{"x": []string{"1", "2", "3"}}},
				&record{lineno: 2, raw: "4 5 6", parsed: map[string]any{"x": []string{"4", "5", "6"}}},
				&record{lineno: 3, raw: "7 8 9", parsed: map[string]any{"x": []string{"7", "8", "9"}}},
			},
		},
		{
			name:           "SSV parsing with maxNumOfFields",
			input:          "1 2 3 4\n5 6 7 8\n9 10 11 12",
			maxNumOfFields: 3,
			expectedOutput: []Record{
				&record{lineno: 1, raw: "1 2 3 4", parsed: map[string]any{"x": []string{"1", "2", "3 4"}}},
				&record{lineno: 2, raw: "5 6 7 8", parsed: map[string]any{"x": []string{"5", "6", "7 8"}}},
				&record{lineno: 3, raw: "9 10 11 12", parsed: map[string]any{"x": []string{"9", "10", "11 12"}}},
			},
		},
		{
			name:           "SSV parsing `ps` output with maxNumOfFields",
			input:          `PID TTY           TIME CMD
 2858 ttys001    0:00.17 /bin/zsh -il
 2859 ttys003    0:00.17 /bin/zsh -il
 2860 ttys005    0:00.17 /bin/zsh -il
 7855 ttys008    0:00.58 /bin/zsh -il
 4973 ttys010    0:00.22 -zsh`,
			maxNumOfFields: 4,
			expectedOutput: []Record{
				&record{lineno: 1, raw: "PID TTY           TIME CMD", parsed: map[string]any{"x": []string{"PID", "TTY", "TIME", "CMD"}}},
				&record{lineno: 2, raw: " 2858 ttys001    0:00.17 /bin/zsh -il", parsed: map[string]any{"x": []string{"2858", "ttys001", "0:00.17", "/bin/zsh -il"}}},
				&record{lineno: 3, raw: " 2859 ttys003    0:00.17 /bin/zsh -il", parsed: map[string]any{"x": []string{"2859", "ttys003", "0:00.17", "/bin/zsh -il"}}},
				&record{lineno: 4, raw: " 2860 ttys005    0:00.17 /bin/zsh -il", parsed: map[string]any{"x": []string{"2860", "ttys005", "0:00.17", "/bin/zsh -il"}}},
				&record{lineno: 5, raw: " 7855 ttys008    0:00.58 /bin/zsh -il", parsed: map[string]any{"x": []string{"7855", "ttys008", "0:00.58", "/bin/zsh -il"}}},
				&record{lineno: 6, raw: " 4973 ttys010    0:00.22 -zsh", parsed: map[string]any{"x": []string{"4973", "ttys010", "0:00.22", "-zsh"}}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := &SsvReader{MaxNumOfFields: tc.maxNumOfFields}
			records, err := reader.Read(strings.NewReader(tc.input))
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			var output []Record
			for record := range records {
				output = append(output, record)
			}

			if len(output) != len(tc.expectedOutput) {
				t.Fatalf("Expected %d records, got %d", len(tc.expectedOutput), len(output))
			}

			for i, expected := range tc.expectedOutput {
				actual := output[i]
				if actual.LineNo() != expected.LineNo() {
					t.Errorf("Record %d: expected lineno %d, got %d", i, expected.LineNo(), actual.LineNo())
				}
				if actual.String() != expected.String() {
					t.Errorf("Record %d: expected raw %q, got %q", i, expected.String(), actual.String())
				}
				expectedFields := expected.Parsed()["x"].([]string)
				actualFields := actual.Parsed()["x"].([]string)
				if len(actualFields) != len(expectedFields) {
					t.Errorf("Record %d: expected %d fields, got %d", i, len(expectedFields), len(actualFields))
				}
				for j, expectedField := range expectedFields {
					if actualFields[j] != expectedField {
						t.Errorf("Record %d, Field %d: expected %q, got %q", i, j, expectedField, actualFields[j])
					}
				}
			}
		})
	}
}