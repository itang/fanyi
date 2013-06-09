package fanyi

import "fmt"
import "net/http"
import "io/ioutil"
import "sync"

const ApiURL = "http://translate.google.cn/translate_a/t?client=t&hl=zh-CN&sl=%s&tl=%s&ie=UTF-8&oe=UTF-8&q=%s"

var initCtx sync.Once
var defaultFanyiServer *FanyiServer

/////////////////////////////////////////////////////////////////
type FanyiError struct {
	Cause error
}

func (this *FanyiError) Error() string {
	return "翻译出错:" + this.Cause.Error()
}

func NewFanyiError(cause error) *FanyiError {
	return &FanyiError{cause}
}

func Fanyi(q string, sl string, tl string) (string, error) {
	return DefaultFanyiServer().Fanyi(q, sl, tl)
}

func DefaultFanyiServer() *FanyiServer {
	initCtx.Do(func() {
		defaultFanyiServer = NewFanyiServer()
	})
	return defaultFanyiServer
}

func NewFanyiServer() *FanyiServer {
	return &FanyiServer{http.DefaultClient}
}

// @Mutable
type FanyiServer struct {
	httpClient *http.Client
}

func (this *FanyiServer) Fanyi(q string, sl string, tl string) (string, error) {
	url := fmt.Sprintf(ApiURL, sl, tl, q)

	resp, err := this.HttpClient().Get(url)
	if err != nil {
		return "", NewFanyiError(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", NewFanyiError(err)
	}

	content := string(body)
	return content, nil
}

func (this *FanyiServer) HttpClient() *http.Client {
	return this.httpClient
}

func (this *FanyiServer) SetHttpClient(httpClient *http.Client) *FanyiServer {
	this.httpClient = httpClient
	return this
}
