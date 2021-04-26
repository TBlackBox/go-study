package ffmpegutil

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//命令路径
var command string

//SetCommand 设置命令
func SetCommand(cmd string) {
	command = cmd
}

//GetCommand 获取命令
func GetCommand() string {
	return command
}

//Test 测试
func Test() {
	fmt.Println("进入命令行测试。。。")

	dir, _ := os.Getwd()

	SetCommand(dir + "\\ffmpegutil\\ffmpeg\\ffmpeg")

	//fmt.Println("当前路径",wd)
	cmdArguments := []string{
		"-protocol_whitelist",
		"file,http,crypto,tcp,https",
		"-i",
		`C:\work\code\go-study\src\ffmpegutil\test\test.mp4`,
		"-c",
		"copy",
		`C:\work\code\go-study\src\ffmpegutil\test\test.m3u8`,
	}

	execResult := ExecuteCommand(GetCommand(), cmdArguments...)
	if execResult.IsSuccess {
		fmt.Println("命令执行错误:", execResult.err)
	}

	fmt.Println(execResult.out)
}

//ExecResult 命令执行结果返回
type ExecResult struct {
	IsSuccess bool
	PID       int
	ExitCode  uint32
	CmdInfo   string
	err       string
	out       string
}

//ExecuteCommandVoid 不需要输出
func ExecuteCommandVoid(cmdPath string, arg ...string) error {
	cmd := exec.Command(cmdPath, arg...)
	//启动命令
	if err := cmd.Start(); err != nil {
		fmt.Println("启动命令错误", err)
		return err
	}

	// //等待命令执行完成
	// if err := cmd.Wait(); err != nil {
	// 	fmt.Println("等待命令完成错误", err.Error())
	// 	return "", err
	// }
	return nil

}

//ExecuteCommand 执行命令
func ExecuteCommand(cmdPath string, arg ...string) *ExecResult {
	cmd := exec.Command(cmdPath, arg...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	var execResult = ExecResult{}
	if err != nil {
		execResult.IsSuccess = false
		execResult.err = fmt.Sprint(err) + ": " + stderr.String()
		return &execResult
	}

	execResult.IsSuccess = true
	execResult.PID = cmd.ProcessState.Pid()
	execResult.err = stderr.String()
	execResult.out = out.String()
	execResult.ExitCode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitCode
	execResult.CmdInfo = cmd.String()
	return &execResult

}

//ExecuteCommandOut 执行命令 返回输出消息和错误
func ExecuteCommandOut(cmdPath string, arg ...string) (string, error) {
	cmd := exec.Command(cmdPath, arg...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	fmt.Println("原始命令信息：", cmd.String())
	err := cmd.Run()
	if err != nil {
		return "", errors.New(err.Error() + ": " + stderr.String())
	}

	return out.String(), nil

}

// //ExecuteCommand 执行命令
// func ExecuteCommand(cmdPath string, arg ...string) (string, error) {
// 	cmd := exec.Command(cmdPath, arg...)

// 	// result, err := cmd.CombinedOutput()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return "", err
// 	// }

// 	// fmt.Println(string(result))

// 	//标准输出
// 	// stdout, err := cmd.StdoutPipe()
// 	// if err != nil {
// 	// 	fmt.Println("标准输出绑定错误", err)
// 	// 	return "", err
// 	// }

// 	//标准的错误输出
// 	// stderr, err := cmd.StderrPipe()
// 	// if err != nil {
// 	// 	return "", err
// 	// }
// 	// outinfo := bytes.Buffer{}
// 	// cmd.Stdout = &outinfo

// 	var out bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr

// 	//启动命令
// 	// if err := cmd.Start(); err != nil {
// 	// 	fmt.Println("启动命令错误", err)
// 	// 	return "", err
// 	// }

// 	// //等待命令执行完成
// 	// if err := cmd.Wait(); err != nil {
// 	// 	fmt.Println("等待命令完成错误", err.Error())
// 	// 	return "", err
// 	// }

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
// 		return "", err
// 	}

// 	fmt.Println("进程PID:", cmd.ProcessState.Pid())
// 	fmt.Println("进程退出码：", cmd.ProcessState.Sys().(syscall.WaitStatus).ExitCode)
// 	fmt.Println("cmd信息:", cmd.String())

// 	fmt.Println("输出信息：\n", out.String())
// 	fmt.Println("错误信息：\n", stderr.String())

// 	//获取执行结果的字符串
// 	result := ""

// 	return result, nil

// }
