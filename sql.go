package mifer

import (
	"fmt"
	"os"
	"path/filepath"
)

type SQLFile struct {
	Content       []byte
	IsContentRead bool
	FilePath      string
	Name          string
}

// Read read file content
func (sf *SQLFile) Read() error {
	if sf.IsContentRead {
		return nil
	}
	content, err := os.ReadFile(sf.FilePath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	sf.Content = content
	sf.IsContentRead = true
	return nil
}

func ReadSQLs(dirPath string) ([]*SQLFile, error) {
	files, err := extractFiles(dirPath)
	if err != nil {
		return nil, err
	}
	sqls := make([]*SQLFile, 0, len(files))
	for _, sqlFile := range files {
		if ext := filepath.Ext(sqlFile); ext != ".sql" {
			continue
		}
		sqls = append(sqls, &SQLFile{Name: sqlFile, FilePath: fmt.Sprintf("%s%s", dirPath, sqlFile)})
	}
	return sqls, nil
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
		return nil, NewErr(SqlErr, "no file included")
	}
	return fileNames, nil
}
