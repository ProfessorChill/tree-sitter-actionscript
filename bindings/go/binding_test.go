package tree_sitter_actionscript_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-actionscript"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_actionscript.Language())
	if language == nil {
		t.Errorf("Error loading Actionscript grammar")
	}
}
