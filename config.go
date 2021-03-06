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

var _configConfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xd0\x4f\x4b\xc3\x40\x10\x05\xf0\xfb\x7e\x8a\x47\x0b\xde\xd2\x7a\xf1\x12\x10\x69\xab\x60\xc1\x36\xa2\xf1\xcf\x75\xd3\x8c\xe9\xb6\x93\x9d\xb2\x7f\x1a\xe2\xa7\x97\xd4\x92\x08\xd2\x5b\xf8\xbd\x97\x59\x66\x0a\xed\x29\x3a\x06\x90\x62\x1b\xc2\xc1\xa7\xd3\x69\xd3\x34\x13\x63\x4b\xa2\x72\xb2\x91\x7a\xba\x93\xc2\xdf\xa9\xae\xc8\xa6\x36\x01\x29\x46\x57\xa7\xaf\xdb\x9b\xeb\x91\x52\x5d\x9c\xaa\x04\x3b\x29\x52\x05\xec\xa9\x6d\xc4\x95\x29\x2a\x61\x6d\x2b\x05\xb0\x6c\x74\x30\x62\xbb\x14\x48\xb0\xd6\x7e\x7b\x34\xcc\x84\x7c\x7d\xa6\xb9\x44\x2e\xc9\x61\x91\x9d\x60\x8c\x04\x2b\xe3\xbd\x44\xd6\x58\xe5\xbd\x2d\x84\x63\x5d\x44\x8f\xec\x71\x30\xed\x6a\x62\x2c\xd7\xbd\x64\xdc\xd6\x07\xa3\xf1\x31\xeb\x69\x2e\xdf\x54\x6b\x8b\xe5\x7d\x4f\xaf\x9a\x03\x9e\xf4\x9e\xb0\x30\xa1\xc5\xdb\xf0\xca\xb3\xb8\x80\x5c\x1a\xeb\xc9\x96\x7f\xa7\xcc\x1c\x1b\x5b\x05\xb1\xc8\x3f\x87\xd1\x64\x83\xd8\xdf\x7d\x66\x2f\x3d\x3f\xc4\x8a\x2c\x21\x1b\x64\x1e\xfb\xdf\xdf\x73\xa5\x3a\x3a\x9f\x6c\xfc\xef\x68\x17\xd2\xaf\xc8\x0c\x1f\xf4\x66\x8f\x92\x8e\xc4\x72\x20\x77\xb1\xeb\xc4\x86\xa4\xdb\x60\xa8\xfe\x04\x00\x00\xff\xff\xf2\x66\x9e\x52\xed\x01\x00\x00")

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

	info := bindataFileInfo{name: "config/config.yml", size: 493, mode: os.FileMode(493), modTime: time.Unix(1615051586, 0)}
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
