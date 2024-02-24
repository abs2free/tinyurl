package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/abs2free/tinyurl/base62"
	"github.com/abs2free/tinyurl/store"
	"github.com/abs2free/tinyurl/store/model"
)

// 测试插入1亿数据
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	mysql, err := store.NewMysql(dsn)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urlsChan := make(chan []model.TinyURL, 20000)
	// 监测
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		var mem runtime.MemStats

		for {
			select {
			case <-ticker.C:
				fmt.Printf("urlsChan的 数量: %d \n", len(urlsChan))
				fmt.Printf("goroutine 数量: %d \n", runtime.NumGoroutine())
				runtime.ReadMemStats(&mem)
				fmt.Printf("TotalAlloc = %v MB", mem.Alloc/1024/1024/8)
			}
		}
	}()

	var wg sync.WaitGroup
	// 定义workder
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for urls := range urlsChan {
				_, err := mysql.BatchAddTinyURL(ctx, urls)
				if err != nil {
					cancel()
					log.Fatal(err)
				}
			}
		}()
	}

	// 性能分析
	go func() {
		fmt.Println("pprof start...")
		fmt.Println(http.ListenAndServe(":9876", nil))
	}()

	c := make(chan os.Signal, 1)
	// Trigger graceful shutdown on SIGINT or SIGTERM.
	// The default signal sent by the `kill` command is SIGTERM,
	// which is taken as the graceful shutdown signal for many systems, eg., Kubernetes, Gunicorn.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("%s received.\n", sig.String())
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var urls []model.TinyURL
		for i := 1; i < 1000000000; i++ {
			url := model.TinyURL{
				LongURL:   fmt.Sprintf("this is a very very long url, please remember it ,the number is:%d", i),
				ShortURL:  base62.Encode(i),
				ExpiredAt: time.Now().AddDate(1, 2, 2),
				CreatedAt: time.Now(),
			}
			urls = append(urls, url)

			if i%100 == 0 {
				// fmt.Printf("第%d个加入到chan\n", i)
				urlsChan <- urls
				urls = []model.TinyURL{}

				select {
				case <-ctx.Done():
					close(urlsChan)
					return
				default:
				}
			}
		}
		urlsChan <- urls
		close(urlsChan)
	}()

	wg.Wait()
	fmt.Println("done....")
}
