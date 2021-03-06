// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/config.yml
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

var _configConfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\xcd\x4e\xc2\x40\x14\x85\xf7\xf3\x14\x27\x2c\xdc\x10\xc0\x8d\x9b\x26\xc6\x14\x34\x91\x44\xac\xd1\xfa\xb3\x9d\xd2\x9b\x3a\x70\x7b\x2f\x99\x1f\x1b\x7c\x7a\x53\xd3\xc2\x86\xdd\x9c\x6f\xce\xf9\x6e\x65\x03\x25\xcf\x00\x32\x7c\xc7\x78\x08\xd9\x62\xd1\x75\xdd\xdc\x49\x4d\x54\xcf\xb7\xda\x2e\x76\x5a\x85\x3b\xd3\x17\xd9\xb5\x2e\x22\xc3\xe4\xea\xff\x75\x7b\x73\x3d\x31\xa6\xff\xce\xcc\x0c\x3b\xad\x32\x03\xec\xe9\xd8\xa9\xaf\x33\x6c\x55\xa2\x57\x0e\x20\x69\x9c\x10\x79\x03\xb0\x6e\x6d\x74\x2a\x7d\x11\x98\x61\xa5\xac\xde\xd6\x3a\xc4\xdc\xef\xad\x04\x1b\xcc\x05\x5d\xa3\x6c\xa5\xb9\xe0\x58\x6a\xe2\x9a\x3c\x56\xc5\x00\x36\x2e\x04\x4d\x6c\xb1\x29\xcf\x67\x52\x5b\xa5\x80\xe2\x71\x24\xd6\xb7\xc4\x58\x3f\x0f\xb9\xe0\x63\x7b\x70\x16\x9f\xf9\xc9\xfa\x4b\xad\x15\xac\xef\x07\xf0\x66\x39\x4e\x9f\xec\x9e\xa6\x2b\x17\x8f\x78\x1f\xdd\x2f\xea\xe3\xb4\xd4\x4e\x02\x49\x7d\xde\xe7\x9e\x9d\x34\x51\x05\xe5\xd7\xa8\x24\x89\x2a\x3f\x8e\x99\x90\xbf\x0e\xf0\x21\x35\x24\x84\x62\xcc\xcb\x74\x1a\x7e\x94\xc6\xfc\x05\x00\x00\xff\xff\x78\x92\xc2\x5c\x9f\x01\x00\x00")

func configConfigYmlBytes() ([]byte, error) {
	return bindataRead(
		_configConfigYml,
		"config/config.yml",
	)
}

func configConfigYml() (*asset, error) {
	bytes, err := configConfigYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.yml", size: 415, mode: os.FileMode(493), modTime: time.Unix(1614978235, 0)}
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
	"config/config.yml": configConfigYml,
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
	"config": &bintree{nil, map[string]*bintree{
		"config.yml": &bintree{configConfigYml, map[string]*bintree{}},
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
