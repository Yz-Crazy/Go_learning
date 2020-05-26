package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup
	c := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			m.Lock()         // 先加锁
			defer m.Unlock() // 最后解锁
			//fmt.Println(i)
			fmt.Println(c)
			rand.Seed(time.Now().UnixNano())    // 设置随机数种子，这里用纳秒来做随机数种子
			time.Sleep(time.Millisecond*time.Duration(rand.Intn(10)))
			// time.Duration 因为time.Sleep 的参数是Duration类型数据，所以这里我们要把随机数的值转换成Duration类型
			c++
			wg.Done()
		}()
	}
	wg.Wait()
}
