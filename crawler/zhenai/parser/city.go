package parser


/*
解析 city 页面里的用户 URL
*/
import (
	"Go_learning/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	// 正则
	re := regexp.MustCompile(cityRe)
	// 查找所有的子匹配文本
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name:=string(m[2])
		result.Items = append(
			result.Items, "User "+name)
		result.Requests = append(
			result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(
				c []byte) engine.ParseResult {
				return ParseProfile(
					c, name)
			},
		})

	}
	return result
}

