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

var _configConfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x4d\x6b\xe3\x30\x10\x40\xef\xfa\x15\x43\xc2\xee\xc5\x24\x59\x16\xf6\x62\x08\x4b\xe2\x06\x5a\x68\xea\xd2\xba\x1f\x57\x39\x16\xae\x92\xb1\x26\x95\xe4\x9a\xf4\xd7\x17\xe9\xa2\x89\x93\x9b\xfd\xde\x68\x06\x5e\x2d\x9d\xea\x2d\x02\x40\x0e\x1f\xde\x1f\x5d\xbe\x58\x0c\xc3\x30\xd7\xa6\x51\xaa\x99\xef\xa8\x5b\xec\xa9\x76\xff\x45\x18\x44\xdd\x69\x0f\x39\x4c\x7e\xc7\xaf\xe5\xbf\x3f\x13\x21\x82\xce\xc5\x0c\xf6\x54\xe7\x02\xe0\xa0\x4e\x03\xd9\x26\x87\xcf\x65\x41\xc6\x5b\x42\x97\x6d\x4c\xab\x8d\x52\x56\x00\x20\xed\xa4\xd7\x64\xc2\x28\xc0\x0c\x70\xb9\xa6\x1e\x1b\x65\x7f\xfd\x2d\xb2\xa2\x8c\x74\x1a\xf9\x56\x3b\x47\x3d\xca\x20\xb6\x15\x13\x05\x61\xdf\xd5\xbd\x0b\xa2\xbc\xe5\x42\xda\x4e\x61\xc0\x77\x0f\x0c\x97\x78\xea\x8e\x3a\xee\x79\x5b\x31\xbe\xa6\x6f\xd5\x49\x13\xe7\x6f\x18\x7f\x96\xe8\xb3\x7b\x79\x50\x59\xa1\xfd\x29\xe8\x17\x7e\xfe\x91\xac\xcf\x2a\x1a\x8c\x53\xa6\xb9\x58\xba\xb2\xa8\x4d\xeb\x29\xae\xad\xde\xf9\x39\x65\x3c\x99\x2f\x8d\xa8\x82\x5b\x3d\x31\xb7\xe9\x5b\x65\x22\x2e\x39\x5e\xf7\x7c\xdb\x6b\x25\x44\xe0\x31\xf4\x14\xce\x52\xb7\x84\xd2\xb4\x91\xa6\xc2\xe1\xef\x5a\xe3\xc4\x47\x8d\x93\x18\x35\x66\x82\x37\x4e\xf8\xbc\x31\x3f\xcc\x1b\x27\x7e\xb5\x71\xd2\xd7\x1a\x27\x3b\x6e\xcc\xce\x5d\x34\x4e\xee\xac\x31\x7b\x32\x6e\xfc\x13\x00\x00\xff\xff\x0d\xb0\x4e\x09\x12\x03\x00\x00")

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

	info := bindataFileInfo{name: "config/config.yml", size: 786, mode: os.FileMode(509), modTime: time.Unix(1614861200, 0)}
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
