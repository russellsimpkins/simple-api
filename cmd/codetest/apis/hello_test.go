package apis

import (
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	runAPITests(t, []apiTestCase{
		{"Test hello World", "GET", "/hello/:who", "/hello/world", "", GetHello, http.StatusOK, ""},
		{"Test hello Russ", "GET", "/hello/:who", "/hello/russ", "", GetHello, http.StatusOK, ""},
	})
}
