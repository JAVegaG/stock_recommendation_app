package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type responseCaptureWriter struct {
	http.ResponseWriter
	statusCode int
	buf        *bytes.Buffer
}

func (response *responseCaptureWriter) WriteHeader(statusCode int) {
	response.statusCode = statusCode
}

func (response *responseCaptureWriter) Write(b []byte) (int, error) {
	return response.buf.Write(b)
}

// ResponseWrapper middleware wraps JSON responses in a "data" object
func ResponseWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// Capture response
		capture := &responseCaptureWriter{
			ResponseWriter: response,
			statusCode:     http.StatusOK,
			buf:            &bytes.Buffer{},
		}

		next.ServeHTTP(capture, request)

		// Detect content type
		contentType := response.Header().Get("content-type")

		if contentType == "" {
			contentType = http.DetectContentType(capture.buf.Bytes())
			response.Header().Set("Content-Type", contentType)
		}

		// Wrap JSON only if response is JSON and 200
		if capture.statusCode == http.StatusOK &&
			contentType == "application/json" {

			var raw interface{}
			_ = json.Unmarshal(capture.buf.Bytes(), &raw)

			response.WriteHeader(http.StatusOK)
			resp := map[string]interface{}{
				"data": raw,
			}
			_ = json.NewEncoder(response).Encode(resp)
		} else {
			// Fallback to original response
			response.WriteHeader(capture.statusCode)
			io.Copy(response, capture.buf)
		}
	})
}
