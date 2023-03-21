// Code generated for package config by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/.env.example
// assets/bluetooth.toml
// assets/bluetooth2.toml
// assets/config.example.toml
package config

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

var _assetsEnvExample = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0c\xf0\x0c\x76\xf5\x0b\x76\x8d\x77\xf5\x75\xf4\xf4\x89\x0f\x0d\x76\x0d\xb2\xcd\x4d\x75\x48\xad\x48\xcc\x2d\xc8\x49\xd5\x4b\xce\xcf\xe5\x42\x53\x12\xe0\x18\x1c\x6c\x5b\x5c\x5a\x90\x5a\xa4\x5b\x9c\x9a\x5c\x94\x5a\xa2\x57\x90\x58\x5c\x5c\x9e\x5f\x94\x02\x08\x00\x00\xff\xff\x47\xd5\xa7\xc1\x4c\x00\x00\x00")

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

	info := bindataFileInfo{name: "assets/.env.example", size: 76, mode: os.FileMode(438), modTime: time.Unix(1679427268, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsBluetoothToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x5b\x8b\xdc\x46\x13\x7d\xd7\xaf\x28\x34\x98\xdd\x85\x59\x8d\x76\x0c\xdf\xc3\xe0\xf9\xc0\x09\x76\x1e\x12\x4c\xc0\xc6\x81\x2c\xcb\xd2\x23\x95\x46\xed\x74\x77\x69\xbb\x4b\x73\x89\xd9\xff\x1e\xaa\x75\x99\xcb\xee\x82\x63\x92\x3c\x0c\x23\x75\x57\xd5\x39\x55\x75\xaa\x34\x81\x4f\x35\x82\x53\x16\x81\x2a\xe0\x1a\x01\x5d\xd9\x90\x76\x0c\x97\x6d\xc0\x12\xb4\x8b\xa7\x1e\x1b\xf2\x7c\x95\x44\xcb\x25\xa4\x3f\x98\x16\x99\x88\xeb\x34\x49\xba\x18\x15\x79\xab\x78\x88\xe2\x31\x34\xe4\x02\xc2\xe5\x97\x40\x0e\xc8\xc3\xce\x9a\xab\xa4\x37\x5a\x42\x2a\xc7\xa3\xef\x88\xc9\x04\x85\x32\x26\x83\xcf\xca\x6b\xb5\x32\x18\xa0\x50\x0e\x56\x08\xda\x31\xfa\x86\x8c\x62\x2c\x61\xab\xb9\x86\xaf\x5f\x21\xfb\x55\x79\x65\x3f\x08\xa7\xc7\x47\xb8\xfc\x89\x80\xd1\x36\x62\x03\x61\xef\x58\xed\xae\x3a\x00\x8f\xa0\xe4\x67\x02\x01\x6f\x09\x56\xad\x36\xac\x1d\x54\xad\x2b\x58\x93\x0b\x0b\xf8\x40\x5b\x50\xae\x84\x77\x6e\x93\x4c\xe2\x9b\x0e\x10\x2b\x10\xc1\x38\x66\x68\x0c\x6d\xb5\x5b\xf7\xc1\x17\x91\x82\x98\xa6\x6f\xba\xc4\xfe\x9f\x0a\x8f\x3d\xb5\x91\x75\xa5\x5d\xd9\x3b\xc6\xac\x85\xc7\x02\x6a\xe6\x26\x2c\x66\xb3\xe6\x8f\x75\xb6\xa6\xac\xc4\xcd\x8c\xb5\xc5\xc9\x27\x6d\x31\x7b\x1f\x2d\x93\x89\xd0\xf8\x56\x02\x62\x9a\xbe\x41\xb7\xb9\xdf\x28\x7f\x2f\x0d\x8a\x34\x92\x95\x0a\x78\xdd\x7a\x03\x4b\xb8\x18\x40\x2d\xad\xb4\xd1\xbc\xcf\x54\xa3\x33\x6a\xd0\x95\x8a\x55\xdd\xae\xb2\x82\xec\x6c\x23\xd1\xde\x36\xfa\x33\xfa\xa0\xc9\xc1\xe3\xe3\xac\x32\x8a\x67\x63\xaf\x3f\xb2\x92\x6a\x1d\x0e\x5e\xcd\xf3\x1f\xa9\x75\xfc\x6a\x9e\x7b\x2c\xc8\x97\xb3\x79\x3e\xbf\xb9\xbe\x99\x5f\xbf\xbe\xf9\x34\x7f\xbd\xc8\xf3\x45\x9e\x67\x79\x9e\xff\x2e\x17\x73\xb9\x98\x9f\x5d\x5c\x88\x06\x7e\xd1\x21\x2a\x47\xc8\x07\x79\xd8\x28\xa3\x4b\xc5\xe4\x03\x70\xad\x18\x42\x4d\xad\x29\xc1\x11\x8b\x16\x7c\xeb\x32\xf8\x19\xb1\x11\x75\xda\xae\xca\x8a\x01\x77\x8c\xde\x29\x73\xec\x2d\x6d\x97\xa8\x7d\x15\x07\x93\xec\x4d\xac\x53\x32\x11\xc5\xad\x90\x19\x3d\x84\xbd\x5d\x91\xd1\x7f\x62\x17\x8d\x6b\xdc\x47\xf7\xc1\x27\xc1\x5d\x61\xda\x12\xcb\xeb\xa3\xf8\x4b\xb8\x4d\x00\x00\x52\xaf\xdc\x1a\xd3\xe4\x6e\x90\xf4\x43\x8b\x7e\x0f\x8d\xe8\x13\x19\x25\x0f\x82\x46\x85\xf8\x7f\x3c\x66\xff\xa0\xd4\x6f\x6f\x23\xe8\xdd\x5d\xf2\x74\xaa\xcf\xe8\x8c\x73\x6c\xb4\xd5\x9c\xf6\x0e\x1b\x65\xda\x17\x3d\xba\xcb\x25\xa4\xd7\x37\x69\x72\x84\x35\x44\x2a\x75\x60\xed\x0a\x4e\x0f\x96\xec\x5b\x7c\xd6\x36\xa0\xc1\x13\x4b\x1b\x9f\xa6\x36\x56\x56\xc6\x61\x6a\x1b\xf4\x9a\xca\x67\xdd\xb7\x32\x49\x07\xef\x8b\x50\x50\x89\x19\x3e\x64\x69\x1c\xc8\xae\x46\xe9\x10\x42\x2e\x5e\xff\x2f\xcf\xa7\x41\x15\xac\x37\xd1\x50\x98\x45\xe5\x0d\xc5\x87\x12\x2b\xed\x74\x5c\x06\xc9\x04\xde\x93\x07\x65\x0c\x38\x72\x50\x90\x0b\xac\x1c\xc3\x66\x68\xd4\x34\x16\xc8\xb5\x76\x85\xbe\x57\x6b\x8b\x01\x6c\x1b\xa2\x3c\xe5\x32\x08\x09\xd9\x27\x5b\x6d\x8c\x1c\x0e\xcb\x94\x7c\x89\x7e\xd8\x73\x67\x0b\x37\xb4\xab\x2f\x58\x74\x5b\xb0\x16\x3d\x49\x4f\x07\xd4\x67\xdb\xfa\xb4\xa1\x92\xbd\xf4\xf3\xb7\x1a\xb9\x16\xa4\x63\x2b\x59\x29\x63\x3a\xe4\x65\x9e\x92\xf1\x7d\x09\x95\x32\x01\x7b\x94\x86\x42\xd0\x52\x98\x3e\xb9\x27\x80\xfd\xf9\x38\x00\x6f\xd7\xba\xb9\x7f\xa7\x5d\x60\xd4\x2e\x9d\xc6\xc3\x0b\xa5\xd4\x85\xcc\xc4\x49\x1a\x03\xd5\xc3\xaa\x49\x8f\x59\x48\x6f\x9e\x44\x9f\xf7\xa3\xf5\xb1\xa8\xd1\xaa\xa3\x6e\x1d\x7d\x70\x5a\xc3\xc9\xed\x6d\xf7\x90\xa1\xe3\x17\x46\xa1\xd2\x68\xca\x91\x04\x55\x55\xc0\x71\x02\x78\xdf\x9c\xda\xc1\x65\x60\xaf\xdd\x7a\x1a\xa7\x72\x8d\x7e\x0a\x95\x21\xc5\x53\x58\x11\x19\x54\x6e\x0a\x14\x9b\x76\x95\x44\xdf\x25\xa4\xbd\xe1\x10\xd2\x6a\xa7\x6d\x6b\x4f\x87\xab\x8f\x4d\xce\xec\xe5\xeb\x30\x04\x8f\x8a\x89\xf1\xaf\x20\x60\x14\x82\x23\x87\xa0\xab\xb8\xfa\x54\xd3\x18\x5d\x48\x15\x13\xab\x9d\x60\xc9\xed\x08\xa4\x76\xdf\x03\x94\x58\xb5\x3b\x0e\x75\x2c\x9c\xce\x5b\x07\xf0\xf8\xd0\x6a\x8f\xe5\x20\x9a\xf1\xbd\x6f\x57\xc7\xa0\xa8\xb5\x29\x3d\xba\x97\xc0\x87\x4a\xc5\xf3\xd8\xdc\x28\x8d\xb3\x8e\x9d\xad\xa6\x27\x65\x3d\x49\xfd\x84\xfc\x39\xab\x6f\xc3\x91\x4f\xe0\x01\x46\x79\xaf\xf6\x7f\x07\xe4\x2c\x70\xd6\x81\x1e\xc5\x1f\x37\xd9\xbf\x93\xc9\x33\x80\x51\x01\xff\x2d\x5e\xb7\xb5\x0f\x98\xdd\xd4\x7c\x2f\xe4\x5f\x01\x00\x00\xff\xff\x06\x04\xea\x60\x96\x0a\x00\x00")

func assetsBluetoothTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsBluetoothToml,
		"assets/bluetooth.toml",
	)
}

