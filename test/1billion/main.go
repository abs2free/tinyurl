package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

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

	urlsChan := make(chan []model.TinyURL, 20)

	var wg sync.WaitGroup
	// 定义workder
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for urls := range urlsChan {
				select {
				case <-ctx.Done():
					return
				default:
					_, err := mysql.BatchAddTinyURL(ctx, urls)
					if err != nil {
						cancel()
						log.Fatal(err)
					}
				}
			}
			wg.Done()
		}()
	}

	var urls []model.TinyURL
	for i := 1; i < 100000000; i++ {
		url := model.TinyURL{
			LongURL:   fmt.Sprintf("this is a very very long url, please remember it ,the number is:%d", i),
			ShortURL:  base62.Encode(i),
			ExpiredAt: time.Now().AddDate(1, 2, 2),
			CreatedAt: time.Now(),
		}
		urls = append(urls, url)

		if i%100 == 0 {
			fmt.Printf("第%d个加入到chan\n", i)
			urlsChan <- urls
			urls = []model.TinyURL{}
		}
	}
	urlsChan <- urls
	wg.Wait()
	fmt.Println("done....")
}
