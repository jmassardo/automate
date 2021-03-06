package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/chef/automate/components/license-control-service/pkg/keys"
	"github.com/chef/automate/lib/io/fileutils"
)

var genTemplate = `// Code generated by tools/gen-keys. DO NOT EDIT.
//
// Last updated at {{.Date}}
package keys

var BuiltinKeyData = PublicKeysData{
	Keys: []string{
{{range .Keys}}		{{$.Backtick}}{{ . }}{{$.Backtick}}{{end}},
	},
}
`

func fatalf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		fatalf("Unexpected number of arguments (%d). Usage: gen-keys KEYS_URL OUTPUT_PATH", len(os.Args))
	}

	url := os.Args[1]
	path := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		fatalf("could not fetch keys from %s: %s", url, err.Error())
	}
	defer resp.Body.Close() // nolint: errcheck

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fatalf("could not read response")
	}

	keys := keys.PublicKeysData{}
	err = json.Unmarshal(data, &keys)
	if err != nil {
		fatalf("failed to unmarshal response")
	}

	t, err := template.New("keys").Parse(genTemplate)
	if err != nil {
		fatalf("failed to parse template")
	}
	err = fileutils.AtomicWriter(path, func(w io.Writer) error {
		return t.Execute(w, struct {
			Keys     []string
			Backtick string
			Date     string
		}{
			Keys:     keys.Keys,
			Backtick: "`",
			Date:     time.Now().Format(time.RFC1123),
		})
	})
	if err != nil {
		fatalf("failed to write to %s: %s", path, err.Error())
	}
}
