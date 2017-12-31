package test

import (
	"testing"
	"runtime"
	"path/filepath"
	"github.com/astaxie/beego"
	"io"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"net/url"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func makeHttpRequest(method string, path string, body io.Reader) (req *http.Request, write *httptest.ResponseRecorder, err error) {
	if body == nil {
		body = ioutil.NopCloser(strings.NewReader(url.Values{}.Encode()))
	}
	r, e := http.NewRequest(method, path, body)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	return r, w, e
}

func TestPingCheck(t *testing.T){
		r, w, _ := makeHttpRequest("GET", "/cmdb/ping", nil)
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		t.Log(w.Body.String())
}