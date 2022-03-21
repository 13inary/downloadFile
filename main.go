package main

import (
	"bufio"
	"downloadFile/config"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	// load config
	err := config.InitConfig()
	if err != nil {
		return
	}
	url := config.Conf.Url
	fmt.Println(config.Conf.Title)
	// get day
	var day string
	fmt.Scanln(&day)
	fmt.Printf("time.Now() = %+v\n\n", time.Now())
	// set tasks
	tasks := make(map[string]bool, 0)
	for _, v := range config.Conf.Tasks {
		tasks[v] = false
	}
	// begin
	for i, v := range tasks {
		if v == false {
			dowload(url, day, i, tasks)
		}
	}
	// check result
	for _, v := range tasks {
		if v == false {
			fmt.Println("==== EOORR ====")
			return
		}
	}
	fmt.Println("==== SUCCESS ====")
}

func dowload(url string, day string, name string, tasks map[string]bool) {
	// http get
	pres, err := http.Get(url + "/" + day + "/" + name)
	if err != nil {
		fmt.Println("get ERROR " + day + " " + name)
		panic(err)
	}
	defer pres.Body.Close()
	// check if exist
	if pres.Header["Content-Type"][0] != config.Conf.Type {
		fmt.Println("dowload ERROR " + day + " " + name)
		return
	}
	// init
	err = os.MkdirAll("./dowload", 0777)
	if err != nil {
		return
	}
	os.Remove("./dowload" + name)
	pfile, err := os.Create("./dowload/" + name)
	if err != nil {
		fmt.Println("create ERROR " + day + " " + name)
		panic(err)
	}
	// save
	_, err = io.Copy(pfile, pres.Body)
	if err != nil {
		fmt.Println("copy ERROR " + day + " " + name)
		panic(err)
	}
	fmt.Println("dowload and write SUCCESS " + day + " " + name)
	tasks[name] = true
}

func dowload2(url string, term string, name string) {
	filePath := "./"
	fileName := path.Base(url)
	pres, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer pres.Body.Close()

	reader := bufio.NewReaderSize(pres.Body, 32*1024)

	file, err := os.Create(filePath + fileName)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	writen, _ := io.Copy(writer, reader)
	fmt.Println(writen)
}
