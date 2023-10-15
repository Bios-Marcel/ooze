//go:build !windows

package fsrepository_test

import (
	"os"
	"syscall"
	"testing"

	"github.com/gtramontina/ooze/internal/fsrepository"
	"github.com/stretchr/testify/assert"
)

func TestFSTemporaryRepository_NonWindows(t *testing.T) {
	t.Run("an existing regular file", func(t *testing.T) {
		dir := t.TempDir()
		assert.NoError(t, os.WriteFile(dir+"/file.txt", []byte("original data"), 0o600))

		repository := fsrepository.NewTemporary(dir)
		repository.Overwrite("file.txt", []byte("new data"))

		data, err := os.ReadFile(dir + "/file.txt")
		assert.NoError(t, err)
		assert.Equal(t, []byte("new data"), data)

		stat, err := os.Stat(dir + "/file.txt")
		assert.NoError(t, err)
		mask := syscall.Umask(0)
		defer syscall.Umask(mask)
		assert.Equal(t, os.ModePerm^os.FileMode(mask), stat.Mode())
	})
}
