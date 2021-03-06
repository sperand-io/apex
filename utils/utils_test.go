package utils_test

import (
	"testing"

	"github.com/apex/apex/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Sha256(t *testing.T) {
	s := utils.Sha256([]byte("Hello World"))
	assert.Equal(t, "pZGm1Av0IEBKARczz7exkNYsZb8LzaMrV7J32a2fFG4=", s)
}

func Test_LoadFiles(t *testing.T) {
	files, _ := utils.LoadFiles("_fixtures/fileAndDir", []string{"testfile"})
	assert.Equal(t, "testdir/indir", files[0])
	assert.Equal(t, 2, len(files))
}

func Test_ReadIgnoreFile_found(t *testing.T) {
	patterns, err := utils.ReadIgnoreFile("_fixtures")
	assert.NoError(t, err)
	assert.Equal(t, []string{"*.go", "*.log", "whatever"}, patterns)
}

func Test_ReadIgnoreFile_missing(t *testing.T) {
	patterns, err := utils.ReadIgnoreFile("_fixtures/fileAndDir")
	assert.NoError(t, err)
	assert.Equal(t, []string(nil), patterns)
}
