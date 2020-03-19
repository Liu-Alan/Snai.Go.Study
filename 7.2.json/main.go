package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ServerSettings struct {
	Servers []server `json:"servers"`
}

type server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

/*
//读取
func main() {
	file, err := os.Open("servers.json")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	serversSettings := ServerSettings{}
	err = json.Unmarshal(data, &serversSettings)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	fmt.Println(serversSettings)
}
*/
/*
//输出
func main() {
	serversSettings := ServerSettings{}
	serversSettings.Servers = append(serversSettings.Servers, server{ServerName: "Shanghai_VPN", ServerIP: "192.168.0.2"})
	serversSettings.Servers = append(serversSettings.Servers, server{ServerName: "Beiging_VPN", ServerIP: "192.168.0.2"})

	data, err := json.MarshalIndent(&serversSettings, "  ", "    ")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	os.Stdout.Write(data)
}
*/
/*
//写入
func main() {
	file, err := os.OpenFile("servers.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	defer file.Close()

	serversSettings := ServerSettings{}
	serversSettings.Servers = append(serversSettings.Servers, server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	serversSettings.Servers = append(serversSettings.Servers, server{ServerName: "Beiging_VPN", ServerIP: "127.0.0.1"})

	data, err := json.MarshalIndent(&serversSettings, "  ", "    ")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
}
*/
//读取并写入
func main() {
	filename := "servers.json"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	serversSettings := ServerSettings{}
	err = json.Unmarshal(data, &serversSettings)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	for i, _ := range serversSettings.Servers {
		serversSettings.Servers[i].ServerIP = "192.168.0.1"
	}

	data, err = json.MarshalIndent(&serversSettings, "  ", "    ")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
}
