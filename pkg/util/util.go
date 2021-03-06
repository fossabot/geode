package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/geode-lang/geode/pkg/util/log"
	"github.com/shibukawa/configdir"
)

// RunCommand executes a command and returns stdout from it.
func RunCommand(command string, args ...string) ([]byte, error) {
	var out []byte
	var err error

	tmpcmd := command + " " + strings.Join(args, " ")
	maxLen := 500
	if len(tmpcmd) > maxLen {
		tmpcmd = tmpcmd[:maxLen-3] + "..."
	}
	title := fmt.Sprintf("Command Execution (%s)", tmpcmd)
	log.Timed(title, func() {
		cmd := exec.Command(command, args...)
		out, err = cmd.CombinedOutput()
	})

	if err != nil {
		return out, err
	}
	return out, err
}

// RunCommandStr is a wrapper around RunCommand that returns a string instead
func RunCommandStr(command string, args ...string) (string, error) {
	b, e := RunCommand(command, args...)
	return string(b), e
}

// StdLibDir returns the stdlib directory path
func StdLibDir() string {
	libpath := os.Getenv("GEODELIB")
	if libpath == "" {
		libpath = "/usr/local/lib/geodelib"
	}
	return libpath
}

// StdLibFile takes a path in the stdlib and
// joins it to the directory path
func StdLibFile(p string) string {
	return path.Join(StdLibDir(), p)
}

// HomeDir will return the home directory of the current user.
func HomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// GetCacheDir get's the OS Specific cache directory
func GetCacheDir() string {
	configDirs := configdir.New("nick-wanninger", "geode-lang")
	cache := configDirs.QueryCacheFolder()
	return cache.Path
}

var tmpdir string

// GetTmp returns a temporary directory
func GetTmp() string {

	if tmpdir == "" {
		pth, err := ioutil.TempDir(path.Join(HomeDir(), ".geode/tmp"), "")
		if err != nil {
			log.Fatal("Unable to get temp directory\n")
		}
		os.MkdirAll(pth, os.ModePerm)

		tmpdir = pth
	}

	return tmpdir

}

// PurgeCache -
func PurgeCache() {
	cacheDir := GetCacheDir()

	fmt.Println(cacheDir)
	os.MkdirAll(cacheDir, os.ModePerm)

	files, _ := ioutil.ReadDir(cacheDir)
	// if err != nil {
	// 	log.Fatal("Unable to search cache for files\n")
	// }

	now := time.Now()

	cacheInvalidationTimeout := 10 * time.Minute

	for _, f := range files {
		if now.Sub(f.ModTime()) > cacheInvalidationTimeout {
			os.Remove(path.Join(cacheDir, f.Name()))
		}
	}

}

// RandomHex returns a random hex string of n bytes in length
func RandomHex(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// HashFile takes a path and hashes it efficiently into sha1
func HashFile(path string) string {
	var returnMD5String string
	file, err := os.Open(path)
	if err != nil {
		return returnMD5String
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String

}
