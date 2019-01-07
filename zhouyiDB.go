package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/PhoenixTeng/universe1/zhouyi"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nosixtools/solarlunar"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func readDb() []zhouyi.Gua {
	db, err := sql.Open("mysql", "root:123@tcp(35.235.71.36:3306)/zhouyi")
	check(err)

	rows, err := db.Query("select * from zhouyi.zhouyiguaci")
	check(err)

	defer rows.Close()

	var record [64][9]string
	yi := make([]zhouyi.Gua, 64)

	for m := 0; rows.Next(); m++ {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		err = rows.Scan(scanArgs...)

		//存储当前列彖传
		yi[m].Tuan = string(values[1].([]byte))

		//把当前列存到数组
		for i, col := range values {
			if col != nil {
				record[m][i] = string(col.([]byte))
				if a := i - 2; a >= 0 && a < 6 {
					yi[m].Xi[a] = record[m][i]
				}
			}

		}

		//打印出当前列
		//fmt.Println(record[m])
	}
	//打印二维数组，（数据表table)
	//fmt.Println(record)
	return yi

}

func main() {

	ch := make(chan int, 30)
	i := 0
	var yi []zhouyi.Gua

	go func() {
		yi = readDb()
		ch <- 1
	}()

	print("waiting")

	go func() {
		for {
			i = i + 1
			time.Sleep(0.15e9)
			print(".")

		}
	}()

	<-ch
	println()

	m, n, l, f := universe()
	a := zhouyi.ReGuaNu(m, n, l, f, yi)
	writeToFile(f, a, n, yi)
	writeToTerminal(a, n, yi)
	f.Close()
}

func universe() (m, n, l []uint, f *os.File) {
	m = []uint{0, 0, 0, 0, 0, 0}
	n = []uint{0, 0, 0, 0, 0, 0}
	l = []uint{0, 0, 0, 0, 0, 0}
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	day := strconv.Itoa(now.Day())
	if len(day) == 1 {
		day = "0" + day
	}
	date := year + "-" + month + "-" + day
	universeTime := solarlunar.SolarToChineseLuanr(date)
	r := rand.New(rand.NewSource(now.UnixNano()))

	for i := 0; i < 6; i++ {
		a := r.Intn(2) + 2
		b := r.Intn(2) + 2
		c := r.Intn(2) + 2
		m[i] = uint(a + b + c)
		if m[i] == 6 || m[i] == 9 {
			n[i] = m[i]
		}
		if m[i] == 7 || m[i] == 9 {
			l[i] = 1
		}
	}

	//输出文本文件
	f, err := os.OpenFile("./"+date, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file error: %v\n,creating...\n", err)
		f, err = os.Create("./" + date)
	}
	//defer f.Close()

	fmt.Println(m)
	appendToFile(f, fmt.Sprintln(m))
	fmt.Println(n)
	appendToFile(f, fmt.Sprintln(n))
	fmt.Print(l, " ", date, " ", now.Hour(), ":", now.Minute(), "  ", universeTime)
	appendToFile(f, fmt.Sprint(l, " ", date, " ", now.Hour(), ":", now.Minute(), "  ", universeTime))

	return
}
func writeToTerminal(a uint, n []uint, yijing []zhouyi.Gua) {
	fmt.Printf(" 第%d卦\n", a)
	fmt.Println(yijing[a-1].Tuan)
	//动爻 判断
	for i := 0; i < 6; i++ {
		fmt.Println(yijing[a-1].Xi[i])
	}
	fmt.Println("\n<－－动爻－－>")
	for i := 0; i < 6; i++ {
		if n[i] != 0 {
			fmt.Println(yijing[a-1].Xi[i])
		}
	}
}
func writeToFile(f *os.File, a uint, n []uint, yijing []zhouyi.Gua) {

	appendToFile(f, fmt.Sprintf(" 第%d卦\n", a))
	appendToFile(f, fmt.Sprintln(yijing[a-1].Tuan))
	//动爻 判断
	for i := 0; i < 6; i++ {
		appendToFile(f, fmt.Sprintln(yijing[a-1].Xi[i]))
	}
	appendToFile(f, fmt.Sprintln("\n<－－动爻－－>"))
	for i := 0; i < 6; i++ {
		if n[i] != 0 {
			appendToFile(f, fmt.Sprintln(yijing[a-1].Xi[i]))
		}
	}

}

func appendToFile(f *os.File, content string) {
	// 以只写的模式，打开文件

	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	f.WriteAt([]byte(content), n)

}
