package log

import (
	"testing"
)

func init() {
	Init("example", Feishu("token"))
}

func TestError(t *testing.T) {
	o, _ := Init("example", Feishu("token"))
	defer o.Sync()
	With("q", 1).With("w", 2).Warnf("%s=%d", "e", 5)
	Warnf("%s=%d", "e", 5)
}
