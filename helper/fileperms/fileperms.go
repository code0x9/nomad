package fileperms

import (
	"fmt"
	"os"
)

// Mode represents a simplified octal representation of a Unix file permissions
// mode interface provided by os.FileMode.
//
// The use of this package is intended to prevent accidental use of incorrect
// permissions caused by improper octal notation in Go's default decimal base.
//
// For example, the code `os.Chmod(655)` produces file permissions | todo
type Mode = os.FileMode

// 0) no permission
// 1) execute
// 2) write
// 3) execute + write
// 4) read
// 5) read + execute
// 6) read + write
// 7) read + write + execute
const (
	Oct600 Mode = 0600
	Oct655 Mode = 0655
	Oct777 Mode = 0777
)

// Check returns true if the os.FileMode permission bits of f match the
// expected permissions in exp.
func Check(f *os.File, exp Mode) error {
	info, err := f.Stat()
	if err != nil {
		return err
	}

	perm := info.Mode().Perm()
	if perm != exp {
		return fmt.Errorf("file mode expected %o, got %o", exp, perm)
	}

	return nil
}