func assetsBluetoothToml() (*asset, error) {
	bytes, err := assetsBluetoothTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/bluetooth.toml", size: 2710, mode: os.FileMode(438), modTime: time.Unix(1679427268, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsBluetooth2Toml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xcf\x6e\xdb\x30\x0c\xc6\xef\x7a\x0a\x41\x40\xe1\x4b\x62\x2b\x2e\xb0\x43\x80\x1c\xd6\x61\x3b\xee\xb2\x9d\x16\x04\x05\x63\xd1\x09\x07\xfd\x71\x24\x2a\x6b\x5a\xf4\xdd\x07\x39\x6d\x93\x16\xd9\x61\x03\xba\x9b\x69\x7e\xfc\x7e\xa4\x48\x0f\x0e\xe5\x42\xaa\x1b\x9b\x91\x43\xe0\x6d\xab\xc4\x1a\x12\x4e\x73\xb4\x72\x21\xab\x2d\xf3\x90\xe6\x4d\xe3\xc2\x9a\x2c\xf1\xa1\x86\x81\xea\x30\xa0\x37\xc0\xb0\xcd\xeb\xba\x0b\xae\xd9\xb7\x4d\x6f\x81\x9b\x17\x93\x6f\x0c\x4c\xc1\x9f\x7e\x5c\xb5\xfa\x53\xc8\x9e\xaf\x5a\x1d\xb1\x0b\xd1\x34\xad\x6e\x67\xd3\x59\x3b\xbd\x9e\x7d\x6f\xaf\xe7\x5a\xcf\xb5\xae\xb5\xd6\x3f\x4a\xa2\x2d\x89\xf6\x4d\xa2\x12\x78\xd7\xd9\x6c\xd0\x4c\xf7\x60\xc9\x00\x87\x98\xe4\x42\x2e\x57\xa2\x0f\xd1\x01\x97\x31\x7e\xa6\xe0\x95\x10\xcb\xe5\x2e\x63\x3c\xac\x56\xe2\x79\x3e\x4b\x8e\x58\x89\x3d\xd8\x3c\xc6\xd3\xd9\x45\x99\xa1\xc4\xe4\xbb\x73\x25\xc7\x8c\x17\xb5\x09\x2d\xbe\x52\xba\xf1\x6b\xe2\xc6\xf6\x98\x1c\x4e\xdc\x80\x91\x82\xb9\x58\xfe\x6b\x8b\x11\x4f\xd5\x55\xea\x82\xc1\x1a\x77\xb5\x7a\x78\x90\xf5\xd7\xa2\x7a\x7c\x54\xcf\x16\x25\x71\xfd\x41\xeb\x49\x82\x8e\x69\x3f\x0a\x4b\x67\x55\xb1\xde\x43\x24\x58\x5b\x3c\x73\x2f\xf5\x4a\x74\xc1\x27\x06\x5f\x9e\xa6\x07\x9b\xf0\x48\x1b\x5f\x4d\x48\x29\xa5\xfa\xb8\xa1\xe1\xf6\x33\xf9\xc4\x48\x5e\x4d\xc6\x9f\xd5\x4d\x0e\xf7\xf7\x74\xfb\x85\x30\x42\x25\x56\x85\x10\x31\x65\xcb\x35\x7a\x7e\x35\x43\xe8\xfb\x84\xac\x04\x1f\x86\x31\x26\xcf\xb8\xc1\xa8\x84\x23\x5f\x62\x1f\x3c\x2a\xe1\xe0\xee\x14\x44\xdc\x65\x8a\x68\xe4\x42\x96\xfe\x45\x4f\x68\xcd\xd3\x1e\xff\x0c\x7a\x5a\xdf\xbb\x73\xca\x55\x9f\x30\x10\x23\x1c\xfe\x06\xf2\xc6\xb8\x3e\x42\xcf\xfc\x5f\x0e\xe2\x7d\x26\xb9\x00\x1c\x37\xfe\x7f\x79\xc7\xe3\x3f\x31\x13\x47\xf2\x9b\x7f\x45\xfe\x0e\x00\x00\xff\xff\xeb\xb4\xc4\xeb\xa2\x04\x00\x00")

