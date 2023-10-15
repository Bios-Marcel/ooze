package verboserepository_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/gtramontina/ooze/internal/oozetesting/fakelogger"
	"github.com/gtramontina/ooze/internal/oozetesting/fakerepository"
	"github.com/gtramontina/ooze/internal/verboserepository"
	"github.com/stretchr/testify/assert"
)

func TestVerboseRepository(t *testing.T) {
	t.Run("logs when listing source files", func(t *testing.T) {
		logger := fakelogger.New()

		verboserepository.New(
			logger,
			fakerepository.New(fakerepository.FS{
				"file_a.go":     []byte("contents a"),
				"file_b.go":     []byte("contents b"),
				"dir/file_c.go": []byte("contents c"),
			}),
		).ListGoSourceFiles()

		assert.Equal(t, []string{
			"listing go source files…",
			fmt.Sprintf("found 3 source files: [%s %s %s]", filepath.Clean("dir/file_c.go"), "file_a.go", "file_b.go"),
		}, logger.LoggedLines())
	})

	t.Run("logs when linking to temporary path", func(t *testing.T) {
		logger := fakelogger.New()

		verboserepository.New(
			logger,
			fakerepository.New(
				fakerepository.FS{
					"file_a.go": []byte("contents a"),
					"file_b.go": []byte("contents b"),
				},
				fakerepository.NewTemporaryAt("dummy"),
			),
		).LinkAllToTemporaryRepository("some-path")

		assert.Equal(t, []string{
			"linking all files to temporary path 'some-path'…",
			"linked all files to temporary path 'some-path'",
		}, logger.LoggedLines())
	})
}

func TestTemporaryRepository(t *testing.T) {
	t.Run("logs when overwriting paths", func(t *testing.T) {
		logger := fakelogger.New()

		repository := verboserepository.New(
			logger,
			fakerepository.New(
				fakerepository.FS{},
				fakerepository.NewTemporaryAt("some-path"),
			),
		)

		temporary := repository.LinkAllToTemporaryRepository("some-path")
		logger.Clear()

		temporary.Overwrite("source.go", []byte("dummy"))

		assert.Equal(t, []string{
			"overwriting 'source.go'…",
		}, logger.LoggedLines())
	})

	t.Run("logs when removing", func(t *testing.T) {
		logger := fakelogger.New()

		repository := verboserepository.New(
			logger,
			fakerepository.New(
				fakerepository.FS{},
				fakerepository.NewTemporaryAt("some-path"),
			),
		)

		temporary := repository.LinkAllToTemporaryRepository("some-path")
		logger.Clear()

		temporary.Remove()

		assert.Equal(t, []string{
			"removing 'some-path'…",
		}, logger.LoggedLines())
	})
}
