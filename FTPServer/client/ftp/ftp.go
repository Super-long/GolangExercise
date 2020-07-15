package client

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"path"
	"strings"
	"exercise/FTPServer/ftp"
)

type FtpClient struct {
	ftp.FtpConn
}

func (ftpCon *FtpClient) WriteCommand(cmdid uint8, args []string) error {
	if cmdid == ftp.Commands["put"] {
		return ftpCon.WritePut(cmdid, args[0])
	}

	var length uint32
	argstr := strings.Join(args, "")
	length = uint32(binary.Size(length)+binary.Size(cmdid)) + uint32(len(argstr))

	err := binary.Write(ftpCon.Con, binary.LittleEndian, length)
	if err != nil {
		return err
	}
	err = binary.Write(ftpCon.Con, binary.LittleEndian, cmdid)
	if err != nil {
		return err
	}
	err = binary.Write(ftpCon.Con, binary.LittleEndian, ftp.Str2sbyte(argstr))
	if err != nil {
		return err
	}
	return nil
}

func (ftpCon *FtpClient) WritePut(cmdid uint8, filePath string) error {
	filePath = strings.Replace(filePath, "\\", "/", -1)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// 发送命令与文件名
	var length uint32
	fileName := ftp.Str2sbyte(path.Base(filePath))
	length = uint32(binary.Size(length)+binary.Size(cmdid)) + uint32(len(fileName))

	err = binary.Write(ftpCon.Con, binary.LittleEndian, length)
	if err != nil {
		return err
	}
	err = binary.Write(ftpCon.Con, binary.LittleEndian, cmdid)
	if err != nil {
		return err
	}
	err = binary.Write(ftpCon.Con, binary.LittleEndian, fileName)
	if err != nil {
		return err
	}

	// 发送文件长度
	fileInfo, err := f.Stat()
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		return errors.New("put 命令不支持发送文件夹，请尝试putdir命令")
	} else {
		err = binary.Write(ftpCon.Con, binary.LittleEndian, fileInfo.Size())
		if err != nil {
			return err
		}
	}

	// 发送文件内容
	buf := make([]byte, 4096)
	bufReader := bufio.NewReader(f)
	for {
		n, err := bufReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		err = binary.Write(ftpCon.Con, binary.LittleEndian, buf[0:n])
		if err != nil {
			return err
		}
	}
	return nil
}

// get只有一个元素,第一个参数为服务端目标文件,取路径最末尾的文件名问客户端存储的文件名
func (ftpCon *FtpClient) HandleGet(filePath string) error {
	// 返回路径的最后一个元素
	// 为空返回. 全部都是斜杠则返回'/'
	fileName := path.Base(filePath)
	f, err := os.Create(fileName)
	if err != nil {
		if os.IsExist(err) {
			err = f.Truncate(0)
			if err != nil {
				return err
			}
		} else {
			return err
		}
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
		if length-total < bufSize {
			buf = buf[0 : length-total]
		}
	}
	return nil
}

