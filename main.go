package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/v2fly/domain-list-community/release/geolocation-!cn.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	f, err := os.Create("autoproxy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("[AutoProxy 0.2.9]\n")
	w.WriteString("! Last Modified: $time\n")
	w.WriteString("! Expires: 24h\n")
	//w.WriteString("! HomePage: https://github.com/sunshineplan/autoproxy\n")
	//w.WriteString("! GitHub URL: https://raw.githubusercontent.com/sunshineplan/autoproxy/release/autoproxy.txt\n")
	//w.WriteString("! jsdelivr URL: https://cdn.jsdelivr.net/gh/sunshineplan/autoproxy@release/autoproxy.txt\n")
	w.WriteString("\n")
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		i := scanner.Text()
		if strings.HasSuffix(i, "@cn") {
			continue
		}
		i = strings.ReplaceAll(i, ":@ads", "")
		switch {
		case strings.HasPrefix(i, "domain:"):
			w.WriteString(strings.Replace(i, "domain:", "||", 1) + "\n")
		case strings.HasPrefix(i, "full:"):
			w.WriteString(strings.Replace(i, "full:", "|http://", 1) + "\n")
			w.WriteString(strings.Replace(i, "full:", "|https://", 1) + "\n")
		case strings.HasPrefix(i, "keyword:"):
			w.WriteString(strings.Replace(i, "keyword:", "", 1) + "\n")
		case strings.HasPrefix(i, "regexp:"):
			w.WriteString(strings.Replace(i, "regexp:", "/", 1) + "/" + "\n")
		default:
			log.Println("unknow format:", i)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	w.Flush()
}
