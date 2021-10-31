package fs

import (
	"go/build"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func Exist(p string) bool {
	_, err := os.Stat(p)
	return err == nil || !os.IsNotExist(err)
}

func NotExist(p string) bool {
	_, err := os.Stat(p)
	return err != nil && os.IsNotExist(err)
}

func IsFile(p string) bool {
	fi, _ := os.Stat(p)
	return fi != nil && fi.Mode().IsRegular()
}

func IsDir(p string) bool {
	fi, _ := os.Stat(p)
	return fi != nil && fi.IsDir()
}

func IsEmptyDir(p string) (bool, error) {
	f, err := os.Open(p)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	} else {
		return false, err
	}
}

func DetectDir(dir string, features ...string) string {
	for ; dir != `/`; dir = filepath.Dir(dir) {
		if hasAllFeatures(dir, features) {
			return dir
		}
	}
	return ``
}

func hasAllFeatures(dir string, features []string) bool {
	for _, feature := range features {
		if matches, err := filepath.Glob(filepath.Join(dir, feature)); err != nil {
			log.Panic(err)
		} else if len(matches) == 0 {
			return false
		}
	}
	return true
}

func GoPath() string {
	gopath := os.Getenv(`GOPATH`)
	if gopath == `` {
		gopath = build.Default.GOPATH
	}
	return gopath
}

func GoModPath() string {
	return filepath.Join(GoPath(), `pkg`, `mod`)
}

func GoSrcPath() string {
	return filepath.Join(GoPath(), `src`)
}

func SourceDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

func SourceFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.FromSlash(filename)
}
