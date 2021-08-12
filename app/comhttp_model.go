package app

import (
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/utl"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type COMHTTP struct {
	ID        string
	TLD       string
	Slug      string
	Titile    string
	ProtoURL  string
	BodyClass string
}

func parseTemplates(base string, t *template.Template) *template.Template {
	log.Println("cfg.Pathcfg.Pathcfg.Pathssssssssssss", cfg.Path)
	log.Println("ssssssssssss", cfg.Path+"tpl/"+base)
	err := filepath.Walk(cfg.Path+"tpl/"+base, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".gohtml") {
			_, err = t.ParseFiles(path)
			if err != nil {
				utl.ErrorLog(err)
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
