package tree_sitter_gritql_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_gritql "github.com/tree-sitter/tree-sitter-gritql/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_gritql.Language())
	if language == nil {
		t.Errorf("Error loading GritQL grammar")
	}
}
