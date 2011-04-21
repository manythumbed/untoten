package untoten

import (
	"testing"
	"os"
	"fmt"
	"html"
)

func TestSomething(t *testing.T) {
	f, err := os.Open("test-data/html5.html", os.O_RDONLY, 0)
	if err != nil {
		t.Errorf("Unable to open file: %v", err)
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		t.Errorf("Error parsing html:- %v", err)
	}
	var fn func(*html.Node, int)
	fn = func(n *html.Node, i int) {
		switch n.Type {
		default:
			fmt.Println(i, "Unexpected type")
		case html.TextNode:
			fmt.Println(i, "Text Node", n.Data, n.Parent)
		case html.DocumentNode:
			fmt.Println(i, "Document Node")
		case html.ElementNode:
			fmt.Println(i, "Element Node", n.Data, n.Parent)
		case html.CommentNode:
			fmt.Println(i, "Element Node")
		}
		for _, c := range n.Child {
			fn(c, i + 1)
		}
	}
	fn(doc, 0)
}
