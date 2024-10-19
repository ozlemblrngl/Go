package main

import "fmt"

func main(){
	fmt.Println("Hello")
	fmt.Println("World!")

	var text string = "Hello World!"
	fmt.Println(text)

	var number int = 30
	fmt.Println(number)
	fmt.Println(100+number/10)

	var number2 float32 = 0.15
	fmt.Println(number2)

	text2 := "Mustafa"
	fmt.Println(text2)
	fmt.Printf("text2 type is : %T", text2)

}

// go büyük küçük harf duyarlıdır ve kullanmadığımız bir import olursa çalışmaz hata verir.
// string ifadeler çift tırnakla yazılır, tek tırnağı kabul etmez.
// := yazdığımızda tür yazmadan direkt değişken oluşturup değer atayabiliyoruz. Go'ya özel bir işlemdir.
// Printf => print format tır. Println yazarsak text2 nin değerini ya<dırır Printf yazarsak tipini yani formatını yazdırır.