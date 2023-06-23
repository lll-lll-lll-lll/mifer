package mifer

import (
	"fmt"
	"os"
	"path/filepath"
)

type MigrationFile struct {
	Content       []byte
	IsContentRead bool
	FilePath      string
	Name          string
}

// Read read file content
func (mf *MigrationFile) Read() error {
	if mf.IsContentRead {
		return nil
	}
	content, err := os.ReadFile(mf.FilePath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	mf.Content = content
	mf.IsContentRead = true
	return nil
}

func ReadMigrations(dirPath string) ([]*MigrationFile, error) {
	sqls, err := extractFiles(dirPath)
	if err != nil {
		return nil, err
	}
	migrations := make([]*MigrationFile, 0, len(sqls))
	for _, sqlFile := range sqls {
		if ext := filepath.Ext(sqlFile); ext != ".sql" {
			continue
		}
		migrations = append(migrations, &MigrationFile{Name: sqlFile, FilePath: fmt.Sprintf("%s%s", dirPath, sqlFile)})
	}
	return migrations, nil
}

func extractFiles(dirPath string) ([]string, error) {
	es, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	fileNames := make([]string, 0, len(es))
	for _, f := range es {
		if f.IsDir() {
			continue
		}
		fileNames = append(fileNames, f.Name())
	}
	if len(fileNames) == 0 {
		return nil, fmt.Errorf("%w", Error(MigrationErr, "no file included"))
	}
	return fileNames, nil
}
