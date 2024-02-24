package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/abs2free/tinyurl/base62"
	"github.com/abs2free/tinyurl/store"
	"github.com/abs2free/tinyurl/store/model"
	"github.com/stretchr/testify/assert"
)

var mysql *store.Mysql

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	mysql, err = store.NewMysql(dsn)
	if err != nil {
		panic(err)
	}
}

func TestAddTinyURL(t *testing.T) {
	logURL := "this  is a very long url remenber it"
	shortURL := base62.Encode(100000000)

	ctx := context.Background()
	result, err := mysql.AddTinyURL(ctx, logURL, shortURL, time.Now().AddDate(0, 2, 2))
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, result)
}

func TestBatchAddTinyURL(t *testing.T) {
	var urls []model.TinyURL
	for i := 1; i <= 100; i++ {
		longURL := fmt.Sprintf("this is a very long url remenber it ,this number is:%d", i)
		shortURL := base62.Encode(1000000 + i)
		url := model.TinyURL{
			LongURL:   longURL,
			ShortURL:  shortURL,
			ExpiredAt: time.Now().AddDate(1, 2, 2),
			CreatedAt: time.Now(),
		}
		urls = append(urls, url)
	}

	ctx := context.Background()
	result, err := mysql.BatchAddTinyURL(ctx, urls)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, result)
}
