package our

import (
	"fmt"
	"github.com/comhttp/jorm/pkg/cfg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Page struct {
	ID       string
	TLD      string
	Slug     string
	Titile   string
	ProtoURL string

	BodyClass string
}

func parseTemplates(base string, t *template.Template) *template.Template {
	fmt.Println("cfg.Pathcfg.Pathcfg.Pathssssssssssss", cfg.Path)
	fmt.Println("ssssssssssss", cfg.Path+"tpl/"+base)
	err := filepath.Walk(cfg.Path+"tpl/"+base, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".gohtml") {
			_, err = t.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}
	return t
}

func parseFiles(base, tpl string) (*template.Template, error) {
	t := new(template.Template)
	return parseTemplates(base, t).ParseFiles(cfg.Path + "tpl/" + tpl)
}
