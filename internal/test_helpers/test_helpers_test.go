package testhelpers

import ("testing")

func TestAssert(t *testing.T) {
	t.Run("checking if Assert function works", func(t *testing.T) {
		got := 1
		want := 1
		Assert(t, got, want)
	})
}