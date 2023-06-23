package mifer_test

import (
	"testing"

	"github.com/lll-lll-lll-lll/mifer"
	"gotest.tools/assert"
)

func Test_extractFiles_NoFile(t *testing.T) {
	t.Parallel()
	dirPath := "./testdata"
	if _, err := mifer.ExtractFiles(dirPath); err != nil {
		assert.Equal(t, err.Error(), "migration error no file included")
	}
}

func Test_extractFiles(t *testing.T) {
	t.Parallel()
	dirPath := "./testdata/migrations"
	files, err := mifer.ExtractFiles(dirPath)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, len(files), 2)
}

func Test_ReadMigrations(t *testing.T) {
	t.Parallel()
	wantFiles := []string{"create_update_users.sql", "create_users.sql"}
	wantDirPath := []string{"./testdata/migrations/create_update_users.sql", "./testdata/migrations/create_users.sql"}
	dirPath := "./testdata/migrations/"
	migrations, err := mifer.ReadMigrations(dirPath)
	if err != nil {
		t.Error(err)
	}
	for i, m := range migrations {
		assert.Equal(t, m.FilePath, wantDirPath[i], "filepath missed")
		assert.Equal(t, m.Name, wantFiles[i], "wrong file name")
	}
	assert.Equal(t, len(migrations), 2)
}

func Test_MigrationFileRead(t *testing.T) {
	t.Parallel()
	filePath := "./testdata/migrations/create_users.sql"
	mf := &mifer.MigrationFile{
		FilePath: filePath,
	}
	if err := mf.Read(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, mf.IsContentRead, true)
}
