package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/itang/fanyi"
	"github.com/itang/gotang"
)

var httpProxyUrl string = ""

func init() {
	flag.StringVar(&httpProxyUrl, "proxy", httpProxyUrl, "http proxy url")
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("请输入要翻译的词")
		return
	}

	var q, sl, tl string = args[0], "auto", "zh-CN"
	switch len(args) {
	case 2:
		tl = args[1]
	case 3:
		sl = args[1]
		tl = args[2]
	}
	result, err := fanyiServer().Fanyi(q, sl, tl)
	gotang.AssertNoError(err, "")

	prettyOutput(result)
}

func fanyiServer() *fanyi.FanyiServer {
	fanyiServer := fanyi.DefaultFanyiServer()
	if httpProxyUrl != "" {
		proxyUrl, err := url.Parse("http://" + httpProxyUrl)
		gotang.AssertNoError(err, "")
		fanyiServer.SetHttpClient(
			&http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}},
		)
	}
	return fanyiServer
}

//TODO 优化
func prettyOutput(result string) {
	rets := strings.Split(result, "]],")
	for i, v := range rets {
		fmt.Printf("%d: %s\n", i, v)
	}
}
