package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"path/filepath"
//	"strings"
//
//	"github.com/k0kubun/pp"
//
//	"github.com/E-commerce-hapo/backend/pkg/errorx"
//
//	"github.com/go-openapi/spec"
//	"github.com/gorilla/mux"
//)
//
//func main() {
//	dir, _ := os.Getwd()
//	pp.Println(dir)
//	//pathItems := map[string]spec.PathItem{
//	//	"/api/CreateCategory": spec.PathItem{
//	//		PathItemProps: spec.PathItemProps{
//	//			Post: &spec.Operation{
//	//				OperationProps: spec.OperationProps{
//	//					Description: "DescriptionDescriptionDescriptionDescriptionDescription",
//	//					Tags:        []string{"tag 1", "tag 2", "tag 3"},
//	//					Summary:     "SummarySummarySummarySummary",
//	//					Parameters: []spec.Parameter{
//	//						{
//	//							ParamProps: spec.ParamProps{
//	//								Name:     "body",
//	//								In:       "body",
//	//								Required: true,
//	//								Schema: &spec.Schema{
//	//									SchemaProps: spec.SchemaProps{
//	//										Ref: spec.Ref{Ref: jsonreference.MustCreateRef("prefixDefRequest" + "id")},
//	//									},
//	//								},
//	//							},
//	//						},
//	//					},
//	//					Responses: &spec.Responses{
//	//						ResponsesProps: spec.ResponsesProps{
//	//							StatusCodeResponses: map[int]spec.Response{
//	//								200: {
//	//									ResponseProps: spec.ResponseProps{
//	//										Description: "A successful response",
//	//										Schema: &spec.Schema{
//	//											SchemaProps: spec.SchemaProps{
//	//												Ref: spec.Ref{Ref: jsonreference.MustCreateRef("prefixDefResponse" + "id")},
//	//											},
//	//										},
//	//									},
//	//								},
//	//							},
//	//						},
//	//					},
//	//				},
//	//			},
//	//		},
//	//	},
//	//}
//	//paths := &spec.Paths{
//	//	Paths: pathItems,
//	//}
//	//
//	//definitions := map[string]spec.Schema{
//	//	"customer_name": spec.Schema{
//	//		SchemaProps: spec.SchemaProps{
//	//			Type: spec.StringOrArray{
//	//				"string",
//	//			},
//	//		},
//	//	},
//	//	"deleted": spec.Schema{
//	//		SchemaProps: spec.SchemaProps{
//	//			Type: spec.StringOrArray{
//	//				"boolean",
//	//			},
//	//		},
//	//	},
//	//}
//	//
//	//tags := []spec.Tag{
//	//	{
//	//		TagProps: spec.TagProps{
//	//			Name: "tag swagger 1, swagger 2",
//	//		},
//	//	},
//	//}
//	//err := writeSwaggerFile("", &spec.SwaggerProps{
//	//	Consumes:    []string{"application/json"},
//	//	Produces:    []string{"application/json"},
//	//	Schemes:     []string{"http", "https"},
//	//	Swagger:     "2.0",
//	//	Info:        nil,
//	//	Paths:       paths,
//	//	Definitions: definitions,
//	//	Tags:        tags,
//	//})
//	//pp.Println(err)
//
//	r := mux.NewRouter()
//	r.Handle("/doc/njv", RedocHandler())
//	r.Handle("/doc/njv/swagger.json", SwaggerHandler("/swagger.json"))
//	svr := &http.Server{
//		Addr:    ":8080",
//		Handler: r,
//	}
//	if err := svr.ListenAndServe(); err != nil {
//		panic(err)
//	}
//}
//
//func RedocHandler() http.HandlerFunc {
//	const tpl = `<!DOCTYPE html>
//<html>
//	<head>
//	<title>API Documentation</title>
//	<meta charset="utf-8"/>
//	<meta name="viewport" content="width=device-width, initial-scale=1">
//	<style>body { margin: 0; padding: 0 }</style>
//	</head>
//	<body>
//	<redoc spec-url='%v'></redoc>
//	<script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"></script>
//	</body>
//</html>`
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		path := filepath.Join(r.URL.Path, "swagger.json")
//		_, _ = fmt.Fprintf(w, tpl, path)
//	}
//}
//
//func writeSwaggerFile(docPath string, swaggerDoc *spec.SwaggerProps) (_err error) {
//	//dir := filepath.Join(gen.ProjectPath(), "doc", docPath)
//	filename := filepath.Join("/Users/inmac/workspace/pet-projects/HAPO-Backend/scripts/test/generateSwagger", "swagger.json")
//	//err := os.MkdirAll(dir, 0755)
//	//if err != nil {
//	//	return err
//	//}
//	f, err2 := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
//	if err2 != nil {
//		return err2
//	}
//	defer func() {
//		err3 := f.Close()
//		if _err == nil {
//			_err = err3
//		}
//	}()
//	encoder := json.NewEncoder(f)
//	encoder.SetIndent("", "  ")
//	if err4 := encoder.Encode(swaggerDoc); err4 != nil {
//		return errorx.Errorf(300, nil, "generate swagger: %v", err4)
//	}
//	return nil
//}
//
//func SwaggerHandler(docFile string) http.Handler {
//	data, err := Asset(docFile)
//	if err != nil {
//		panic(err)
//	}
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		data, err = Asset(docFile)
//		if err != nil {
//			panic(err)
//		}
//		_, _ = w.Write(data)
//	})
//}
//
//func Asset(name string) ([]byte, error) {
//	base := "/Users/inmac/workspace/pet-projects/HAPO-Backend/scripts/test/generateSwagger"
//	if strings.Contains(name, "..") {
//		panic(fmt.Sprintf("invalid name (%v)", name))
//	}
//	return ioutil.ReadFile(filepath.Join(base, name))
//}
