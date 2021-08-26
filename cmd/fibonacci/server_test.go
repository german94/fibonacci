package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestComputeRequest(t *testing.T) {
	cases := []struct {
		endpointInput      string
		expected           string
		expectedStatusCode int
	}{
		{
			endpointInput:      "/?n=0",
			expected:           "0",
			expectedStatusCode: 200,
		},
		{
			endpointInput:      "/?n=1",
			expected:           "1",
			expectedStatusCode: 200,
		},
		{
			endpointInput:      "/?n=8",
			expected:           "21",
			expectedStatusCode: 200,
		},
		{
			endpointInput:      "/",
			expected:           missingInput,
			expectedStatusCode: 400,
		},
		{
			endpointInput:      "/?n=p",
			expected:           wrongInput,
			expectedStatusCode: 400,
		},
		{
			endpointInput:      "/?t=p",
			expected:           missingInput,
			expectedStatusCode: 400,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(calcHandler))
	defer server.Close()

	for _, testCase := range cases {
		fullURL := server.URL + testCase.endpointInput
		response, err := server.Client().Get(fullURL)

		if err != nil {
			t.Errorf("There was an error when making the request: %s", err.Error())
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("There was an error when reading the response body: %s", err.Error())
		}

		actual := string(body)
		if actual != testCase.expected {
			t.Errorf("Error for input (%s): expected result %s, actual %s", fullURL, testCase.expected, actual)
		}

		if response.StatusCode != testCase.expectedStatusCode {
			t.Errorf("Error for input (%s): expected status code %d, actual %d",
				fullURL, testCase.expectedStatusCode, response.StatusCode)
		}
	}
}
