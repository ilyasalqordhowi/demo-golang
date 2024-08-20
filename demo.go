package main

import "fmt"

func calc(r1 float32) float32 {
	var phi float32 = 3.14
	// var result int = num1 + num2
	result := phi * (r1 * r1)
	return result
}
func kelilinglingkaran(r1 float32) float32 {
	var phi float32 = 3.14
	// var result int = num1 + num2
	result := 2 * phi * (r1)
	return result
}
func luasDanKeliling(r float32) (float32, float32) {
	var luas float32
	var keliling float32
	if int(r)%7 == 0 {
		phi := 22 / 7
		luas = float32(phi * int(r) * int(r) / 7)
		keliling = float32(2 * phi * int(r) / 7)
	} else {
		luas = 3.14 * r * r
		keliling = 2 * 3.14 * r
	}
	return luas, keliling
}

func main() {
	var r float32 = 4
	// var  num  float32 =  calc(4)
	// var num2  float32 = keliling(4)

	var luas, keliling = luasDanKeliling((r))
	fmt.Printf("hello %s, your age is %d\n\n", "world", 12)
	fmt.Printf("hitung sebuah lingkaran dengan jari-jari %f\n", r)
	fmt.Printf(" hasil dari luas lingkaran adalah %f\n", luas)
	fmt.Printf(" keliling lingkaran adalah %f", keliling)
	// fmt.Printf("%t",)

	// var num int = 123
	// var name string = "hello guys"
	// var isMaried bool = false
	// result := "hello"

	// fmt.Printf("hello world %d\n\n\n%s is im maried \n\n%t\n\n%s",num,name,isMaried,result)
}
