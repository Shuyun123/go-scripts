package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"scripts/utils"
	"strings"
)

func main() {
	var dest = flag.String("o", GetCurrentPath(), "download dir")
	var urls = flag.String("l", "", "download urls, f or l params must input")
	var file = flag.String("f", "", "download urls file path, f or l params must input")
	var concurrent = flag.Int("n", 3, "concurrent number")
	var help = flag.String("h", "", "help")
	flag.Parse()
	if *help == "h" || *help == "help" {
		flag.PrintDefaults()
		return
	} else {
		if *urls == "" && *file == "" {
			flag.PrintDefaults()
			return
		}
	}

	var data = strings.Split(*urls, ",")
	if *file != "" {
		data = readUrls(*file)
	}
	_, err := os.Stat(*dest)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(*dest, os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	downloader := utils.NewDownloader(*dest)
	for _, v := range data {
		if v != "" {
			requestUrl, _ := url.QueryUnescape(v)
			name := requestUrl[strings.LastIndex(requestUrl, "/")+1:]
			str := strings.Replace(name, "\n", "", -1)
			downloader.AppendResource(str, strings.Replace(v, "\n", "", -1))
		}
	}
	downloader.Concurrent = *concurrent
	_ = downloader.Start()
}

func readUrls(path string) (result []string) {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Print(inputError)
		return // exit the function on error
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		result = append(result, inputString)
		if readerError == io.EOF {
			return
		}
	}
}

//get current path
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
