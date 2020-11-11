package main

import (
	"errors"
	"fmt"
	"log"
	config "mymod/common"
	"os"
)

func GetConfig( section,url string) string  {
	Url:= config.GetConfig(section,url)
	return Url

}

func division( x,y float32) (float32,error)  {
	if y == 0{
		return 0,errors.New("can't division by zero")
	}

	return x/y,nil
}
//
//func main()  {
//	var x float32 = 16
//	var y float32 = 0
//
//	res,err:= division(x,y)
//	if err!=nil{
//		log.Print(err)
//	}
//
//	fmt.Println(res)
//}
func main() {
	var x float32 = 11
	var y float32

	res, exception := division(x, y)
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(file)

	if err != nil {
		log.Fatal(err)
	}

	if exception != nil {
		log.Print(exception)
	}

	defer file.Close()

	fmt.Println(res)
}