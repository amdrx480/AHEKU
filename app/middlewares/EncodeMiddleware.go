package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// EncodeCommaMiddleware adalah middleware untuk mengubah respons agar encoding koma.
func EncodeCommaMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil response writer dari konteks
		rw := c.Response().Writer

		// Buat struct ResponseWriter dengan wrapper untuk memodifikasi respons
		modifiedResponseWriter := &responseWriterWithEncoder{ResponseWriter: rw, encodeComma: true}

		// Override response writer di konteks dengan yang sudah dimodifikasi
		c.Response().Writer = modifiedResponseWriter

		// Jalankan handler berikutnya
		err := next(c)

		// Restore original response writer
		c.Response().Writer = rw

		return err
	}
}

// responseWriterWithEncoder adalah struct untuk memodifikasi response writer dengan encoding koma.
type responseWriterWithEncoder struct {
	http.ResponseWriter
	encodeComma bool
}

// Write adalah method untuk menulis data ke response writer dengan mengganti koma jika diperlukan.
func (rw *responseWriterWithEncoder) Write(data []byte) (int, error) {
	if rw.encodeComma {
		// Encode koma
		data = []byte(strings.ReplaceAll(string(data), ",", "%2C"))
	}
	return rw.ResponseWriter.Write(data)
}

// WriteHeader adalah method untuk menulis header response.
func (rw *responseWriterWithEncoder) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
}
