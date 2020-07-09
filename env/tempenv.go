package tempenv
// 同一个目录中不能有来两个不同的package

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func init(){
	fmt.Println("应该在程序预处理阶段执行")
}

func init(){
	fmt.Println("应该在程序预处理阶")
}
// 任何文件可以包含任意数量的init函数，当程序启动的时候按照声明的顺序自动执行
func (c Celsius) String() string    { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gC", f) }

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
