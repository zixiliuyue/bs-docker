

package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

var pidfile string

// SavePid common func for save proc pid
func SavePid(processConfig conf.ProcessConfig) error {
	pidPath := filepath.Join(processConfig.PidDir, filepath.Base(os.Args[0])+".pid")
	if fi, err := os.Stat(pidPath); err == nil && !fi.IsDir() {
		os.Remove(pidPath)
	} else if !os.IsNotExist(err) {
		return err
	}
	SetPidfilePath(pidPath)
	if err := WritePid(); err != nil {
		return fmt.Errorf("write pid file failed. err:%s", err.Error())
	}

	return nil
}

// SetPidfilePath sets the pidfile path.
func SetPidfilePath(p string) {
	pidfile = p
}

// WritePid the pidfile based on the flag. It is an error if the pidfile hasn't
// been configured.
func WritePid() error {
	if pidfile == "" {
		return fmt.Errorf("pidfile is not set")
	}

	if err := os.MkdirAll(filepath.Dir(pidfile), os.FileMode(0755)); err != nil {
		return err
	}

	file, err := AtomicFileNew(pidfile, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("error opening pidfile %s: %s", pidfile, err)
	}
	defer file.Close() // in case we fail before the explicit close

	_, err = fmt.Fprintf(file, "%d", os.Getpid())
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

// ReadPid the pid from the configured file. It is an error if the pidfile hasn't
// been configured.
func ReadPid() (int, error) {
	if pidfile == "" {
		return 0, fmt.Errorf("pidfile is empty")
	}

	d, err := ioutil.ReadFile(pidfile)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(bytes.TrimSpace(d)))
	if err != nil {
		return 0, fmt.Errorf("error parsing pid from %s: %s", pidfile, err)
	}

	return pid, nil
}
