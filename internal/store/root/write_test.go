package root

import (
	"testing"

	"github.com/quite/fork-gopasspw--gopass/internal/config"
	"github.com/quite/fork-gopasspw--gopass/pkg/ctxutil"
	"github.com/quite/fork-gopasspw--gopass/pkg/gopass/secrets"
	"github.com/quite/fork-gopasspw--gopass/tests/gptest"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithHidden(ctx, true)

	rs, err := createRootStore(ctx, u)
	require.NoError(t, err)

	sec := secrets.NewAKV()
	sec.SetPassword("foo")
	_, err = sec.Write([]byte("bar"))
	require.NoError(t, err)
	require.NoError(t, rs.Set(ctx, "zab", sec))

	err = rs.Set(ctx, "zab2", sec)
	require.NoError(t, err)
}
