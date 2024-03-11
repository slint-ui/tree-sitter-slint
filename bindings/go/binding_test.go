package tree_sitter_slint_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-slint"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_slint.Language())
	if language == nil {
		t.Errorf("Error loading Slint grammar")
	}
}
