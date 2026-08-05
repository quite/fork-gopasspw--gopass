package root

import (
	"testing"

	"github.com/quite/fork-gopasspw--gopass/internal/config"
	"github.com/quite/fork-gopasspw--gopass/pkg/ctxutil"
	"github.com/quite/fork-gopasspw--gopass/tests/gptest"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithHidden(ctx, true)

	rs, err := createRootStore(ctx, u)
	require.NoError(t, err)

	_, err = rs.Get(ctx, "foo")
	require.NoError(t, err)
}
