package main

import (
	"fmt"
	"reflect"
)

type (
	iDatabase interface {
		GetData(string) string
		PutData(string, string)
	}

	Database struct {
		database map[string]string
	}

	mongoDb struct {
		Database
	}

	sql struct {
		Database
	}

	iFile interface {
		CreateFile(string, string)
		FindFile(string) FileInfo
	}

	FileInfo struct {
		name    string
		content string
	}

	File struct {
		files map[string]FileInfo
	}

	ntfs struct {
		File
	}

	ext4 struct {
		File
	}

	Factory func(string) interface{}
)

func (d Database) GetData(query string) string {
	if _, ok := d.database[query]; !ok {
		return ""
	}
	return d.database[query]
}

func (d Database) PutData(query, data string) {
	d.database[query] = data
}

func (file File) CreateFile(name, path string) {
	fileInfo := FileInfo{content: name, name: path}
	file.files[path] = fileInfo
	fmt.Println(reflect.ValueOf(file))
}

func (file File) FindFile(path string) FileInfo {
	if _, ok := file.files[path]; !ok {
		return FileInfo{}
	}

	return file.files[path]
}

func FileSystemFactory(env string) interface{} {
	switch env {

	case "production":
		return ntfs{File{
			files: make(map[string]FileInfo)}}

	case "development":
		return ext4{File{
			files: make(map[string]FileInfo),
		}}
	default:
		return nil
	}
}

func DatabaseFactory(env string) interface{} {
	switch env {
	case "production":
		return mongoDb{Database{
			database: make(map[string]string),
		}}

	case "development":
		return sql{Database{
			database: make(map[string]string),
		}}

	default:
		return nil
	}
}

func AbstractFactory(factory string) Factory {
	switch factory {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FileSystemFactory
	default:
		return nil
	}
}

func Setup(env string) (iDatabase, iFile) {
	databaseFactory := AbstractFactory("database")
	fileSystemFactory := AbstractFactory("filesystem")

	return databaseFactory(env).(iDatabase), fileSystemFactory(env).(iFile)
}

func main() {
	env1 := "production"
	env2 := "development"

	db1, fs1 := Setup(env1)
	db2, fs2 := Setup(env2)

	db1.PutData("key", "mongo")
	fmt.Println(db1.GetData("key"))

	db2.PutData("key", "sqlite")
	fmt.Println(db2.GetData("key"))

	fs1.CreateFile("ntfs", "../example/testntfs.txt")
	fmt.Println(fs1.FindFile("../example/testntfs.txt"))

	fs2.CreateFile("ext", "../example/testext4.txt")
	fmt.Println(fs2.FindFile("../example/testext4.txt"))

}
