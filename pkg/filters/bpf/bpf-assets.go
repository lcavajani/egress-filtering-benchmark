// Code generated by go-bindata.
// sources:
// datapath/bpf.o
// DO NOT EDIT!

package bpf

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _datapathBpfO = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\xcd\xca\xd3\x50\x10\x3d\xb9\x69\x6d\x6b\xb5\xd8\x56\xb0\x8a\x14\xb1\x5d\x88\x60\x28\x4a\xdd\x09\xb5\xe0\x0f\x18\xb0\x8a\xc5\xba\x0a\x69\xb8\xfd\x81\x44\x6b\x92\x85\x8a\xa0\x0b\x05\x5f\xc0\x17\x70\xe5\x1b\x74\x59\x1f\xc3\xa5\x4b\x97\xba\xb2\x0b\x21\x1f\xb9\x99\xdb\x86\xdb\x94\x0c\x34\x93\x73\x3a\x73\x67\xce\x49\xf2\xe1\x9e\x79\x9f\x69\x1a\x64\x68\xf8\x87\x3d\xda\xc7\x4f\xb6\xbf\x1f\xd0\xf5\x34\x34\xac\x89\xb3\x1b\xa3\x24\xd7\x4d\x91\x37\xcd\x84\x2f\xe9\x40\x0d\xc0\x8d\xce\x85\x03\xfe\xaa\xe0\xcf\x0b\xfc\xba\x71\x56\xe4\x39\x03\xca\x31\x6e\x9e\x11\xf8\xf1\xcd\xa4\xfe\x39\x03\xa2\x08\x18\xb3\x9a\xf8\x7f\xcd\x80\x2b\x00\x9c\xeb\xdb\x28\x99\xdb\x16\x75\xce\xa5\xff\x02\x6f\xbe\xd1\x1c\x06\x6c\xa3\x28\x6a\x29\xa2\x3e\x0b\xad\xc0\x86\xf8\x35\xe1\x31\x79\x21\x75\xcd\x01\xc4\xd2\xbf\x12\xae\x22\xd9\xaf\x00\xe0\xe3\xa7\xf6\xce\xab\x94\x3d\x78\x30\x32\x71\x4e\xfa\xf9\xee\x29\xca\xef\xab\x5a\xac\xa6\x45\x3f\x19\x7f\x32\x7c\x56\xe3\x91\x98\xf5\x37\x52\xf9\x85\xb8\xea\xf8\xa5\xf0\x36\xf1\xbf\x33\xce\xd2\xa1\x1f\x70\xb7\x00\xd4\x71\x6a\x87\x0b\x94\xaf\x09\xbe\xb8\xe3\xbb\x94\x2f\x02\x68\xa4\xce\x91\x1a\xbe\x53\x8e\xfd\x88\x4f\xbb\x9c\xc2\xe2\x4c\x23\xe4\x6f\x42\x18\x3e\x77\x67\x4b\x37\xe4\xbe\xc5\xe7\x3e\x0f\x02\x84\x8e\xbc\xf3\xec\x55\x00\x77\xe5\x59\x49\x01\x2c\xcb\x5d\x3a\xfc\x65\xc0\x45\x97\xc1\x17\xd6\xcc\xb7\x3d\x8e\xe9\x6a\x66\x38\x30\x82\xd0\x0f\xed\x29\x8c\xe0\xad\x17\x67\x73\x38\xec\x59\xb7\x93\xd4\xcf\xb5\x35\x37\x9e\x08\xbf\x0e\x63\x41\x0f\xfa\x95\xc2\xab\xdf\x8c\x96\xf2\x22\x1d\x83\x23\xf3\x0a\x0a\xae\xe6\xf4\xab\xef\x4e\x59\xc1\x25\x00\x95\x8c\x39\x0f\x69\x7f\xf9\x7e\x56\x48\xa7\xec\x97\x7c\x87\xe6\xab\x1e\xf4\x48\x68\x37\x67\xff\xfe\x91\xfe\x67\x5a\x76\xbd\xea\xdf\x5d\xe2\x98\xc2\x4f\xa8\xb0\xa7\xf0\xaa\xfe\x3b\x47\xf4\x4f\x32\xf4\x97\x32\xf4\xbf\xc8\x98\x1d\xc7\x17\x9a\xff\x23\xb5\x77\x31\xd5\x2f\xbf\xef\x93\x00\x00\x00\xff\xff\x47\xd3\x52\x41\x58\x05\x00\x00")

func datapathBpfOBytes() ([]byte, error) {
	return bindataRead(
		_datapathBpfO,
		"datapath/bpf.o",
	)
}

func datapathBpfO() (*asset, error) {
	bytes, err := datapathBpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "datapath/bpf.o", size: 1368, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"datapath/bpf.o": datapathBpfO,
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
	"datapath": &bintree{nil, map[string]*bintree{
		"bpf.o": &bintree{datapathBpfO, map[string]*bintree{}},
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
