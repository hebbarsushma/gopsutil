// +build !darwin,!linux,!freebsd,!openbsd,!solaris,!windows

package mem

import (
	"context"

	"github.com/hebbarsushma/gopsutil/new/common"
)

func VirtualMemory() (*VirtualMemoryStat, error) {
	return VirtualMemoryWithContext(context.Background())
}

func VirtualMemoryWithContext(ctx context.Context) (*VirtualMemoryStat, error) {
	return nil, common.ErrNotImplementedError
}

func SwapMemory() (*SwapMemoryStat, error) {
	return SwapMemoryWithContext(context.Background())
}

func SwapMemoryWithContext(ctx context.Context) (*SwapMemoryStat, error) {
	return nil, common.ErrNotImplementedError
}
