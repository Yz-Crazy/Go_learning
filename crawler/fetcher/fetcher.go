package fetcher
/*
fetcher 的功能是把接收到的 URL 请求之后返回网页内容
*/
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Fetch(url string) ([]byte, error) {
	//fmt.Printf("Fetching %s\n", url)

	// 判断 URL 是不是用户页的 URL ，如果是，返回 设置本地存储好的页面 这么做的原因是原网址加密了，无法解决，所以用这个方法来写
	if strings.Contains(url, `album.zhenai.com/u/`) {
		fmt.Println("打开City_data.html")
		// 读取出 profile_test_data.html 中的数据，contents 是[]byte类型
		contents, err := ioutil.ReadFile("src/Go_learning/crawler/fetcher/profile_test_data.html")
		if err != nil{
			panic(err)
		}
		//fmt.Println(string(contents))
		return contents,nil
	}
	// http 请求
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// 延迟关闭请求到的内容通道
	defer resp.Body.Close()
	// 判断 http 请求状态码
	if resp.StatusCode != http.StatusOK {
		all, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", all)
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}
	// 自动识别网页编码

	// e:=determineEncoding(resp.Body)

	// 网页编码转换
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// resp.Body 是内存地址，这里需要用 ioutil.ReadAll 把内容全部读出来，返回的是一个[]byte
	return ioutil.ReadAll(resp.Body)

}

// 自动识别网页编码
//func determineEncoding(r io.Reader) encoding.Encoder{
//	bytes,err=:bufio.NewReader(r).Peek(1024)
//	if err != nil{
//      log.Printf("Fetcher error: %v",err)
//		return unicode.UTF8
//	}
//	e,_,_:=charset.DetermineEncoding(bytes,"")
//	return e
//}
