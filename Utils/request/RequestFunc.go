package rq

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/2/8 17:00
 * @Desc:   Grace under pressure
 */

import (
	"fmt"
	"ginDemoProject/Models"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func HttpRequest(request *Models.Request) (float64, string, int) {
	method := request.Method
	url := request.URL
	body := request.GetBody()
	headers := request.Headers

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return 0, "", 0
	}
	client := &http.Client{Timeout: 5 * time.Second}
	for key, header := range headers {
		req.Header.Set(key, header)
	}

	begin := time.Now()
	resp, err := client.Do(req)
	respTime := time.Since(begin).Milliseconds()

	//defer resp.Body.Close()

	resCode := resp.StatusCode
	if resCode != 200 {
		fmt.Println("请求失败:", err)
		return 0, "", 0
	}
	data, err := ioutil.ReadAll(resp.Body)

	return float64(respTime), string(data), resCode
}

func SendHttp(id int, request *Models.Request, ch chan *Models.RequestResults, count int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	var isSucceed bool
	for i := 0; i < count; i++ {
		rt, data, code := HttpRequest(request)
		if code == 200 {
			isSucceed = true
		} else {
			isSucceed = false
		}
		requestResults := &Models.RequestResults{
			ID:       id,
			RT:       rt,
			Succeed:  isSucceed,
			RespData: data,
		}
		ch <- requestResults
	}
}
