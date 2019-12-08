package main

import (
	"strconv"
	"sync"
)

var(
	waitGroup sync.WaitGroup
)

func main031() {
	baseUrl := "https://www.duotoo.com/zt/rbmn/index"
	for i := 1; i < 23; i++ {
		var url string
		if i != 1 {
			url = baseUrl + "_" + strconv.Itoa(i) + ".html"
		} else {
			url = baseUrl + ".html"
		}
		DownladPageImgs(url)
	}
	waitGroup.Wait()
}

func DownladPageImgs(url string) {
	imginfos := GetPageImginfos(url)
	for _, imginfoMap := range imginfos {
		DownloadImgAsync2(imginfoMap["url"], imginfoMap["filename"],&waitGroup)
	}
}

func DownloadImgAsync2(url, filename string,wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		chSem <- 123
		DownloadImg(url, filename)
		<-chSem
		wg.Done()
	}()
}
