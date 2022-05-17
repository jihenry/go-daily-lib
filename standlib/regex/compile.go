package regex

import (
	"fmt"
	"regexp"
)

func compile(refer string) bool {
	regexpRule, _ := regexp.Compile(`^https://servicewechat.com/[\w]+/[\w.]+/(page-frame|index).html`)
	url := regexpRule.FindString(refer)
	fmt.Printf("refer url:%s\n", url)
	return url != ""
}
