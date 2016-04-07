package main

import "net/http"

type AdminResponder struct {
	BaseResponder
}

func NewAdminResponder(linksDB map[string]string) *AdminResponder {
	handler := new(AdminResponder)
	handler.BaseResponder = MakeBaseResponder(linksDB)
	return handler
}

func (responder *AdminResponder) AddShortlinkFormResponse(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	if req.Method != "GET" {
		return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
	}
	form_html := `<html>
<form action="/admin/post" method="POST">
URL: <input type="text" name="url"><br>
Short Code: <input type="text" name="code"><br>
<input type="submit" value="Submit">
</form>
</html>
`
	return http.StatusOK, nil, []byte(form_html)
}

func (responder *AdminResponder) PostResponse(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	if req.Method != "POST" {
		return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
	}

	url_path := req.URL.Path
	switch url_path {
	case "/admin/post":
		responder.linksDB[req.FormValue("code")] = req.FormValue("url")
		output := LinkAddedTemplateOutput(req.FormValue("url"), req.FormValue("code"))
		// normally you want to redirect instead of returning direction...
		// otherwise the user can reload and unintentionally post the same data multiple times.
		return http.StatusOK, nil, output
	default:
		return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
	}
}
