package common

import (
	"strings"
	"log"
	"io/ioutil"
	"net/http"
	"io"
	"time"
)

func main2() {
	b, err := ioutil.ReadFile("resources/icons.html")
	var urls []string
	if err == nil {
		content := string(b)
		//log.Print(content)
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			line = strings.Trim(line, "")
			if strings.Contains(line, "data-src=") {
				//log.Print(line)
				start := strings.Index(line, "data-src=\"")
				tmp := len("data-src=\"")
				end := strings.Index(line[start+tmp:], "\"")
				urlStr := line[start+tmp:start+end]
				//log.Println(urlStr)
				urls = append(urls, urlStr)
			}
		}
	}
	var client = &http.Client{
		//Timeout: clientTimeout,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	//
	for _, u := range urls {
		log.Println(u)

		last := strings.LastIndex(u, ".png")
		first := strings.LastIndex(u[0:last], "/")
		name := u[first+1:last] + ".png"
		log.Println(name)

		req, _ := http.NewRequest("GET", u, nil)

		var response *http.Response
		if response, err = client.Do(req); response != nil && response.Body != nil && err == nil {
			var bytes []byte
			defer func() { //把body读完，然后关闭，不然会内存泄露
				io.Copy(ioutil.Discard, response.Body)
				response.Body.Close()
			}()
			bytes, _ = ioutil.ReadAll(response.Body)
			ioutil.WriteFile("icons/"+name, bytes, 0644)
			log.Println("run " + name + " over~")
		}

	}

}
