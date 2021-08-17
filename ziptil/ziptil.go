package ziptil

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/TruthHun/gotil/filetil"
	"github.com/TruthHun/gotil/strtil"
)

//解压zip文件
//@param			zipFile			需要解压的zip文件
//@param			dest			需要解压到的目录
//@return			err				返回错误
func Unzip(zipFile, dest string) (err error) {
	dest = strings.TrimSuffix(dest, "/") + "/"
	// 打开一个zip格式文件
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()
	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, f := range r.File {
		if !f.FileInfo().IsDir() { //非目录，且不包含__MACOSX
			//不要用defer来关闭，如果文件太多的话，会报too many open files 的错误
			fname := f.Name
			if f.Flags == 0 {
				if fnameConv, errConv := strtil.EncodeUTF8(fname, "gbk"); errConv == nil {
					fname = fnameConv
				}
			}

			savePath := filepath.Join(dest, strings.TrimLeft(strings.TrimSpace(fname), "./"))
			folder := filepath.Dir(savePath)
			if strings.Contains(folder, "__MACOSX") {
				continue
			}
			os.MkdirAll(folder, 0777)
			fcreate, errCreate := os.Create(savePath)
			if errCreate != nil {
				return
			}

			rc, errOpen := f.Open()
			if errOpen != nil {
				fcreate.Close()
				return errOpen
			}

			io.Copy(fcreate, rc)
			rc.Close()
			fcreate.Close()
		}
	}
	return nil
}

//压缩指定文件或文件夹
//@param			dest			压缩后的zip文件目标，如/usr/local/hello.zip
//@param			filepath		需要压缩的文件或者文件夹
//@return			err				错误。如果返回错误，则会删除dest文件
func Zip(dest string, filepath ...string) (err error) {
	if len(filepath) == 0 {
		return errors.New("lack of file")
	}
	//创建文件
	fzip, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fzip.Close()

	var filelist []filetil.FileList
	for _, file := range filepath {
		if info, err := os.Stat(file); err == nil {
			if info.IsDir() { //目录，则扫描文件
				if f, _ := filetil.ScanFiles(file); len(f) > 0 {
					filelist = append(filelist, f...)
				}
			} else { //文件
				filelist = append(filelist, filetil.FileList{
					IsDir: false,
					Name:  info.Name(),
					Path:  file,
				})
			}
		} else {
			return err
		}
	}
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range filelist {
		if !file.IsDir {
			if fw, err := w.Create(strings.TrimLeft(file.Path, "./")); err != nil {
				return err
			} else {
				if filecontent, err := ioutil.ReadFile(file.Path); err != nil {
					return err
				} else {
					if _, err = fw.Write(filecontent); err != nil {
						return err
					}
				}
			}
		}
	}
	return
}
