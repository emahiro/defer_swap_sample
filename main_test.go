package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	t.Run("rewrite test", func(t *testing.T) {
		defer rewriteString("fuga")()
	})
}

func rewriteString(s string) func() {
	var tmp string
	tmp, rewriter := rewriter, s
	return func() {
		rewriter = tmp
	}
}
