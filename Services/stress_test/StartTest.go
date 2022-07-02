package stress

import (
	"ginDemoProject/Models"
	rq "ginDemoProject/Utils/request"
	"sync"
	"time"
)

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/7/2 13:47
 * @Desc:   Grace under pressure
 */

func StartTest(concurrency int, count int, request *Models.Request) {
	// 设置接收数据缓存
	ch := make(chan *Models.RequestResults, concurrency*count)
	var (
		wg          sync.WaitGroup // 发送数据完成
		wgReceiving sync.WaitGroup // 数据处理完成
	)
	wgReceiving.Add(1)
	go ReceivingResults(concurrency, ch, &wgReceiving)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go rq.SendHttp(i, request, ch, count, &wg)
	}
	// 等待所有的数据都发送完成
	wg.Wait()
	// 延时1毫秒 确保数据都处理完成了
	time.Sleep(1 * time.Millisecond)
	close(ch)
	// 数据全部处理完成了
	wgReceiving.Wait()
	return
}
