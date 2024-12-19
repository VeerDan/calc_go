package application

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"encoding/json"
)

type Response struct {
	Result float64 `json:"result"`
}

func TestRequestHandlerSuccessCase(t *testing.T) {
	testCasesSuccess := []struct {
		name string
		body []byte
		expected float64
	}{
		{
			name: `without devision and breaks, expected result is int`,
			body: []byte(`curl --location 'localhost/api/v1/calculate'
			--header 'Content-Type: application/json'
			--data '{
			  "expression": "2+2*2-1"
			}'`),
			expected: float64(5),
		},
		{
			name: `without devision, expected result is int`,
			body: []byte(`curl --location 'localhost/api/v1/calculate'
			--header 'Content-Type: application/json'
			--data '{
			  "expression": "(2+2-1)*2"
			}'`),
			expected: float64(6),
		},
		{
			name: `devision without breaks, expected result is float`,
			body: []byte(`curl --location 'localhost/api/v1/calculate'
			--header 'Content-Type: application/json'
			--data '{
			  "expression": "1/2*3"
			}'`),
			expected: float64(1)/2*3,
		},
		{
			name: `devision with breaks, expected result is float`,
			body: []byte(`'{"expression": "(1+2)/4"}'`),
			expected: (1 + 2) / 4,
		}, 
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(testCase.body))
			w := httptest.NewRecorder()
			CalcHandler(w, req)
			res := w.Result()
			defer res.Body.Close()
			if res.StatusCode != 200 {
				t.Fatalf("expected code: 200 in case\n %s\n, got: %d", string(testCase.body), res.StatusCode)
			}
			var p Response
			err := json.NewDecoder(res.Body).Decode(&p)
			if err != nil {
				t.Fatalf("Encoding error")
			}
			if p.Result != testCase.expected {
				t.Fatalf("%f should be equal %f", p.Result, testCase.expected)
			}
		})
	}
	// expected := "Hello John"
	// req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", bytes.NewReader(body))
	// w := httptest.NewRecorder()
	// RequestHandler(w, req)
	// res := w.Result()
	// defer res.Body.Close()
	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	t.Errorf("Error: %v", err)
	// }

	// if string(data) != expected {
	// 	t.Errorf("Expected Hello John but got %v", string(data))
	// }

	// if res.StatusCode != http.StatusOK {
	// 	t.Errorf("wrong status code")
	// }
}

func TestRequestHandlerBadRequestCase(t *testing.T) {


	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	w := httptest.NewRecorder()
	CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("wrong status code")
	}
}

func TestRequestInvalidExpression(t *testing.T) {

} 