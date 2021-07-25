//Package ffmpegutil 用于处理视频相关信息
package ffmpegutil

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//https://blog.csdn.net/AugustDY/article/details/83005690

//TestVideo 测试视频
func TestVideo() {
	dir, _ := os.Getwd()
	SetCommand(dir + "\\ffmpegutil\\ffmpeg\\win64\\bin\\ffmpeg")
	// SetCommand("\\ffmpeg")

	// input := dir + "\\ffmpegutil\\test\\test.mp4"
	input := "./test.mp4"

	output := dir + "\\ffmpegutil\\test\\m3u8\\test.m3u8"
	keyinfo := dir + "\\ffmpegutil\\test\\m3u8\\keyinfo.keyinfo"
	ConvertVideoM3u8(input, output, 2, 0, 0, keyinfo)

	//播放
	// fmt.Println("进入播放测试,out:", output)
	// SetCommand(dir + "\\ffmpegutil\\ffmpeg\\ffplay")
	// cmdArg := []string{output}
	// ExecuteCommandVoid(GetCommand(), cmdArg...)
}

// 切分视频并生成M3U8文件
// ffmpeg -i input.mp4 -c:v libx264 -c:a aac -strict -2 -f hls -hls_time 20 -hls_list_size 0 -hls_wrap 0 output.m3u8
// 相关参数说明：

// -i 输入视频文件
// -c:v 输出视频格式
// -c:a 输出音频格式
// -strict -2 指明音频使有AAC。
// -f hls 输出视频为HTTP Live Stream（M3U8）
// -hls_time 设置每片的长度，默认为2，单位为秒
// -hls_list_size 设置播放列表保存的最多条目，设置为0会保存所有信息，默认为5
// -hls_wrap 设置多少片之后开始覆盖，如果设置为0则不会覆盖，默认值为0。这个选项能够避免在磁盘上存储过多的片，而且能够限制写入磁盘的最多片的数量。
// 注意，播放列表的sequence number对每个segment来说都必须是唯一的，而且它不能和片的文件名（当使用wrap选项时，文件名可能会重复使用）混淆。

//ConvertVideoM3u8 视频切片
//input 输入文件
//output 输出文件
//hlsTime 设置每片的长度，默认为2，单位为秒
//hleListSize 设置播放列表保存的最多条目，设置为0会保存所有信息，默认为5
//hlsWrap 设置多少片之后开始覆盖，如果设置为0则不会覆盖，默认值为0。
func ConvertVideoM3u8(input string, output string, hlsTime int, hleListSize int, hlsWrap int, keyinfo string) {

	//fmt.Println("当前路径",wd)
	cmdArguments := []string{
		"-i",
		input,
		"-c:v",
		"copy",
		"-c:a",
		"copy",
		// "-strict -2",
		// "-f hls",
		"-hls_time",
		strconv.Itoa(hlsTime),
		"-hls_list_size",
		strconv.Itoa(hleListSize),
		"-hls_wrap",
		strconv.Itoa(hlsWrap),
		// "-hls_playlist_type vod",
		"-hls_key_info_file",
		keyinfo,
		output,
	}

	fmt.Println("命令路径：", GetCommand())
	out, err := ExecuteCommandOut(GetCommand(), cmdArguments...)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(out)
}
