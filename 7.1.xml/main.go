package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//读取
/*
type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
	//Description string   `xml:",innerxml"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}

	v := Servers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}

	fmt.Println(v)
}
*/
//输出
/*
type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}
type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:serverIP`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beiging_VPN", "127.0.0.1"})

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
*/
//写入
/*
type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}
type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:serverIP`
}

func main() {
	file, err := os.OpenFile("servers.xml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	defer file.Close()

	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.2"})
	v.Svs = append(v.Svs, server{"Beiging_VPN", "127.0.0.2"})

	data, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	_, err = file.Write(append([]byte(xml.Header), data...))
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
}
*/

//读取并写入
type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	filename := "servers.xml"
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

	v := Servers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	for i, _ := range v.Svs {
		v.Svs[i].ServerIP = "192.168.0.1"
	}

	data, err = xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	err = ioutil.WriteFile(filename, append([]byte(xml.Header), data...), 0666)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
}
