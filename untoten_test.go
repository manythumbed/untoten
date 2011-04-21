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
	var fn func(*html.Node)
	fn = func(n *html.Node) {
		switch n.Type {
		default:
			fmt.Println("Unexpected type")
		case html.TextNode:
			fmt.Println("Text Node", n.Data)
		case html.DocumentNode:
			fmt.Println("Document Node")
		case html.ElementNode:
			fmt.Println("Element Node", n.Data)
		case html.CommentNode:
			fmt.Println("Element Node")
		}
		for _, c := range n.Child {
			fn(c)
		}
	}
	fn(doc)
}
