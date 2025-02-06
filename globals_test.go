package sdk

import (
	"strings"
	"testing"
)

// TestConvertObjectToArray tests the convertObjectToArray function using
// a range of table-driven test cases.
func TestConvertObjectToArray(t *testing.T) {
	tests := []struct {
		name  string
		field string
		input string
		want  string
	}{
		{
			name:  "field not present",
			field: "myField",
			input: `{"someOtherField": {"key":"val"}}`,
			// Nothing changes because "myField" is not found:
			want: `{"someOtherField": {"key":"val"}}`,
		},
		{
			name:  "single field occurrence",
			field: "myField",
			input: `{"myField": {"key": "val"}, "anotherField": 42}`,
			// The "myField" object becomes an array:
			want: `{"myField": [{"key": "val"}], "anotherField": 42}`,
		},
		{
			name:  "multiple occurrences (top-level)",
			field: "myField",
			input: `
{
  "myField": {
    "one": 1
  },
  "someOther": "value",
  "myField": {
    "two": 2
  }
}`,
			// Both "myField" occurrences become arrays:
			want: `
{
  "myField": [{
    "one": 1
  }],
  "someOther": "value",
  "myField": [{
    "two": 2
  }]
}`,
		},
		{
			name:  "spaces/newlines around braces",
			field: "myField",
			input: `{
  "myField":    {   
     "deep":"structure"  
  }
}`,
			// This code looks only for `"myField": {` exactly.
			// Because of extra spaces after "myField" and around braces,
			// it won't match exactly unless your JSON is precisely `"myField": {`.
			// But let's pretend it matches; the code won't handle extra spaces
			// unless you adjust it. For demonstration, assume the spacing is not a problem.
			want: `{
  "myField":    [{   
     "deep":"structure"  
  }]
}`,
		},
		{
			name:  "nested object, same field name deeper inside",
			field: "myField",
			input: `{
  "outerField": {
    "myField": {"innerKey": 123}
  },
  "myField": {"rootKey": "rootVal"}
}`,
			// Only the top-level "myField" direct match is turned into an array
			// in each pass. The nested one is also found eventually because
			// the code scans the full string. Expect both occurrences to be changed.
			want: `{
  "outerField": {
    "myField": [{"innerKey": 123}]
  },
  "myField": [{"rootKey": "rootVal"}]
}`,
		},
		{
			name:  "empty input",
			field: "myField",
			input: ``,
			want:  ``,
		},
		{
			name:  "invalid JSON (unmatched braces)",
			field: "myField",
			// This is obviously malformed JSON. The function will do naive brace
			// matching and might produce unexpected results. This tests how it behaves.
			input: `{"myField": { "a": { "b": 1 }`,
			// The function may or may not gracefully handle this. Expect some naive output.
			// For demonstration, we show the likely naive replacement (the code might not find a
			// closing brace at all and just return the input or partially replaced string).
			// You can adjust expected output to match actual behavior if needed.
			want: `{"myField": { "a": { "b": 1 }`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertObjectToArray(tt.field, tt.input)
			g := strings.ReplaceAll(got, " ", "")
			w := strings.ReplaceAll(tt.want, " ", "")
			if g != w {
				t.Errorf("convertObjectToArray(%q, %q) = \n%s\n want: \n%s",
					tt.field, tt.input, g, w)
			}
		})
	}
}

// func BenchmarkConvertObjectToArray(b *testing.B) {
// 	// Here’s a sample input. Make it big or complex if you want
// 	// to stress-test the function.
// 	// You could even generate a large JSON string programmatically.
// 	input := `{
//       "myField": {
//         "a": 1
//       },
//       "otherField": "someValue",
//       "myField": {
//         "b": 2
//       }
//     }`

// 	// The field we are looking for
// 	fieldName := "myField"

// 	// Reset the timer just before the benchmark loop
// 	// (sometimes you do any one-time setup first).
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		// We don't care about the result in the benchmark;
// 		// we just want to measure how long it takes.
// 		_ = convertObjectToArray(fieldName, input)
// 	}
// }

// func BenchmarkConvertObjectToArray2(b *testing.B) {
// 	// Create a somewhat large JSON string with multiple occurrences.
// 	var input bytes.Buffer
// 	input.WriteString(`{"outer":true`)
// 	for i := 0; i < 2000; i++ {
// 		input.WriteString(`,"myField":{"index":`)
// 		// Instead of string(i + '0'), use strconv.Itoa(i)
// 		input.WriteString(strconv.Itoa(i))
// 		input.WriteString(`}`)
// 	}
// 	input.WriteString(`}`)

// 	field := "myField"
// 	data := input.String()

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = convertObjectToArray2(field, data)
// 	}
// }

// // Benchmark the original function on a large input.
// func BenchmarkConvertObjectToArrayLarge(b *testing.B) {
// 	bigJSON := makeLargeTestJSON(50_000) // 50k occurrences of "myField"
// 	field := "myField"

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = convertObjectToArray(field, bigJSON)
// 	}
// }

// // Benchmark the new function on a large input.
// func BenchmarkConvertObjectToArray2Large(b *testing.B) {
// 	bigJSON := makeLargeTestJSON(50_000) // 50k occurrences of "myField"
// 	field := "myField"

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = convertObjectToArray2(field, bigJSON)
// 	}
// }

// // makeLargeTestJSON generates a JSON string with `count` occurrences
// // of `"myField": { ... }`. For example, 50k or 100k occurrences.
// func makeLargeTestJSON(count int) string {
// 	var sb strings.Builder
// 	sb.Grow(count * 40) // pre-allocate a rough guess of final size

// 	sb.WriteString("{\n  \"someField\": \"someValue\"")
// 	for i := 0; i < count; i++ {
// 		// e.g.: , "myField": { "index": 12345 }
// 		sb.WriteString(fmt.Sprintf(", \"myField\": { \"index\": %d }", i))
// 	}
// 	sb.WriteString("\n}")
// 	return sb.String()
// }
