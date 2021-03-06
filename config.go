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

var _configConfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x4d\x4f\xc2\x40\x10\x86\xef\xfb\x2b\xde\x70\xf0\x06\x78\xf1\xd2\xc4\x18\x40\x13\x49\xc4\x1a\xad\x1f\xd7\x6d\x3b\xd6\x85\xe9\x0c\xd9\x0f\x1a\xfc\xf5\xa6\xa6\x85\x0b\xb7\x9d\x67\x9e\xf7\xcd\x66\x4a\x1b\x28\x79\x06\x90\xe1\x27\xc6\x7d\xc8\xe6\xf3\xae\xeb\x66\x4e\x6a\xa2\x7a\x56\x69\x3b\xdf\x6a\x19\xee\x4c\x2f\xb2\x6b\x5d\x44\x86\xc9\xd5\xff\xeb\xf6\xe6\x7a\x62\x4c\xbf\xce\xcc\x14\x5b\x2d\x33\x03\xec\xe8\xd8\xa9\xaf\x33\x54\x2a\xd1\x2b\x07\x90\x34\x4e\x88\xbc\x01\x58\x2b\x1b\x9d\x4a\x2f\x02\x53\x2c\x35\x71\x4d\x1e\xab\x7c\x00\x1b\x17\x82\x26\xb6\xd8\x14\x03\x59\x29\xa7\xb6\x4c\x01\xf9\xe3\x48\xac\x6f\x89\xb1\x7e\x1e\xe6\x9c\x8f\xed\xde\x59\x7c\x2e\x4e\xad\xbf\xd4\x5a\xc1\xfa\x7e\x00\x6f\x96\x23\x9e\xec\x8e\xb0\x72\xf1\x88\xf7\xb1\xfb\x45\x7d\x44\xa1\x9d\x04\x92\xfa\x9c\x5f\x78\x76\xd2\x44\x15\x14\x5f\x63\x25\x49\x54\x39\x38\x66\xc2\xe2\x75\x80\x0f\xa9\x21\x21\xe4\xe3\xbc\x4c\xa7\xe0\x47\x61\x2e\x9c\xa4\x51\xb6\xd2\x5c\xda\x7c\x27\x66\x84\x68\xab\x1d\x6a\x3a\x10\xeb\x9e\xfc\x45\xcf\xab\xc4\x69\xff\xdb\xb3\xf6\x17\x00\x00\xff\xff\xd4\x17\x73\xce\xc3\x01\x00\x00")

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

	info := bindataFileInfo{name: "config/config.yml", size: 451, mode: os.FileMode(493), modTime: time.Unix(1615034268, 0)}
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
