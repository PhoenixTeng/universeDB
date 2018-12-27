package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type text struct {
	tuan string
	xi   [6]string
}

var yijing [64]text

//var yijing []text = make([]text.64)
//var yijing  = make([]text, 64)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func readDb() {
	db, err := sql.Open("mysql", "root:123@tcp(35.235.71.36:3306)/zhouyi")
	check(err)

	rows, err := db.Query("select * from zhouyi.zhouyiguaci")
	check(err)

	defer rows.Close()

	var record [64][9]string
	var m = 0
	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		err = rows.Scan(scanArgs...)

		//存储当前列彖传
		yijing[m].tuan = string(values[1].([]byte))

		//把当前列存到数组
		for i, col := range values {
			if col != nil {
				record[m][i] = string(col.([]byte))
				if a := i - 2; a >= 0 && a < 6 {
					yijing[m].xi[a] = record[m][i]
				}
			}

		}

		m = m + 1
		//打印出当前列
		//fmt.Println(record[m])
	}
	//打印二维数组，（数据表table)
	//fmt.Println(record)

}

func work() {
	var a = 0
	var m = []int{0, 0, 0, 0, 0, 0}
	var n = []int{0, 0, 0, 0, 0, 0}
	var l = []int{0, 0, 0, 0, 0, 0}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 6; i++ {
		a := r.Intn(2) + 2
		b := r.Intn(2) + 2
		c := r.Intn(2) + 2
		m[i] = a + b + c
		if m[i] == 6 || m[i] == 9 {
			n[i] = m[i]
		}
		if m[i] == 7 || m[i] == 9 {
			l[i] = 1
		}
	}
	fmt.Println(m)
	fmt.Println(n)
	fmt.Print(l)

	switch {
	//乾卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 1
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[0].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[0].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[0].xi[i])
			}
		}

	//坤卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 2
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[1].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[1].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[1].xi[i])
			}
		}

	//屯卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 3
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[2].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[2].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[2].xi[i])
			}
		}

	//蒙卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 4
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[3].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[3].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[3].xi[i])
			}
		}

	//需卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 5
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[4].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[4].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[4].xi[i])
			}
		}

	//讼卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 6
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[5].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[5].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[5].xi[i])
			}
		}

	//师卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 7
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[6].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[6].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[6].xi[i])
			}
		}

	//比卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 8
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[7].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[7].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[7].xi[i])
			}
		}

	//小畜卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 9
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[8].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[8].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[8].xi[i])
			}
		}

		//履卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 10
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[9].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[9].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[9].xi[i])
			}
		}

		//泰卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 11
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[10].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[10].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[10].xi[i])
			}
		}

		//否卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 12
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[11].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[11].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[11].xi[i])
			}
		}

		//同人卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 13
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[12].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[12].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[12].xi[i])
			}
		}

		//大有卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 14
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[13].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[13].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[13].xi[i])
			}
		}

		//谦卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 15
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[14].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[14].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[14].xi[i])
			}
		}

		//豫卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 16
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[15].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[15].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[15].xi[i])
			}
		}

	//随卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 17
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[16].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[16].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[16].xi[i])
			}
		}

		//蛊卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 18
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[17].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[17].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[17].xi[i])
			}
		}

		//临卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 19
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[18].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[18].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[18].xi[i])
			}
		}

		//观卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 20
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[19].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[19].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[19].xi[i])
			}
		}

		//噬嗑卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 21
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[20].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[20].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[20].xi[i])
			}
		}

		//贲卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 22
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[21].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[21].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[21].xi[i])
			}
		}

		//剥卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 23
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[22].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[22].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[22].xi[i])
			}
		}

		//复卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 24
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[23].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[23].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[23].xi[i])
			}
		}

		//无妄卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 25
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[24].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[24].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[24].xi[i])
			}
		}

		//大畜卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 26
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[25].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[25].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[25].xi[i])
			}
		}

		//颐卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 27
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[26].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[26].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[26].xi[i])
			}
		}

		//大过卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 28
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[27].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[27].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[27].xi[i])
			}
		}

		//坎卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 29
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[28].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[28].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[28].xi[i])
			}
		}

		//离卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 30
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[29].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[29].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[29].xi[i])
			}
		}

		// .....下经......
		//咸卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 31
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[30].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[30].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[30].xi[i])
			}
		}

		//恒卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 32
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[31].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[31].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[31].xi[i])
			}
		}

		//遁卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 33
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[32].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[32].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[32].xi[i])
			}
		}
		//大壮卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 34
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[33].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[33].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[33].xi[i])
			}
		}

		//晋卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 35
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[34].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[34].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[34].xi[i])
			}
		}

		//明夷卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 36
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[35].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[35].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[35].xi[i])
			}
		}

		//家人卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 37
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[36].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[36].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[36].xi[i])
			}
		}

		//睽卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 38
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[37].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[37].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[37].xi[i])
			}
		}

		//蹇卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 39
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[38].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[38].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[38].xi[i])
			}
		}

		//解卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 40
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[39].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[39].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[39].xi[i])
			}
		}

		//损卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 41
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[40].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[40].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[40].xi[i])
			}
		}

		//益卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 42
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[41].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[41].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[41].xi[i])
			}
		}

		//夬卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 43
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[42].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[42].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[42].xi[i])
			}
		}

		//姤卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 44
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[43].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[43].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[43].xi[i])
			}
		}

		//萃卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 45
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[44].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[44].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[44].xi[i])
			}
		}

		//升卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 46
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[45].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[45].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[45].xi[i])
			}
		}

		//困卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 47
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[46].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[46].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[46].xi[i])
			}
		}

		//井卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 48
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[47].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[47].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[47].xi[i])
			}
		}

		//革卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 49
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[48].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[48].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[48].xi[i])
			}
		}

		//鼎卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 50
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[49].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[49].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[49].xi[i])
			}
		}

		//震卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 51
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[50].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[50].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[50].xi[i])
			}
		}

		//艮卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 52
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[51].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[51].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[51].xi[i])
			}
		}

		//渐卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 53
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[52].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[52].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[52].xi[i])
			}
		}

		//归妹卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 54
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[53].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[53].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[53].xi[i])
			}
		}

		//丰卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 55
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[54].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[54].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[54].xi[i])
			}
		}

		//旅卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 56
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[55].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[55].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[55].xi[i])
			}
		}

		//巽卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 57
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[56].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[56].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[56].xi[i])
			}
		}

		//兑卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 58
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[57].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[57].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[57].xi[i])
			}
		}

		//涣卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 59
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[58].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[58].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[58].xi[i])
			}
		}

		//节卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 60
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[59].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[59].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[59].xi[i])
			}
		}

		//中孚卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 61
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[60].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[60].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[60].xi[i])
			}
		}

		//小过卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 62
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[61].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[61].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[61].xi[i])
			}
		}

		//既济卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 63
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[62].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[62].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[62].xi[i])
			}
		}

		//未济卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 64
		fmt.Printf(" 第%d卦\n", a)
		fmt.Println(yijing[63].tuan)

		//动爻 判断
		for i := 0; i < 6; i++ {
			fmt.Println(yijing[63].xi[i])
		}
		fmt.Println("\n<－－动爻－－>")
		for i := 0; i < 6; i++ {
			if n[i] != 0 {
				fmt.Println(yijing[63].xi[i])
			}
		}

	}
}

func main() {

	ch := make(chan int, 30)
	i := 0

	go func() {
		readDb()
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

	if flg() == -1 {
		work()
	}

}

func flg() int {
	//fptr := flag.String("fpath", "test.txt", "file path to read from")
	fprt := flag.String("l", "-1", "第几卦")
	flag.Parse()

	a, error := strconv.Atoi(*fprt)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
		panic(error)
	}

	if a == -1 {
		return -1
	}

	fmt.Printf(" 第%d卦\n", a)
	fmt.Println(yijing[a-1].tuan)

	//动爻 判断
	for i := 0; i < 6; i++ {
		fmt.Println(yijing[a].xi[i])
	}
	return 0
}
