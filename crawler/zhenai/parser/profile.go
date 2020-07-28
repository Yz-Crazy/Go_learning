package parser

import (
	"Go_learning/crawler/engine"
	"Go_learning/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`"age":(\d+),"`)
var marriageRe = regexp.MustCompile(`"marriageString":"([^<]+)","memberID"`)
var heightRe = regexp.MustCompile(`"heightString":"(\d+)cm","hideVerifyModule"`)
var incomeRe = regexp.MustCompile(`"salaryString":"([^<]+)","showHighVipPic"`)

//var weightRe = regexp.MustCompile(`"salaryString":"([^<]+)",`)

func ParseProfile(contents []byte,name string) engine.ParseResult {
	//fmt.Printf("%s",contents)
	profile := model.Profile{}
	profile.Name=name
	age, err := strconv.Atoi(extractSting(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractSting(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	profile.Income = extractSting(contents, incomeRe)
	profile.Marriage = extractSting(contents, marriageRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractSting(contents []byte, re *regexp.Regexp) string {
	match := marriageRe.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
