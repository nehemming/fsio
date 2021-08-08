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

package fsio

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestNewFileModes(t *testing.T) {
	fm := NewFileModes()

	if fm.FileMode != os.FileMode(0666) {
		t.Errorf("Unexpected file mode %v", fm.FileMode)
	}

	if fm.DirMode != os.FileMode(0777) {
		t.Errorf("Unexpected dir mode %v", fm.DirMode)
	}
}

func TestExpandFilePathEmpty(t *testing.T) {
	_ = SetHome()
	defer ResetHome()

	p, e := ExpandFilePath("")
	if e != nil {
		t.Errorf("Exact test: Err %v", e)
	} else if p != "" {
		t.Errorf("Exact test: Path %v", p)
	}
}

func TestExpandFilePath(t *testing.T) {
	h := SetHome()
	defer ResetHome()

	p, e := ExpandFilePath(h)
	if e != nil {
		t.Errorf("Exact test: Err %v", e)
	} else if p != h {
		t.Errorf("Exact test: Path %v", p)
	}

	p, e = ExpandFilePath("~")
	if e != nil {
		t.Errorf("Exact test: Err %v", e)
	} else if p != h {
		t.Errorf("Exact test: Path %v", p)
	}
}

func TestCreateFileDirectory(t *testing.T) {
	_ = SetHome()
	defer ResetHome()

	fm := NewFileModes()

	file := filepath.Join(testRoot, "/deep/fileA")
	dir := filepath.Join(testRoot, "/deep")

	if err := CreateFileDirectory(file, fm.DirMode); err != nil {
		t.Errorf("CreateFileDirectory: Err %v", err)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("CreateFileDirectory: Does not exist %v", dir)
	}

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		t.Errorf("CreateFileDirectory: File exists %v", file)
	}

	// repeat
	if err := CreateFileDirectory(file, fm.DirMode); err != nil {
		t.Errorf("CreateFileDirectory x 2: Err %v", err)
	}
}

func TestReadWriteFileToPath(t *testing.T) {
	_ = SetHome()
	defer ResetHome()

	fm := NewFileModes()

	file := filepath.Join(testRoot, "/deep/fileA")

	if err := WriteFileToPath(file, []byte("hello world"), fm); err != nil {
		t.Errorf("WriteFileToPath: Err %v", file)
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("WriteFileToPath: File not exist %v", file)
	}

	if err := WriteFileToPath(file, []byte("hello world 2"), fm); err != nil {
		t.Errorf("WriteFileToPath: Err %v", file)
	}

	buf, err := ReadFileFromPath(file)
	if err != nil {
		t.Errorf("ReadFileToPath: Err %v", file)
	}

	if string(buf) != "hello world 2" {
		t.Errorf("ReadFileToPath: Wrong contents %v", string(buf))
	}
}

func TestMakeAbsFromRelativeToFile(t *testing.T) {
	_ = SetHome()
	defer ResetHome()

	file := filepath.Join(testRoot, "/deep/fileA")

	abPath, err := MakeAbsFromRelativeToFile("../other/altfile", file)
	if err != nil {
		t.Errorf("MakeAbsFromRelativeToFile: Err %v", err)
	}

	expected := filepath.Join(testHome, testRoot, "other/altfile")

	if abPath != expected {
		t.Errorf("MakeAbsFromRelativeToFile mismatch: %v  %v", abPath, expected)
	}
}

func TestMakeAbsFromRelativeToFileWithAbsPath(t *testing.T) {
	h := SetHome()
	defer ResetHome()

	file := filepath.Join(testRoot, "/deep/fileA")

	abPath, err := MakeAbsFromRelativeToFile(h, file)
	if err != nil {
		t.Errorf("MakeAbsFromRelativeToFile: Err %v", err)
	}

	if abPath != h {
		t.Errorf("MakeAbsFromRelativeToFile mismatch: %v  %v", abPath, h)
	}
}

const homeEnv = "HOME"

var (
	homeDir    string
	testHome   string
	needsReset bool
)

var testRoot = "./testdata"

func init() {
	// determin home dir
	var err error

	homeDir, err = homedir.Dir()

	if err != nil {
		panic("cannot get homedir for tests")
	}

	// Reset home if its set
	if os.Getenv(homeEnv) != "" {
		needsReset = true
	}
}

func SetHome() string {
	testHome, _ = os.Getwd()
	os.Setenv(homeEnv, testHome)

	homedir.Reset()

	// clear test data
	os.RemoveAll(filepath.Join(testHome, testRoot))

	return testHome
}

func ResetHome() {
	if homeDir == "" {
		return
	}

	if needsReset {
		os.Setenv(homeEnv, homeDir)
	} else {
		os.Setenv(homeEnv, "homeDir")
	}

	// clear test data
	os.RemoveAll(filepath.Join(testHome, testRoot))

	homedir.Reset()
}
