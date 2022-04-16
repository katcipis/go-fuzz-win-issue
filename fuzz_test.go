//go:build go1.18
// +build go1.18

package jtoh_test

import (
	"testing"
)

func FuzzIssue(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b string) {
	})
}
