package server_info

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Local struct {
}

func (l *Local) GetCpuName() ([]string, error) {
	stats, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	res := make([]string, len(stats))
	for _, s := range stats {
		res = append(res, s.ModelName)
	}
	return res, nil
}

func (l *Local) GetDiskUsage(path string) (float64, error) {
	stat, err := disk.Usage(path)
	if err != nil {
		return 0, err
	}
	return stat.UsedPercent, nil
}

func (l *Local) GetPathSize(path string) (float64, error) {
	var size float64
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	if !info.IsDir() {
		return float64(info.Size()) / KB, nil
	}
	files, err := walkDir(path)
	if err != nil {
		return 0, err
	}
	for _, file := range files {
		size += file.size
	}
	return size, nil
}

const ( //文件大小单位
	_  = iota
	KB = 1 << (10 * iota)
	MB
)

type fileInfo struct { //文件信息
	name string
	size float64
}

//dirname：目录名
func walkDir(dirname string) ([]fileInfo, error) {
	op, err := filepath.Abs(dirname) //获取目录的绝对路径
	if nil != err {
		fmt.Println(err)
		return nil, err
	}
	files, err := ioutil.ReadDir(op) //获取目录下所有文件的信息，包括文件和文件夹
	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	var fileInfos []fileInfo //返回值，存储读取的文件信息
	for _, f := range files {
		if f.IsDir() { // 如果是目录，那么就递归调用
			fs, err := walkDir(op + `/` + f.Name()) //路径分隔符，linux 和 windows 不同
			if nil != err {
				return nil, err
			}
			fileInfos = append(fileInfos, fs...) //将 slice 添加到 slice
		} else {
			fi := fileInfo{op + `/` + f.Name(), float64(f.Size()) / KB}
			fileInfos = append(fileInfos, fi) //slice 中添加成员
		}
	}

	return fileInfos, nil
}
func (l *Local) ReadDir(dir string) ([]string, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return []string{info.Name()}, nil
	}
	files := []string{}
	fs, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range fs {
		files = append(files, f.Name())
	}
	return files, nil
}
