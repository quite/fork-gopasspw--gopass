package root

import (
	"testing"

	"github.com/fatih/color"
	"github.com/quite/fork-gopasspw--gopass/internal/config"
	"github.com/quite/fork-gopasspw--gopass/pkg/ctxutil"
	"github.com/quite/fork-gopasspw--gopass/tests/gptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCrypto(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithHidden(ctx, true)
	color.NoColor = true

	rs, err := createRootStore(ctx, u)
	require.NoError(t, err)

	assert.NotNil(t, rs.Crypto(ctx, ""))
}
