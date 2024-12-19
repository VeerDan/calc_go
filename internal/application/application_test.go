package application

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"io"
)


func TestRequestHandlerSuccessCase(t *testing.T) {
	testCasesSuccess := []struct {
		name string
		requestBody Request
		expectedCode int
		expectedResult string
	}{
		{
			name: `without devision and breaks, expected result is int`,
			requestBody: Request{
				Expression: "2+2*2-1",
			},
			expectedCode: 200,
			expectedResult: fmt.Sprintf(`{"result":"%f"}`, float64(5)),
		},
		{
			name: `without devision, expected result is int`,
			requestBody: Request{
				Expression: "(2+2-1)*2",
			},
			expectedCode: 200,
			expectedResult: fmt.Sprintf(`{"result":"%f"}`, float64(6)),
		},
		{
			name: `devision without breaks, expected result is float`,
			requestBody: Request{
				Expression: "1/2*3",
			},
			expectedCode: 200,
			expectedResult: fmt.Sprintf(`{"result":"%f"}`, float64(1)/2*3),
		},
		{
			name: `devision with breaks, expected result is float`,
			requestBody: Request{
				Expression: "(1+2)/4",
			},
			expectedCode: 200,
			expectedResult: fmt.Sprintf(`{"result":"%f"}`, float64(1 + 2) / 4),
		}, 
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			body_json := fmt.Sprintf(`{"expression":"%s"}`, testCase.requestBody.Expression)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader([]byte(body_json)))
			w := httptest.NewRecorder()
			CalcHandler(w, req)
			res := w.Result()
			defer res.Body.Close()
			if res.StatusCode != 200 {
				t.Fatalf("expected code: 200 in case\n %s\n, got: %d", string(testCase.requestBody.Expression), res.StatusCode)
			}
			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if string(data) != testCase.expectedResult {
				t.Errorf("expected: %s, got: %s", testCase.expectedResult, string(data))
			}
		})
	}
}

func TestRequestHandlerBadRequestCase(t *testing.T) {
	testCasesSuccess := []struct {
		name string
		requestBody Request
		expectedCode int
		expectedResult string
	}{
		
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			body_json := fmt.Sprintf(`{"expression":"%s"}`, testCase.requestBody.Expression)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader([]byte(body_json)))
			w := httptest.NewRecorder()
			CalcHandler(w, req)
			res := w.Result()
			defer res.Body.Close()
			if res.StatusCode != 200 {
				t.Fatalf("expected code: 200 in case\n %s\n, got: %d", string(testCase.requestBody.Expression), res.StatusCode)
			}
			//data, err := io.ReadAll(res.Body)
			//if err != nil {
			//	t.Errorf("Error: %v", err)
			//}
			
		})
	}
}

func TestRequestInvalidExpression(t *testing.T) {

} 