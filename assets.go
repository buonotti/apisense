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

	info := bindataFileInfo{name: "assets/.env.example", size: 66, mode: os.FileMode(420), modTime: time.Unix(1673507198, 0)}
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

var _assetsConfigExampleToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x4d\x6b\xe4\x46\x13\xbe\xf7\xaf\x28\xa4\xcb\xfb\x86\x99\x1e\xf0\x12\x12\x16\x04\x59\xec\x4d\xd6\x21\xf6\x84\xd8\xc9\xc5\x2c\xa1\xad\x2e\x8d\x8a\xb4\xba\x44\x77\xc9\xda\xc9\x21\xbf\x3d\x74\x4b\x9a\x95\x3f\xb2\x38\x61\x2e\xa3\x52\xd5\x53\xcf\x53\x5f\x2a\xe1\xb6\xa5\x08\x14\x41\x5a\x04\xfc\x64\xba\xde\x21\xd4\xec\x1b\x3a\x0c\xc1\x08\xb1\x87\x86\x1c\x42\xc3\x21\xbb\xec\x2f\x3e\xc0\x85\x11\x03\x57\xec\x49\x38\x68\xa5\xca\xb2\x84\x8b\x77\xef\xaf\xf6\xd7\x70\xbe\xbf\xfe\xfe\xf2\x87\x5f\x7f\x79\x77\x7b\xb9\xbf\x86\xb2\x2c\xd5\x9d\x35\xd8\xb1\xff\xa8\x4a\xb8\x41\xc9\x10\xe4\x05\xc3\x83\x71\x40\x1e\x22\xd6\xec\x6d\x4c\x7f\xc7\x96\xea\x36\x3b\x4c\x21\x10\x5b\x1e\x9c\x85\x30\x78\x0d\xb7\x6d\xa6\xd0\x19\x49\x5c\xeb\xc0\x1e\xfe\xd7\x8a\xf4\xf1\xed\x6e\x37\x8e\xa3\x1e\xe9\x0f\xea\xd1\x92\xd1\x1c\x0e\xbb\xf4\xb4\x3b\x0f\xec\xff\xaf\x4e\xc9\x2a\x28\xbe\x82\xf9\x57\xa8\x89\x4e\x84\xb1\x45\x69\x31\xbc\x9c\x16\x0c\x3c\x18\x47\x76\x2a\x43\x32\x0c\x7d\x72\x10\x13\x64\xe8\xd5\xfc\x0e\xb7\xec\xb7\xb3\x0d\x2a\x90\x30\xa0\x5a\xa9\xb5\xd8\x90\xa7\x8c\x60\x29\x60\x2d\x1c\x8e\xfa\x51\xd5\x4f\xe6\xc4\x26\x60\xb6\x99\x9e\x56\x91\x11\x4c\x40\x88\xc2\x01\xad\x56\x2b\xfb\xd6\x52\x48\xca\xfe\xda\xb1\x6d\xb7\xd6\x88\xd9\x76\x53\x5f\x76\x2b\x2f\x6d\xb3\xe0\x8b\xcf\x4c\x52\x4b\x53\x72\x23\x93\x1a\x18\x49\x52\xed\x29\x42\x1f\xb0\xa1\x4f\x39\xa1\x67\x01\xf2\xb5\x1b\x2c\x5a\x45\x07\xcf\x01\xb7\xf3\xeb\x0a\x8a\xdf\x8b\xb5\xcc\x80\x3d\x07\x01\x1e\xa4\x1f\xe4\xdf\x28\x9d\x02\x1f\x2b\x9c\x6d\x5f\x50\x37\x7b\x14\x6a\x19\x30\xed\xf8\xb0\x1e\x32\xc7\x07\x70\xf8\x80\x4e\xc3\x6f\xa9\x4f\xa9\x93\x03\xe6\x34\x6f\xc1\xe2\xfd\x70\xd8\x00\xf9\x86\x37\x30\x9a\xe0\x37\x80\x21\x70\xd8\x40\x63\xc4\xb8\x0d\xf4\xc6\x53\xad\x72\x7c\x22\x90\x1c\x1f\xa9\x4d\xe8\xa9\x86\x1a\x7e\x42\xf3\x80\x80\x5d\x2f\x47\x10\xce\x2f\x84\x21\x8a\xe5\x41\xb4\x2a\xf3\xee\x54\x50\xec\x5a\xee\x70\x67\x86\x80\x8e\x78\xa7\xa7\x05\x7b\x2e\xeb\xa9\x21\xa9\x2a\xd4\x02\x92\x29\xbc\xf7\xe6\xde\x61\x6a\x93\xc8\x31\xe7\x9b\x6a\x3e\x57\x7a\x24\xe7\xa0\x66\xc7\x81\xfe\xc4\x13\xd7\xb9\x2d\xc6\x5b\xe8\x03\x79\x01\x4a\x9d\x05\x03\x01\x8d\xcd\x78\xd3\x72\x69\x55\xc2\x65\x03\x31\xc9\x64\x68\x8c\x8b\xb8\x79\x0a\x92\x33\xdc\xa7\x35\x86\x1f\x6f\xf6\xd7\xa7\xc8\x99\xd1\xe7\x0d\x38\x4f\x2c\x96\xa8\x34\x04\x79\x82\x2c\x50\x93\xe0\x0e\xe4\x73\xa9\x92\xb8\xa5\xb2\x9c\x63\xb3\x95\x43\x8d\x93\x8e\xc5\x59\xc3\xde\xbb\x23\xb4\x26\x82\xf1\x80\x4d\x83\xb5\x24\xac\x39\x2f\xc5\x85\x76\xc2\xd0\x2a\x23\x6c\x27\x84\x85\xd3\x32\x2b\x9e\x85\x1a\xaa\xf3\x5e\x2f\x43\xf3\xcf\x97\x20\xa2\xb7\xb0\x0e\x89\x5a\x61\xee\x82\x85\x6a\x2a\x92\xfa\x0f\x20\xc0\x49\x0e\x35\xeb\xf5\xa9\xd9\x8b\x21\x1f\xa7\x71\x8c\x5a\x25\x9f\x74\x5f\xf2\xf3\x0b\xd7\x05\x3b\x43\x0e\x8c\xb5\x01\x63\xcc\x93\xf7\x3c\x8f\xb0\x56\x93\x5f\x05\x45\x87\xdf\xcd\x57\x5e\xd7\xdc\x3d\x1a\xea\x9b\xab\xdb\x9f\x21\x62\x78\x48\x02\x18\x86\x38\x5d\xfd\x84\x98\x7a\xf5\xa4\x02\xb1\x93\x7e\x3b\x7b\x57\x50\xa4\x47\xfd\x45\xe4\x2c\xf0\x95\xb8\xd9\xb7\x82\xaf\xbf\xfd\xe6\x19\xcc\x10\x31\x78\xd3\xe1\x6b\xa1\x4e\xfe\x15\x14\x33\xbf\x17\xb8\x99\x18\x47\x0e\xf6\xd5\xfc\x16\xff\x0a\x8a\xe5\x7f\x31\x7d\x09\x6f\x6e\x3e\xbc\xf4\x19\x8c\xb1\x5d\x9f\xa7\xe4\xd5\x72\x94\x67\x07\xc4\x38\xc7\x23\x18\x7f\xcc\xaf\x93\xa5\x66\xef\xb1\x16\x95\x9f\x97\x2b\xb0\xc6\x49\xc5\xd2\x6a\x2e\xd9\xd9\x9b\xb3\x37\x67\x4a\xdd\xc9\x40\xeb\x7c\x69\xcb\x46\x23\x75\x1a\xce\xe8\x10\x7b\x10\xea\x50\x29\x75\x67\x7a\xfa\xf8\x77\x00\x00\x00\xff\xff\x29\xba\x58\x35\x06\x08\x00\x00")

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

	info := bindataFileInfo{name: "assets/config.example.toml", size: 2054, mode: os.FileMode(420), modTime: time.Unix(1673446088, 0)}
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
