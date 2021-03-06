// +build !windows

package hidden

import (
	"os"
	"path/filepath"
	"strings"
)

func hide(filename string) error {
	if !strings.HasPrefix(filepath.Base(filename), ".") {
		newName := filepath.Join(filepath.Dir(filename), "."+filepath.Base(filename))
		err := os.Rename(filename, newName)
		if err != nil {
			return err
		}
	}
	return nil
}

func isHidden(filename string) (bool, error) {
	if len(filepath.Base(filename)) > 0 && filepath.Base(filename)[0:1] == "." {
		return true, nil
	}

	return false, nil
}
