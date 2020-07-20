// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1593601729_initial_schema.down.sql (144B)
// 1593601729_initial_schema.up.sql (1.753kB)
// doc.go (382B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __1593601729_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x28\x2d\xce\x88\xcf\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x4f\xce\xc9\x4c\xcd\x2b\x89\x2f\x4e\x2d\x2a\x4b\x2d\x2a\xb6\xe6\x22\x46\x71\x66\x5e\x5a\x3e\x54\xa5\xa7\x9f\x8b\x6b\x84\x42\x66\x4a\x45\x3c\x5e\xd5\xf1\x05\xa5\x49\x39\x99\xc9\xf1\xd9\xa9\x95\xd6\x5c\x80\x00\x00\x00\xff\xff\x6d\xb4\xf8\x65\x90\x00\x00\x00")

func _1593601729_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1593601729_initial_schemaDownSql,
		"1593601729_initial_schema.down.sql",
	)
}

func _1593601729_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1593601729_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1593601729_initial_schema.down.sql", size: 144, mode: os.FileMode(0644), modTime: time.Unix(1594896579, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa, 0x95, 0x55, 0x64, 0x38, 0x40, 0x16, 0xbf, 0x8b, 0x1c, 0x18, 0xb4, 0xc5, 0x7f, 0xd0, 0xb8, 0xf0, 0x3c, 0xa2, 0x82, 0xf8, 0x8d, 0x5a, 0xd3, 0xb6, 0x6e, 0xa3, 0xb4, 0xc, 0x9, 0x33, 0x0}}
	return a, nil
}

var __1593601729_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xc1\x6e\xe3\x20\x10\xbd\xfb\x2b\xe6\xd8\x48\x39\xec\xbd\x27\x27\x25\x2b\x4b\x08\xef\x26\x44\xca\x0d\xb1\x78\x5a\xa3\x78\x71\x17\x70\xb5\xf9\xfb\x15\x76\xe2\x26\xc5\x8b\xab\xb6\x17\x4b\x1e\x1e\xa3\x79\x6f\x1e\x6f\xbd\x25\x39\x27\xc0\xf3\x15\x25\x50\x6c\x80\x95\x1c\xc8\xa1\xd8\xf1\x1d\x3c\x77\xae\x16\xa6\xf5\xfa\x51\x2b\xe9\x75\x6b\x84\x6a\x34\x1a\x2f\x1c\xda\x17\xb4\x0e\xee\x32\x80\xe7\xee\x57\xa3\x95\x38\xe2\x09\x56\xb4\x5c\xf5\xf7\xd9\x9e\xd2\x65\x06\x60\xf1\x49\x3b\x8f\x16\x2b\x58\x95\x25\x25\x39\x83\x07\xb2\xc9\xf7\x94\xc3\x26\xa7\x3b\x72\x8b\x11\xd2\x43\xc1\xf8\xd8\x61\xc4\x7e\x0b\xb8\x46\x3a\x2f\x2c\x7a\xab\xe7\x90\x01\x74\x12\xaa\xed\x4c\x0a\x25\x95\x42\xe7\x84\x6f\x8f\x68\x80\x93\x03\x0f\xc5\x3d\x2b\x7e\xee\xc9\xdd\x2b\xa7\x05\x94\x0c\xd6\x25\xdb\xd0\x62\xcd\x61\x4b\x7e\xd0\x7c\x4d\xb2\xc5\x7d\x96\x7d\x44\xb7\x3f\x1d\x5a\x8d\xf3\xba\x0d\xb8\x88\xe6\xe5\xe8\x24\x74\x15\x5f\x8a\x66\x5f\x5e\xb0\x5f\x4b\x42\x9b\xc7\x76\x96\xc1\xe0\x10\x91\x82\x68\xe3\xbc\x6c\x9a\xa1\xb7\xae\xfa\x1d\xdc\x00\xa2\x0d\xbd\xf1\x56\xb0\xc2\xcb\xb4\x4a\xc1\x9d\xba\x35\x51\x3d\xd6\xe8\xed\x18\xcb\x78\xf4\xaf\x95\xcf\x5b\xa9\x8e\x58\x89\xdf\xe8\x9c\x7c\x3a\x9b\xe1\xfc\x33\xb9\x57\x55\x4b\x3f\xa9\xcf\xa5\xd3\x04\xff\x33\xcf\xd7\xb6\xb7\x1c\x8a\xef\xac\xdc\x92\x0c\xe0\xa3\x24\x5c\xf8\x5c\x1f\xcc\xd3\x48\x59\xa1\x96\xae\xc6\xea\x73\x6e\xe9\xf3\x61\x32\x1d\xde\x9f\x09\xae\xeb\x2d\x37\x86\x55\x84\x1a\x53\x0b\xad\x6d\x6d\xa2\x53\xb4\x80\x25\x24\x4c\xb7\xf8\xc4\x2a\x86\xfc\xb4\x57\x5b\xb8\x64\xea\x50\x8b\xa5\x04\x50\xad\xf1\x52\x05\x57\xb9\xfe\x78\xa8\xba\x93\xf1\x35\x7a\xad\x82\xbe\xff\xa7\x36\x92\xbb\xc6\xcf\xbe\x91\x82\x3d\x90\x03\xe8\xea\xaf\x48\x06\xcb\xb5\x07\x4a\x96\x0e\xa1\xd4\x33\x5e\xdc\x67\xff\x02\x00\x00\xff\xff\x39\x9f\x6b\x23\xd9\x06\x00\x00")

