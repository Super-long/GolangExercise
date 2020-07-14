package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"exercise/FTPServer/client/ftp"
	"exercise/FTPServer/ftp"
)

func printHelp() {
	log.Println("Help:\t[command] [args]\ncd [path]\n")
}

/*
 * para@ ftpCon  服务端套接字
 * para@ command 发送的命令
 * para@ args    命令参数
 */
func handleCommand(ftpCon *client.FtpClient, command string, args []string) (err error) {
	cmdid, ok := ftp.Commands[command]
	if !ok {
		return errors.New("unsupported command\n")
	}
	err = ftpCon.WriteCommand(cmdid, args)
	if err != nil {
		return err
	}

	if cmdid == ftp.Commands["get"] {
		err = ftpCon.HandleGet(args[0])
		if err != nil {
			return err
		}
	}

	var length uint32
	err = binary.Read(ftpCon.Con, binary.LittleEndian, &length)
	if err != nil {
		return err
	}
	if length == 0 {
		fmt.Printf("\n%s:", ftpCon.Cwd)
		return nil
	}

	res := make([]byte, length-uint32(binary.Size(length)))
	err = binary.Read(ftpCon.Con, binary.LittleEndian, res)
	if err != nil {
		return err
	}
	if cmdid == ftp.Commands["cd"] {
		// 改变执行目录并换行显示
		ftpCon.Cwd = ftp.Sbyte2str(res)
		fmt.Printf("\n%s:", ftpCon.Cwd)
		return nil
	}
	if cmdid == ftp.Commands["exit"] {
		ftpCon.Exit = true
		fmt.Printf("%s\n", ftp.Sbyte2str(res))
		return nil
	}
	// 其他的命令就直接输出就可以了
	fmt.Printf("%s\n%s:", ftp.Sbyte2str(res), ftpCon.Cwd)
	return
}

func main() {
	// 获取用户身份信息与ftp服务器host信息
	if len(os.Args) < 2 {
		fmt.Println("无法通过身份认证")
		return
	}
	arg := os.Args[1]
	if !strings.Contains(arg, "@") {
		fmt.Println("无法通过身份认证")
		return
	}
	args := strings.Split(arg, "@")
	user := args[0]
	host := args[1]
	fmt.Print("Password:")
	var pwd string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		pwd = input.Text()
	}

	// 连接到ftp服务器
	con, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()
	ftpCon := ftp.FtpConn{
		Con: con,
	}
	ftpClient := client.FtpClient{
		ftpCon,
	}

	// 身份验证
	err = ftpClient.Write(ftp.Str2sbyte(user))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ftpClient.Write(ftp.Str2sbyte(pwd))
	if err != nil {
		fmt.Println(err)
		return
	}

	var res uint32
	err = binary.Read(con, binary.LittleEndian, &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res == 0 {
		fmt.Println("认证失败")
		return
	}

	// 得到服务器返回的家目录
	cwd := make([]byte, res)
	err = binary.Read(con, binary.LittleEndian, cwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	ftpClient.Cwd = ftp.Sbyte2str(cwd)
	ftpClient.Home = ftpCon.Cwd
	fmt.Println(ftpClient.Cwd, ":")

	// 监听命令行输入
	for input.Scan() && !ftpClient.Exit {
		argstr := input.Text()
		args := strings.Split(strings.TrimSpace(argstr), " ")
		if len(args) == 0 {
			printHelp()
			continue
		}
		command := args[0]
		if len(args) > 1 {
			args = args[1:]
		} else {
			args = nil
		}
		err = handleCommand(&ftpClient, command, args)
		if err != nil {
			log.Println(err)
		}
	}

}