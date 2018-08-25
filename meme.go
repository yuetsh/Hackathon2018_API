package main

import (
	"os/exec"
	"os"
	"io/ioutil"
	"html/template"
	"errors"
)

type Meme struct {
	Name  string   `json:"name" validate:"required"`
	Subs  []string `json:"subs" validate:"required"`
	hash  string
	paths Paths
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
	ZhenXiang = "wangjingze"
	Sorry     = "sorry"
	DaGong    = "dagong"
	JinKeLa   = "jinkela"
	Touche    = "diandongche"
	KongMing  = "kongming"
	Marmot    = "marmot"
)

var (
	ErrFound   = errors.New("found")
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
	if val, ok := NameLenMap[m.Name]; !ok {
		return ErrName
	} else {
		if len(m.Subs) != val {
			return ErrSubsLen
		}
		return nil
	}
}

func (m *Meme) initPaths() {
	m.hash = NewMd5(m.Subs)
	m.paths.template.mp4 = "./templates/" + m.Name + "/template.mp4"
	m.paths.template.ass = "./templates/" + m.Name + "/template.ass"
	m.paths.output.name = "./dist/" + m.Name
	m.paths.output.ass = "./dist/" + m.Name + "/" + m.hash + ".ass"
	m.paths.output.gif = "./dist/" + m.Name + "/" + m.hash + ".gif"
	m.paths.output.mp4 = "./dist/" + m.Name + "/" + m.hash + ".mp4"
}

func (m *Meme) renderAss() error {
	if _, err := os.Stat(m.paths.output.name); os.IsNotExist(err) {
		os.Mkdir(m.paths.output.name, os.ModePerm)
	}

	if _, err := os.Stat(m.paths.output.ass); os.IsNotExist(err) {
		text := ""
		if buf, err := ioutil.ReadFile(m.paths.template.ass); err != nil {
			return err
		} else {
			text = string(buf)
		}
		if newSub, err := template.New("ASS File").Parse(text); err != nil {
			return err
		} else {
			if file, err := os.Create(m.paths.output.ass); err != nil {
				return err
			} else {
				data := map[string][]string{
					"sentences": m.Subs,
				}
				if err = newSub.Execute(file, data); err != nil {
					return err
				}
			}
			return nil
		}
	} else {
		return ErrFound
	}
}

func (m *Meme) renderGif() error {
	cmd := exec.Command("ffmpeg",
		"-i", m.paths.template.mp4,
		"-vf", "ass="+m.paths.output.ass+",scale=300:-2",
		"-r", "8",
		"-y", m.paths.output.gif)

	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func (m *Meme) renderMp4() error {
	cmd := exec.Command("ffmpeg",
		"-i", m.paths.template.mp4,
		"-vf", "ass="+m.paths.output.ass,
		"-an",
		"-y", m.paths.output.mp4)

	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func (m *Meme) New() error {
	if err := m.check(); err != nil {
		return err
	}
	m.initPaths()
	err := m.renderAss()
	switch err {
	case ErrFound:
		return nil
	case nil:
		m.renderGif()
		m.renderMp4()
		return nil
	default:
		return err
	}
}
