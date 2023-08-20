package media

import (
	"fmt"
	"testing"
)

func TestDownVideo(t *testing.T) {
	err := DownFile("test.mp4", "https://v26-artist.vlabvod.com/3080551af1889ccce87ca9c87d7c5a97/631493a5/video/tos/cn/tos-cn-v-436d67/af01b5b79fea4da787924fff857362be/?a=4066&ch=0&cr=0&dr=4&er=0&cd=0%7C0%7C0%7C0&br=1989&bt=1989&cs=0&ds=4&ft=.N~IVQnnrThWH6qcOf-bmo&mime_type=video_mp4&qs=0&rc=aWY1Nzw0aWc3aDM7PDk3NUBpM25tcTM6ZjVwZDMzNDZlM0AwY2ItYDRhXmIxLi1hMDU1YSNxamUzcjQwZG5gLS1kYC9zcw%3D%3D&l=2022090419013801019811413010D06F41")

	if err != nil {
		fmt.Println(err)
	}
}
