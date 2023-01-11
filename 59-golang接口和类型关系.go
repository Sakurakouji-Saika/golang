// 一个类型实现多个接口
package main

import "fmt"

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile_2 struct {
}

func (m Mobile_2) playMusic() {
	fmt.Printf("play music\n")

}

func (m Mobile_2) playVideo() {
	fmt.Printf("play Video\n")
}

func main() {
	m := Mobile_2{}

	m.playMusic()
	m.playVideo()
}
