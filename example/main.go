package main

import (
	"fmt"
	"github.com/DLChenMR/goconfig"
	"os"
)

type Properties struct {
	AppName   string `prop:"APP_NAME"`
	AppPort   int    `prop:"APP_PORT"`
	TimeStamp int64  `prop:"TIMESTAMP"`
	MySQL     struct {
		Port int    `prop:"MYSQL_PORT"`
		HOST string `prop:"MYSQL_HOST"`
	}
	IPs      []string  `prop:"IPS" separator:","`
	Weights  []float32 `prop:"WEIGHTS" separator:";"`
	Booleans []bool    `prop:"BOOLEANS"` //default: separator:","
}

func main() {
	properties := &Properties{}

	os.Setenv("MYSQL_PORT", "3307")
	os.Setenv("BOOLEANS", "true,true,true")

	err := goconfig.Init("example/app.ini", properties)
	if err != nil {
		fmt.Println("[ERROR]", err)
	} else {
		fmt.Println("AppName:", properties.AppName)
		fmt.Println("AppPort:", properties.AppPort)
		fmt.Println("TimeStamp:", properties.TimeStamp)
		fmt.Println("MySQL_PORT:", properties.MySQL.Port)
		fmt.Println("MySQL_HOST:", properties.MySQL.HOST)
		fmt.Println("IPs:", properties.IPs)
		fmt.Println("Weights:", properties.Weights)
		fmt.Println("Booleans:", properties.Booleans)
	}

}
