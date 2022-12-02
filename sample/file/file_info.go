package file

import (
	"fmt"
	"os"
)

const FILE_PATH = "data/big_data.txt"

type myFile struct {
	file *os.File
}

func newMyFile(file *os.File) *myFile {
	f := myFile{
		file: file,
	}
	return &f
}

func (f *myFile) showStat() {
	stat, err := f.file.Stat()
	if err != nil {
		panic(err)
	}
	/*
		type FileInfo interface {
		Name() string       // base name of the file
		Size() int64        // length in bytes for regular files; system-dependent for others
		Mode() FileMode     // file mode bits
		ModTime() time.Time // modification time
		IsDir() bool        // abbreviation for Mode().IsDir()
		Sys() any           // underlying data source (can return nil)
	*/
	fmt.Println("===========  STAT  =============")
	fmt.Printf("Name: %s\n", stat.Name())
	fmt.Printf("Size: %d\n", stat.Size())
	fmt.Printf("Mode: %v\n", stat.Mode())
	fmt.Printf("IsDir: %t\n", stat.IsDir())
	fmt.Printf("ModTime: %v\n", stat.ModTime())
	fmt.Printf("ModTime.Local: %v\n", stat.ModTime().Local())
	fmt.Printf("ModTime.UTC: %v\n", stat.ModTime().UTC())

	// if runtime.GOOS == "windows" {
	// 	sys := stat.Sys().(*syscall.Win32FileAttributeData)
	// 	lastAccessTime := time.Since(time.Unix(0, sys.LastAccessTime.Nanoseconds()))
	// 	creationTime := time.Since(time.Unix(0, sys.CreationTime.Nanoseconds()))
	// 	lastWriteTime := time.Since(time.Unix(0, sys.LastWriteTime.Nanoseconds()))

	// 	fmt.Printf("LastAccessTime: %v\n", lastAccessTime)
	// 	fmt.Printf("creationTime: %v\n", creationTime)
	// 	fmt.Printf("lastWriteTime: %v\n", lastWriteTime)
	// }
	// else {
	// 	sys := stat.Sys()
	// 	aTime := sys.(*syscall.Stat_t).Atim
	// 	cTime := sys.(*syscall.Stat_t).Ctim
	// 	mTime := sys.(*syscall.Stat_t).Mtim
	// 	lastAccessTime := time.Since(time.Unix(aTime.Sec, aTime.Nsec))
	// 	creationTime := time.Since(time.Unix(cTime.Sec, cTime.Nsec))
	// 	lastWriteTime := time.Since(time.Unix(mTime.Sec, mTime.Nsec))
	// }
	fmt.Println("============================")
}

func (f *myFile) close() {
	f.file.Close()
}

func FileInfoMain() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	myFile := newMyFile(file)
	myFile.showStat()
}

/*
===========  STAT  =============
Name: big_data.txt
Size: 289998059
Mode: -rw-rw-rw-
IsDir: false
ModTime: 2022-07-20 20:28:46.6977735 +0900 KST
ModTime.Local: 2022-07-20 20:28:46.6977735 +0900 KST
ModTime.UTC: 2022-07-20 11:28:46.6977735 +0000 UTC
LastAccessTime: 51h22m46.2213526s
creationTime: 52h14m28.1290086s
lastWriteTime: 51h22m46.2213526s
============================
*/

func IsFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
