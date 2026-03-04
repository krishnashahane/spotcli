package cli

import (
	"context"
	"errors"
	"testing"

	"github.com/krishnashahane/spotcli/internal/output"
	"github.com/krishnashahane/spotcli/internal/spotify"
	"github.com/krishnashahane/spotcli/internal/testutil"
)

func TestSearchTrackCmd(t *testing.T) {
	ctx, out, _ := testutil.NewTestContext(t, output.FormatPlain)
	mock := &testutil.SpotifyMock{
		SearchFn: func(ctx context.Context, kind, query string, limit, offset int) (spotify.SearchResult, error) {
			return spotify.SearchResult{Type: kind, Total: 1, Items: []spotify.Item{{ID: "t1", Name: "Song", Type: "track"}}}, nil
		},
	}
	ctx.SetSpotify(mock)
	cmd := SearchTrackCmd{SearchArgs: SearchArgs{Query: "song", Limit: 1, Offset: 0}}
	if err := cmd.Run(ctx); err != nil {
		t.Fatalf("run: %v", err)
	}
	if out.String() == "" {
		t.Fatalf("expected output")
	}
}

func TestSearchTrackCmdError(t *testing.T) {
	ctx, _, _ := testutil.NewTestContext(t, output.FormatPlain)
	mock := &testutil.SpotifyMock{
		SearchFn: func(ctx context.Context, kind, query string, limit, offset int) (spotify.SearchResult, error) {
			return spotify.SearchResult{}, errors.New("boom")
		},
	}
	ctx.SetSpotify(mock)
	cmd := SearchTrackCmd{SearchArgs: SearchArgs{Query: "song", Limit: 1, Offset: 0}}
	if err := cmd.Run(ctx); err == nil {
		t.Fatalf("expected error")
	}
}
