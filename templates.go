// Code generated by go-bindata.
// sources:
// templates/Caddyfile
// templates/docker-compose.yml
// templates/setup.sh
// DO NOT EDIT!

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

var _templatesCaddyfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x41\x6f\xdb\x30\x0c\x85\xef\xfa\x15\x0f\xd0\x75\x8e\x0f\xcd\x29\xb7\x02\xf1\xb6\x02\x43\x03\x34\xe9\xbd\x9a\xcd\x44\x42\x55\x29\x10\xe9\xba\x85\xe1\xff\x3e\x44\x76\xb3\x35\xa9\xb1\x93\xc0\x47\xf2\x3d\x7e\x90\xc6\x7a\x83\xfb\xcd\x0e\xd5\xfa\x6e\x87\xef\x77\xbf\xaa\x6f\xb8\x7d\xdc\x6d\x7e\x54\xf7\xd5\xc3\xed\xae\x5a\x2f\x94\x56\x1a\x0f\xd4\x32\x41\x2c\xc1\xf0\x73\xe1\x02\x8b\xf1\x1e\x66\x2f\x94\x40\x8d\x13\x17\x0e\xb9\x5b\xc7\xb0\x77\x07\xec\x9d\x27\x98\xd0\x20\x51\x91\xda\x70\xb9\xa8\x34\x24\x46\x8f\xce\x89\xc5\x53\xc1\x4f\xab\x1c\xf2\xc9\xbb\x60\xa5\x95\xea\x7b\x2c\x7e\x46\x96\x60\x5e\x08\xc3\x80\x5e\x29\x40\x9f\x06\x1b\x1c\x53\x7c\x7b\x57\x18\x5f\x94\x27\xad\x84\x15\x39\xae\xca\xa9\xe8\x15\x80\x9c\x12\x5b\x19\x27\xb2\x22\xc9\x04\x3e\x9a\x44\x41\x14\x30\x8c\x9e\xe4\x9f\x6d\x4c\xe1\xd2\x76\x92\xcf\xce\x53\xbd\x5a\x2e\x97\xcb\xab\x84\xa9\x39\x1f\x62\x5a\xb1\x57\x87\xb7\x62\xff\x1e\x9e\x8b\xcb\xc3\x5b\xb1\x73\x9e\x7d\x0f\xb7\x47\x88\x82\xc5\x23\xd3\xf6\x06\xc5\x30\xe4\x28\x16\x23\xae\xce\x5f\xc1\x0a\x48\x31\x0a\xca\x57\x93\xca\xae\xeb\xf2\x5e\x01\xf2\x4c\xe7\xf9\xb5\x63\xf3\xdb\x53\x03\xc3\xd8\xde\xa0\x33\x0c\x0a\x59\xc9\xed\x2b\x43\xfd\xb5\x65\x68\x30\x4c\xb0\xb5\xf9\x87\xf2\x03\xb0\x36\xef\x1f\x7c\x33\xa7\x9f\x56\xb7\x94\x5e\x89\xd1\xb9\xe6\x40\x02\x2b\x2f\x3e\x37\xe8\xad\xa6\xa3\xa0\x1c\x75\x9e\x5c\x3e\x81\xfc\x17\x65\x84\x99\x09\xd0\xf3\x11\x67\xb0\x2f\xbe\x61\x50\x7f\x02\x00\x00\xff\xff\x3e\x32\xb0\x75\x48\x03\x00\x00")

func templatesCaddyfileBytes() ([]byte, error) {
	return bindataRead(
		_templatesCaddyfile,
		"templates/Caddyfile",
	)
}

