//go:build mutation

package ooze_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/gtramontina/ooze"
	"github.com/gtramontina/ooze/internal/viruses/floatincrement"
	"github.com/gtramontina/ooze/internal/viruses/integerdecrement"
	"github.com/gtramontina/ooze/internal/viruses/integerincrement"
	"github.com/gtramontina/ooze/internal/viruses/loopbreak"
)

func TestMutation(t *testing.T) {
	color.NoColor = false

	ooze.Release(
		t,
		ooze.WithRepositoryRoot("."),
		ooze.WithTestCommand("make test MAKEFLAGS="),
		ooze.WithMinimumThreshold(0.5),
		ooze.Parallel(),
		ooze.WithViruses(
			integerincrement.New(),
			integerdecrement.New(),
			floatincrement.New(),
			loopbreak.New(),
		),
	)
}
