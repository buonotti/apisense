// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/.env.example
// assets/bluetooth.definition.toml
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

var _assetsEnvExample = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\x77\xf1\x8d\x77\xf5\x75\xf4\xf4\x89\x0f\x0d\x76\x0d\xb2\xcd\x4d\x75\x48\xad\x48\xcc\x2d\xc8\x49\xd5\x4b\xce\xcf\xe5\x42\xc8\x06\x38\x06\x07\xdb\x16\x97\x16\xa4\x16\xe9\x16\xa7\x26\x17\xa5\x96\xe8\x15\x24\x16\x17\x97\xe7\x17\xa5\x00\x02\x00\x00\xff\xff\x4d\x5f\x22\xbe\x42\x00\x00\x00")

func assetsEnvExampleBytes() ([]byte, error) {
	return bindataRead(
		_assetsEnvExample,
		"assets/.env.example",
	)
}

func assetsEnvExample() (*asset, error) {
	bytes, err := assetsEnvExampleBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/.env.example", size: 66, mode: os.FileMode(420), modTime: time.Unix(1673517048, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsBluetoothDefinitionToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x5b\x6b\xe4\x46\x13\x7d\xd7\xaf\x28\x34\x2c\x63\xc3\x58\x23\xcf\xc2\xf7\x30\x78\x3e\xd8\x84\xdd\x3c\x24\x2c\x81\x35\x1b\x88\x31\xa6\x47\x2a\x8d\x7a\xd3\xdd\x25\x77\x97\xe6\x92\xc5\xff\x3d\x54\x4b\x9a\x9b\x6d\xd8\x2c\x49\x1e\x8c\xd5\xdd\x55\xe7\xd4\xe5\x54\xcd\x08\x6e\x6b\x04\xa7\x2c\x02\x55\xc0\x35\x02\xba\xb2\x21\xed\x18\x2e\xda\x80\x25\x68\x17\x6f\x3d\x36\xe4\xf9\x32\x89\x96\x0b\x48\x7f\x30\x2d\x32\x11\xd7\x69\x92\x74\x18\x7b\x3f\x26\x28\x94\x31\x19\x7c\x56\x5e\xab\xa5\xc1\x00\x85\x72\xb0\x44\xd0\x8e\xd1\x37\x64\x14\x63\x09\x1b\xcd\x35\x7c\xfd\x0a\xd9\xaf\xca\x2b\xfb\x51\x70\x9f\x9e\xe0\xe2\x27\x02\x46\xdb\x88\x0d\x84\x9d\x63\xb5\xbd\xec\x08\x3c\x82\x92\x3f\x13\x08\x78\x43\xb0\x6c\xb5\x61\xed\xa0\x6a\x5d\xc1\x9a\x5c\x98\xc3\x47\xda\x80\x72\x25\xbc\x77\xeb\x64\x14\x4f\x3a\x40\xcc\x22\x92\x49\x1e\x15\x19\x43\x1b\xed\x56\x3d\xf8\x3c\x86\x20\xa6\xe9\x4d\x45\xde\x2a\xfe\x7f\x2a\x71\xec\xa8\x8d\x51\x57\xda\x95\xbd\xa3\x3c\x82\xc4\x31\x87\x9a\xb9\x09\xf3\xe9\xb4\xf9\x63\x95\xad\x28\x2b\x71\x3d\x65\x6d\x71\x74\xab\x2d\x66\x1f\xa2\x65\x32\x92\x30\xbe\x35\x00\x31\x4d\x6f\xd0\xad\x1f\xd6\xca\x3f\x48\x91\x63\x18\xc9\x52\x05\xbc\x6a\xbd\x81\x05\x8c\x07\x52\x4b\x4b\x6d\x34\xef\x32\xd5\xe8\x8c\x1a\x74\xa5\x62\x55\xb7\xcb\xac\x20\x3b\x5d\x0b\xda\xbb\x46\x7f\x46\x1f\x34\x39\x78\x7a\x9a\x56\x46\xf1\x74\xdf\xaf\x4f\xac\xa4\x5a\x87\x8b\x37\xb3\xfc\x47\x6a\x1d\xbf\x99\xe5\x1e\x0b\xf2\xe5\x74\x96\xcf\xae\xaf\xae\x67\x57\x6f\xaf\x6f\x67\x6f\xe7\x79\x3e\xcf\xf3\x2c\xcf\xf3\xdf\xe5\x61\x26\x0f\xb3\xb3\x87\xb1\x68\xe0\x17\x1d\x58\x34\x24\xc1\x07\xf9\x58\x2b\xa3\x4b\xc5\xe4\x03\x70\xad\x18\x42\x4d\xad\x29\xc1\x11\x8b\x16\x7c\xeb\x32\xf8\x19\xb1\x11\x85\xd9\xae\xca\x8a\x01\xb7\x8c\xde\x29\x73\xec\x2d\x6d\x17\xd4\xbe\x8a\x83\x49\x76\x13\xeb\x94\x8c\x44\x71\x4b\x64\x46\x0f\x61\x67\x97\x64\xf4\x9f\xd8\xa1\x71\x8d\xbb\xe8\x3e\xf8\x24\xb8\x2d\x4c\x5b\x62\x79\x75\x84\xbf\x80\xbb\x04\x00\x20\xf5\xca\xad\x30\x4d\xee\x07\x49\x3f\xb6\xe8\x77\xd0\x88\x3e\x91\x51\xf2\x20\x68\x54\x88\xff\x8f\x47\xe5\x1f\x94\xfa\xdd\x5d\x24\xbd\xbf\x4f\x9e\x4f\xe6\x59\x38\xfb\x59\x34\xda\x6a\x4e\x7b\x87\xb5\x32\xed\xab\x1e\xdd\xe3\x02\xd2\xab\xeb\x34\x39\xe2\x1a\x90\x4a\x1d\x58\xbb\x82\xd3\x83\x25\xfb\x16\x5f\xb4\x0d\x68\xf0\xc4\xd2\xc6\xaf\x89\x8d\x95\x95\x71\x98\xd8\x06\xbd\xa6\xf2\x45\xf7\x8d\x4c\xd2\xc1\x7b\x1c\x0a\x2a\x31\xc3\xc7\x2c\x8d\x03\xd9\xd5\x28\x1d\x20\xe4\xe1\xed\xff\xf2\x7c\x12\x54\xc1\x7a\x1d\x0d\x25\xb2\xf1\xd0\xaa\x7e\x3c\xfb\xbc\x3d\x86\x86\x5c\x40\xb8\xf8\x12\xc8\x01\x79\xd8\x5a\x73\x99\xf4\x46\x0b\x48\xe5\x3a\x6e\xae\xa1\x71\x50\x62\xa5\x9d\x8e\x8b\x24\x19\xc1\x07\xf2\xa0\x8c\x01\x47\x0e\x0a\x72\x81\x95\x63\x58\x0f\x4d\x9e\x44\x12\xd7\xda\x25\xfa\x5e\xe9\x2d\x06\xb0\x6d\x88\xd2\x96\xc7\x20\x09\xc8\x2e\xda\x68\x63\xe4\x72\x58\xa6\xe4\x4b\xf4\xc3\x8e\x3c\x5b\xb8\xa1\x5d\x7e\xc1\xa2\xdb\xa0\xb5\x68\x51\xf4\x30\xb0\xbe\x28\x89\xe7\x62\x90\xca\x89\x16\x7e\xab\x91\x6b\x61\x3a\xb6\x92\x75\xb4\x4f\x87\xbc\xcc\x62\xb2\x3f\x2f\xa0\x52\x26\x60\xcf\xd2\x50\x08\x5a\x0a\xd3\x27\xf7\x8c\xb0\xbf\xdf\x0f\xcf\xbb\x95\x6e\x1e\xde\x6b\x17\x18\xb5\x4b\x27\xf1\x72\xac\x94\x1a\xcb\x3c\x9d\xa4\x31\x84\x7a\x58\x53\xe9\x71\x14\xd2\xd7\x67\xe8\xb3\x7e\x2c\x3f\x15\x35\x5a\x75\xd4\xad\xa3\x96\xb7\x86\x93\xbb\xbb\xee\x23\x43\xc7\xaf\x8c\x51\xa5\xd1\x94\xfb\x20\xa8\xaa\x02\xee\xa7\x87\x77\xcd\xa9\x1d\x5c\x04\xf6\xda\xad\x26\x71\xa2\x57\xe8\x27\x50\x19\x52\x3c\x81\x25\x91\x41\xe5\x26\x40\xb1\x69\x97\x49\xf4\x5d\x40\xda\x1b\x0e\x90\x56\x3b\x6d\x5b\x7b\x3a\x98\x3d\x36\x39\xb3\x13\xe9\x0e\xe0\x51\x31\x11\xff\x12\x02\x46\x21\x38\x72\x08\xba\x8a\x6b\x53\x35\x8d\xd1\x85\x54\x31\xb1\xda\x09\x97\xbc\xee\x89\xd4\xf6\x7b\x88\x12\xab\xb6\xc7\x50\xc7\xc2\xe9\xbc\x75\x00\x8f\x8f\xad\xf6\x58\x0e\xa2\xd9\x9f\xfb\x76\x75\x11\x14\xb5\x36\xa5\x47\xf7\x1a\xf9\x50\xa9\x78\x1f\x9b\x1b\xa5\x71\xd6\xb1\xb3\xb5\xf6\xac\xac\x27\xa9\x9f\x04\x7f\x1e\xd5\xb7\xf1\xc8\xcf\xe7\x81\x46\x79\xaf\x76\x7f\x87\xe4\x0c\x38\xeb\x48\x8f\xf0\xf7\x5b\xf0\xdf\xc9\xe4\x05\xc2\xa8\x80\xff\x96\xaf\xdb\xf8\x07\xce\x6e\x6a\xbe\x97\xf2\xaf\x00\x00\x00\xff\xff\x0c\xd6\xf0\xff\x96\x0a\x00\x00")

func assetsBluetoothDefinitionTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsBluetoothDefinitionToml,
		"assets/bluetooth.definition.toml",
	)
}

func assetsBluetoothDefinitionToml() (*asset, error) {
	bytes, err := assetsBluetoothDefinitionTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/bluetooth.definition.toml", size: 2710, mode: os.FileMode(420), modTime: time.Unix(1673360661, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsConfigExampleToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x51\x6f\xe4\x34\x10\x7e\xf7\xaf\x18\x25\x2f\x80\xb6\x5e\xa9\x27\x04\x3a\x29\x12\xa7\xf6\xe0\x8a\x68\x17\xb1\x85\x97\xea\x84\xdc\x64\xb2\x19\xe1\x78\x22\x7b\xb2\xb9\xf2\xc0\x6f\x47\xe3\x24\xab\xe5\x5a\x50\x51\x5e\x62\xfb\x9b\x6f\xbe\x19\x7f\x9e\x12\xee\x3b\x4a\x40\x09\xa4\x43\xc0\x4f\xae\x1f\x3c\x42\xcd\xa1\xa5\xc3\x18\x9d\x10\x07\x68\xc9\x23\xb4\x1c\x33\x64\x77\xfd\x01\xae\x9d\x38\xb8\xe5\x40\xc2\xd1\x1a\x53\x96\x25\x5c\xbf\x7b\x7f\xbb\xbb\x83\xab\xdd\xdd\xf7\x37\x3f\xfc\xfa\xcb\xbb\xfb\x9b\xdd\x1d\x94\x65\x69\x1e\x1a\x87\x3d\x87\x8f\xa6\x84\x3d\x4a\xa6\xa0\x20\x18\x8f\xce\x03\x05\x48\x58\x73\x68\x92\xfe\x4e\x1d\xd5\x5d\x06\xcc\x21\x90\x3a\x1e\x7d\x03\x71\x0c\x16\xee\xbb\x2c\xa1\x77\xa2\x5a\xeb\xc8\x01\xbe\xe8\x44\x86\xf4\x76\xbb\x9d\xa6\xc9\x4e\xf4\x07\x0d\xd8\x90\xb3\x1c\x0f\x5b\x5d\x6d\xaf\x22\x87\x2f\xcd\x29\x59\x05\xc5\x57\xb0\x7c\x85\x99\xe5\x24\x98\x3a\x94\x0e\xe3\xcb\x69\xc1\xc1\xd1\x79\x6a\xe6\x36\xe8\xc6\x38\x28\x40\x5c\x94\x71\x30\xcb\x19\x5e\x70\xb8\x58\xf6\xa0\x82\xd6\xf9\x84\xca\x7f\x8d\x2d\x05\x3a\x75\x50\x3b\xec\x64\x0e\x86\x89\x44\x4b\xa5\x04\x43\xc4\x96\x3e\x81\x8b\x08\x81\x05\x28\xd4\x7e\x6c\xb0\x31\x74\x08\x1c\xf1\x62\x39\xae\xa0\xf8\xbd\x30\x6b\x33\xad\xe7\xc3\x79\x43\x3d\x1f\xc0\xe3\x11\xbd\x85\xdf\x54\x93\xaa\x1e\x31\x29\xe9\x5b\x68\xf0\x71\x3c\x6c\x80\x42\xcb\x1b\x98\x5c\x0c\x1b\xc0\x18\x39\x6e\xa0\x75\xe2\xfc\x06\x06\x17\xa8\x36\x39\x5e\xf3\x28\x70\xed\xcf\x89\x5d\x0b\xb0\xf0\x13\xba\x23\x02\xf6\x83\x3c\x81\x70\x3e\x10\x86\x24\x0d\x8f\x62\x4d\xb6\x49\x05\xc5\x5f\x5b\x37\x50\xc2\x90\xf0\xf4\xa3\x82\x0b\x53\xae\x88\x4c\xff\x3e\xb8\x47\x8f\x5a\xbf\xc8\x53\xe6\xe2\x51\x86\x51\xec\xec\xc8\x89\xbc\x87\x9a\x3d\x47\xfa\x13\x4f\x3a\x66\x08\xb8\xd0\xc0\x10\x29\x08\x90\xb6\x0c\x1c\x44\x74\x4d\xe6\x9b\x4d\x62\x4d\x09\x37\x2d\x24\x2d\x81\xe7\x2b\xd9\x7c\x4e\x92\x33\x3c\xaa\x1d\xe1\xc7\xfd\xee\xee\x14\xb9\x28\xaa\x40\xe2\x98\x2f\xf2\x4a\x55\xac\x51\xfa\x58\xf2\xd5\x34\x40\xad\xd2\x1d\x28\xe4\x36\x68\x71\x6b\xd7\x38\xc7\xe6\x5d\x8e\x35\xce\x75\xac\x60\x0b\xbb\xe0\x9f\xa0\x73\x09\x5c\x00\x6c\x5b\xac\x45\xb9\x96\xbc\x94\x56\xd9\xca\x61\x4d\x66\xb8\x98\x19\x56\x4d\xab\x0f\x02\x0b\xb5\x54\x67\x7f\xae\x86\xf8\x77\x47\x27\x0c\x0d\x9c\x87\x24\x6b\x30\xdf\x42\x73\xee\xdb\xff\x4b\x02\xac\xe5\x50\x9b\xc1\x11\x07\x8e\xa2\x03\x44\x1c\x85\x34\x5b\x2d\x59\xa3\x18\x7d\x27\x79\x7d\xd6\xdb\xd5\x64\xd8\x3b\xf2\xe0\x9a\x26\x62\x4a\xd9\x55\xcf\xf3\x08\x5b\x33\xe3\x2a\x28\x7a\xfc\x6e\x99\x56\xb6\xe6\xfe\x1f\x86\xdd\xdf\xde\xff\x0c\x09\xe3\x51\x0b\x60\x18\xd3\x3c\xbd\x94\x51\xef\xea\xb3\x0e\xa4\x5e\x86\x8b\x05\x5d\x41\xa1\x4b\xfb\x9f\xcc\xb9\xc0\x57\xf2\x66\x6c\x05\x5f\x7f\xfb\xcd\x33\x9a\x31\x61\x0c\xae\xc7\xd7\x52\x9d\xf0\x15\x14\x8b\xbe\x17\xb4\xb9\x94\x26\x8e\xcd\xab\xf5\xad\xf8\x0a\x8a\xf5\xbf\x98\x27\xfa\x7e\xff\xe1\xa5\x71\x9e\x52\x77\x3e\x7a\x14\xd5\x71\x92\x67\xc3\xc1\x79\xcf\x13\xb8\xf0\x94\x8f\x75\xa7\xe6\x10\xb0\x16\x93\xd7\xeb\x14\x38\xe7\xd1\x66\x59\xb3\xb4\xec\xf2\xcd\xe5\x9b\x4b\x63\x1e\x64\xa4\xf3\x7c\xfa\xca\x26\x27\xb5\x9a\x33\x79\xc4\x01\x84\x7a\x34\xc6\x3c\xb8\x81\x3e\xfe\x1d\x00\x00\xff\xff\xb4\xbc\xbd\x17\xce\x06\x00\x00")

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

	info := bindataFileInfo{name: "assets/config.example.toml", size: 1742, mode: os.FileMode(420), modTime: time.Unix(1673954273, 0)}
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
	"assets/.env.example":              assetsEnvExample,
	"assets/bluetooth.definition.toml": assetsBluetoothDefinitionToml,
	"assets/config.example.toml":       assetsConfigExampleToml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
		".env.example":              &bintree{assetsEnvExample, map[string]*bintree{}},
		"bluetooth.definition.toml": &bintree{assetsBluetoothDefinitionToml, map[string]*bintree{}},
		"config.example.toml":       &bintree{assetsConfigExampleToml, map[string]*bintree{}},
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
