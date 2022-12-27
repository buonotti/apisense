// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/config.example.toml
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsConfigExampleToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\xb1\xae\x1a\x31\x10\x45\x7b\x7f\xc5\x15\xb4\x1b\x7f\x40\x24\x3a\x22\x25\x51\x12\x0a\xa2\x34\x51\x8a\x61\x77\xbc\x8c\x64\x3c\xab\xf1\x2c\x84\xf7\xf5\x4f\xeb\x7d\x50\x50\x5a\xbe\xf7\xcc\xb9\x5b\xfc\x3e\x4b\x85\x54\xf8\x99\xc1\xff\xe9\x32\x65\x46\xaf\x25\xc9\x38\x1b\xb9\x68\x41\x92\xcc\x48\x6a\x2d\x72\xd8\x7f\xc5\x9e\x9c\xf0\x53\x8b\xb8\x5a\x0c\xe1\x6f\xd6\xf1\x5f\xd8\xe2\xc8\xde\x22\x59\x47\x64\xbe\x72\x8e\xf8\x43\x59\x06\x5c\x29\xcf\x5c\x41\xc6\x9f\x31\xf0\x69\x1e\x3b\x48\x49\xda\xe1\x46\x56\x3a\xb0\x99\x5a\x87\x44\x4e\xb9\xc3\x44\x45\xfa\xd0\xfa\xd8\x61\xb3\x04\x37\xe1\x85\xbe\x18\x45\xfc\x60\xba\x32\xf8\x32\xf9\x1d\xae\xed\xc3\x15\xd5\x07\x9d\x3d\x86\x66\xbd\xc3\xa6\x95\xbf\x14\x3a\x65\xc6\x64\xec\x7e\x6f\x49\x9d\x7d\x9a\x3d\xae\xf3\x6f\x92\x33\x7a\xcd\x6a\xf2\xc6\xcf\x2b\x6b\x04\x54\x06\x4c\x26\xc5\x21\x0e\x29\x20\x18\xd3\xd0\x78\x49\xed\x42\x1e\xc3\x16\xdf\x12\xea\x22\xa8\x48\x94\x2b\x77\xaf\x90\x76\xe1\xc4\x4b\xff\xfb\xf1\xf0\xeb\xd9\xfc\x30\xda\xc1\x6d\xe6\xe7\x4c\x6d\xcf\x46\x53\xeb\x79\x55\x5b\x70\xa3\x94\x31\xe2\x50\xf2\x1d\x67\xaa\xa0\x02\x4e\x89\x7b\x87\xa4\xc7\x38\xa9\x0f\x93\x85\x11\x43\x23\x7c\x5a\x09\xeb\x99\xf7\x00\x00\x00\xff\xff\x56\x7f\x67\x1d\xf4\x01\x00\x00")

func assetsConfigExampleTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsConfigExampleToml,
		"assets/config.example.toml",
	)
}

func assetsConfigExampleToml() (*asset, error) {
	bytes, err := assetsConfigExampleTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/config.example.toml", size: 500, mode: os.FileMode(420), modTime: time.Unix(1671709788, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/config.example.toml": assetsConfigExampleToml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"config.example.toml": &bintree{assetsConfigExampleToml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
