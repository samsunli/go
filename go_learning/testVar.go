package main
import(
	"fmt"
	"unsafe"
)

//匿名变量：下划线"_"，通常我们用匿名接收必须接收，但是又不会用到的值。

//变量声明
func f1(){
	var v1 string = "v1"
	var v2 = "v2"
	v3 := "v3"	//只能用于函数内部
	v4, v5 := "v4", "v5"
	var (
		v6 string = "v6"
		v7 string = "v7"
	)	//一般用于声明全局变量
	fmt.Println(v1, v2, v3, v4, v5, v6, v7)
}

//整型
func f2(){
	var(
		v1 int = 1
		v2 int8 = 1
		v3 int16 = 1
		v4 int32 = 1
		v5 int64 = 1
		v6 int = 0b1010
		v7 int = 0o12
		v8 int = 0xa
	)
	fmt.Printf("%T %d\n", v1, unsafe.Sizeof(v1))
	fmt.Printf("%T %d\n", v2, unsafe.Sizeof(v2))
	fmt.Printf("%T %d\n", v3, unsafe.Sizeof(v3))
	fmt.Printf("%T %d\n", v4, unsafe.Sizeof(v4))
	fmt.Printf("%T %d\n", v5, unsafe.Sizeof(v5))
	fmt.Printf("%d\n", v6)
	fmt.Printf("%d\n", v7)
	fmt.Printf("%d\n", v8)
}

//浮点型
func f3()  {
	//float32精确到后面6位,有效数字为7位 ?
	//float64精确到后面15位,有效数字为16位 ?
	var(
		v1 = 1.5
		v2 float32 = 0.123456789123456789
		v3 float64 = 0.123456789123456789
	)

	fmt.Printf("%.20f %T %d\n", v2, v2, unsafe.Sizeof(v2))
	fmt.Printf("%.20f %T %d\n", v3, v3, unsafe.Sizeof(v3))
	fmt.Printf("%f %T %d\n", v1, v1, unsafe.Sizeof(v1))
}

//byte、rune
func f4() {
	//在GO中，单引号用来表示字符，双引号用来表示字符串
	//byte与unit8本质上无区别，它表示的是 ACSII 表中的一个字符
	//rune与unit32本质上无区别，它表示的是 Unicode 表中的一个字符，故中文字符应该用rune表示

	var(
		v1 byte = 65
		v2 uint8 = 66
		v3 byte = 'A'
		v4 uint = 'B'
	)

	fmt.Printf("%T %d\n", v1, unsafe.Sizeof(v1))

	fmt.Printf("%d %d\n", v1, v2)
	fmt.Printf("%c %c\n", v1, v2)

	fmt.Printf("%d %d\n", v3, v4)
	fmt.Printf("%c %c\n", v3, v4)

	fmt.Println(string(v1), string(v2))		//byte转string，string转byte？
}

//字符串
func f5(){
	//string 的本质，其实是一个 byte数组
	//用反引号定义原生字符，转义符会当做普通字符处理
	//Go 语言的 string 是用 uft-8 进行编码，英文字母占用一个字节，而中文字母占用 3个字节，换行符占用一个字节

	var(
		v1 string = "hello"
		v2 [5]byte = [5]byte{104, 101, 108, 108, 111}
		v3 string = "world\n你好\n"
		v4 string = `world\n你好\n`
	)

	fmt.Printf("%d\n", len(v3))
	fmt.Printf("%d\n", len(v4))
	fmt.Printf("%s\n", v1)
	fmt.Printf("%s\n", v2)
	fmt.Print(v3)
	fmt.Print(v4)
}

func main(){
	f1()
	f2()
	f3()
	f4()
	f5()
}