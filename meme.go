package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Meme struct {
	Name   string   `json:"name"`
	Format string   `json:"format"`
	Subs   []string `json:"subs"`
	hash   string
	paths  Paths
}

type Paths struct {
	template struct {
		mp4 string
		ass string
	}
	output struct {
		name string
		ass  string
		gif  string
		mp4  string
	}
}

const (
	ZhenXiang = "zhenxiang"
	Sorry     = "sorry"
	DaGong    = "dagong"
	JinKeLa   = "jinkela"
	Touche    = "diandongche"
	KongMing  = "kongming"
	Marmot    = "marmot"
)

var (
	ErrSubsLen = errors.New("subs length is wrong")
	ErrName    = errors.New("name is wrong")
	NameLenMap = map[string]int{
		ZhenXiang: 4,
		DaGong:    6,
		Sorry:     9,
		JinKeLa:   6,
		Touche:    6,
		KongMing:  2,
		Marmot:    2,
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
		return nil
	}
}

func (m *Meme) isExist() bool {
	m.hash = NewMd5(m.Subs)
	m.paths.template.mp4 = "./templates/" + m.Name + "/template.mp4"
	m.paths.template.ass = "./templates/" + m.Name + "/template.ass"
	m.paths.output.name = "./dist/" + m.Name
	m.paths.output.ass = "./dist/" + m.Name + "/" + m.hash + ".ass"
	m.paths.output.gif = "./dist/" + m.Name + "/" + m.hash + ".gif"
	m.paths.output.mp4 = "./dist/" + m.Name + "/" + m.hash + ".mp4"
	if _, err := os.Stat(m.paths.output.name); os.IsNotExist(err) {
		os.Mkdir(m.paths.output.name, os.ModePerm)
	}
	if _, err := os.Stat(m.paths.output.ass); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
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

func (m *Meme) renderGif(wg *sync.WaitGroup) error {
	defer wg.Done()
	cmd := exec.Command("ffmpeg",
		"-i", m.paths.template.mp4,
		"-vf", "ass="+m.paths.output.ass+",scale=300:-2",
		"-r", "8",
		"-y", m.paths.output.gif)

	if _, err := cmd.CombinedOutput(); err != nil {
		panic(err)
	}
	return nil
}

func (m *Meme) renderMp4(wg *sync.WaitGroup) error {
	defer wg.Done()
	cmd := exec.Command("ffmpeg",
		"-i", m.paths.template.mp4,
		"-vf", "ass="+m.paths.output.ass,
		"-an",
		"-y", m.paths.output.mp4)

	if _, err := cmd.CombinedOutput(); err != nil {
		panic(err)
	}
	return nil
}

func (m *Meme) New() error {
	if err := m.check(); err != nil {
		return err
	}
	if m.isExist() {
		return nil
	} else {
		if err := m.renderAss(); err != nil {
			return err
		}
		var wg sync.WaitGroup
		wg.Add(2)
		go m.renderGif(&wg)
		go m.renderMp4(&wg)
		wg.Wait()
		return nil
	}
}
