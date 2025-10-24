package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sunshineplan/utils/txt"
)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/v2fly/domain-list-community/release/geolocation-!cn.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	res, err := txt.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("autoproxy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := txt.NewWriter(f)
	w.WriteLine("[AutoProxy 0.2.9]")
	w.WriteLine("! Last Modified: $time")
	w.WriteLine("! Expires: 24h")
	//w.WriteLine("! HomePage: https://github.com/sunshineplan/autoproxy")
	//w.WriteLine("! GitHub URL: https://raw.githubusercontent.com/sunshineplan/autoproxy/release/autoproxy.txt")
	//w.WriteLine("! jsdelivr URL: https://cdn.jsdelivr.net/gh/sunshineplan/autoproxy@release/autoproxy.txt")
	w.WriteLine("")
	for _, i := range res {
		if strings.HasSuffix(i, "@cn") {
			continue
		}
		i = strings.ReplaceAll(i, ":@ads", "")
		switch {
		case strings.HasPrefix(i, "domain:"):
			w.WriteLine(strings.Replace(i, "domain:", "||", 1))
		case strings.HasPrefix(i, "full:"):
			w.WriteLine(strings.Replace(i, "full:", "|http://", 1))
			w.WriteLine(strings.Replace(i, "full:", "|https://", 1))
		case strings.HasPrefix(i, "keyword:"):
			w.WriteLine(strings.Replace(i, "keyword:", "", 1))
		case strings.HasPrefix(i, "regexp:"):
			w.WriteLine(strings.Replace(i, "regexp:", "/", 1) + "/")
		default:
			log.Println("unknow format:", i)
		}
	}
	w.Flush()
}
