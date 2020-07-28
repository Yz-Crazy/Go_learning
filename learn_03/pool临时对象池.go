package main

//var (
//	count int32
//	initFunc = func() interface{} {
//		return atomic.AddInt32(&count, 1)
//	}
//
//	pool = sync.Pool{New:initFunc}
//
//)
//func main() {
//	debug.SetGCPercent(debug.SetGCPercent(-1))
//
//	v1 := pool.Get()
//	fmt.Printf("value 1: %v\n", v1)
//	pool.Put(10)
//	pool.Put(11)
//	pool.Put(12)
//	v2 := pool.Get()
//	fmt.Printf("value 2: %v\n", v2)
//
//	runtime.GC()
//
//	v3 := pool.Get()
//	fmt.Printf("value 3: %v\n", v3)
//	pool.New = nil
//	v4 := pool.Get()
//	fmt.Printf("value 4: %v\n", v4)
//}

//func get(pool sync.Pool) {
//	for {
//		val := pool.Get()
//		if val == nil {
//			break
//		}
//		fmt.Println(val)
//	}
//}
//
//func put() sync.Pool {
//	var pool sync.Pool
//	for i := 0; i <= 100; i++ {
//		pool.Put(i)
//	}
//	return pool
//}
//func main() {
//	//var pool sync.Pool
//	pool := put()
//
//	time.Sleep(time.Millisecond)
//
//	get(pool)
//
//}

func main() {
	//var pool sync.Pool
	//pool.Put(10)
	//pool.Put(1)
	//pool.Put(2)
	//pool.Put(3)
	//pool.Put(4)
	//pool.Put(5)
	//runtime.GC()
	////time.Sleep(time.Millisecond)
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//a := "BRcHsSCslT9B4GHuIFt7Yw=="
	//decodeBytes, err := base64.StdEncoding.DecodeString(a)
	//if err != nil{
	//	fmt.Println("解码出错")
	//}
	//fmt.Println(string(decodeBytes))
}
