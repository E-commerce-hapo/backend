package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/go-openapi/spec"
	"github.com/kiem-toan/pkg/errorx"
)

func writeSwaggerFile(docPath string, swaggerDoc *spec.SwaggerProps) (_err error) {
	//dir := filepath.Join(gen.ProjectPath(), "doc", docPath)
	filename := filepath.Join("/Users/inmac/workspace/pet-projects/HAPO-Backend/scripts/test/generateSwagger", "swagger.json")
	//err := os.MkdirAll(dir, 0755)
	//if err != nil {
	//	return err
	//}
	f, err2 := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err2 != nil {
		return err2
	}
	defer func() {
		err3 := f.Close()
		if _err == nil {
			_err = err3
		}
	}()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err4 := encoder.Encode(swaggerDoc); err4 != nil {
		return errorx.Errorf(300, nil, "generate swagger: %v", err4)
	}
	return nil
}
