package leaf

import (
	"context"

	"github.com/quite/fork-gopasspw--gopass/internal/out"
	"github.com/quite/fork-gopasspw--gopass/internal/store"
	"github.com/quite/fork-gopasspw--gopass/pkg/ctxutil"
	"github.com/quite/fork-gopasspw--gopass/pkg/debug"
	"github.com/quite/fork-gopasspw--gopass/pkg/gopass"
	"github.com/quite/fork-gopasspw--gopass/pkg/gopass/secrets"
	"github.com/quite/fork-gopasspw--gopass/pkg/gopass/secrets/secparse"
)

// Get returns the plaintext of a single key.
func (s *Store) Get(ctx context.Context, name string) (gopass.Secret, error) {
	p := s.passfile(ctx, name)

	ciphertext, err := s.storage.Get(ctx, p)
	if err != nil {
		debug.Log("File %s not found: %s", p, err)

		return nil, store.ErrNotFound
	}

	content, err := s.crypto.Decrypt(ctx, ciphertext)
	if err != nil {
		out.Errorf(ctx, "Decryption failed: %s\n%s", err, string(content))

		return nil, store.ErrDecrypt
	}

	if !ctxutil.IsShowParsing(ctx) {
		debug.Log("secrets parsing is disabled. parsing as AKV")

		return secrets.ParseAKV(content), nil
	}

	debug.Log("secrets parsing is enabled")

	return secparse.Parse(content)
}
