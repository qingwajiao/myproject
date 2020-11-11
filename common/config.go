package config

import (
	"fmt"
	"github.com/go-ini/ini"
	)

func GetConfig( section string ,key string)string  {
	cfg,err := ini.Load("/server/config.ini")
	if err != nil{
		fmt.Println("加载配置文件异常")

	}
	sec,err :=cfg.GetSection(section)
	if err != nil{
		fmt.Println("获取配置文件异常")
	}
	val := sec.Key(key)

	str:= val.String()
	//fmt.Println(str)
	return str
}
