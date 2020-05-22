package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
)

var conf Config

type Config struct {
	DirectoryImages     string `yaml:"directoryImages"`
	DirectoryComponents string `yaml:"directoryComponents"`
	SnipPath            string `yaml:"snipPath"`
}

type Image struct {
	Filename      string
	Basename      string
	MaxWidth      int
	ComponentName string
}

func loadConf() {
	f, err := os.OpenFile(".gatigen", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadConf()
	tmpl := template.Must(template.ParseGlob("./.gatigen_templates/*"))
	var images []Image
	err := filepath.Walk(conf.DirectoryImages, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, conf.SnipPath, "", 1)
		ext := strings.ToUpper(filepath.Ext(path))
		if ext != ".JPEG" && ext != ".JPG" && ext != ".GIF" && ext != ".PNG" {
			return nil
		}
		componentName := strcase.ToCamel(strings.Replace(path, "/", "_", -1))
		baseName := filepath.Base(path)

		i := Image{
			Filename:      path,
			Basename:      baseName[:len(baseName)-len(ext)],
			MaxWidth:      1024,
			ComponentName: componentName[:len(componentName)-len(ext)+1] + strings.ToUpper(componentName[len(componentName)-len(ext)+1:]),
		}
		images = append(images, i)

		file, err := os.OpenFile(filepath.Join(conf.DirectoryComponents, fmt.Sprintf("%s.tsx", i.ComponentName)), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.ExecuteTemplate(file, "imageTSX", i)
		if err != nil {
			return err
		}
		file.Close()
		return nil
	})
	file, err := os.OpenFile(filepath.Join(conf.DirectoryComponents, "index.tsx"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.ExecuteTemplate(file, "indexTSX", images)
	if err != nil {
		fmt.Println(err)
	}
}