func templatesCaddyfile() (*asset, error) {
	bytes, err := templatesCaddyfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/Caddyfile", size: 840, mode: os.FileMode(420), modTime: time.Unix(1476823100, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x51\x6f\xea\x38\x13\x7d\xcf\xaf\xb0\xca\xc3\x7d\xf9\x20\xbd\x5f\xaf\x56\xab\x48\x3c\xa4\xe0\xa5\x11\xb9\x04\xc5\xa1\xbd\x7d\xca\x35\xce\x00\x59\x52\x1b\xd9\x0e\x2c\xaa\xf8\xef\x2b\x27\x01\x42\xa0\x14\x55\xdb\xaa\x15\x66\xce\x8c\xc7\x33\xe7\x8c\xdd\x42\xfd\x00\x8d\x82\x08\xe1\xbe\x17\xa1\xbf\x3c\x1f\xff\x0f\xb9\x93\x28\x18\xe0\x11\x0e\xdd\x08\xf7\x3b\x56\xcb\x6a\xa1\x10\x72\x05\x48\x2f\x00\x51\xb5\x6c\xa7\x5c\x69\x9a\x65\x88\xce\x34\x48\x04\x49\xaa\x53\x3e\x2f\xac\x4c\xf0\x59\x3a\x47\xb3\x34\x03\x44\x79\x82\x24\xb4\x65\xce\x9b\x8e\x56\x0b\x69\x21\x32\xb4\x49\xf5\x02\xfd\x6e\xab\xdf\x4e\xb1\xc9\x49\xec\xb6\xb2\x5a\x96\xb5\x06\xa9\x52\xc1\x1d\xf4\xed\xff\xdf\x2c\xab\xf5\xe1\x8f\xd5\x42\x04\xe4\x3a\x65\xa0\xae\xa1\x2c\x55\x81\x1c\xcb\x42\x88\xd1\xad\x63\x21\x84\x50\xfa\x46\xe7\xe0\x20\x26\x24\xcd\x56\x52\xfc\x0d\x4c\xdb\x8c\x6e\x0b\x9b\x04\xa5\xa9\xd4\x0e\x12\xbc\x3d\xa3\x69\x96\x4b\x70\xbe\xdf\x17\x26\xe0\xeb\x54\x0a\xfe\x06\x5c\x97\x71\x10\x6a\xa3\xbb\x81\xeb\xf5\xbb\xef\xef\xa8\x33\x10\x62\x9e\x81\xcb\x69\xb6\xd5\x29\x53\x5e\x1f\xed\x76\x77\x47\x9c\x3b\x89\x9e\xe2\x9e\xef\xe1\x51\x14\x7b\xfd\x2e\xa3\xdb\xa6\xd1\xfc\x0b\x42\x2f\x7a\x2d\xc2\x85\x42\xe8\x49\xe8\xa3\xdd\xce\xa6\xb9\x5e\xd8\x4c\x70\x0e\x4c\xd7\x9c\xb0\x3f\x7c\x0a\xc2\x51\x3c\x09\xfd\xa6\x07\x64\xcb\x85\x90\xbc\xbe\x03\x19\x5e\x02\x52\xb5\x4c\x4a\x14\x07\xbd\x11\x72\xa9\x8e\x47\x9b\x52\xb6\x6c\xab\x05\x64\x33\x53\x3d\x83\xfc\xb8\x7c\xc6\xfa\xb5\xfa\x99\xc4\x9e\x02\x12\x75\xef\x3b\xc5\xaf\xf3\xe7\x7d\x23\x6d\x3f\x18\x0c\xbc\xd1\x20\xf6\xf1\x33\xf6\xbb\xdf\x1b\xd6\x9f\xc1\x68\x10\xc4\x93\xd0\xeb\xbe\x09\x3e\x17\xc9\xd4\xb1\xed\x22\xe9\x76\xb1\xb6\x35\x28\xdd\x70\xc1\x23\xf7\xd1\xc7\x71\x2f\x08\x49\x37\x0a\x27\xb8\x61\x2e\x9a\x31\x9e\x3c\xfa\x5e\x2f\x1e\xe2\xb2\x1b\x6e\xae\x17\xe3\x7c\x9a\xa5\x6c\x08\xdb\x46\x63\xc9\x30\x0e\x71\xcf\x1d\x47\xbd\x27\x37\x26\xb8\x17\xe2\xa8\x2c\x33\x30\xba\xd2\x6c\x41\x09\x30\x09\xfa\xe0\x75\xb5\xd2\xe6\x8b\x04\x56\xc0\x13\x15\x0b\x7e\x84\xd4\x8e\x64\xba\x51\x35\xf8\xe3\x86\x54\x80\xaf\xf7\xc4\x90\x65\xa1\xf5\xca\xb1\x6b\x24\x29\xac\xe4\x21\x7e\x9c\xf4\x86\xd5\x29\xc9\xc3\x63\xce\x96\xb5\xe3\x95\x11\x5e\x48\x1c\xe2\x81\x17\x8c\xca\xf2\xbd\x90\x10\xe6\xa9\xe0\xe7\x28\xb7\xd7\xc3\x84\x98\x42\xc7\x95\x90\xdc\x17\xe2\x32\x06\x4a\x0d\x61\x7b\xa6\xa2\x13\x8f\x33\xf8\x29\xf8\xd0\x95\xab\xed\x78\x7f\x6f\xa3\x74\x86\xb8\xd0\xa8\x33\x51\x40\x1e\xd0\x6e\x77\x8c\x51\xf0\x80\x3c\xe1\x7e\xfc\xe8\x12\x7c\x49\x42\x9b\x34\x99\x83\x56\x76\x19\x6d\x2d\xb2\xfc\x0d\x6a\xbd\xed\xec\x5b\xe1\xd8\xb9\x92\xb6\x92\xcc\xa6\xab\x95\x6d\x26\x13\xc8\xbd\x73\x95\x08\x02\x9e\xec\x77\xff\x12\x4b\x0a\x19\x1a\xb1\xe6\x7a\x71\x65\xd6\x99\x45\xdb\x60\xbe\x46\x8f\x5e\x10\xba\x7e\x25\x93\x20\x8c\xba\x27\x8a\xad\x19\xa3\x70\x42\xa2\x78\x1c\x06\xbf\x5e\x9b\x3a\xab\xa3\x82\x21\x1e\xc5\xf8\xd7\xd8\x0b\x5f\xe3\xc8\xfb\x89\xbb\x0f\x7f\xa0\x85\xbc\x0c\xde\xab\xdd\xaf\xa9\xdd\x1c\xe4\xb2\xd8\xeb\x89\x86\xde\xb3\x1b\xe1\x53\x41\xcb\x74\x4d\x35\x9c\xf1\xa6\xee\x76\xf3\x18\xa8\x39\x11\x4c\x88\x17\x8c\xea\xb3\x80\x80\x32\x17\x5b\x83\x7a\x4d\x47\xd7\xf7\x83\x17\xdc\xaf\xae\x09\x62\x2e\x09\xd4\x60\x1b\xa3\x59\x66\x88\x70\x39\x40\x18\x04\xd1\xc5\x31\x9f\xeb\xc5\xd5\xe1\x73\x20\xc3\x45\x56\x1d\x0a\x6c\xb8\x75\xe5\xbe\x45\xa8\x85\xfa\x54\xd3\x29\x55\xa0\xae\x43\x2d\x54\x1f\x69\x27\x5c\x2d\x77\xfa\x84\x99\x67\x3a\xab\x47\xb3\x13\xaa\xa9\x9d\x4c\x6f\xbb\xd8\x0e\xa7\xfb\x0f\xb2\xa8\x05\xbb\x25\x89\xa2\xec\x9f\xd7\xf4\x05\xa6\xc5\x7b\x07\xe4\xe7\x55\x65\x34\x49\x4e\xdf\x39\x74\x9a\x0a\x25\x66\xe6\x8d\x93\x24\xdb\x8f\xa6\x54\xcf\x18\xcd\x43\xce\xb1\x41\xb3\xe3\xb2\x82\x5c\x9d\x92\xb5\x19\xb7\xa6\xd2\xde\x6c\x36\x27\x83\xad\x39\xda\x8c\x4b\x91\x4b\x9b\x81\xd4\xca\xb1\xa5\x10\xda\xee\x1c\xd3\x5b\x09\xf3\xf5\x91\xe1\xfb\xa7\x81\x61\xf5\x58\x48\x23\xa0\xfa\xe7\xbb\xd3\x1c\x8b\xfc\x88\x7f\x32\xc6\xf7\x11\x7e\xfc\x78\x30\x7f\x77\x17\xf3\xba\xd4\x24\xa5\x53\xc1\x3f\xd0\xca\xad\x43\xb9\x8e\x3f\x3c\x94\xcc\xa2\x7e\x49\x97\x45\xd9\xde\xf6\xea\x45\xa3\x2a\xd5\xab\xaf\xdf\xfa\x79\xaa\x93\x94\x1f\xf7\x29\x1f\x56\xe5\x7d\x71\xdb\xd6\xcf\x25\x77\xae\xee\x7c\xe0\xd7\x9e\x90\x55\xa7\xcb\x4b\xe5\x1f\x0d\x92\xd3\xcc\x41\x33\x9a\x29\xb8\x20\xc1\x73\xc4\xd9\xac\x68\x42\xfe\x0d\x00\x00\xff\xff\xf0\xce\xf5\xc8\xd9\x0c\x00\x00")

func templatesDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDockerComposeYml,
		"templates/docker-compose.yml",
	)
}

func templatesDockerComposeYml() (*asset, error) {
	bytes, err := templatesDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/docker-compose.yml", size: 3289, mode: os.FileMode(420), modTime: time.Unix(1476824060, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSetupSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\xc1\x4e\xc3\x30\x10\x44\xef\xfe\x8a\xc1\x3d\x3b\xbd\xf7\x0a\x5c\xab\x4a\x15\x1f\xb0\x38\x5b\xb2\x22\xb1\x23\xaf\xd3\x80\xa2\xfc\x3b\x72\x42\x44\x5b\x4e\xb6\x3c\xb3\xde\x37\xb3\x7b\xda\xbf\x4b\xd8\x6b\x63\x8c\x72\x86\x63\x63\xd8\x37\x11\xf6\x39\x31\x65\x09\x1f\xc8\x0d\x43\x39\x5d\xc5\xb3\x22\xc4\xf1\x60\x37\x8b\x35\x66\x87\x73\xa6\x94\xef\x4c\x95\xa9\xa3\xff\xe4\xe4\x7c\xec\xfa\xa8\x8c\xa1\x87\xab\xff\x86\xd6\xf3\xbc\x7d\xe9\xcb\x22\xae\x2b\x6b\xa6\x09\x72\x41\x75\x22\xd5\x31\xa6\x1a\xf3\xfc\x30\x73\x8c\xe3\x6a\x2f\x5c\x83\x72\x3a\xd8\x5b\x94\x05\x99\x17\x96\x22\xfe\xe3\x48\x43\x80\x73\xa9\x03\x0d\xb9\x81\x6f\xe5\x77\x37\xdc\x05\xce\x71\x47\xd2\xc2\x4e\x13\xaa\xd7\xe5\x3a\xcf\x16\xce\xf5\x1b\xcd\xa2\xdc\xb0\x15\x31\x50\xc7\xab\xf0\x22\xda\xb7\xf4\x7d\x2c\x0f\xf3\x6c\x1f\xd3\xbe\x29\x27\x8c\x74\x9b\xf6\xde\x70\x6a\x99\x0a\x22\x77\xf1\x5a\x22\x88\xe2\x22\x2d\x83\x14\x92\xe1\x63\xc8\x24\x41\x97\x6c\x7d\x4b\x12\x32\x7f\x65\x6c\x6c\x6b\x77\x1c\x96\xca\x7e\x02\x00\x00\xff\xff\xc0\x19\x53\xa3\xd2\x01\x00\x00")

func templatesSetupShBytes() ([]byte, error) {
	return bindataRead(
		_templatesSetupSh,
		"templates/setup.sh",
	)
}

func templatesSetupSh() (*asset, error) {
	bytes, err := templatesSetupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/setup.sh", size: 466, mode: os.FileMode(420), modTime: time.Unix(1476814501, 0)}
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
	"templates/Caddyfile": templatesCaddyfile,
	"templates/docker-compose.yml": templatesDockerComposeYml,
	"templates/setup.sh": templatesSetupSh,
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
	"templates": &bintree{nil, map[string]*bintree{
		"Caddyfile": &bintree{templatesCaddyfile, map[string]*bintree{}},
		"docker-compose.yml": &bintree{templatesDockerComposeYml, map[string]*bintree{}},
		"setup.sh": &bintree{templatesSetupSh, map[string]*bintree{}},
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

