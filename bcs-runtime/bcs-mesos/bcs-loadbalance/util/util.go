

// Package util xxx
package util

import (
	"crypto/md5"
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/deckarep/golang-set"
)

// ExeCommand execute shell command, command like:
// haproxy -f config.cfg -c
// haproxy -f haproxy.cfg -p haproxy.pid -sf $(cat haproxy.pid)
func ExeCommand(command string) (string, bool) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), false
	}
	return string(output), true
}

// Md5SumForFile calculate md5sum of file
func Md5SumForFile(filename string) (string, error) {
	filePtr, err := os.Open(filename)
	defer func() {
		err = filePtr.Close()
		if err != nil {
			blog.Warnf("close file %s failed, err %s", filename, err.Error())
		}
	}()
	if err != nil {
		blog.Errorf("Open file %s failed: %s", filename, err.Error())
		return "", fmt.Errorf("Open file %s failed: %s", filename, err.Error())
	}
	md5Block := md5.New()
	_, err = io.Copy(md5Block, filePtr)
	if err != nil {
		blog.Errorf("do io.Copy failed when calculate file %s md5, err %s", filename, err.Error())
		return "", fmt.Errorf("do io.Copy failed when calculate file %s md5, err %s", filename, err.Error())
	}
	md5Str := string(md5Block.Sum([]byte("")))
	return md5Str, nil
}

// ReplaceFile xxx
func ReplaceFile(oldFile, curFile string) error {
	// backup file first
	src, sErr := os.Open(curFile)
	defer func() {
		err := src.Close()
		if err != nil {
			blog.Warnf("close curFile %s failed, err %s", curFile, err.Error())
		}
	}()
	if sErr != nil {
		blog.Errorf("Read new config file [%s] failed: %s", curFile, sErr.Error())
		return sErr
	}
	dst, dErr := os.OpenFile(oldFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer func() {
		err := dst.Close()
		if err != nil {
			blog.Warnf("close oldFile %s failed, err %s", oldFile, err.Error())
		}
	}()
	if dErr != nil {
		blog.Errorf("Read old config file %s failed: %s", oldFile, dErr.Error())
		return dErr
	}
	// mv file
	_, err := io.Copy(dst, src)
	if err != nil {
		blog.Errorf("Copy new nginx.cfg failed: %s", err.Error())
		return err
	}
	blog.Infof("Replace config file %s success", oldFile)
	return nil
}

// GetSubsection return slice come from first - second
func GetSubsection(first, second []string) (sub []string) {
	if len(first) == 0 {
		return []string{}
	}
	if len(second) == 0 {
		return first
	}
	var lsInf, rsInf []interface{}
	for _, s := range first {
		lsInf = append(lsInf, s)
	}
	for _, s := range second {
		rsInf = append(rsInf, s)
	}
	ls := mapset.NewSetFromSlice(lsInf)
	rs := mapset.NewSetFromSlice(rsInf)
	subSet := ls.Difference(rs)
	for _, s := range subSet.ToSlice() {
		sub = append(sub, s.(string))
	}
	return sub
}

// TrimSpecialChar trim special char
func TrimSpecialChar(src string) string {
	// trim special char
	validPath := strings.Replace(src, "/", "", -1)
	validPath = strings.Replace(validPath, " ", "", -1)
	validPath = strings.Replace(validPath, "~", "", -1)
	validPath = strings.Replace(validPath, "*", "", -1)
	validPath = strings.Replace(validPath, ".", "", -1)
	validPath = strings.Replace(validPath, "\\", "", -1)

	return validPath
}

// GetValidZookeeperPath 去除/等特殊字符，/转换为_
func GetValidZookeeperPath(src string) string {
	if src == "/" {
		return ""
	}
	// trim special char
	validPath := strings.Replace(src, "/", "_", -1)
	validPath = strings.Replace(validPath, " ", "", -1)
	validPath = strings.Replace(validPath, "~", "", -1)
	validPath = strings.Replace(validPath, "*", "", -1)
	validPath = strings.Replace(validPath, ".", "", -1)
	validPath = strings.Replace(validPath, "\\", "", -1)

	return validPath
}

// GetValidTargetGroupSub 去除/等特殊字符，/转换为-,targetGroup名字特殊字符只能是-
func GetValidTargetGroupSub(src string) string {
	if src == "/" {
		return ""
	}
	// trim special char
	validPath := strings.Replace(src, "/", "-", -1)
	validPath = strings.Replace(validPath, " ", "", -1)
	validPath = strings.Replace(validPath, "~", "", -1)
	validPath = strings.Replace(validPath, "*", "", -1)
	validPath = strings.Replace(validPath, ".", "", -1)
	validPath = strings.Replace(validPath, "\\", "", -1)

	return validPath
}
