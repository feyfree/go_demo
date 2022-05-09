package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	// 这地方 127.0.0.1: 或者是 127.0.0.1:0 都会随机分配端口
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())
}
