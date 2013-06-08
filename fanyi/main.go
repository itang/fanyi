package main

import "fmt"
import "net/http"
import "net/url"
import "flag"
import "github.com/itang/fanyi"

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
	fanyi.CheckError(err)

	fmt.Println(result)
}

func fanyiServer() *fanyi.FanyiServer {
	fanyiServer := &fanyi.FanyiServer{}

	fmt.Println("httpProxyUrl:" + httpProxyUrl)

	if httpProxyUrl != "" {
		proxyUrl, err := url.Parse("http://" + httpProxyUrl)
		fanyi.CheckError(err)
		fmt.Println("Use proxy!")
		fanyiServer.SetHttpClient(&http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}})
	}

	return fanyiServer
}