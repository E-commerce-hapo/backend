// +build !release

package doc

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Asset(name string) ([]byte, error) {
	base := filepath.Join("github.com/kiem-toan", "doc")
	if strings.Contains(name, "..") {
		panic(fmt.Sprintf("invalid name (%v)", name))
	}
	return ioutil.ReadFile(filepath.Join(base, name))
}
