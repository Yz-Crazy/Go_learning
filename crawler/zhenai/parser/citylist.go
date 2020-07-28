package parser

import (
	"Go_learning/crawler/engine"
	"regexp"
)
// 正则抽取需要的 URL 和 City 列表
const cityListRe=`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult {


	re := regexp.MustCompile(cityListRe)
	mateches := re.FindAllSubmatch(contents, -1)
	result :=engine.ParseResult{}
	limit:=10

	for _, m := range mateches {
		result.Items=append(result.Items,"City "+string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
		limit --
		if limit ==0{
			break
		}
	}
	return result

}
