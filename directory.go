/*
The MIT License (MIT)

Copyright © 2020-2021 The fsio Authors (Neil Hemming)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package fsio provides so file system helper functions and types
package fsio

import (
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// FileModes contains filer and director permissions.
type FileModes struct {
	FileMode os.FileMode
	DirMode  os.FileMode
}

// NewFileModes creates a new FileModes struct with default
// file permissions.
func NewFileModes() FileModes {
	return FileModes{
		FileMode: os.FileMode(0666),
		DirMode:  os.FileMode(0777),
	}
}

// ExpandFilePath expadas a file path to an absolute path.
// This function handles home directory '~' specifications.
func ExpandFilePath(path string) (string, error) {
	// If we have an empty path return empty (Abs expands to cwd)
	if path == "" {
		return "", nil
	}
	p, err := homedir.Expand(path)
	if err != nil {
		return path, err
	}

	p, err = filepath.Abs(p)
	return p, err
}

// CreateFileDirectory creates a file's parent directories as necessary.
func CreateFileDirectory(filename string, dirMode os.FileMode) error {
	path, err := ExpandFilePath(filename)
	if err != nil {
		return err
	}
	dir, _ := filepath.Split(path)

	return os.MkdirAll(dir, dirMode)
}

// WriteFileToPath writes a file to the filename specified.
// Function expand the path (including ~ substitution) and creates
// as necessary the parent directoried.
func WriteFileToPath(filename string, buf []byte, modes FileModes) error {
	filename, err := ExpandFilePath(filename)
	if err != nil {
		return err
	}

	if err = CreateFileDirectory(filename, modes.DirMode); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf, modes.FileMode)
}

// ReadFileFromPath expands the directory path and then reads the file specified.
func ReadFileFromPath(filename string) ([]byte, error) {
	filename, err := ExpandFilePath(filename)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(filename)
}

// MakeAbsFromRelativeToFile makes path into an absolute path, using relfile as
// sibling file location to the relative base of the path.
func MakeAbsFromRelativeToFile(path, relFile string) (string, error) {
	path, err := homedir.Expand(path)
	if err != nil {
		return "", err
	}
	relFile, err = homedir.Expand(relFile)
	if err != nil {
		return "", err
	}

	if filepath.IsAbs(path) {
		return path, nil
	}

	a, err := filepath.Abs(relFile)
	if err != nil {
		return "", err
	}

	dir, _ := filepath.Split(a)

	return filepath.Join(dir, path), nil
}
