package version

import (
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	got, err := New("1.2.3")
	require.NoError(t, err)
	autogold.ExpectFile(t, got)
}
