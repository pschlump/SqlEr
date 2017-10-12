package SqlEr

import "testing"

func Test_SqlEr0(t *testing.T) {
	// t.Errorf("[%d] got an error when non-expected, NErrors= %d\n", ii, NErrors)
	x := ValidateInstallModel("test1.jx")
	_ = x
}
