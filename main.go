package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	server  = ""
	tips    = ""
	cfg     config
	day     string
	failure []string
)

// config
type config struct {
	Url   string   `json:"url"`
	Files []string `json:"files"`
}

// init
func init() {
	failure = make([]string, 0)

	rsp, err := http.Get(server)
	if err != nil {
		log.Println("get config error")
		panic(err)
	}
	defer rsp.Body.Close()

	if rsp.Header["Content-Type"][0] != "application/json; charset=utf-8" {
		log.Println("get config error")
	}

	if err := json.NewDecoder(rsp.Body).Decode(&cfg); err != nil {
		log.Println("jsom unmarshal error")
	}
	fmt.Println(cfg)
	fmt.Println("")
}

func main() {
	fmt.Println(tips)
	fmt.Scanln(&day)
	fmt.Printf("time.Now() = %+v\n\n", time.Now())

	for _, v := range cfg.Files {
		url := cfg.Url + "/" + day + "/" + v
		dowload(url, v)
	}

	for _, v := range failure {
		fmt.Printf("=== ERROR : %s ===", v)
	}
	fmt.Println("==== SUCCESS ====")
}

func dowload(url string, file string) {
	pres, err := http.Get(url)
	if err != nil {
		failure = append(failure, file)
		fmt.Printf("get %s ERROR\n", url)
		return
	}
	defer pres.Body.Close()

	if pres.Header["Content-Type"][0] != "image/jpeg" {
		failure = append(failure, file)
		fmt.Printf("dowload %s ERROR\n", file)
		return
	}

	err = os.MkdirAll("./dowload", 0777)
	if err != nil {
		failure = append(failure, file)
		fmt.Println("mkdir error ")
		return
	}
	os.Remove("./dowload" + file)
	pfile, err := os.Create("./dowload/" + file)
	if err != nil {
		failure = append(failure, file)
		fmt.Printf("create %s ERROR\n", file)
		return
	}

	_, err = io.Copy(pfile, pres.Body)
	if err != nil {
		failure = append(failure, file)
		fmt.Printf("copy %s ERROR\n", file)
		return
	}

	fmt.Printf("dowload and write %s SUCCESS\n", file)
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
