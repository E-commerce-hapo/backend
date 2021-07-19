package doc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func SwaggerHandler(docFile string) http.Handler {
	data, err := ioutil.ReadFile(filepath.Join("root", "doc"))
	if err != nil {
		panic(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, _ = w.Write(data)
	})
}

func RedocHandler() http.HandlerFunc {
	const tpl = `<!DOCTYPE html>
<html>
	<head>
	<title>API Documentation</title>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>body { margin: 0; padding: 0 }</style>
	</head>
	<body>
	<redoc spec-url="localhost:8080/swagger.json""></redoc>
	<script src="https://cdn.jsdelivr.net/npm/redoc@2.0.0-rc.54/bundles/redoc.standalone.js"></script>
sa
	</body>
</html>`

	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, tpl)
	}
}

