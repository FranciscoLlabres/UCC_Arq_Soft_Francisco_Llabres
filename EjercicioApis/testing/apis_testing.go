package testing

import (
	"testing"

	"github.com/FranciscoLlabres/ej-apis/apiCall"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Breeds     []interface{} `json:"breeds"`
	Categories []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`

	ExpectedError bool
}

func TestDivision(t *testing.T) {
	for _, testCase := range []TestCase{
		{
			ExpectedError: false,
		},
		{
			ExpectedError: true,
		},
	} {
		t.Run(testCase.URL, func(t *testing.T) {
			result, err := apiCall.ApiCall()
			if testCase.ExpectedError {
				assert.NotNil(t, err)
				return
			}
			assert.EqualValues(t, testCase.URL, result)
		})
	}
}
