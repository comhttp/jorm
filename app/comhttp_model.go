package app

import (
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
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
	log.Print("ssssssssssss", "tpl/"+base)
	err := filepath.Walk("tpl/"+base, func(path string, info os.FileInfo, err error) error {
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
	return j.parseTemplates(base, t).ParseFiles("tpl/" + tpl)
}

//
//func (j *JORM) parseTemplates(path string) {
//	t := new(template.Template)
//	log.Print("cfg.Pathcfg.Pathcfg.Pathssssssssssss")
//	log.Print("ssssssssssss", path)
//	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
//		if strings.Contains(path, ".gohtml") {
//			log.Print("WalkWalkWalk", path)
//			_, err = j.goHTML.ParseFiles(path)
//			if err != nil {
//				utl.ErrorLog(err)
//			}
//		}
//
//		return err
//	})
//
//	if err != nil {
//		panic(err)
//	}
//	j.goHTML = *t
//	return
//}
//
//func (j *JORM) parseFiles(path string) (*template.Template, error) {
//	return j.goHTML.ParseFiles(path)
//}
