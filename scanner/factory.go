package scanner

import "runtime"

// NewScanner returns a platform-appropriate Scanner implementation.
func NewScanner() Scanner {
	switch runtime.GOOS {
	case "windows":
		return &WindowsScanner{}
	case "darwin":
		return &MacScanner{}
	default:
		return &LinuxScanner{}
	}
}
