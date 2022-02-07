package mockserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
)

// NewMockServer creates and returns a mock server
// 1. paramater the response that you want to return
// 2. paramater the status code that you want to return
// 3. paramater the headers that you want to return
func NewMockServer(res []byte, statusCode int, headers []map[string]interface{}) *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		for _, header := range headers {
			for key, value := range header {
				w.Header().Set(key, value.(string))
			}
		}
		w.WriteHeader(statusCode)
		_, _ = w.Write(res)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

// NewMockEchoServer ...
func NewMockEchoServer(target, method, responseBody string, headers []map[string]string) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	request := httptest.NewRequest(method, target, bytes.NewReader([]byte(responseBody)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	for _, header := range headers {
		for key, value := range header {
			request.Header.Set(key, value)
		}
	}
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	return context, request, recorder
}
