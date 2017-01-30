// Code generated by go-bindata.
// sources:
// assets/cluster.yml
// assets/pipeline.yml
// assets/repo.yml
// assets/service.yml
// assets/vpc.yml
// DO NOT EDIT!

package templates

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

var _assetsClusterYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1a\xfd\x6f\xdb\x36\xf6\xf7\xfc\x15\xaf\xbe\x1e\xb2\x0d\x93\xed\x38\xeb\xb0\x09\xd7\x0d\x8e\xe3\xb5\xc6\x9c\xd6\x88\xdd\x16\x58\x3b\x14\xb4\xf4\x6c\xf3\x2a\x91\x3a\x92\x4a\xe2\xf5\xfa\xbf\x1f\x48\xea\x83\x94\xe4\x34\xe9\x7a\xb7\xdb\xe1\x5c\x04\xb5\xc9\xc7\xf7\xfd\x49\x29\x08\x82\xa3\xf1\xab\xe5\x0a\xd3\x2c\x21\x0a\x7f\xe2\x22\x25\xea\x25\x0a\x49\x39\x0b\xe1\x78\x34\x3c\x19\x06\xc3\xef\x83\xe1\xf7\xc7\x47\x0b\x22\x48\x8a\x0a\x85\x0c\x8f\x00\x66\x4c\x2a\xc2\x22\x5c\xed\x33\xd4\xbf\x01\xcc\x37\x58\x2a\x41\xd9\xd6\x2c\x9c\xa3\x8c\x04\xcd\x94\x41\x55\xc2\x83\xda\x67\x08\x8a\x43\x2e\xb1\x5f\x80\x6d\x48\x9e\xa8\x10\xd4\xa8\x9f\xd2\x48\xf0\x23\x73\x94\x0a\x8c\x27\x24\x23\x11\x55\x7b\x97\xc0\xb3\x3c\x5d\xa3\xf0\x4f\x1e\x9f\x1c\xb7\x29\x5a\x40\xe0\x1b\xa0\x05\x6d\xa9\xe9\x26\x24\x67\xd1\x0e\x28\x83\x3d\xcf\x05\x4c\x27\x4b\x88\x92\x5c\x2a\x83\xf3\x82\xdc\x2c\xe9\x6f\xf8\x51\x7a\xa3\x0e\x7a\x17\xe4\x86\xa6\x79\x0a\xac\x8b\xee\x8e\x28\x88\x08\x83\x35\x16\x0c\x60\x7c\x80\x85\x9f\x71\xff\x8c\xa4\x77\xd2\x69\x01\xaa\xa5\x22\x52\xf2\x88\x12\x85\x70\x4d\xd5\x0e\xae\xb9\x78\x87\xa2\x66\xa0\x0f\x30\x47\x72\x85\xb0\x4e\x08\x7b\xa7\x0f\xc4\x54\x92\x75\x82\xb0\x5c\x3e\x05\x12\x45\x28\x65\xc3\x1a\xc7\x5a\xc4\xa5\xdc\x8d\x93\x84\x5f\x87\x6d\xe2\xcb\x7c\xcd\x50\xc1\x46\xf0\x14\xae\x77\x34\xda\x19\x36\x34\x70\x0b\x67\x4b\x8a\x0b\xca\xe6\xc8\xb6\x6a\x17\xc2\xf1\xf7\x56\x95\x17\xe4\xa6\x5a\x3a\xf9\xee\xd8\xe7\x65\xd8\x37\xff\x06\x43\xb3\x6c\x38\xc2\x78\x41\x94\x42\xc1\x42\xe8\x7d\xf1\xe6\x4d\xfc\xfe\xe4\xeb\xd3\x0f\x5f\xbe\x79\xd3\xbf\xcb\x8f\x41\xf1\x75\xf4\xe1\xcb\x9e\x41\x39\xe1\x4c\x2a\x41\x28\x53\x9e\x8c\xc7\x69\x2e\x95\xb6\x19\x81\x2b\x92\xd0\x18\x26\xb3\xf3\x4b\x58\x27\x3c\x7a\x17\xc2\x4d\xdf\xfc\x1b\xdc\xf4\x8d\xa6\x22\x92\xe0\xf3\x5c\xad\x76\x02\xe5\x8e\x27\x71\x87\xca\xaa\x3d\x20\xaa\xd0\x19\x01\xa9\x0f\x02\xcf\x15\xe0\x15\x32\x05\xd7\x34\x49\x34\x49\xca\xa8\xd2\x06\x8d\x3f\xe6\x8b\xdf\x0d\x2b\xfa\x33\xf6\x89\xe4\x29\xfb\x54\xea\xa7\x86\xfa\x2c\x25\x5b\x9c\x75\x11\xd5\xce\x3d\xbe\x98\xd5\x91\x77\xc8\xaf\x5d\xb7\x7b\x99\x45\x25\xb2\xdb\x23\xc0\xb8\x3f\xdf\x80\xda\xa1\xb6\x50\x6e\x62\x81\xa6\x19\x17\x0a\x36\x5c\x98\x75\x83\xec\x08\x60\x91\xaf\x13\x1a\x59\xb7\x1d\xff\x72\xf2\xf9\x08\x8c\x7f\x39\x01\x69\xa3\x81\xb6\x09\x8d\x3e\x27\xa1\x91\x47\xa8\xa9\x36\x9f\xf0\xe9\xe7\x24\x7c\x7a\x0b\xe1\x0b\x54\x24\x26\x8a\x68\x6a\xe3\x57\xcb\x30\x9c\x24\x3c\x8f\x6d\x29\xd1\x24\xc2\x19\x53\x28\x36\x24\x2a\x92\x5a\x55\x48\x9e\x08\x9e\x67\xd2\x2e\x02\x04\x30\x27\x6b\x4c\xca\x9f\xfa\x13\x97\x54\x7a\x55\xf9\x98\x70\xb6\xa1\xdb\x5c\x18\xd4\xbd\x0a\xd6\x2f\x4e\xe5\x27\xf0\xca\x94\xb7\x51\xe4\x4e\x6f\xad\xcc\x76\x77\x61\x68\x9c\x2b\x6e\x82\x8e\xb2\xed\x7d\x99\x6a\x54\x37\x6f\xaf\xa8\x40\xbe\xa2\x0c\x1f\x15\x92\x76\xe9\x3d\xa0\xab\xb2\xd4\xda\xc0\xfb\xb1\x64\xcc\xab\x30\xfe\xd1\x9f\x71\xaf\x0f\x6c\x05\x61\xca\x49\xe3\xf0\x85\xad\x1b\xda\x1f\x18\x67\xf8\x65\x85\xcb\x2f\x10\x3e\xb2\x3a\x59\x76\xe1\xac\x50\x74\xd6\x7a\x1f\x53\x01\xe2\x56\xca\xaa\xb6\x41\xc4\x73\xa6\x2a\x6c\x5e\x05\xf7\xb1\x94\x05\xfa\x56\x2c\x13\xce\x62\xaa\xcd\x68\xd4\xfd\x94\x48\x4f\x5b\xbd\x9f\x58\x18\x3e\xe3\xaa\x57\x3b\xad\x59\x9a\xfe\x23\x27\x89\xec\x85\xf0\xfa\xc1\x25\x6e\x4a\x0d\x7f\x0d\xc7\xc7\xbf\x5a\x2c\x8d\xe4\x73\x2f\x6c\xad\xc4\x75\x10\xef\xe8\x77\xe0\x1d\xdd\x82\xf7\xf4\x77\xe0\x3d\x2d\xf1\x5e\xa2\xe4\xb9\x88\xd0\x28\x76\x1a\xc9\x89\x35\x81\x9b\xa3\x4c\xf6\x98\x4e\x4c\x0a\x29\xfb\xa1\xe9\x64\xa9\x63\xad\x08\x35\x93\x32\x5a\x47\x1c\x00\xef\x87\x81\x2e\xf2\x55\x86\x2c\x96\xcf\x59\x08\xaf\x7f\xb5\xc1\x25\x78\x86\x42\x51\xac\xe2\xea\xe5\x62\xf2\x0b\x67\x38\x8b\x91\x29\xba\xa1\x25\x6b\x5a\x4c\x2d\xe5\x6c\x53\x3b\x55\xd0\x61\x53\x67\xd3\x80\x9b\x14\xfa\x52\x67\xd4\x10\x1e\x2c\xf3\x35\x3c\x7c\xdf\xb2\xe4\x07\xe7\x90\xd1\x9d\x11\xe7\x19\x37\xc7\xee\x43\x7d\x74\x6f\xea\xa3\xcf\x48\xfd\xf4\xde\xd4\x4f\xef\x46\x7d\x6e\x32\x97\x97\x5e\x4d\x30\xda\x03\x13\xce\x14\xa1\x0c\x45\x99\xf1\x64\x99\x04\x28\x33\x49\xa0\x1a\x0c\xea\xbc\x60\x4f\xba\x59\xb6\x9d\x81\x2c\x4c\x77\x96\x5e\x91\xad\x53\xb1\x7e\xc6\xbd\xad\xa0\x95\x2c\xa5\xd0\x95\x40\x4b\x45\xa2\x77\x1e\x88\xf6\x3c\xb2\x25\x0a\xc7\xca\xca\x17\x82\x12\x85\xc8\x13\x81\x46\xca\x05\x4f\x68\x54\xe5\xc2\x32\x74\x96\x74\xcb\x88\x53\x90\x56\x34\x45\x9e\xab\x10\x16\xab\x93\x47\x17\x66\xf9\x45\x16\x13\x85\xfe\x71\x27\x22\x2e\x79\xa2\xff\xb3\x50\x35\xa2\x0b\xca\x2a\x1d\xce\xd8\x12\xc5\x15\x8d\x3c\xf5\x19\x05\x9e\x11\x15\xed\x9a\x8a\xd5\x65\x2a\x97\xa8\x59\x71\xf9\xd0\x9f\x57\x84\xaa\xe7\xcc\x67\x5e\x86\x70\xac\xa5\x75\x7b\x66\x97\xdb\x43\x51\x5d\x7c\xb1\xa0\x07\x02\x78\x1c\xff\x3d\x97\x2a\x45\xa6\x2c\x96\xc9\x8e\xb0\x2d\xce\x58\xc3\x84\xcd\x04\xe1\x78\x54\x47\xb2\x29\x0e\x4d\x38\x4f\x62\x7e\xcd\x42\x38\x1d\x0e\xcb\xe2\x67\xc1\x6a\xb2\x21\x9c\xd4\xad\xf8\xff\x90\x54\x81\x16\xeb\x02\x53\x2e\xf6\xe3\x84\x88\xf4\x29\xdd\xee\x5a\x82\x99\xae\xef\x95\x76\x91\x30\x34\x50\x87\xe4\xd1\x7b\x5e\x0b\x6a\x32\x84\xd1\x5a\xa0\xe7\x20\xba\x29\x68\xc1\x0f\xf0\xf0\x7d\x6b\xb0\xfa\xf0\x57\xd3\x87\x3c\x82\x94\xb2\x5c\xd5\x31\x8f\x4a\xd0\xc8\x4a\x6d\x8f\x5f\xa2\x44\x71\x65\xc2\xa9\x80\xd1\xbb\x32\xd3\x8d\xa8\x66\x79\x30\x9d\x2c\x4b\x99\x15\x51\x54\x2a\x1a\x85\x30\xbe\x42\x41\xb6\x65\xb4\x2e\x50\x50\x1e\xbb\xea\x99\xea\x26\xd9\xc6\xa8\xd9\x93\xd6\xe8\x46\x19\xd5\xf0\x65\xf5\xde\x62\xdd\xd5\xc0\x38\xaa\xba\x0c\x27\x07\xfa\x21\x51\xe6\x27\x9a\x22\x93\x3e\xb4\x15\xb4\x28\x93\x07\x33\x50\x5d\x66\x2b\x83\xa7\x19\x11\x54\x72\xf6\x3c\x43\x41\x14\x17\x21\x3c\xd1\x39\x07\xc5\x6a\x47\x98\xcb\xa9\x63\xef\x79\xd9\xe0\x7d\x6e\x73\x53\xe6\x58\xfb\x6f\xa5\xb5\x9d\x31\xf6\x4f\x66\xec\x19\xbb\xaf\xad\xcb\x3c\xf1\x9f\x31\xf5\x1c\xa5\x6c\xda\xb9\x5d\x45\x6f\xcf\x59\x1d\x35\xf9\xa8\xb0\x49\x35\x06\x1a\xc9\xbb\x47\x41\xaa\xea\xb2\x13\x19\x24\xee\x70\x14\xf1\x34\x25\x2c\xf6\x06\x26\x80\xe1\xc9\x5b\x12\xc7\x6f\xcb\x66\xfd\xad\xe2\x6f\x23\xb7\x7b\x6c\x9d\x2f\x9c\xec\x9f\x8d\x5d\x80\xbf\x3c\x18\xac\x29\x1b\xac\x89\xdc\xb5\xf6\x30\xda\x71\x9d\x2b\xdf\x4e\xe6\x2f\x96\xab\xe9\xe5\xe3\x87\xef\x6b\xa5\x7e\x00\xf8\xe1\x07\x18\xa0\x8a\x06\x18\x49\xfd\xd7\xb7\xdc\x3b\x68\x36\x34\xc1\x06\xe7\x3d\x73\x22\xda\x30\xfd\x17\xec\xf2\xcc\x9c\xea\xb5\xd9\x66\xca\x64\xda\x03\x6c\xbf\x4e\x09\x65\xbf\xb6\x96\xa5\xee\x2d\x1e\x3f\x7c\x5f\x37\x1a\x6e\x5b\x55\x7e\x04\x6e\x29\x67\x25\xd8\xa5\xf9\xd5\x84\x4a\x79\xac\xab\xfa\x70\x38\xfc\x66\x38\x3c\x6e\x6c\xf2\x6b\x86\x22\x04\xc1\xb9\x6a\xec\x6c\x4d\x3b\xde\xde\xa9\xc5\xde\x71\xfe\x4e\xf6\x63\x23\x3e\xc9\x15\x0f\x04\x26\x9c\xc4\x28\x3e\x51\x11\x2d\x3c\x81\xa6\xd0\x56\x8d\x12\x74\xbb\x45\x21\x1f\x67\x5c\xaa\x7e\x6e\xfa\x9d\x16\x50\x46\xd4\xee\x71\x35\x97\xf4\xdb\x91\xd0\x2f\x9d\xba\x7f\xd0\x9b\x5b\x48\x89\x09\xf6\xc7\x03\x9e\xa9\x01\xb9\x96\xc6\xdf\x34\xd7\x94\x51\x05\xc1\x15\x04\x81\x31\x1b\xb8\x66\xd3\x51\xfd\x01\x82\x40\x14\xbc\x74\x04\xa5\xd9\xd5\xa6\x83\x5b\x0d\x09\x20\x72\x46\xe4\xe3\x86\x49\xa4\x6d\xe9\x1a\xde\x29\xf7\xf2\x8a\x7a\x11\x59\x58\xc1\xfa\x6a\x73\x19\x00\x19\x59\x27\x18\x3b\x3d\x5c\x73\x5f\xe6\x02\x2f\x73\xc6\x74\xaa\x38\x04\xd5\x11\x27\x60\xa7\xc9\xee\x68\xb9\x15\xf2\x23\x0e\x76\xa0\x2a\x95\x17\x96\x36\x77\x16\xbf\xca\xfa\x80\x51\x2e\xa8\xda\x17\xd7\x53\xf0\xda\x02\x3d\xe5\x52\x2d\x9f\x40\xe9\x6a\xde\x2d\x4c\x81\xa6\x7d\xd9\x34\x23\x69\xb9\xba\x10\x5c\x0b\x5e\xf5\x64\xa3\xc6\x46\x71\xa2\xbc\x6f\x80\x07\xb3\x0d\xbc\x76\x6e\x20\xbe\x06\xff\x6e\xc1\xfc\xea\xb9\x43\x53\xaf\xe4\xed\x85\x44\x71\xee\xa4\x62\x30\x53\xd9\x19\x91\xf8\xed\x37\xae\xde\x3b\x82\xcc\x49\x90\x10\xdc\xf8\x21\xb3\xcf\x53\x7b\x61\x92\x24\x10\xec\x81\x5c\xcb\x40\x6b\x7d\xcd\xb9\x92\x4a\x90\xcc\x03\xfe\x43\xfc\xbf\x45\x54\x9a\xa1\x03\x02\x84\x87\x3f\xde\x8d\x72\x47\xb3\x7c\x0b\xe9\xb6\x19\x5b\xc5\x73\x36\xbe\xd0\x99\xa2\x6d\xeb\xb6\x57\x2e\x88\xda\x85\xd0\x1b\x94\x1e\x7f\xc9\x9d\x40\x09\x2a\xc7\xd1\xcb\x96\xb6\xfe\xd6\x4d\xb0\x80\xe9\xec\xc8\xa4\xcc\x53\xd4\x00\xb6\xf5\x38\xe7\x51\x6e\xfa\xfc\x4a\x95\xba\x43\x42\x7f\x29\x80\xe9\x66\x83\x91\x0a\xc1\xbd\x22\xb5\x04\x28\x8b\x68\x46\x12\x3f\xa2\xcb\x21\xf2\xc8\x0f\x5c\x8c\x46\x7d\x92\x92\xdf\x38\x23\xd7\xba\x84\xa6\xce\xbe\x6d\x93\xfc\xbb\x52\xa9\x64\x58\x33\x7c\x40\x4f\x46\x0e\xea\xaa\xca\x4a\x66\x03\x09\x23\x19\x14\xf9\xaf\x9e\x59\x0f\x48\xde\x29\xfb\x6d\xd2\x77\x71\x6d\xe5\x94\xc6\x4d\x74\xef\xd0\x72\xe6\x0e\xd8\x73\x14\xf7\x81\xa6\x32\xe2\x57\x28\x16\x3c\x49\xa6\x2c\xce\x38\x65\xaa\x03\x6c\xa9\x88\x50\x2b\x4c\x30\x45\x25\xf6\x4b\x94\xb2\x6e\x8f\x3d\xb8\x7c\x9d\x52\xf5\x55\x6b\x47\x84\xed\x35\x19\x6a\xa2\xde\x72\x59\x41\x43\xe8\x7d\xa5\x4d\x62\x33\x65\xc7\xc5\xde\x28\x0c\xbd\xe4\x7a\xc0\x41\x9d\xc7\x3f\x50\xe4\xae\xae\x1b\x25\x03\x56\xc6\xbf\xc1\xd7\x7a\xf2\xa4\x39\xa9\xf2\x39\xb8\x93\xaf\xc7\xc7\x8c\x6d\x05\x4a\xc7\x7d\x66\xd9\x42\x70\xc5\x23\x9e\x84\xa0\xa2\x3a\xb1\xfd\x24\x78\xba\xe0\xc2\x3c\x02\x1e\xd5\x85\x6d\xc5\x3b\x16\x27\x34\x16\xb3\xac\x9c\x0e\xea\x47\x0b\xd3\x64\xfd\x5f\xa0\x9c\xf9\xd9\xbf\x49\x2f\xdf\x0d\x3b\xf4\xe2\x2e\x96\x7a\x71\x1f\xea\x4e\xe7\x67\x23\x6d\xab\xcb\xbc\x23\x9f\xb5\x55\x53\xf0\x75\xa8\xb6\x77\x32\xe9\xb0\x58\x31\x53\xf1\xf7\xed\xa3\x47\xa7\x8f\xca\xd5\xa5\xbd\xae\xf2\x08\xea\x4e\xe1\x09\xaa\xb1\x52\xd6\x7e\xfd\x62\xd9\x55\xb0\x0b\x64\x43\xc0\x81\xd2\x0b\xa3\xe9\xfc\xec\xcf\x20\x61\x8b\xf9\x4e\x11\x9b\x7a\x98\x46\x72\x9a\xac\xdb\xb2\x25\x44\xcf\xda\x73\x4e\xe2\x33\x92\x10\x16\x51\xb6\x7d\x39\x0a\xc3\x7a\xa1\x98\x58\xdb\x62\xda\xcb\x62\xf9\xff\x6b\xf8\x3f\xfa\x1a\xbe\xd1\x11\x37\x1a\x12\xed\x07\x95\xfd\xe7\xba\x86\xb1\xae\xc7\x3a\x87\xfc\xa0\x38\x70\xc0\x07\x5c\x37\x19\x0b\x56\x5f\x76\x4c\x93\x75\x01\x52\x3c\x8b\x6e\x5d\xb2\x58\xe2\x1b\x2e\xae\x89\x88\xeb\x9c\x44\xc4\x16\x95\x91\xa4\x89\xaf\x40\xe4\x40\x54\xfd\x45\x23\x8b\xd5\xe1\xf7\x74\xb5\x5a\x54\xc2\xb7\x11\xdc\x59\x0d\x4d\xa2\x1d\xcd\x61\xc9\xc4\x2d\x6c\xc0\xbd\x0b\xc4\xf3\x5c\x65\xb9\x8d\x31\x3d\x1f\xbc\x10\x45\x1b\xe7\x02\xef\x94\xca\xc2\xc1\xc0\x5c\x87\x4c\x93\x75\xff\xfc\xd9\xd2\xb4\xcd\x83\x23\x68\xbd\xee\x31\x3f\x83\x17\x97\x73\xff\x1d\x0e\x8b\xeb\xe8\x4e\xfc\xb4\x30\xbe\x5c\x4c\x60\x76\x5e\xbd\x8a\xe0\xbf\x2c\x05\x30\xbd\xd1\xa8\x4a\xe4\xc5\xe4\x64\x31\x36\x9a\xfc\xa0\x7c\x11\xc4\x77\x54\xed\x04\x9e\xc4\xb5\x3f\x78\xae\xe9\x31\x35\x16\xac\x7c\x67\x42\x4b\x5c\x02\xf6\xef\xcb\x52\x8b\x95\x8e\x47\xa3\xb7\x5c\xf2\x1d\x7c\x91\xc3\xd1\xd2\xa7\xf0\x54\xd2\x38\xfa\x57\x00\x00\x00\xff\xff\xc6\x3d\x5c\x8f\x1f\x28\x00\x00")

func assetsClusterYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsClusterYml,
		"assets/cluster.yml",
	)
}

func assetsClusterYml() (*asset, error) {
	bytes, err := assetsClusterYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/cluster.yml", size: 10271, mode: os.FileMode(420), modTime: time.Unix(1485238853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsPipelineYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x57\xeb\x6f\xdb\x36\x10\xff\xee\xbf\xe2\x9a\x15\x08\x10\xc0\x76\x9c\x47\xdb\x10\x18\x06\x25\xf1\x32\x03\x6e\x6c\x58\x49\xfb\x61\x18\x0a\x5a\x3e\x5b\x6c\x24\x52\x20\x4f\xf1\xb2\xa2\xff\xfb\x40\xea\x61\xca\x8f\xcc\x49\x31\x0c\xf3\x27\xeb\xee\xc7\xbb\x1f\xef\x45\xb2\xdd\x6e\xb7\x82\xcf\xe1\x1d\xa6\x59\xc2\x09\x7f\x55\x3a\xe5\xf4\x09\xb5\x11\x4a\x32\x38\x3c\x39\xee\x1d\xb7\x8f\x2f\xda\xc7\x17\x87\xad\x31\xd7\x3c\x45\x42\x6d\x58\x0b\xe0\x46\xd0\x6f\xf9\xf4\xde\xa0\xb6\x5f\x00\x77\x4f\x19\x32\x08\x49\x0b\xb9\x70\x82\x6b\x34\x91\x16\x19\x39\x43\x05\x1a\x2c\xbc\x5e\x3b\xc1\x4c\xbd\x60\xad\x85\x03\x29\xc8\xf2\x24\x81\xb9\x56\x69\x07\x46\x32\x79\x02\x8a\x11\x6e\x79\x8a\x1d\x90\x8a\xdc\xd7\xfd\x64\x58\x3b\xb9\xd4\x5c\x46\xf1\x0b\xdc\x14\x0b\x4a\xf5\x9c\xe7\x09\x31\x38\x48\xb9\x21\xd4\x07\xb5\xd5\x3b\xf5\x80\xb2\x30\x7a\xab\xfa\x51\xac\x18\x90\xce\x71\x0f\x2f\x21\x46\x1a\xa9\x03\x03\x82\x54\x2c\x62\x82\x44\xa9\x07\x30\x2a\x45\x8a\x85\x5c\x40\x22\x1e\x10\x2e\xa6\xbd\x0f\x17\xbc\xf7\xee\xfc\xec\xdd\xd9\xe9\xf9\xc9\xc9\xf9\xbb\xde\xfc\xfd\xf4\x14\xa7\xb3\xb3\x33\xde\x3b\x3f\xed\xf1\xb3\x93\x0f\xef\xf9\x1c\x46\x41\x4e\xb1\x23\x03\x4b\x41\x31\xf0\x28\x42\x63\x6c\x94\x6c\xb4\x3a\x70\xe3\x22\x16\x13\x65\x86\x75\xbb\x0b\x41\x71\x3e\xed\x44\x2a\xed\x1a\x24\x12\x72\x61\xba\x64\x17\x9b\x16\xc0\x65\x2e\x92\x99\xe3\xbe\x6b\x17\x55\x30\x12\x21\xf3\x3f\xaf\x94\x24\x2e\x64\x11\x14\xb7\xf6\x4a\xa5\x59\x4e\xb8\x97\x89\xcb\xfb\xc1\xf0\xfa\xcb\x4d\xff\xb6\x3f\x09\x86\xbd\x2f\xe1\xc7\x60\x38\xac\x0d\x0d\x52\xbe\xf8\x67\x13\x7c\x69\xba\x91\x9a\xe1\xd4\x2e\xe9\x7e\xe5\x8f\x9c\xa9\x0c\xe5\xd7\xd9\x43\xfb\x83\x35\x75\xa5\x66\x38\x16\x19\x26\x42\xe2\x65\x1e\x3d\x20\xed\x53\x04\x05\xd2\x06\x2d\x37\x08\x73\xa5\x1b\x76\x80\x6b\x12\x73\x1e\x91\x69\x4d\xd0\xa8\x5c\x47\xe8\x7a\xc1\x62\x1c\xf5\x89\x4a\x1a\xcc\x83\xcf\x21\x63\x83\xe0\x23\x63\x56\xe3\x14\x63\xad\x32\xd4\x24\x8a\x95\xf6\x17\x18\x93\xa7\x68\x01\x63\x95\x88\xe8\xe9\x5a\x45\x79\x8a\x92\x2a\x3d\x40\x48\x9c\xb0\x29\x6a\x43\x7f\x3e\xc7\x88\x18\x04\x49\xa2\x96\xb5\xdc\x3a\x10\x32\x12\x19\x4f\x98\x27\x04\x08\x51\x3f\x8a\x08\x9b\xc2\x36\xd4\x21\xec\xf0\x94\xff\xa5\x24\x5f\x1a\x5b\x20\x1e\x2a\x88\x5c\x68\x5a\xfe\x2a\x43\x86\xad\x68\x97\xaa\x31\xa7\x98\xc1\x41\xf7\xa0\xfa\xb6\xbb\xf1\xf6\xd9\x2e\x24\x4f\xb6\x59\xd9\xca\x71\xdb\x14\xcc\x6a\x07\xbb\xa2\xb0\x35\x0e\xcf\x45\xa2\xe6\x0e\x07\x47\x07\x0d\x79\x95\xbd\x75\x4d\x63\xf0\x9d\xb4\x7b\xc7\xed\xde\xfb\xc3\xb5\x62\xfa\x7f\xe6\x38\x2b\xd9\xff\x17\x69\xae\x7c\xff\x78\xa6\x37\x49\x36\x6a\x98\x1d\xed\x95\x65\xd8\x19\xd7\x5d\x1e\xcc\x29\xbb\x41\x1a\x4d\xbf\x62\x44\xcf\xa8\xca\xf2\xd9\x8a\x28\x26\x4b\x89\xa8\x66\xcf\xbf\xc2\x74\x9c\x6f\x63\x5a\xbb\x58\x5b\xc0\xb5\x64\x7c\x69\x98\x39\x65\x8c\xf9\xd9\x3a\xfa\x51\x1e\x47\xeb\x69\x4a\x54\x3e\x9b\xbb\x0b\x86\x5d\xb0\xae\xc6\xc8\x6c\xc8\x04\x4f\xd9\x98\x1b\xe3\x95\xdf\x6b\x03\xb6\xab\xb5\xdd\xdc\x1e\x6b\x65\x03\xb6\xd1\xd6\xf5\x60\x67\xac\x84\x94\x47\x46\x86\x72\x66\x46\x92\x35\x47\xff\x8e\xe6\x2f\x3a\xe1\xcd\x04\xe7\x85\xd5\x90\x78\xf4\x60\x85\xa5\x7e\xed\x04\x12\xc9\x0c\x78\x96\x25\x22\x72\x81\x2a\x41\x65\x6f\xbb\xe9\x03\x6f\x6e\x90\x02\xa2\xa6\xf7\x4e\xa0\x2b\x70\x50\x9d\x53\xab\xb4\x14\xbb\xba\x1a\x5d\xf7\xc7\x83\x71\x7f\x38\xb8\xed\x97\xaa\xbe\x7c\x14\x5a\xc9\x66\xb3\x15\x68\x47\xb9\xbe\x18\xd4\x4a\xff\xa4\xf7\x30\x9e\xb8\x86\x16\x67\xb9\x07\x72\x82\x6a\x4b\x6b\x05\xb9\x93\xe2\x9d\x48\x51\xe5\x34\x90\x1f\x85\xcc\x09\x0d\x83\xde\x71\x0b\xa0\x9a\xc8\x5b\xd3\x56\x2b\x59\xf5\x6f\x3d\x75\x7e\xe6\x77\x64\xce\x86\x35\xd0\xb2\x19\x70\xff\x20\xf0\x62\x1e\x12\x5f\xf8\x53\xb0\xc8\x7a\xb1\xc7\x7a\x8b\x45\xb3\x18\x7f\xc6\x0f\x64\x96\xd3\x2a\x61\xf0\xfb\x1f\x5e\xd5\x6e\x35\x52\x99\xb1\x3b\x1e\xcc\x9a\x9d\x77\xc5\x09\x17\x4a\x3f\x6d\x59\x04\x30\x5a\x4a\xd4\x0c\xee\x62\xa1\x67\x63\xae\xe9\xa9\xa1\x5d\x75\x48\xef\xb0\xa1\x18\x6b\xf5\x28\x66\x76\x65\x71\xfb\xf5\x94\xa3\x9c\x1a\xec\x1b\x93\xc1\x27\x5f\x00\x3d\xf5\x95\x92\x73\xb1\xc8\x35\xdf\x9c\x1e\x25\x4d\x57\x34\xab\x67\xc6\x5a\xff\x67\xaa\x81\xb0\x82\x06\xa2\xbc\xf8\xfb\x18\xef\x6a\x5f\xbb\xaa\x2f\xd0\x0d\xa4\x93\x78\xc0\x49\x2e\x47\xda\x45\xa0\xb7\x96\x5f\x5b\xf4\x22\x79\x36\xc1\x0d\xe0\xab\x52\xe9\x2a\x75\x5b\x88\x82\xcf\xe1\x4b\x53\x58\x8f\x0c\x4f\xbf\x56\x82\xfb\x27\x71\x8f\xf4\x97\xfb\x7e\x49\xfe\xcb\x9e\xf4\xc6\xe6\x46\xab\xee\x4a\x4b\x45\x25\x24\xa5\x37\x26\x4b\x78\x5a\x0b\x86\xaa\x18\xae\xa5\xfd\xcd\xc7\x42\xab\x20\xdc\xbc\xde\xdf\xeb\xf2\xc2\xf5\x89\x27\xb9\xe5\x16\xe6\xd3\xfa\x75\x15\x29\x69\xec\x44\xb0\x57\xab\xe2\x92\xe5\x5e\x5a\xab\x37\x4a\xac\x52\xfc\x45\xe3\x42\x28\xf9\xf3\xdb\x6f\x6e\x50\x4d\xdc\xd7\xf7\x9f\xba\x59\xb1\x37\xd3\x7d\xfb\xcd\xdf\xeb\xf7\xee\xa3\xc0\xe5\xe6\x33\xa5\x66\x54\xbe\x75\x7d\xfe\xaf\x25\x59\x9d\xfb\xcf\xf2\xb4\x74\xba\x6f\xbf\x55\xbe\xbe\x6f\xa7\x56\xbf\x97\x2c\xbb\xd6\xdf\x01\x00\x00\xff\xff\x12\xec\x66\x74\x63\x10\x00\x00")

func assetsPipelineYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsPipelineYml,
		"assets/pipeline.yml",
	)
}

func assetsPipelineYml() (*asset, error) {
	bytes, err := assetsPipelineYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/pipeline.yml", size: 4195, mode: os.FileMode(420), modTime: time.Unix(1485758542, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRepoYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x92\x4f\x6f\xda\x40\x10\xc5\xef\x7c\x8a\x49\x14\x29\x52\x24\x13\xd3\x4b\x9b\xbd\xb9\x84\x44\x48\x51\x6b\x61\x48\xce\x93\x65\x80\x15\xfb\xc7\x1a\xcf\x96\xd2\x88\xef\x5e\x79\x43\xa8\xa9\x5b\x9f\xac\xf7\x7e\xf3\xbc\xfb\x3c\x59\x96\x0d\x8a\x97\x6a\x4e\xae\xb6\x28\xf4\x10\xd8\xa1\x3c\x13\x37\x26\x78\x05\xd7\x9f\xf2\x51\x9e\xe5\x77\x59\x7e\x77\x3d\x28\x91\xd1\x91\x10\x37\x6a\x00\x30\xa3\x3a\x7c\x43\x47\xed\x3b\xc0\x7c\x5f\x93\x82\x4a\xd8\xf8\x75\x12\xee\xa9\xd1\x6c\x6a\x49\x31\x2d\x0b\x1e\x1d\x0d\x66\xd4\x84\xc8\x9a\x52\xc4\x44\x37\xad\xd3\x4d\x28\x5e\x2a\xa5\x26\xe3\x99\x52\xad\xd3\x18\x09\xbc\x4f\x76\xc9\xa1\x26\x16\xf3\x3e\xd9\x3e\x7f\x80\x74\x0c\xb8\x98\xd1\xea\x74\xaa\x1e\x53\x06\x6b\xf4\x7e\x4e\x3f\xe5\x63\x1e\xa0\x7b\xcd\xfc\x4b\x36\xca\xb3\xd1\xe7\xeb\x93\x5b\x09\x0a\x39\xf2\x9d\x81\x0c\x2a\xb3\x54\x50\x58\x1b\x76\x65\x6c\x36\x65\xb4\xf6\x64\x02\x4c\x56\x2b\xd2\x72\xf4\x3b\x7a\xc9\xc6\x6b\x53\xa3\x55\x1d\x11\xd2\x5d\xe1\xf2\xe6\xb2\x23\x8e\x83\x5f\x9a\xd4\xd9\x39\xc9\xfe\xc9\x6c\xe9\x5c\x04\xc0\x5d\xa3\xaa\x54\x67\xc1\x5e\xc1\x45\x15\x5f\x01\xd9\xab\x56\x37\xe8\xd4\xd5\x5b\xaa\xb3\xd0\x3a\x44\x2f\xd3\xe5\x41\xc5\x86\xf8\xf6\xa6\x13\x53\xe8\xbf\xbf\x96\x01\x69\x56\x8f\x24\xf7\x61\xe7\x6d\xc0\xe5\x82\xed\x43\xe0\x27\xdc\x13\xf7\xb0\xaf\x28\x7a\xf3\x48\x32\x75\xb8\xa6\x7f\xbb\xe3\x0d\xe9\x6d\x9a\x2e\x7e\xa0\xb1\xf8\x6a\xac\x91\x7d\x8f\x2d\xe3\x7f\x42\xa6\xde\x88\x41\xa1\x14\xb1\xa8\xdb\x13\xf5\x98\x77\x39\x11\x25\xb2\xf4\xfc\x71\x70\xb5\xa5\xf3\x8c\xef\x51\xea\x28\xa7\x55\x5e\xf0\xf1\xef\x9c\x2d\xee\x82\x2d\x84\x15\xc8\x86\x80\xa9\x0e\x09\x78\x46\x1b\xe9\xd8\x76\xaf\xe1\xe1\x72\xcb\x43\xd2\x3c\x3c\x3a\x33\x5a\x9b\xe0\x0f\x43\x74\xf8\x2b\x78\xdc\x35\x43\x1d\xdc\xed\xd5\xdb\xc7\x9e\x1e\x06\xbf\x03\x00\x00\xff\xff\xa7\xe7\x4b\xa3\x81\x03\x00\x00")

func assetsRepoYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRepoYml,
		"assets/repo.yml",
	)
}

func assetsRepoYml() (*asset, error) {
	bytes, err := assetsRepoYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/repo.yml", size: 897, mode: os.FileMode(420), modTime: time.Unix(1485238853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsServiceYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\x98\x1a\x07\xe4\xae\x38\xa7\x4e\x8a\x3b\x34\x7c\x73\x6d\x67\xe3\xdd\x38\x30\x62\x27\x7d\x28\x8a\x82\xa6\xc6\x16\x51\x8a\xd4\xf2\x4f\x5b\x77\xd1\xef\xbe\x20\x69\x49\x94\xad\x66\x9d\xee\xe6\x29\xe6\xfc\x38\xbf\xf9\xcb\x19\x0d\x06\x83\xde\xe8\xdd\x72\x85\x45\x29\xa8\xc5\x6b\xa5\x0b\x6a\x1f\x51\x1b\xae\x24\x81\xb3\xcb\xe1\xc5\x70\x30\xbc\x1a\x0c\xaf\xce\x7a\x0b\xaa\x69\x81\x16\xb5\x21\x3d\x80\x25\xea\xcf\x9c\xe1\x1d\x2d\xd0\xff\x04\x58\xed\x4a\x24\xb0\xb4\x9a\xcb\x6d\x38\x98\xa0\x61\x9a\x97\x36\x68\xf2\x38\x50\x1b\x30\xf1\x5a\x0f\x60\x56\xd0\x2d\x3e\x68\x71\xca\xed\x89\x62\x9f\x50\xc7\x2b\xf0\x70\x7f\xdb\xf0\x2f\x94\xb6\xa7\x68\xf0\x38\xc0\xaf\xa5\x32\x98\x01\x97\x89\x1d\x1e\xb8\xa1\x4e\x58\x02\x67\x6f\x86\x6f\x86\x67\x8d\xf2\x1b\xa4\xc2\xe6\x53\x99\x95\x8a\xcb\x93\x68\x2a\x2c\x58\x05\x16\x8d\xad\x78\x20\x0f\xaa\x0e\xe8\x5e\xc5\xd3\x84\x71\x5c\xba\x53\x68\xc6\x8b\x07\x70\x92\x5b\xe3\x79\x34\x7a\x12\x84\x8d\xd2\xc0\x94\xb4\x94\x4b\xd4\x07\x4c\x17\xa9\x5b\x73\x2c\x94\xde\x9d\xc2\x13\x91\x9e\x84\x0a\xa1\x18\xb5\xe8\xff\xf7\x24\x9c\x4a\xd4\xf0\x6f\x2e\x61\xce\xdf\xfe\xe7\x80\xed\xf5\x30\xa5\x9b\xa0\xe1\x1a\xb3\xb1\x72\xed\x18\xde\xb9\x62\x7d\x64\xe8\xe5\x59\x47\xed\x04\xa0\xaf\x9e\xda\xbd\xe0\xb8\xa0\x4e\xb2\xdc\x67\x73\xa7\x9c\x86\xe9\x78\x59\x85\xfb\xbc\x07\xb0\xa0\x36\x5f\x50\x6b\x51\xcb\x94\x75\xac\x8a\x82\x4e\x50\xf0\x82\x5b\xcc\x6e\xb9\xb1\xc7\x84\xfe\xd4\xd3\x95\xd4\xe6\x31\xc4\xca\x45\xd7\x6d\x8e\x29\x47\x6a\xb9\x37\xdc\x5f\x44\x89\xfa\xde\x09\x5c\x68\xae\x34\xb7\xbb\x1f\xba\x9c\x10\xae\x72\x84\x72\x8f\xf7\xc4\x9e\x46\x3b\x81\xb0\x46\x2e\xb7\x40\xb3\x0c\xb3\x8a\x5e\xec\x39\xda\xf4\x17\x3d\x80\xc7\x92\xcd\xb2\xe7\x34\xa3\x57\xf7\x99\x0a\x17\x5c\xe3\x45\xe9\x7b\xc4\x57\x91\x3f\x0f\xca\x7a\x00\x53\x66\xc6\xc2\x19\x8b\xfa\x9f\xd0\x3c\x65\x06\xf6\xea\xbc\x24\xc3\x52\x28\x5f\x5f\xe7\x91\x69\x2a\xd6\x55\x08\x47\xed\xb4\xfd\x1d\x57\x46\x5a\x56\x90\xe9\xed\xdb\x3a\x80\xa1\xac\xad\xa5\x2c\x0f\x22\x4b\xf5\x16\x2d\x6c\xb5\x72\x65\xb0\x68\xac\x64\xc6\x3d\x4b\x78\xef\x6e\xa8\x39\xaa\xa8\xfe\xb5\x24\xe4\x4e\xd9\x7e\xfc\x09\x30\x88\x47\xd3\xdf\x1d\x15\xa6\x3e\xad\xcf\x7f\x55\x5c\xf6\x09\xbc\x87\x7e\xff\xbf\xf0\xe2\x1e\x37\x69\x91\x7e\x48\xd1\xfd\xde\x3d\x1a\xe5\x34\xc3\x40\x3e\x65\x66\xdf\x4c\x69\x4c\x46\xef\x96\x84\x4c\xc7\x4b\x42\x96\xc9\x6b\xb6\xd0\xaa\x44\x6d\x79\xbc\xe9\xff\x5a\xf9\xf3\x7f\xde\x96\x59\x88\xd1\xa3\x0f\x19\x81\x17\x4b\xb7\x86\x7f\xfd\xd1\xe4\xfa\xfb\x1e\xdb\xea\xdd\x68\x72\x47\x57\xd7\x60\x9f\xcd\x02\xa5\x1d\x2b\xb9\xe1\x5b\xa7\x69\x48\x52\xcd\x3b\xa7\x5f\x79\xe1\x8a\x05\x6a\x86\x5e\xdf\xc5\x70\xd8\xc8\xb8\xf4\xb2\xf8\xe4\xee\x6a\x48\x05\xb8\x55\x34\x7b\x4b\x05\x95\x6c\x3f\x7f\xaa\x50\x05\x5f\x36\xcd\x89\x3f\x6b\xe7\xaa\x25\x1a\x57\x2f\x48\x18\x5c\x2d\x8f\xfc\x49\x82\x85\x06\x1b\x86\x4c\x0b\xeb\x4f\x5a\xd8\x55\xa8\x9e\x5f\x7c\xf1\xf8\xca\x8d\xe0\x58\xd1\x89\xa8\x65\x4a\x80\x84\x1c\xde\xa9\x90\x87\xbd\xf4\x5e\x09\x3c\xc8\xd5\xe9\xfe\x55\xbc\x7b\x3b\xbd\xae\x93\x48\x57\xd4\x7c\x9a\xe0\x86\x4b\x1e\xfb\x2a\xe0\xe6\x9c\x69\xb5\x7f\xee\xda\x80\x56\x4d\x36\x06\x27\x75\x39\x1b\xcd\x09\xa9\xe9\xeb\x5e\x22\x5d\xc6\x1f\x97\xec\xc8\x18\x57\x04\xc5\x0b\x25\x38\xdb\x4d\x14\x73\xbe\xb0\x9a\x38\x2c\x2d\xb5\xd8\x3e\x1a\xc0\x74\xb3\x41\x66\x09\x8c\x84\x50\x5f\x12\xbf\x17\x9a\x4b\xc6\x4b\x2a\x48\x2b\x69\xad\x9e\x6a\xb4\x20\x33\xe7\xb4\xa0\xdf\x94\xa4\x5f\xcc\x39\x53\x45\x22\x1f\xb1\x76\x49\x7b\xbc\xb1\x86\x34\x06\xef\x45\xde\x47\x02\xfd\x57\xfd\xea\xb7\xf7\x23\xf1\x70\x10\x4f\x76\xb1\x0e\x91\x99\x41\xba\x92\xd4\x37\x3a\x3c\xef\xf4\xfd\x29\xef\xbb\xac\x8e\x7e\x5e\x92\x91\xb3\xb9\xd2\xfc\x1b\x2e\x91\x39\x3f\x7e\x42\x9d\xce\xe4\x56\xa3\x31\x1d\xf8\xf8\xf8\xae\xf1\xe5\xa1\x4c\x50\x63\x39\x13\x8a\x66\xeb\xd0\xa6\x5c\x6e\xc9\x04\x35\x6e\xfd\x6b\xab\x67\xd2\x58\xdf\xbb\xe6\x5a\xab\x22\xed\xe5\xd3\xd4\x3c\x83\xf3\xfe\x90\xf1\x1d\xb7\xf9\xf3\x19\x2b\xc3\x63\xf3\x1e\x45\xe2\x29\x33\x93\x7e\xff\x89\x7b\x37\xcd\xa2\x78\xa2\xa7\x5d\x26\x56\xd3\x83\x40\xff\xa5\xaf\xc0\x1f\x77\x72\xf7\x40\x39\xea\xf6\xce\xb9\x52\xbd\x8f\x0d\x32\x29\xef\xbf\x78\x60\xfd\xa6\xdb\x92\x8e\x4b\x57\x0b\xa7\xc6\xa0\xdf\x31\x05\x81\x33\xab\x1d\x9e\xd5\x92\xf0\x09\xb0\xbf\x58\x7d\x41\x34\x33\x24\x2e\xb6\x2d\xb5\xf1\x2c\x69\x2a\x6d\xe7\xb4\x2c\xb9\xdc\xb6\xc6\xc8\x8d\x32\x36\x3e\xf3\xc3\x24\x90\xa7\x4d\x80\x47\x25\x5c\x81\x86\xc0\xfb\x0f\x47\x6b\x8c\xdf\x04\x8f\x43\x1c\xd3\xd9\x54\x25\x97\xdb\xc7\x4b\x42\xd2\x4b\x3f\xf7\x6c\xb2\x83\x24\x44\xd2\x8d\xd2\x5f\xa8\xce\x6a\xc7\x9e\x35\xac\xda\x6b\x50\x54\x7b\xcd\x51\x64\x24\xec\xc7\x83\xf2\x60\x0a\x85\xa9\x62\xc8\xd1\x7e\x53\x0d\xf2\xc3\xfd\x0e\x9e\x5e\x48\xda\x2b\x61\xb5\x97\xd4\xcb\x75\xa4\xe9\xda\xbb\xeb\x54\x24\x2e\x9d\x9c\x89\xc3\x30\x3c\x37\x11\xb1\x8b\xc7\x39\xb2\x4f\x33\x69\x51\x7f\xa6\x62\x89\x4c\xc9\xcc\x10\xf8\xdf\x31\x26\x0e\x8b\xb4\xbc\xda\x9f\x9e\x1d\x37\xb4\xb2\x8a\x29\x41\xe0\x66\xb5\x5a\x1c\xcb\x57\xbc\x40\xe5\x6c\x4d\xfa\xba\x05\xd9\xad\x72\x8d\x26\x57\xa2\xda\xec\x2e\xf7\xe2\x39\xb5\x2c\x4f\x57\xc5\x1b\x6b\xcb\xb1\xca\x90\xc0\xe5\x70\x38\xb8\xbc\xba\xea\x35\x7d\xf4\xc3\x86\xe8\x34\x2e\x2d\x3a\x6b\x35\x5f\x3b\x9b\x4e\xc3\xdf\x70\x47\x20\xdb\xbf\xbb\x71\x71\xfc\x98\xa1\xa0\xbb\x73\x1b\x5d\xf9\x68\xa2\x2f\xed\x42\x23\xf0\xff\xaa\x61\x1f\x64\xde\xed\x5c\x15\xf0\xe4\xfb\xe8\xa9\xa2\x0b\xb0\xef\xbd\x3f\x03\x00\x00\xff\xff\x3b\xd6\xf7\xe2\x1a\x11\x00\x00")

func assetsServiceYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsServiceYml,
		"assets/service.yml",
	)
}

func assetsServiceYml() (*asset, error) {
	bytes, err := assetsServiceYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/service.yml", size: 4378, mode: os.FileMode(420), modTime: time.Unix(1485238853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsVpcYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x99\x4b\x6f\xdb\x38\x10\xc7\xef\xf9\x14\xd3\x60\x01\x6d\x81\x3a\xb5\xa5\xa4\x68\x78\xf3\x3a\x4e\x63\x6c\x9b\x18\x96\xe1\x02\x69\xf6\x40\x4b\xe3\x58\x88\x44\x6a\x49\x2a\x0f\x14\xf9\xee\x0b\x51\x92\x1f\xb4\xe4\xfa\x15\xaf\x8b\x5c\x6a\x72\xc8\x99\xff\x4f\x43\xce\x48\xad\xd5\x6a\x47\xcd\xef\x6e\x1f\xa3\x38\xa4\x0a\x2f\xb9\x88\xa8\x1a\xa0\x90\x01\x67\x04\x2c\xbb\xde\xa8\xd7\xea\xe7\xb5\xfa\xb9\x75\xd4\xa5\x82\x46\xa8\x50\x48\x72\x04\xd0\x61\x52\x51\xe6\x61\x1f\x19\x65\xde\x4b\x3a\x04\x70\x81\xd2\x13\x41\xac\xf4\xe2\xc2\x02\x54\x66\x02\x8a\x43\x22\x11\x46\x5c\xc0\xa0\xdb\xd2\x0b\xfa\x2f\x31\x12\x70\x95\x08\xd8\xbd\x1e\x68\x86\x21\x7f\x42\x7f\x40\xc3\x04\x65\xb6\x69\x0d\x7c\x1c\xd1\x24\x54\x93\x5f\x7e\xe0\x51\x85\x7e\xee\x52\xcf\x91\x19\x23\x57\x8e\xf5\x36\x25\x31\xb9\xc9\x90\xa1\x82\x91\xe0\x11\x3c\x8d\x03\x6f\x9c\x06\x45\x53\x63\x70\xdd\x2b\xa0\x9e\x87\x52\x9e\x94\x87\xf6\x2d\x60\x5f\x91\xdd\xab\x31\x01\xeb\xdc\xca\x86\xe8\xf3\x64\xa8\xf1\xd9\x9a\x0f\xa8\x7e\xa2\xff\x3e\xd6\x67\x85\x75\xa9\x52\x28\x18\x81\xe3\x3f\xef\xee\xfc\x9f\x8d\x0f\xce\xeb\xfb\xbb\xbb\x93\x55\x7e\x7c\xcc\xff\x69\xbf\xbe\x3f\xd6\x5b\xb6\x38\x93\x4a\xd0\x80\xa9\x39\x8d\x56\x94\x48\x05\x43\x04\x0a\x8f\x34\x0c\x7c\x68\x75\x2e\x7a\x30\x0c\xb9\xf7\x40\xe0\xf9\x44\xff\x7d\x7c\x3e\x49\xa3\x1d\xc4\x5e\x2b\xf0\xc5\x5f\x7a\xae\x92\x96\x5e\xba\xfc\xb1\xad\xcb\xa6\x51\xc0\x69\x7c\x3a\x5c\x3a\xdd\x64\x18\x06\x5e\x06\xa1\x79\xdb\x58\x8b\x54\xf3\xb6\xb1\x63\x52\xf6\xe9\xef\x42\xca\x5e\x93\x94\xbd\x43\x52\x8d\xdf\x8a\x94\xb3\x26\x29\x67\x87\xa4\xec\x83\x26\xd5\xe2\xcc\x0f\xd2\x35\xba\x08\x5c\x51\x69\x1c\xc6\x8c\xd7\xf1\x25\x23\xe4\x9a\xab\xe3\xec\x67\x5a\x1d\xf4\x50\xfb\xdf\x84\x86\xf2\x98\xc0\x8f\x77\x3d\x1c\x55\x1e\xe4\x0f\x60\x59\xff\x94\x6d\x6f\x6f\xb1\xbd\xfd\xeb\xed\x9d\x2d\xb6\x77\x8c\xed\x7b\x28\x79\x22\xbc\xac\x58\x0e\xba\x2d\x32\x93\x22\xcd\xef\x2e\x21\xed\x96\x4d\x48\x71\x71\x77\x05\x8f\x51\xa8\xa0\xa8\xad\x00\xd3\x0c\x04\xed\x6d\xb6\x24\xe4\x26\x6d\x46\x87\x21\x5e\x30\xe9\x26\x71\xcc\x85\x22\x60\x29\x91\xa0\x65\x4e\x5f\x71\xa9\x18\x8d\x50\x1a\x06\x66\xab\x90\x39\x32\x46\x73\xdb\x3e\xbd\x97\x53\x1c\x7f\xe3\x0b\x81\x6b\x1a\x61\x3e\x02\xa0\x1b\x03\x02\xef\xdc\x64\x08\x7f\xfc\xd4\x02\x5d\x45\xbd\x87\xd4\xe8\x75\xf1\xce\x2e\xa7\x91\x4d\x17\x79\x9a\xe5\x19\x29\x49\xb2\x0a\x64\x83\xd8\xeb\xf8\x05\xae\x1c\xec\x22\xc8\xaa\xa4\xcb\xcd\xbf\xd1\x38\xb3\xe8\xc4\x37\xec\x2b\x4d\x98\x37\x26\x90\x52\xcb\xe7\x9b\x8f\x34\x08\xe9\x30\x08\x03\xf5\x72\xcb\x99\xd6\x8c\x21\x7a\x0a\x7e\x40\xfd\x03\xbc\xfb\x92\xee\x2a\xf3\x0c\xab\x22\x87\xea\x89\x8b\x07\x13\x5e\xe6\x77\x53\xc8\xb5\x58\x2f\xaf\x35\x16\xef\xfd\xad\x68\xdb\x3b\xa4\x6d\xef\x92\x76\xe3\x10\x68\xdb\x8b\xb5\x63\x2b\xda\xce\x0e\x69\x3b\xbb\xa4\x6d\x1f\x02\x6d\x47\xbf\xe0\xa4\xc5\x10\xd5\x17\xaa\xf0\x89\xbe\x94\xd3\x36\x8c\x2a\xa0\xee\x2b\xfc\xac\x02\xac\x14\xf8\xa0\xdb\xca\xe7\x9b\x4a\x51\x6f\x1c\x21\x53\x6b\xa5\x84\xe1\x65\x62\xb1\x48\x24\xd3\xd4\xe3\x89\xc2\x7e\x5a\x29\xca\x03\x9a\xce\xaf\x15\xc6\x9e\x33\x63\x5e\xce\x12\x25\x79\xc3\x15\x23\xf3\xe5\x0d\x23\x25\xcf\xa5\x42\xe6\x94\xc3\x44\xad\x09\x30\xb7\xbc\x40\xa9\x02\x46\xd3\x03\x3e\x73\x3e\xe7\xdf\x3a\x01\x56\x7d\x3e\x93\x42\x35\xf5\xd3\x94\x92\x7b\x81\x76\xb0\xec\xae\x29\x5d\xb0\x71\x71\xcd\xe6\x0d\xed\xf3\x8b\x56\x85\x64\xd4\x84\x3d\x09\xab\xaa\x63\x4b\x85\xd9\x5b\x08\x73\xf6\x24\xac\xaa\x64\x2c\x15\xe6\x6c\x20\x2c\x3f\xc0\x4d\x2f\x2c\x17\x31\x9d\x3f\xf0\xab\xa2\xc3\x86\x3c\x61\x7e\x3b\x1e\x63\x84\x82\x86\x5d\x2e\x94\x29\xb1\xcd\x94\xa8\xb8\xa3\x0d\xa3\x0a\xb1\x53\x2b\x83\xac\x81\x09\xa0\x97\x84\x78\x9d\x44\x43\x14\xe9\x4b\x61\xdd\x29\xfa\xf3\xae\xe0\x8a\x7b\x3c\x24\x60\x7d\xb2\x66\x6c\x9b\x5e\x96\x09\xfa\xfb\x58\xd1\xec\xdf\x0b\x94\x69\x83\x3f\xa2\xa1\x9c\x74\xf8\x4b\xee\x9f\x54\x73\x8f\xb2\x7b\x24\x13\x7a\x97\x82\x47\x3a\x02\xfb\xd4\x9a\x0c\xf6\x79\xea\xfe\xec\xcc\x39\xb3\xa6\xe4\x5c\xf7\xea\x70\x78\x9d\xbe\x09\x2f\x1d\x40\xf1\xc5\xf2\x97\xcc\x6c\xdb\x20\x96\x0d\xe4\xb8\xae\x94\x8a\x0f\x87\xd7\xd9\xff\x9c\x5f\x9f\xeb\x06\xab\x6c\xe0\x26\x51\x1a\xd6\xe1\x80\xaa\x6f\x05\x6a\xf6\x4d\x7b\x23\x4e\x26\xa6\xc9\x21\x34\x6a\xaf\x29\x66\xc5\x62\x53\xba\xe0\x6d\xdb\x83\xd5\x9e\x84\x51\x81\xf7\x2a\x6f\xab\x26\x61\x13\x79\xce\x5e\xe5\x6d\xd5\x2a\xac\x22\xef\x26\x51\x71\xa2\xb2\x6f\x5e\xba\xd8\xe7\xfd\xf6\xcc\xa7\xc6\xfe\x18\x21\xf0\x81\x8f\x40\x8d\x11\x1e\xe3\xac\x9e\x17\x95\x7b\xb6\x35\x68\x3f\xeb\xaf\x5a\x85\x7b\x1a\x55\x97\x76\xed\x6c\xf1\x68\x54\x06\x90\xb5\x02\x20\xb5\x61\xe0\xcf\xfd\xd7\x40\x16\x4a\xee\xf5\x92\x11\xd2\x19\x4d\xdb\x93\x8a\x03\x91\x4e\x2d\x49\xfc\x7c\x52\x47\x7d\xcd\xb5\x83\x75\x15\x2e\x28\x5b\x3c\x29\x6b\xaa\xb5\x37\x50\x6b\x2f\x53\x6b\xbf\x95\x5a\xbb\x44\xad\xb3\xa6\x5a\x67\x03\xb5\xce\x32\xb5\xce\x5b\xa9\x75\x3a\xfe\xd1\x7f\x01\x00\x00\xff\xff\x7b\x24\x36\xe9\xff\x1d\x00\x00")

func assetsVpcYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsVpcYml,
		"assets/vpc.yml",
	)
}

func assetsVpcYml() (*asset, error) {
	bytes, err := assetsVpcYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/vpc.yml", size: 7679, mode: os.FileMode(420), modTime: time.Unix(1484356037, 0)}
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
	"assets/cluster.yml": assetsClusterYml,
	"assets/pipeline.yml": assetsPipelineYml,
	"assets/repo.yml": assetsRepoYml,
	"assets/service.yml": assetsServiceYml,
	"assets/vpc.yml": assetsVpcYml,
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
		"cluster.yml": &bintree{assetsClusterYml, map[string]*bintree{}},
		"pipeline.yml": &bintree{assetsPipelineYml, map[string]*bintree{}},
		"repo.yml": &bintree{assetsRepoYml, map[string]*bintree{}},
		"service.yml": &bintree{assetsServiceYml, map[string]*bintree{}},
		"vpc.yml": &bintree{assetsVpcYml, map[string]*bintree{}},
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

