package fs

import (
	"errors"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

func CreateFileReadStream(path string) (*os.File, error) {
	return os.Open(path)
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
}

func IsAccessible(path string) bool {
	e, err := FileExists(path)
	if err != nil {
		return false
	}
	return e
}

func IsFile(path string) (bool, error) {
	e, err := IsDirectory(path)
	if err != nil {
		return false, err
	}
	return e, nil
}

func IsDirectory(path string) (bool, error) {
	e, err := FileExists(path)
	if !e {
		return false, nil
	} else if err != nil {
		return false, err
	}

	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return s.IsDir(), nil
}

func ReadToString(path string) (string, error) {
	s, err := ioutil.ReadFile(path)
	if err != nil {
		return "", nil
	}
	return string(s), nil
}

func ReadToBytes(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

var (
	ERR_INVALID_DIR = errors.New("invaild directory")
)

func ReadDirRecursive(p string) (map[string]fs.DirEntry, error) {
	if e, err := IsDirectory(p); err != nil || !e {
		return nil, ERR_INVALID_DIR
	}

	entries, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}

	t := make(map[string]fs.DirEntry)
	for _, entry := range entries {
		if entry.IsDir() {
			sub, err := ReadDirRecursive(path.Join(p, entry.Name()))
			if err != nil {
				return nil, err
			}
			for k, v := range sub {
				t[k] = v
			}
		} else {
			t[path.Join(p, entry.Name())] = entry
		}
	}

	return t, nil
}

func CopyFile(src, dst string, overwrite bool) error {
	dst_e, err := FileExists(dst)
	if err != nil {
		return err
	}
	if dst_e && !overwrite {
		return nil
	}

	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}
	return nil
}

func CopyDirRecursive(src, dst string, overwrite bool) error {
	if e, err := IsDirectory(src); err != nil || !e {
		return ERR_INVALID_DIR
	}

	if err := os.MkdirAll(dst, os.FileMode(0755)); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if err = CopyDirRecursive(path.Join(src, entry.Name()), path.Join(dst, entry.Name()), overwrite); err != nil {
				return err
			}
		} else {
			dst_e, err := FileExists(path.Join(dst, entry.Name()))
			if err != nil {
				return err
			}
			if dst_e && !overwrite {
				continue
			}

			if err = CopyFile(path.Join(src, entry.Name()), path.Join(dst, entry.Name()), overwrite); err != nil {
				return err
			}
		}
	}

	return nil
}
