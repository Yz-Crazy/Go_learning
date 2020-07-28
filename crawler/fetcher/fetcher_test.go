package fetcher

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFetch(t *testing.T) {
	//url:="http://www.zhenai.com/zhenghun"
	//if strings.Contains(url,`http://album.zhenai.com/u/`){
	//	fmt.Println(false)
	//}
	data,err:=ioutil.ReadFile("profile_test_data.html")
	if err!=nil{
		panic(err)
	}
	fmt.Println(ioutil.ReadAll(data))
}