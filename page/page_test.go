package page_test

import (
	"github.com/michalsz/scraper/page"
	"testing"
)

func TestBody(t *testing.T) {
	p := page.Page{"http://google.com", ""}
	expected := ""
	actual := p.Body
	if actual != expected {
		t.Error("fail")
	}
}

func TestIsHttps(t *testing.T) {
	p := page.Page{"http://google.com", ""}
	actual := page.IsHttps(p)
	expected := false
	if actual != expected {
		t.Error("Test failed")
	}
}

func BenchmarkIsHttps(b *testing.B) {
	p := page.Page{"http://google.com", ""}
	for i := 0; i < b.N; i++ {
		page.IsHttps(p)
	}
}
