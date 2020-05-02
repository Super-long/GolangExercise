package main

import (
	"fmt"
	"net"
)

/*重要的是不要舍弃错误 函数既然返回了就要去处理
Go没有提供try catch的异常处理方式,而是通过函数的返回值向上抛
这鼓励工程师在代码中显式的检查错误,而非忽略错误,这会避免漏掉本该处理的错误
但是会让代码啰嗦
 */
func main() {
	/*f,err := os.Open("hello.go")
	defer f.Close()
	if err != nil{
		//log.Fatal(err)
		if ins,ok := err.(*os.PathError); ok{
			fmt.Println(ins.Err)
			fmt.Println(ins.Path)
			fmt.Println(ins.Op)
		}
		return
	}
	fmt.Println(f.Name(),"文件打开成功")*/
	//----------------------------

	addrs, err1 := net.LookupHost("www.baidu.com")
	if ins, ok := err1.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("operator timeout.\n")
		} else if ins.Temporary() {
			fmt.Println("临时性错误")
		} else {
			fmt.Println("一般错误")
		}
	}
	fmt.Println(addrs)

	/*err1 := errors.New("创建着玩的")
	fmt.Println(err1)
	fmt.Println(err1.Error()) //两个是一样的

	err2 := fmt.Errorf("错误的信息码:%d", 100)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	//--------------------
	err3 := checkAge(-1)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println("程序没问题")*/
}

func checkAge(age int) error {
	if age < 0 {
		// return errors.New("年龄不合法")
		return fmt.Errorf("给定的年龄不合法: %d", age)
	}
	fmt.Println("年龄是:%d", age)
	return nil
}