func assetsBluetooth2TomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsBluetooth2Toml,
		"assets/bluetooth2.toml",
	)
}

func assetsBluetooth2Toml() (*asset, error) {
	bytes, err := assetsBluetooth2TomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/bluetooth2.toml", size: 1186, mode: os.FileMode(438), modTime: time.Unix(1679427268, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsConfigExampleToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x4d\x6f\xe4\x36\x0c\xbd\xeb\x57\x10\xf6\xa5\x2d\xc6\x9a\x7c\xa0\x68\xb0\x80\x81\x2e\x92\x6d\x37\x45\x93\x29\x36\xd9\x5e\x82\x41\xa0\x58\xf4\x58\x58\x59\x32\x24\x3a\x4e\x7a\xe8\x6f\x2f\x28\x7f\x64\x9a\x9d\x2e\x1a\xf8\x22\x51\x8f\x4f\x8f\x14\x49\xe7\x70\xdb\x98\x08\x26\x02\x35\x08\xf8\xa4\xda\xce\x22\x54\xde\xd5\x66\xd7\x07\x45\xc6\x3b\xa8\x8d\x45\xa8\x7d\x48\x90\xcd\xc5\x47\xb8\x50\xa4\xe0\xca\x3b\x43\x3e\x48\x21\xf2\x3c\x87\x8b\xf7\x1f\xae\x36\xd7\x70\xbe\xb9\xfe\xe5\xf2\xd7\xcf\x9f\xde\xdf\x5e\x6e\xae\x21\xcf\x73\x71\xa7\x15\xb6\xde\x6d\x45\x0e\x37\x48\x89\xc2\x38\xc2\xf0\xa8\x2c\x18\x07\x11\x2b\xef\x74\xe4\xe5\xd0\x98\xaa\x49\x80\xd1\x05\x62\xe3\x7b\xab\x21\xf4\x4e\xc2\x6d\x93\x24\xb4\x8a\x58\x6b\x15\xbc\x83\xef\x1a\xa2\x2e\xbe\x5b\xaf\x87\x61\x90\x83\xf9\x62\x3a\xd4\x46\x49\x1f\x76\x6b\xde\xad\xcf\x83\x77\xdf\x8b\xe5\xb2\x12\xb2\x1f\x60\xfa\x32\x31\xca\x89\x30\x34\x48\x0d\x86\xc3\xd7\x82\x82\x47\x65\x8d\x1e\xd3\xc0\x86\xbe\x63\x00\xa9\x40\x7d\x27\xa6\x33\x2c\xbc\x2b\x26\x1b\x94\x50\x2b\x1b\x91\xf9\x2f\xb0\x36\xce\x2c\x19\xe4\x0c\x2b\x1a\x9d\x61\x30\xc4\xa1\x9a\x08\x5d\xc0\xda\x3c\x81\x0a\x08\xce\x13\x18\x57\xd9\x5e\xa3\x16\x66\xe7\x7c\xc0\x62\x3a\x2e\x21\xbb\xcf\xc4\x9c\x4c\x69\xfd\x6e\x3f\xa1\xd6\xef\xc0\xe2\x23\x5a\x09\x7f\xb2\x26\x56\xdd\x63\x64\xd2\x77\xa0\xf1\xa1\xdf\xad\xc0\xb8\xda\xaf\x60\x50\xc1\xad\x00\x43\xf0\x61\x05\xb5\x22\x65\x57\xd0\x29\x67\x2a\x91\xfc\xf9\x1e\x06\xce\xf9\x59\xd8\x39\x00\x09\xbf\xa3\x7a\x44\xc0\xb6\xa3\x67\x20\x9f\x0e\xc8\x43\x24\xed\x7b\x92\x22\x4f\x75\x52\x42\xf6\xf7\x5a\x75\x26\xa2\x8b\xb8\x2c\x58\x71\x26\x66\x40\xa2\xff\xe0\xd4\x83\x45\x8e\x9f\xe8\x39\x71\xf9\x9e\xba\x9e\xe4\x58\x91\x83\xb1\x16\x2a\x6f\x7d\x30\x7f\xe1\xa2\x63\x84\x80\x72\x1a\xba\x60\x1c\x81\xe1\x94\x81\x82\x80\x4a\x27\xbe\xb1\x48\xa4\xc8\xe1\xb2\x86\xc8\x21\xf8\xf1\x49\x56\xaf\x49\xd2\x0d\x0f\x5c\x8e\xf0\xdb\xcd\xe6\x7a\xf1\x9c\x14\x95\x40\xa1\x4f\x0f\x79\xce\x2a\x66\x2f\x6e\x96\xf4\x34\x1a\x4c\xcd\x74\x3b\xe3\x52\x1a\x38\xb8\x39\x6b\x3e\xf9\x26\xab\x0f\x15\x8e\x71\xcc\x60\x09\x1b\x67\x9f\xa1\x51\x11\x94\x03\xac\x6b\xac\x88\xb9\xa6\x7b\x4d\x9c\x65\x33\x87\x14\x89\xa1\x18\x19\x5e\x34\xcd\x8f\xa3\x4d\xac\x54\xd0\x40\xa6\xc5\x08\xbe\x86\x80\x9d\x0f\x14\x25\x7c\x1a\x17\xe0\xad\x4e\xf5\xad\xdc\x58\x70\x73\xd8\x93\x27\x6a\xb9\x94\xd5\x64\xda\xbe\xbc\xce\x64\xe1\x10\x5f\xb8\x05\xa6\x43\xfd\x0d\x39\x73\xd7\xbd\x49\xc6\x7e\xab\x9e\x1e\xe9\xbd\x7a\x77\x9e\x4c\x6d\xaa\xd4\x87\x73\xe1\xff\x77\xe7\x46\x74\x1a\xf6\x5d\xa2\xdc\x93\xbc\xf4\xe7\x5b\x49\xc0\xf3\xb3\x99\x3a\x81\xc7\x54\xf0\xa0\x24\x65\x5c\x1c\x5b\x2a\x4a\xc1\x18\x9e\x07\x69\x7f\x20\x41\xd8\x2a\x63\x41\x69\x1d\x30\xc6\xd4\x3d\x5f\xdf\x43\x5e\x8a\x11\x57\x42\xd6\xe2\xcf\xd3\x54\x96\x95\x6f\xff\xd5\x98\x37\x57\xb7\x7f\x40\xc4\xf0\xc8\x01\x78\xe8\xe3\x38\xa5\x99\x91\x1f\xec\x55\x06\x62\x4b\x5d\x31\xa1\x4b\xc8\x78\x2b\xbf\xc9\x9c\x02\xfc\x9f\xbc\x09\x5b\xc2\x8f\x67\x3f\x8d\xbf\x83\x9b\x9b\x8f\x87\xfe\x05\x31\x36\xfb\x73\x8b\x51\x8d\x8f\xf4\xd5\x64\x51\xd6\xfa\x01\x94\x7b\x4e\xc7\x6c\xa9\xbc\x73\x58\x91\x48\xfb\x79\x84\xec\xf3\xb0\x02\x29\x26\x1d\x27\xa7\x27\xa7\x27\xa3\x92\xdb\xcf\x97\x87\x94\x50\x6f\xb6\x22\x60\x1d\x30\x36\x50\xc2\xf1\x91\x10\x77\xaa\x33\xdb\x99\xe1\xec\xe8\xec\x48\xe4\x77\x2f\xe3\x7f\x2b\xf2\xbb\xbd\xad\xc4\x27\xc2\xe0\x94\xdd\x6e\x45\xee\x54\x9b\xe6\x9a\x56\xa4\x8a\xc6\x5b\xcc\x44\xde\x29\x62\xe2\x4c\xca\x35\x9b\xef\xd9\x7c\x3f\xf9\xfb\xb0\x26\x15\x76\x48\xeb\x80\x16\x55\xc4\x43\x90\x4c\xe4\x2a\xec\x22\x94\x70\x27\x72\x00\x80\xac\x28\xe6\x0e\xc9\x56\x93\xe9\x78\x59\x15\x45\xef\x0c\x2d\xdb\xc6\xf7\xcc\xb0\x15\x39\x4f\xc6\xa2\x0e\xbe\x2d\x22\x69\xe3\xe6\x9a\xcc\xd3\xf4\x5f\xfa\xe1\x70\x70\x12\x9f\x0c\x15\x95\xd7\x18\x39\x4e\x5e\x40\x09\x47\x22\xf7\x5f\x16\x1e\x8d\xb1\x0a\xa6\x4b\xff\xb9\x12\xb2\x4e\xc5\x88\x3a\x7b\x13\xe1\xf1\x44\x38\x49\x79\xc5\x58\x2b\x63\x51\x67\xff\x04\x00\x00\xff\xff\x26\xb1\x90\xbc\xaa\x08\x00\x00")

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

	info := bindataFileInfo{name: "assets/config.example.toml", size: 2218, mode: os.FileMode(438), modTime: time.Unix(1679427268, 0)}
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
	"assets/.env.example":        assetsEnvExample,
	"assets/bluetooth.toml":      assetsBluetoothToml,
	"assets/bluetooth2.toml":     assetsBluetooth2Toml,
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
		".env.example":        &bintree{assetsEnvExample, map[string]*bintree{}},
		"bluetooth.toml":      &bintree{assetsBluetoothToml, map[string]*bintree{}},
		"bluetooth2.toml":     &bintree{assetsBluetooth2Toml, map[string]*bintree{}},
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
