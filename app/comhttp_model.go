package app

import (
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
	DATA      *Data
}

func (j *JORM) parseTemplates(base string, t *template.Template) *template.Template {
	log.Println("cfg.Pathcfg.Pathcfg.Pathssssssssssss", j.config.Path)
	log.Println("ssssssssssss", j.config.Path+"tpl/"+base)
	err := filepath.Walk(j.config.Path+"tpl/"+base, func(path string, info os.FileInfo, err error) error {
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

func (j *JORM) parseFiles(base, tpl string) (*template.Template, error) {
	t := new(template.Template)
	return j.parseTemplates(base, t).ParseFiles(j.config.Path + "tpl/" + tpl)
}
