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

var _assetsConfigExampleToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\xc1\x4e\xf3\x30\x10\x84\xef\x7e\x8a\x51\x7b\xcd\xef\x07\xf8\xa5\xdc\x8a\x04\x08\xe8\xa1\x88\x0b\xe2\xb0\x4d\xd6\xe9\x4a\xae\x37\xb2\x37\x2d\xe5\xe9\x51\x1c\xda\x43\xaf\xf6\xcc\xb7\xdf\xac\xf1\x7e\x90\x02\x29\xb0\x03\x83\xbf\xe9\x38\x46\x46\xa7\x29\xc8\x30\x65\x32\xd1\x84\x20\x91\x11\x34\xd7\xc8\x76\xf3\x88\x0d\x19\xe1\x55\x93\x98\x66\xef\xdc\x67\xd4\xe1\xcb\xad\xb1\x63\xab\x91\xa8\x03\x22\x9f\x38\x7a\x7c\x50\x94\x1e\x27\x8a\x13\x17\x50\xe6\xff\xe8\x79\x3f\x0d\x0d\x24\x05\x6d\x70\xa6\x9c\x1a\x70\xce\x9a\x1b\x04\x32\x8a\x0d\x46\x4a\xd2\xb9\xda\x47\x8b\xd5\x1c\x5c\xb9\x3b\xfa\x6c\xe4\xf1\xc2\x74\x62\xf0\x71\xb4\x0b\x4c\xeb\x87\x29\x8a\xf5\x3a\x99\x77\xd5\xba\xc5\xaa\x96\x1f\x12\xed\x23\x63\xcc\x6c\x76\xa9\x49\x9d\x6c\x9c\xcc\x2f\xf3\xcf\x12\x23\x3a\x8d\x9a\xe5\x87\x6f\x57\x96\x08\x28\xf5\x18\xb3\x24\x83\x18\x24\x81\x90\x99\xfa\xca\x0b\x9a\x8f\x64\xde\xad\xf1\x14\x50\x66\x41\x45\xa0\x58\xb8\xb9\x87\xd4\x0b\x7b\x9e\xfb\xcf\xbb\xed\xdb\xad\xf9\x67\xd4\x2e\xb5\xdb\x4e\x85\xe5\x89\x2b\x4e\x73\xc7\x8b\xdb\xcc\x1b\x24\x0d\x1e\xdb\x14\x2f\x38\x50\x01\x25\x70\x08\xdc\x19\x24\x5c\xd7\x49\xb9\xaa\xcc\x0c\xef\x2a\xe1\xdf\x42\x68\xeb\xdb\x6f\x00\x00\x00\xff\xff\x91\x67\xf6\x5b\xf5\x01\x00\x00")

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

	info := bindataFileInfo{name: "assets/config.example.toml", size: 501, mode: os.FileMode(420), modTime: time.Unix(1671781026, 0)}
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
