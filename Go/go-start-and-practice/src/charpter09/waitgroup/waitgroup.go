package main
import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = [
		"https://www.github.com/",
		"https://www.baidu.com",
		"https://www.qq.com",
	]

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}
	// 等待所有任务完成
	wg.Wait()
	fmt.Println("over")
}