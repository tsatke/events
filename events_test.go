package events

import "testing"

func TestCorrectInterfaceImplementations(t *testing.T) {
	var _ Dispatcher = (*disp)(nil)
	var _ Consumer = (*cons)(nil)
}
