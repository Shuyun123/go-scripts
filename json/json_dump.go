package main

import (
	_ "bytes"
	"flag"
	"fmt"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var dest = flag.String("o", "", "json output file name")
	var compress = flag.Bool("c", false, "compress json format")
	var indent = flag.String("i", "\t", "indent string, default tab")
	var jsonData = flag.String("d", "", "json string data")
	var file = flag.String("f", "", "need to format json file")
	var help = flag.String("h", "", "help")
	flag.Parse()
	if *help == "h" || *help == "help" {
		flag.PrintDefaults()
		return
	} else {
		if *jsonData == "" && *file == "" {
			flag.PrintDefaults()
			return
		}
	}

	var result []byte
	var Options = &pretty.Options{Width: 80, Prefix: "", Indent: *indent, SortKeys: false}
	if *compress {
		result = pretty.Ugly([]byte(*jsonData))
		fmt.Printf("%s\n", result)
	} else {
		//result = pretty.Color(pretty.PrettyOptions([]byte(*jsonData), Options), pretty.TerminalStyle)
		result = pretty.PrettyOptions([]byte(*jsonData), Options)
		fmt.Printf("%s\n", result)
	}

	if *file != "" {
		if !FileExist(GetCurrentPath() + "/" + *file) {
			fmt.Println("not found " + GetCurrentPath() + "/" + *file)
			return
		}
		json := readFileContent(GetCurrentPath() + "/" + *file)
		if *compress {
			result = pretty.Ugly([]byte(json))
		} else {
			result = pretty.PrettyOptions([]byte(json), Options)
		}
		writeData(GetCurrentPath()+"/"+*file, string(result))
		fmt.Println("it's done!")
	}

	if *dest != "" {
		writeData(*dest, string(result))
		fmt.Println("it's done!")
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

func writeData(path string, data string) {
	fd, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	fdContent := strings.Join([]string{data, "\n"}, "")
	buf := []byte(fdContent)
	_, _ = fd.Write(buf)
	fd.Close()
}

func readFileContent(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
