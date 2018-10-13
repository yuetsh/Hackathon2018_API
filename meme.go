package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type Meme struct {
	Name  string   `json:"name"`
	Subs  []string `json:"subs"`
	hash  string
	paths Paths
}

type Paths struct {
	template struct {
		mp4     string
		ass     string
		palette string
	}
	output struct {
		name string
		ass  string
		gif  string
		mp4  string
	}
}

var (
	ErrSubsLen = errors.New("subs length is wrong")
	ErrNilSubs = errors.New("subs has nil value")
	ErrName    = errors.New("name is wrong")
	NameLenMap = map[string]int{
		"zhenxiang":   4,
		"dagong":      6,
		"sorry":       9,
		"jinkela":     6,
		"diandongche": 6,
		"kongming":    2,
		"marmot":      2,
	}
)

func (m *Meme) check() error {
	m.Name = strings.ToLower(m.Name)
	if val, ok := NameLenMap[m.Name]; !ok {
		return ErrName
	} else {
		if len(m.Subs) != val {
			return ErrSubsLen
		}
		hasNil := false
		for _, v := range m.Subs {
			if v == "" {
				hasNil = true
				continue
			}
		}
		if hasNil {
			return ErrNilSubs
		}
		return nil
	}
}

func (m *Meme) isExist() bool {
	m.hash = NewMd5(m.Subs)
	m.paths.template.mp4 = "./templates/" + m.Name + "/template.mp4"
	m.paths.template.ass = "./templates/" + m.Name + "/template.ass"
	m.paths.template.palette = "./templates/" + m.Name + "/palette.png"
	m.paths.output.name = "./dist/" + m.Name
	m.paths.output.ass = "./dist/" + m.Name + "/" + m.hash + ".ass"
	m.paths.output.gif = "./dist/" + m.Name + "/" + m.hash + ".gif"
	m.paths.output.mp4 = "./dist/" + m.Name + "/" + m.hash + ".mp4"
	if _, err := os.Stat(m.paths.output.name); os.IsNotExist(err) {
		os.Mkdir(m.paths.output.name, os.ModePerm)
	}
	_, err := os.Stat(m.paths.output.ass)
	return !os.IsNotExist(err)
}

func (m *Meme) renderAss() error {
	text := ""
	if buf, err := ioutil.ReadFile(m.paths.template.ass); err != nil {
		return err
	} else {
		text = string(buf)
	}
	if newSub, err := template.New("ASS File").Parse(text); err != nil {
		panic(err)
	} else {
		if file, err := os.Create(m.paths.output.ass); err != nil {
			panic(err)
		} else {
			data := map[string][]string{
				"sentences": m.Subs,
			}
			if err = newSub.Execute(file, data); err != nil {
				panic(err)
			}
		}
		return nil
	}
}

func (m *Meme) renderMp4(c chan bool) error {
	cmd := exec.Command("ffmpeg",
		"-i", m.paths.template.mp4,
		"-vf", "ass="+m.paths.output.ass,
		"-an",
		"-y", m.paths.output.mp4)

	if _, err := cmd.CombinedOutput(); err != nil {
		panic(err)
	}
	c <- true
	return nil
}

func (m *Meme) renderGif(c chan bool) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", m.paths.template.mp4,
		"-i", m.paths.template.palette,
		"-lavfi", "ass="+m.paths.output.ass+",fps=10,scale=300:-1:flags=lanczos[x];[x][1:v]paletteuse",
		"-y", m.paths.output.gif,
	)

	if _, err := cmd.CombinedOutput(); err != nil {
		panic(err)
	}
	c <- true
	return nil
}

func (m *Meme) New() error {
	if err := m.check(); err != nil {
		return err
	}
	if m.isExist() {
		return nil
	} else {
		c := make(chan bool, 2)
		m.renderAss()
		m.renderMp4(c)
		m.renderGif(c)
		<-c
		return nil
	}
}

func NewPalettes() {
	reg := regexp.MustCompile(`templates/(.*)/template.mp4`)
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if reg.MatchString(path) {
			dirname := reg.ReplaceAllString(path, "$1")
			cmd := exec.Command(
				"ffmpeg",
				"-i", "./"+path,
				"-vf", "fps=10,scale=300:-1:flags=lanczos,palettegen",
				"-y", "./templates/"+dirname+"/palette.png",
			)
			if _, err := cmd.CombinedOutput(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
