package file

import (
	"errors"
	"fmt"
	"os"
)

func init(){
	fmt.Print("对文件的操作")

}

// StartFileTest
func StartFileTest(){

}

const (
	FILE_NOT_EXIST = "文件不存在"
)

//CreateFile 创建文件 默认权限o666
func CreateFile(name string) (*os.File,error){
	if len(name) == 0 {
		return nil,errors.New("文件名为空")
	}
	return os.Create(name)
}

//
func NewFile(name string,fd uintptr) (*os.File,error){
	if len(name) == 0 {
		return nil,errors.New("文件名为空")
	}
	return os.NewFile(fd,name),nil
}

//Open 打开文件,自读方式打开
func Open(name string)(file *os.File,err error){
	if len(name) == 0 {
		return nil,errors.New("文件名为空")
	}
	return os.Open(name)
}

//OpenFile 打开文件
//flag  打开方式
//perm 权限
func OpenFile(name string,flag int,perm os.FileMode)(file *os.File,err error){
	return os.OpenFile(name,flag,perm)
}

//WriteStringTooFile 写字符串到文件 n n为写入的长度
func WriteStringTooFile(name string,content string)(n int,err error){
	return WriteByteTooFile(name,[]byte(content))
}

//WriteByteTooFile 写字节数组到文件
func WriteByteTooFile(name string,b []byte)(n int,err error){
	if len(name) == 0 {
		return 0,errors.New("文件名为空")
	}
	file,err := CreateFile(name)
	defer file.Close()
	if err != nil{
		return 0,err
	}
	return file.Write(b)
}


func ReadFile(name string) (string,error){
	if len(name) == 0 {
		return "",errors.New(FILE_NOT_EXIST)
	}
	//TODO
	return "",nil
}