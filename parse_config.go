/*************************
*                        *
*   Copyright 2019 NXP   *
*                        *
*************************/

package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"text/template"
)

func enc(in []byte) []byte {
	out := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i] ^ byte(195-i)
	}
	return out
}

func main() {
	var conf = struct {
		URLPrefix string `yaml:"url_prefix"`
		Name      string `yaml:"name"`
		Password  string `yaml:"password"`
		RootCA    string `yaml:"root_ca"`
	}{}
	data, _ := ioutil.ReadFile("config.yml")
	yaml.Unmarshal(data, &conf)
	var c = struct {
		URLPrefix string
		Name      []byte
		Password  []byte
		RootCA    string
	}{
		URLPrefix: conf.URLPrefix,
		Name:      enc([]byte(conf.Name)),
		Password:  enc([]byte(conf.Password)),
		RootCA:    conf.RootCA,
	}
	var tpl = `package main

var URLPrefix = "{{.URLPrefix}}"
var Name = []byte{ {{ range $n :=  .Name}} {{$n}}, {{end}} }
var Password =[]byte{ {{range $p := .Password}} {{$p}}, {{end}} }
var RootCA = "{{.RootCA}}"
`
	t := template.New("b-est cfg template")
	t.Parse(tpl)
	w, err := os.Create("config_tmp.go")
	checkErr(err)
	defer w.Close()

	err = t.Execute(w, c)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
