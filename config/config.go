package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type configuration struct {
	// 配置信息
	ServerInfo serverInfo
	RedisInfo  redisInfo
}

type serverInfo struct {
	//服务器地址
	Host string
}

type redisInfo struct {
	// redis相关配置
	Host        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var Configuration = configuration{}

func init() {
	filePath := "./config/config.json"
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("Open file error: %v\n", err)
	}

	decoder := json.NewDecoder(file)     //配置文件是json
	err = decoder.Decode(&Configuration) //json-》struct
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
