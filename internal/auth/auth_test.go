package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tcs := []struct {
		name           string
		header         http.Header
		expectedOutput string
		expectedError  error
	}{
		{
			name: "Valid API Key",
			header: http.Header{
				"Authorization": []string{"ApiKey baba-baba-black-sheep"},
			},
			expectedOutput: "baba-baba-black-sheep",
			expectedError:  nil,
		},
		{
			name: "No API Key",
			header: http.Header{
				"Authorization": []string{""},
			},
			expectedOutput: "",
			expectedError:  errors.New("no authorization header included"),
		},
		{
			name: "Invalid API Key",
			header: http.Header{
				"Authorization": []string{"baba-baba-black-sheep baba-baba-black-sheep"},
			},
			expectedOutput: "",
			expectedError:  errors.New("malformed authorization header"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetAPIKey(tc.header)
			// Checking errors
			if tc.expectedError == nil && err != nil {
				t.Errorf("FAILED TEST CASE | Name: %v\n---------------------------------------------------------------------------------\n", tc.name)
				t.Errorf("Expected Output: %#v | Code Output: %#v\n", tc.expectedOutput, result)
				t.Errorf("Expected Error: %#v | Code Error: %#v\n\n\n", tc.expectedError, err)
			}

			if tc.expectedError != nil && err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("FAILED TEST CASE | Name: %v\n---------------------------------------------------------------------------------\n", tc.name)
				t.Errorf("Expected Output: %#v | Code Output: %#v\n", tc.expectedOutput, result)
				t.Errorf("Expected Error: %#v | Code Error: %#v\n\n\n", tc.expectedError, err)
			}

			// Checking outputs
			if result != tc.expectedOutput {
				t.Errorf("FAILED TEST CASE | Name: %v\n---------------------------------------------------------------------------------\n", tc.name)
				t.Errorf("Expected Output: %#v | Code Output: %#v\n", tc.expectedOutput, result)
				t.Errorf("Expected Error: %#v | Code Error: %#v\n\n\n", tc.expectedError, err)
			}

		})
	}
}
