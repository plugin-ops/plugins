package main

import (
	"errors"
	"fmt"
	"github.com/plugin-ops/pedestal/pedestal/plugin"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/plugin-ops/pedestal/pedestal/action"
	"github.com/shirou/gopsutil/disk"
)

type pathInfo struct {
}

func (p *pathInfo) Name() string {
	return "path-info"
}

func (p *pathInfo) Do(params ...*action.Value) (map[string]*action.Value, error) {
	if len(params) == 0 {
		return nil, errors.New("I don't know which path to view the information on")
	}

	resp := map[string]*action.Value{}
	filePath := params[0].String()
	stat, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	size, err := p.GetPathSize(filePath)
	if err != nil {
		return nil, err
	}
	usage, err := p.GetDiskUsage(filePath)
	if err != nil {
		return nil, err
	}
	childs := []string{}
	if stat.IsDir() {
		childs, err = p.ReadDir(filePath)
		if err != nil {
			return nil, err
		}
	}

	resp["0"] = action.NewValue(filePath)
	resp["1"] = action.NewValue(path.Base(filePath))
	resp["2"] = action.NewValue(path.Dir(filePath))
	resp["3"] = action.NewValue(path.Ext(filePath))
	resp["4"] = action.NewValue(stat.IsDir())
	resp["5"] = action.NewValue(size)
	resp["6"] = action.NewValue(usage)
	resp["7"] = action.NewValue(strings.Join(childs, "\n"))
	return resp, nil
}

func cmd(str string) (string, error) {
	cmd := exec.Command("bash", "--noprofile", "--norc", "-c", str)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()

	if err = cmd.Start(); err != nil {
		return "", err
	}
	opBytes, err := ioutil.ReadAll(stdout)
	return string(opBytes), err
}

func (p *pathInfo) Version() float32 {
	return 0.1
}

func (p *pathInfo) Description() string {
	return `
动作功能: 此插件用于获取指定路径的详细信息
传入参数: 动作接受一个参数, 表示需要获取信息的路径
返回参数:
	0. 目标文件完整路径
	1. 目标文件(夹)名
	2. 目标文件父目录完整路劲
	3. 目标文件后缀名
	4. 目标路劲是否是文件夹
	5. 目标路径占用空间(文件夹将会递归获取大小)
	6. 目标路径所在磁盘已使用百分比
	7. 目标下一级文件名列表(如果路径不是文件夹则此项为空字符串), 文件名之间使用'\n'连接
`
}

func (p *pathInfo) GetDiskUsage(path string) (float64, error) {
	stat, err := disk.Usage(path)
	if err != nil {
		return 0, err
	}
	return stat.UsedPercent, nil
}

func (p *pathInfo) GetPathSize(path string) (float64, error) {
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
func (p *pathInfo) ReadDir(dir string) ([]string, error) {
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

func main() {
	plugin.ServePlugin(&pathInfo{})
}
