package wview

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender(t *testing.T) {
	fm := make(template.FuncMap)
	fm["echo"] = func(v string) string {
		return "$" + v
	}

	engine := New(Config{
		Root:         "_examples/test",
		Extension:    ".tpl",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        fm,
		DisableCache: true,
	})

	Use(engine)

	recorder := httptest.NewRecorder()
	expect := "<v>Index</v>"
	err := Render(recorder, http.StatusOK, "index", M{})
	if err != nil {
		t.Errorf("render error: %v", err)
		return
	}
	assertRecorder(t, recorder, http.StatusOK, expect)
}

func assertRecorder(t *testing.T, recorder *httptest.ResponseRecorder, expectStatusCode int, expectOut string) {
	result := recorder.Result()
	if result.StatusCode != expectStatusCode {
		t.Errorf("actual: %v, expect: %v", result.StatusCode, expectStatusCode)
	}
	resultBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Errorf("read result body error: %v", err)
		return
	}
	val := string(resultBytes)
	if val != expectOut {
		t.Errorf("actual: %v, expect: %v", val, expectOut)
	}
}

var cases = []struct {
	Name string
	Data M
	Out  string
}{
	{
		Name: "echo.tpl",
		Data: M{"name": "GoView"},
		Out:  "$GoView",
	},
	{
		Name: "include",
		Data: M{"name": "GoView"},
		Out:  "<v>IncGoView</v>",
	},
	{
		Name: "index",
		Data: M{},
		Out:  "<v>Index</v>",
	},
	{
		Name: "sum",
		Data: M{
			"sum": func(a int, b int) int {
				return a + b
			},
			"a": 1,
			"b": 2,
		},
		Out: "<v>3</v>",
	},
}

func TestViewEngine_RenderWriter(t *testing.T) {
	fm := make(template.FuncMap)
	fm["echo"] = func(v string) string {
		return "$" + v
	}

	gv := New(Config{
		Root:         "_examples/test",
		Extension:    ".tpl",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        fm,
		DisableCache: true,
	})

	for _, v := range cases {
		buff := new(bytes.Buffer)
		err := gv.RenderWriter(buff, v.Name, v.Data)
		if err != nil {
			t.Errorf("name: %v, data: %v, error: %v", v.Name, v.Data, err)
			continue
		}
		val := buff.String()
		if val != v.Out {
			t.Errorf("actual: %v, expect: %v", val, v.Out)
		}
	}
}
