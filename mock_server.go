package mockserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
)

// NewMockServer creates and returns a mock server
func NewMockServer(res []byte, statusCode int, headers map[string]string) *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		for key, value := range headers {
			w.Header().Set(key, value)
		}

		w.WriteHeader(statusCode)
		_, _ = w.Write(res)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// NewMockEchoServer ...
func NewMockEchoServer(target, method, responseBody string) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	request := httptest.NewRequest(method, target, bytes.NewReader([]byte(responseBody)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)

	return context, request, recorder
}
