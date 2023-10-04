package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/go-yaml/yaml"
)

type Section struct {
	Title  string `yaml:"title"`
	Rank   int    `yaml:"rank"`
	Shared string
	Cheats []Cheat
}

type Cheat struct {
	Title  string `yaml:"title"`
	Link   string `yaml:"link"`
	Rank   int    `yaml:"rank"`
	Output string
	Code   string
}

func (c *Cheat) BodyMD() string {
	return c.Code
}
func (c *Cheat) BodyTex() string {
	return c.Code
}

func main() {
	var err error

	sectionDirs, err := os.ReadDir("cheats")
	if err != nil {
		panic(err)
	}

	var sections []Section
	for _, sectionDir := range sectionDirs {
		if !sectionDir.IsDir() {
			continue
		}

		var cheatDirs []os.DirEntry
		cheatDirs, err = os.ReadDir(filepath.Join("cheats", sectionDir.Name()))
		if err != nil {
			panic(err)
		}

		var cheats []Cheat
		for _, cheatDir := range cheatDirs {
			if !cheatDir.IsDir() {
				continue
			}

			var cheatYAMLBs []byte
			cheatYAMLBs, err = os.ReadFile(filepath.Join("cheats", sectionDir.Name(), cheatDir.Name(), "cheat.yaml"))
			if err != nil {
				panic(err)
			}

			var cheat Cheat
			err = yaml.Unmarshal(cheatYAMLBs, &cheat)
			if err != nil {
				panic(err)
			}

			var cheatCodeBs []byte
			cheatCodeBs, err = os.ReadFile(filepath.Join("cheats", sectionDir.Name(), cheatDir.Name(), "cheat.rego"))
			if err != nil {
				panic(err)
			}

			cheat.Code = strings.TrimSpace(strings.Replace(
				string(cheatCodeBs), "package cheat", "", 1,
			))

			cheatOutputPath := filepath.Join("cheats", sectionDir.Name(), cheatDir.Name(), "output.json")
			if _, err = os.Stat(cheatOutputPath); !os.IsNotExist(err) {
				cheatOutputBs, err := os.ReadFile(cheatOutputPath)
				if err != nil {
					panic(err)
				}

				cheat.Output = string(cheatOutputBs)
			}

			cheats = append(cheats, cheat)
		}

		sort.Slice(cheats, func(i, j int) bool {
			return cheats[i].Rank < cheats[j].Rank
		})

		var section Section

		var sectionYAMLBs []byte
		sectionYAMLBs, err = os.ReadFile(filepath.Join("cheats", sectionDir.Name(), "section.yaml"))
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(sectionYAMLBs, &section)
		if err != nil {
			panic(err)
		}
		section.Cheats = cheats

		sharedCodePath := filepath.Join("cheats", sectionDir.Name(), "section.rego")

		if _, err = os.Stat(sharedCodePath); !os.IsNotExist(err) {
			sharedCodeBs, err := os.ReadFile(sharedCodePath)
			if err != nil {
				panic(err)
			}

			section.Shared = strings.TrimSpace(strings.Replace(
				string(sharedCodeBs), "package cheat", "", 1,
			))
		}

		sections = append(sections, section)
	}

	sort.Slice(sections, func(i, j int) bool {
		return sections[i].Rank < sections[j].Rank
	})

	// generate md
	mdTemplateBs, err := os.ReadFile("templates/doc.md")
	if err != nil {
		panic(err)
	}

	mdTemplate, err := template.New("md").Parse(string(mdTemplateBs))
	if err != nil {
		panic(err)
	}

	mdOutputPath := "build/cheatsheet.md"
	if _, err = os.Stat(mdOutputPath); os.IsNotExist(err) {
		_, err = os.Create(mdOutputPath)
		if err != nil {
			panic(err)
		}
	}

	mdOutputFile, err := os.OpenFile(mdOutputPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	err = mdTemplate.Execute(mdOutputFile, struct {
		Sections []Section
	}{Sections: sections})
	if err != nil {
		panic(err)
	}

	// generate tex
	texTemplateBs, err := os.ReadFile("templates/doc.tex")
	if err != nil {
		panic(err)
	}

	texTemplate, err := template.New("tex").Parse(string(texTemplateBs))
	if err != nil {
		panic(err)
	}

	texOutputPath := "build/cheatsheet.tex"
	if _, err = os.Stat(texOutputPath); os.IsNotExist(err) {
		_, err = os.Create(texOutputPath)
		if err != nil {
			panic(err)
		}
	}

	texOutputFile, err := os.OpenFile(texOutputPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	err = texTemplate.Execute(texOutputFile, struct {
		Sections []Section
	}{Sections: sections})
	if err != nil {
		panic(err)
	}

}
