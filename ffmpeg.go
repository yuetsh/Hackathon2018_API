package main

import (
	"os/exec"
	"os"
	"io/ioutil"
	"html/template"
	"crypto/md5"
	"strings"
	"encoding/hex"
)

//常用参数说明：
//主要参数：
//-i 设定输入流
//-f 设定输出格式
//-ss 开始时间
//视频参数：
//-b 设定视频流量，默认为200Kbit/s
//-r 设定帧速率，默认为25
//-s 设定画面的宽与高
//-aspect 设定画面的比例
//-vn 不处理视频
//-vcodec 设定视频编解码器，未设定时则使用与输入流相同的编解码器
//音频参数：
//-ar 设定采样率
//-ac 设定声音的Channel数
//-acodec 设定声音编解码器，未设定时则使用与输入流相同的编解码器
//-an 不处理音频

//func New(name string, subs *Subs, mode string) (hash string, err error) {
//	path := "./templates/" + name
//	sub := path + "/template.ass"
//	mp4 := path + "/template.mp4"
//	outPath := "./dist/" + name
//	hash = subs.Hash()
//
//	if _, err := os.Stat(outPath); os.IsNotExist(err) {
//		os.Mkdir(outPath, os.ModePerm)
//	}
//
//	subOut := outPath + "/" + hash + ".ass"
//	out := ""
//	if mode == "gif" {
//		out = outPath + "/" + hash + ".gif"
//	} else {
//		out = outPath + "/" + hash + ".mp4"
//	}
//
//	if _, err = os.Stat(out); os.IsNotExist(err) {
//		// 不存在相同的表情包
//		// 读取模板字幕文件
//		var assText = ""
//		if buf, err := ioutil.ReadFile(sub); err != nil {
//			return hash, err
//		} else {
//			assText = string(buf)
//		}
//		// 生成新的字幕文件
//		if newSub, err := template.New("ASS File").Parse(assText); err != nil {
//			return hash, err
//		} else {
//			if file, err := os.Create(subOut); err != nil {
//				return hash, err
//			} else {
//				data := map[string][]string{
//					"sentences": subs.subs,
//				}
//				if err = newSub.Execute(file, data); err != nil {
//					return hash, err
//				}
//			}
//		}
//		// 生成新的 Gif/Mp4 文件
//		var cmd *exec.Cmd
//		if mode == "gif" {
//			cmd = exec.Command("ffmpeg",
//				"-i", mp4,
//				"-vf", "ass="+subOut+",scale=300:-1",
//				"-r", "10",
//				"-y", out)
//		} else {
//			cmd = exec.Command("ffmpeg",
//				"-i", mp4,
//				"-vf", "ass="+subOut,
//				"-an",
//				"-y", out)
//		}
//
//		if _, err = cmd.CombinedOutput(); err != nil {
//			return
//		}
//	}
//	return hash, nil
//}

func newAss(name string, subs []string) (hash string, err error) {
	origin := "./templates/" + name + "/template.ass"
	outPath := "./dist/" + name

	cipher := md5.New()
	text := []byte(strings.Join(subs, ","))
	cipher.Write(text)
	hash = hex.EncodeToString(cipher.Sum(nil))

	if _, err = os.Stat(outPath); os.IsNotExist(err) {
		os.Mkdir(outPath, os.ModePerm)
	}

	subOut := outPath + "/" + hash + ".ass"

	if _, err = os.Stat(subOut); os.IsNotExist(err) {
		assText := ""
		if buf, err := ioutil.ReadFile(origin); err != nil {
			return hash, err
		} else {
			assText = string(buf)
		}
		if newSub, err := template.New("ASS File").Parse(assText); err != nil {
			return hash, err
		} else {
			if file, err := os.Create(subOut); err != nil {
				return hash, err
			} else {
				data := map[string][]string{
					"sentences": subs,
				}
				if err = newSub.Execute(file, data); err != nil {
					return hash, err
				}
			}
			return hash, nil
		}
	} else {
		return hash, nil
	}
}

func NewGif(name string, hash string) error {
	subOut := "./dist/" + name + "/" + hash + ".ass"
	origin := "./templates/" + name + "/template.mp4"
	out := "./dist/" + name + "/" + hash + ".gif"

	cmd := exec.Command("ffmpeg",
		"-i", origin,
		"-vf", "ass="+subOut+",scale=300:-1",
		"-r", "10",
		"-y", out)

	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func NewMp4(name string, hash string) error {
	subOut := "./dist/" + name + "/" + hash + ".ass"
	origin := "./templates/" + name + "/template.mp4"
	out := "./dist/" + name + "/" + hash + ".mp4"

	cmd := exec.Command("ffmpeg",
		"-i", origin,
		"-vf", "ass="+subOut,
		"-an",
		"-y", out)

	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func New2(name string, subs []string) (hash string, err error) {
	if hash, err = newAss(name, subs); err != nil {
		return
	} else {
		NewGif(name, hash)
		NewMp4(name, hash)
	}
	return
}
