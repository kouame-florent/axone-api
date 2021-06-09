package main

import (
	"html/template"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		generateRepo(os.Args[i], repoTpl, "../internal/repo/")
	}

}

func generateRepo(name, tpl, pkg string) {
	tt := template.Must(template.New("Entity").Funcs(funcMap).Parse(tpl))
	dest := strings.ToLower(path.Join(pkg, name)) + ".go"
	file, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}

	values := map[string]string{
		"Entity": name,
	}

	tt.Execute(file, values)
}

var funcMap map[string]interface{} = template.FuncMap{
	"ToLower": strings.ToLower,
}
