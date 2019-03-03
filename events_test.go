package events

import "testing"

func TestCorrectInterfaceImplementations(_ *testing.T) {
	var _ Dispatcher = (*disp)(nil)
	var _ Consumer = (*cons)(nil)
}
