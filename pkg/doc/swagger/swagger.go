package swagger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kiem-toan/pkg/config"
)

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
	<redoc spec-url='%v'></redoc>
	<script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"></script>
	</body>
</html>`

	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(r.URL.Path, "swagger.json")
		_, _ = fmt.Fprintf(w, tpl, path)
	}
}

func SwaggerHandler(docFile string) http.HandlerFunc {
	data, err := Asset(docFile)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		data, err = Asset(docFile)
		if err != nil {
			panic(err)
		}
		_, _ = w.Write(data)
	}
}

func Asset(name string) ([]byte, error) {
	base := filepath.Join(config.GetAppConfig().ProjectDir, "docs")
	if strings.Contains(name, "..") {
		panic(fmt.Sprintf("invalid name (%v)", name))
	}
	return ioutil.ReadFile(filepath.Join(base, name))
}
