// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/apiname.definition.toml
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

var _assetsApinameDefinitionToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x41\x4b\x33\x31\x10\x86\xef\xfb\x2b\x96\x85\xb6\x97\x36\x3b\x9b\xc2\x77\x28\x2c\x1f\x2a\xc5\x9b\x17\x7b\xb2\x14\x49\x37\xa3\x3b\xb0\xc9\xc4\x64\xb6\x22\xa5\xff\x5d\xb2\x50\x45\x6f\x9e\x12\xe6\x79\x78\x99\x79\xbd\x71\x58\xb6\x65\x75\x3b\x8c\x28\xcc\xd2\x57\x05\x7a\x1b\x98\xbc\x94\x6d\xb9\xe8\x45\x42\xda\xd4\xb5\xe3\x23\x0d\x24\x1f\xca\x04\x52\x1c\xd0\x5b\x23\xa6\x1f\x8f\xaa\x63\x57\x9f\x74\xfd\x32\x18\xa9\xbf\x32\x1e\xc5\x08\xb1\xff\x1e\xcc\x34\xdc\xf1\xe8\x65\xa6\x21\x62\xc7\xd1\xd6\x1a\x74\xb3\x6a\xf4\x6a\xdd\xec\xf4\x7a\x03\xb0\x01\x50\x00\xf0\x94\x81\xce\x40\xff\x02\xff\x07\x72\x24\xed\xaa\x99\x5b\x4a\x42\xbe\x93\x56\xe2\x88\xf3\x84\x03\x76\xd2\xba\x93\x19\x46\x5c\xe6\x87\xac\x90\xc3\xa5\x0b\x18\x89\xed\xfc\xbd\xc7\x88\x6d\xea\xd8\xa2\xc2\x37\x55\x9d\xcf\xa5\x7a\xc8\x57\x5f\x2e\xd5\x55\xca\x60\xfd\x0f\x60\x99\x4c\x27\x74\x9a\xc4\x9c\xbe\x28\x1c\x4a\xcf\x36\x37\x74\xbf\xdd\x55\x45\xb1\xdf\x07\x13\x8d\x43\xc1\x98\x0e\x87\xe2\x5a\x5f\x0e\xac\x7e\x40\x35\x2d\x94\x9d\xe9\x93\xa5\x9b\x57\x0a\xcf\x5b\xf2\x49\x90\xfc\x5f\xed\xcf\x00\x00\x00\xff\xff\xc2\x86\xde\x6c\xa9\x01\x00\x00")

func assetsApinameDefinitionTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsApinameDefinitionToml,
		"assets/apiname.definition.toml",
	)
}

func assetsApinameDefinitionToml() (*asset, error) {
	bytes, err := assetsApinameDefinitionTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/apiname.definition.toml", size: 425, mode: os.FileMode(420), modTime: time.Unix(1672155208, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsConfigExampleToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\xb1\x4e\x03\x31\x10\x44\x7b\x7f\xc5\x28\x69\x0f\x7f\x00\x52\xba\x20\x01\x02\x52\x04\xd1\x20\x8a\xcd\xdd\xfa\xb2\x92\xe3\x3d\xd9\x7b\x09\xe1\xeb\xd1\xf9\x92\x14\x29\x2d\xcf\xbc\x7d\xb3\xc4\xe7\x5e\x0a\xa4\xc0\xf6\x0c\xfe\xa5\xc3\x10\x19\xad\xa6\x20\xfd\x98\xc9\x44\x13\x82\x44\x46\xd0\x5c\x23\x9b\xf5\x33\xd6\x64\x84\x77\x4d\x62\x9a\xbd\x73\xdf\x51\xfb\x1f\xb7\xc4\x96\xad\x46\xa2\xf6\x88\x7c\xe4\xe8\xf1\x45\x51\x3a\x1c\x29\x8e\x5c\x40\x99\x1f\xd1\xf1\x6e\xec\x1b\x48\x0a\xda\xe0\x44\x39\x35\xe0\x9c\x35\x37\x08\x64\x14\x1b\x0c\x94\xa4\x75\xb5\x8f\x15\x16\x53\x70\xe1\xee\xe8\x93\x91\xc7\x1b\xd3\x91\xc1\x87\xc1\xce\x30\xad\x1f\xa6\x28\xd6\xe9\x68\xde\x55\xeb\x15\x16\xb5\xfc\x94\x68\x17\x19\x43\x66\xb3\x73\x4d\xea\x68\xc3\x68\x7e\x9e\x7f\x92\x18\xd1\x6a\xd4\x2c\x7f\x7c\xbb\x32\x47\x40\xa9\xc3\x90\x25\x19\xc4\x20\x09\x84\xcc\xd4\x55\x5e\xd0\x7c\x20\xf3\x6e\x89\x97\x80\x32\x09\x2a\x02\xc5\xc2\xcd\x3d\xa4\x5e\xd8\xf1\xd4\x7f\xdd\x6e\x3e\x6e\xcd\x8b\xd1\x0a\x96\x47\xbe\xcd\xd4\xfa\xac\x34\xcd\x2d\xcf\x6a\x13\xae\x97\xd4\x7b\x6c\x52\x3c\x63\x4f\x05\x94\xc0\x21\x70\x6b\x90\x70\x1d\x27\xe5\x6a\x32\x31\xbc\xab\x84\x87\x99\x70\x39\xf3\x1f\x00\x00\xff\xff\x2d\x95\xc2\xb0\xf5\x01\x00\x00")

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

	info := bindataFileInfo{name: "assets/config.example.toml", size: 501, mode: os.FileMode(420), modTime: time.Unix(1672153940, 0)}
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
	"assets/apiname.definition.toml": assetsApinameDefinitionToml,
	"assets/config.example.toml":     assetsConfigExampleToml,
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
		"apiname.definition.toml": &bintree{assetsApinameDefinitionToml, map[string]*bintree{}},
		"config.example.toml":     &bintree{assetsConfigExampleToml, map[string]*bintree{}},
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
