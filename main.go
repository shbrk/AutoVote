package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"math"
	"os/exec"
)

func main() {
	var index = 0
	var t = time.Now()
	for {
		time.Sleep(time.Second)
		if time.Since(t) < time.Second * 60 {
			continue
		}
		t = time.Now()
		index += 1
		var ip = GetIp()
		var proxy = "http://" + ip
		fmt.Println(strconv.Itoa(index) + ":" + ip)
		postName, postValue, phpSession := GetPostParam("http://xxx", proxy)
		if postName == "" || postValue == "" || phpSession == "" {
			continue
		}
		session := GetSession("http://xxx", proxy, phpSession)
		if session == "" {
			continue
		}
		for i := 0; i < 3; i++ {
			ret := Vote("http://wxxx", proxy, postName, postValue, phpSession, session)
			time.Sleep( time.Second * time.Duration(rand.Intn(30)))
			if ret == nil {
				break
			}
		}
	}
}

func ParsePostParam(content string) (string, string) {
	if content == "" {
		return "", ""
	}
	beginIndexName := strings.Index(content, `value="`)
	endIndexName := strings.LastIndex(content, `" id="postname"`)
	beginIndexValue := strings.LastIndex(content, `value="`)
	endIndexValue := strings.LastIndex(content, `" id="postvalue"`)
	return content[beginIndexName+7 : endIndexName], content[beginIndexValue+7 : endIndexValue]
}

func GetPostParam(urls string, ip string) (postName string, postValue string, phpSession string) {
	request, _ := http.NewRequest("GET", urls, nil)
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	proxy, err := url.Parse(ip)
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		fmt.Printf("遇到了错误%s", err)
		return
	}
	if response != nil {
		defer response.Body.Close()
		var doc, _ = goquery.NewDocumentFromReader(response.Body)
		if doc == nil {
			return
		}
		pc := doc.Find("#postcontent")
		if pc == nil {
			return
		}
		cmd := exec.Command("node.exe", "E:/Project/js/app2.js", pc.Text())
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		postName, postValue = ParsePostParam(out.String())
		cookies := response.Cookies()
		for _, value := range cookies {
			if value.Name == "PHPSESSID" {
				phpSession = value.Value
				break
			}
		}
	}
	return
}

func GetSession(urls string, ip string, phpSession string) (session string) {
	request, _ := http.NewRequest("GET", urls, nil)
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Cookie", "PHPSESSID="+phpSession)
	proxy, err := url.Parse(ip)
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		fmt.Printf("line-99:遇到了错误")
		return
	}
	if response != nil {
		defer response.Body.Close()
		cookies := response.Cookies()
		for _, value := range cookies {
			if value.Name == "SESSION" {
				session = value.Value
			}
		}
	}
	return
}

func GetIp() string {
	res, err := http.Get("http://api.ip.data5u.com/socks/get.html?xxxx")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	_, err = url.Parse("http://"+string(data[:len(data)-1]))
	if err != nil {
		return ""
	}
	return string(string(data[:len(data)-1]))
}

func Vote(urls string, ip string, postName string, postValue string, phpSession string, session string) *http.Response {
	//var randStr = strconv.FormatFloat(float64(float32(rand.Int31())/float32(math.MaxInt32)), 'f', 16, 64)
	value := url.Values{

	}
	request, _ := http.NewRequest("POST", urls, strings.NewReader(value.Encode()))
	request.Header.Set("User-Agent", getAgent())
	request.Header.Set("Accept-Encoding", "Accept-Encoding")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	request.Header.Set("Accept", "text/html, */*; q=0.01")
	//	request.Header.Set("Content-Length", strconv.Itoa(len(data)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Cookie", "PHPSESSID="+phpSession+"; SESSION="+session)
	request.Header.Set("X-Requested-With", "XMLHttpRequest")

	request.Header.Set("Proxy-Connection", "keep-alive")

	proxy, err := url.Parse(ip)
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		fmt.Printf("line-99:遇到了错误")
		return nil
	}
	if response != nil {
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		data := string(body[3:])
		if data == "1" {
			fmt.Println("投票成功！")
		} else if data == "2" {
			fmt.Println("你已经投过了,请勿重复投票！")
		} else {
			fmt.Println("投票失败,错误码:" + data)
		}
	}
	return response
}

/**
* 随机返回一个User-Agent
 */
func getAgent() string {
	agent := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len1 := len(agent)
	return agent[r.Intn(len1)]
}
