package server

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"exercise/FTPServer/ftp"
)

var Commands = map[string]uint8{
	"cd":    uint8(1),
	"ls":    uint8(2),
	"exit":  uint8(3),
	"mkdir": uint8(4),
	"put":   uint8(5),	// 包含mkdir
	"get":   uint8(6),
}

// 设置用户进入后的默认路径
var DefaultDir = map[string]string{
	"windows": "C:/Users",
	"linux":    "/home/lizhaolong/go/src",
}

type userInfo struct {
	//name string
	pwd  string
	home string
}

var lock sync.Once // 初始化users一次
//var users []userInfo

var users map[string]userInfo

func Init() {
	lock.Do(initUsers)
}

type FtpServer struct {
	ftp.FtpConn
}

// 需要在进入的文件前加'/'
func (ftpCon *FtpServer) HandleCd(args []byte) error {
	cwd := ftp.Sbyte2str(args)
	if strings.HasPrefix(cwd, "/") {
		cwd = path.Join(ftpCon.Cwd, cwd)
	}
	f, err := os.Open(cwd)
	if err != nil {
		ftpCon.Write(ftp.Str2sbyte(err.Error()))
		return nil
	}
	defer f.Close()
	finfo, err := f.Stat()
	if err != nil {
		ftpCon.Write(ftp.Str2sbyte(err.Error()))
		return nil
	}
	if !finfo.IsDir() {
		ftpCon.Write(ftp.Str2sbyte("cd parameter must be directory."))
		return nil
	}
	ftpCon.Cwd = cwd
	return ftpCon.Write(ftp.Str2sbyte(cwd))
}

// 参数为路径 并没有使用系统调用,而是深搜了一遍
func (ftpCon *FtpServer) HandleLs(args []byte) error {
	//cwd := ftp.Sbyte2str(args)
	cwd := ftpCon.Cwd
	if strings.HasPrefix(cwd, "/") {
		cwd = path.Join(cwd, ftp.Sbyte2str(args))
	}
	f, err := os.Open(cwd)
	if err != nil {
		ftpCon.Write(ftp.Str2sbyte(err.Error()))
		return nil
	}
	finfo, err := f.Stat()
	if err != nil {
		ftpCon.Write(ftp.Str2sbyte(err.Error()))
		return nil
	}
	if finfo.IsDir() {
		finfos, err := f.Readdir(0)
		if err != nil {
			ftpCon.Write(ftp.Str2sbyte(err.Error()))
		}
		var res string
		res = fmt.Sprintf("Total:%d\n", len(finfos))
		for _, info := range finfos {
			res = res + fmt.Sprintf("%.30s\t%.10d\t%s\n", info.Name(), info.Size(), info.ModTime())
		}
		err = ftpCon.Write(ftp.Str2sbyte(res))
	} else {
		res := fmt.Sprintf("%.30s\t%.10d\t%s\n", finfo.Name(), finfo.Size(), finfo.ModTime())
		err = ftpCon.Write(ftp.Str2sbyte(res))
	}
	if err != nil {
		err = ftpCon.Write(ftp.Str2sbyte(err.Error()))
	}
	return err
}

func (ftpCon *FtpServer) HandleExit(args []byte) error {
	ftpCon.Exit = true
	ftpCon.Write(ftp.Str2sbyte("Byebye."))
	return nil
}

func (ftpCon *FtpServer) HandleMkdir(args []byte) error {
	dir := ftp.Sbyte2str(args)
	if strings.HasPrefix(dir, "/") {
		dir = path.Join(ftpCon.Home, dir)
	} else {
		dir = path.Join(ftpCon.Cwd, dir)
	}

	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return ftpCon.Write(ftp.Str2sbyte("Ok"))
}

func (ftpCon *FtpServer) HandlePut(args []byte) error {
	fileName := ftp.Sbyte2str(args)
	// 先读取文件名称
	f, err := os.Create(path.Join(ftpCon.Cwd, fileName))
	if err != nil {
		// 创建文件出现问题
		ftpCon.Write(ftp.Str2sbyte(fmt.Sprintf(err.Error())))
		return err
	}
	defer f.Close()

	var length int64
	err = binary.Read(ftpCon.Con, binary.LittleEndian, &length)
	if err != nil {
		return err
	}
	var total, bufSize int64
	if length > 4096 {
		bufSize = 4096
	} else {
		bufSize = length
	}
	buf := make([]byte, bufSize)
	for total < length {
		err = binary.Read(ftpCon.Con, binary.LittleEndian, buf)
		if err != nil {
			return err
		}
		n, err := f.Write(buf)
		if err != nil {
			return err
		}
		total += int64(n)
		if (length - total) < bufSize {
			buf = buf[0 : length-total]
		}
	}

	ftpCon.Write(ftp.Str2sbyte("Ok."))
	return nil
}

func (ftpCon *FtpServer) HandleGet(args []byte) error {
	filePath := ftp.Sbyte2str(args)
	if strings.HasPrefix(filePath, "/") {
		filePath = path.Join(ftpCon.Home, filePath)
	} else {
		filePath = path.Join(ftpCon.Cwd, filePath)
	}
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	finfo, err := f.Stat()
	if err != nil {
		return err
	}
	// TODO 暂不支持下载文件夹
	if finfo.IsDir() {
		return binary.Write(ftpCon.Con, binary.LittleEndian, int64(0))
	}

	err = binary.Write(ftpCon.Con, binary.LittleEndian, finfo.Size())
	if err != nil {
		return err
	}
	bufReader := bufio.NewReader(f)
	buf := make([]byte, 4096)
	for {
		n, err := bufReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		err = binary.Write(ftpCon.Con, binary.LittleEndian, buf[0:n])
		if err != nil {
			return err
		}
	}
	ftpCon.Write(ftp.Str2sbyte("Ok."))
	return nil
}

func initUsers() {
	// runtiom.GOOS获取运行时的操作系统名称
	cwd, ok := DefaultDir[runtime.GOOS]
	if !ok {
		log.Fatal("Unsupported system.")
	}
	// 打开存储用户账号密码的目录,其中存储数据的格式为 name;pwd;起始目录
	f, err := os.Open("/home/lizhaolong/go/src/exercise/FTPServer/server/user")
	if err != nil {
		log.Fatal("failed to load users' information.", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		userinfo := strings.Split(line, ";")
		if len(userinfo) < 3 {
			continue
		}
		home := path.Join(cwd, userinfo[2])
		f, err := os.Open(home)
		if err != nil && os.IsNotExist(err) {
			err = os.Mkdir(home, os.ModePerm)	// 0777 权限
			if err != nil {
				log.Fatal("failed to make directory", home)
			}
		} else {
			f.Close()
		}
		users[userinfo[0]] = userInfo{userinfo[1], home}
	}
}

// 验证用户名和密码，返回验证结果true/false和验证通过后的用户家目录
func Validate(name string, pwd string) (pass bool, home string) {
	if len(users) <= 0 {
		return
	}

	if info, ok := users[name]; ok{
		return true, info.home
	}

	return
}