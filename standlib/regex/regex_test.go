package regex

import (
	"testing"
)

func TestCompile(t *testing.T) {
	cas := []struct {
		refer string
		want  bool
	}{
		{refer: "https://servicewechat.com/wx4e443d9613e66610/31/page-frame.html", want: true},
		{refer: "https://servicewechat.com/wx4e443d9613e66610/31/index.html", want: false},
	}
	for _, ca := range cas {
		compile(ca.refer)
	}
}
