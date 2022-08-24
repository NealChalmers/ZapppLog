package zappplog_test

import (
	"testing"

	"github.com/NealChalmers/zappplog"
)

func TestInit(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5000; j++ {
			zappplog.Infof("%s %d", "dsdsd", i*5000+j)
		}
	}
}
