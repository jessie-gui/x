package xudp_test

import (
	"testing"

	"github.com/jessie-gui/x/xnet/xudp"
	"github.com/stretchr/testify/assert"
)

func TestGetFreePort(t *testing.T) {
	_, err := xudp.GetFreePort()
	if err != nil {
		t.Error(err)
	}
}

func TestGetFreePorts(t *testing.T) {
	ports, err := xudp.GetFreePorts(2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(ports), 2)
}
