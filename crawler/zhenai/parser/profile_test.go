package parser

import (
	"Go_learning/crawler/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	fmt.Println(contents)
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "奈何桥何渡奈何")
	if len(result.Items) != 1 {
		t.Errorf("Items shold contain 1"+"element;but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Age:      33,
		Height:   158,
		Income:   "3000元以下",
		Gender:   "男士",
		Name:     "奈何桥何渡奈何",
		Marriage: "离异",

	}
	if profile != expected {
		t.Errorf("expected %v:but was %v", expected, profile)
	}

}