func _1593601729_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1593601729_initial_schemaUpSql,
		"1593601729_initial_schema.up.sql",
	)
}

func _1593601729_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1593601729_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1593601729_initial_schema.up.sql", size: 1753, mode: os.FileMode(0644), modTime: time.Unix(1595240420, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x28, 0xbf, 0x64, 0xe0, 0x65, 0x53, 0xd3, 0x80, 0xf4, 0x46, 0xce, 0xd6, 0x23, 0x4e, 0xc5, 0x8f, 0x80, 0x4e, 0x91, 0xa7, 0x2e, 0x9, 0x3b, 0xf4, 0x5f, 0xa1, 0xff, 0xfc, 0x6e, 0x4, 0xa2, 0xe7}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x3d\x6e\xec\x30\x0c\x84\x7b\x9d\x62\xb0\xcd\x36\xcf\x52\xf3\xaa\x74\x29\xd3\xe7\x02\x5c\x89\x96\x88\xb5\x24\x43\xa4\xf7\xe7\xf6\x81\x37\x01\xe2\x2e\xed\x87\xf9\x86\xc3\x10\xf0\x59\x44\x31\xcb\xc2\x10\x45\xe3\xc8\xaa\x34\x9e\xb8\x70\xa4\x4d\x19\xa7\x2c\x56\xb6\x8b\x8f\xbd\x06\x35\xb2\x4d\x27\xa9\xa1\x4a\x1e\x64\x1c\x6e\xff\x4f\x2e\x04\x44\x6a\x67\x43\xa1\x96\x16\x7e\x75\x29\xd4\x68\x98\xb4\x8c\xbb\x58\x01\x61\x1d\x3c\xcb\xc3\xe3\xdd\xb0\x30\xa9\xc1\x0a\xd9\x59\x61\x85\x11\x49\x79\xaf\x99\xfb\x40\xee\xd3\x45\x5a\x22\x23\xbf\xa3\x8f\xf9\x40\xf6\x85\x91\x96\x85\x13\xe6\xd1\xeb\xcb\x55\xaa\x8c\x24\x83\xa3\xf5\xf1\xfc\x07\x52\x65\x43\xa3\xca\xba\xfb\x85\x6e\x8c\xd6\x7f\xce\x83\x5a\xfa\xfb\x23\xdc\xfb\xb8\x2a\x48\xc1\x8f\x95\xa3\x71\xf2\xce\xad\x14\xaf\x94\x19\xdf\x39\xe9\x4d\x9d\x0b\x21\xf7\xb7\xcc\x8d\x77\xf3\xb8\x73\x5a\xaf\xf9\x90\xc4\xd4\xe1\x7d\xf8\x05\x3e\x77\xf8\xe0\xbe\x02\x00\x00\xff\xff\x4d\x1d\x5d\x50\x7e\x01\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 382, mode: os.FileMode(0644), modTime: time.Unix(1594896579, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x2f, 0x1e, 0x64, 0x9, 0x93, 0xe4, 0x8b, 0xf2, 0x98, 0x5a, 0x45, 0xe2, 0x80, 0x88, 0x67, 0x7a, 0x2d, 0xd7, 0x4b, 0xd1, 0x73, 0xb6, 0x6d, 0x15, 0xc2, 0x0, 0x34, 0xcd, 0xa0, 0xdb, 0x20}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"1593601729_initial_schema.down.sql": _1593601729_initial_schemaDownSql,

	"1593601729_initial_schema.up.sql": _1593601729_initial_schemaUpSql,

	"doc.go": docGo,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"1593601729_initial_schema.down.sql": &bintree{_1593601729_initial_schemaDownSql, map[string]*bintree{}},
	"1593601729_initial_schema.up.sql":   &bintree{_1593601729_initial_schemaUpSql, map[string]*bintree{}},
	"doc.go":                             &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
