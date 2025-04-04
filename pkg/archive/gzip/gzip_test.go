package gzip

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/goreleaser/goreleaser/v2/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestGzFile(t *testing.T) {
	tmp := t.TempDir()
	f, err := os.Create(filepath.Join(tmp, "test.gz"))
	require.NoError(t, err)
	defer f.Close()
	archive := New(f)
	defer archive.Close()

	require.NoError(t, archive.Add(config.File{
		Destination: "sub1/sub2/subfoo.txt",
		Source:      "../testdata/sub1/sub2/subfoo.txt",
	}))
	require.EqualError(t, archive.Add(config.File{
		Destination: "foo.txt",
		Source:      "../testdata/foo.txt",
	}), "gzip: failed to add foo.txt, only one file can be archived in gz format")
	require.NoError(t, archive.Close())
	require.NoError(t, f.Close())

	f, err = os.Open(f.Name())
	require.NoError(t, err)
	defer f.Close()

	info, err := f.Stat()
	require.NoError(t, err)
	require.Lessf(t, info.Size(), int64(500), "archived file should be smaller than %d", info.Size())

	gzf, err := gzip.NewReader(f)
	require.NoError(t, err)
	defer gzf.Close()

	require.Equal(t, "sub1/sub2/subfoo.txt", gzf.Name)

	bts, err := io.ReadAll(gzf)
	require.NoError(t, err)
	require.Equal(t, "sub\n", string(bts))
}

func TestGzFileCustomMtime(t *testing.T) {
	f, err := os.Create(filepath.Join(t.TempDir(), "test.gz"))
	require.NoError(t, err)
	defer f.Close()
	archive := New(f)
	defer archive.Close()

	now := time.Now().Truncate(time.Second)

	require.NoError(t, archive.Add(config.File{
		Destination: "sub1/sub2/subfoo.txt",
		Source:      "../testdata/sub1/sub2/subfoo.txt",
		Info: config.FileInfo{
			ParsedMTime: now,
		},
	}))
	require.NoError(t, archive.Close())
	require.NoError(t, f.Close())

	f, err = os.Open(f.Name())
	require.NoError(t, err)
	defer f.Close()

	info, err := f.Stat()
	require.NoError(t, err)
	require.Lessf(t, info.Size(), int64(500), "archived file should be smaller than %d", info.Size())

	gzf, err := gzip.NewReader(f)
	require.NoError(t, err)
	defer gzf.Close()

	require.Equal(t, "sub1/sub2/subfoo.txt", gzf.Name)
	require.Equal(t, now, gzf.ModTime)
}
