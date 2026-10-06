package store

import (
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/SumonMSelim/timothy/internal/platform/pgpool"
)

func TestInsertDerivesStatusBeforeDatabaseAccess(t *testing.T) {
	t.Parallel()
	pool := pgpool.New(t.Context(), "", slog.New(slog.NewTextHandler(io.Discard, nil)))
	store := New(pool, slog.New(slog.NewTextHandler(io.Discard, nil)))

	_, err := store.Insert(t.Context(), Memory{Actor: ActorUser, RequireReview: true})
	if err == nil || !errors.Is(err, pgpool.ErrDegraded) {
		t.Fatalf("Insert error = %v, want degraded database error", err)
	}
}
