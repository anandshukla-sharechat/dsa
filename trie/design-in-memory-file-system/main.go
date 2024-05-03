package main

import (
	"fmt"
	"sort"
	"strings"
)

type FileSystem struct {
	children map[string]*FileSystem
	isFile   bool
	content  string
}

func Constructor() FileSystem {
	return FileSystem{children: make(map[string]*FileSystem)}
}

func (this *FileSystem) Ls(path string) []string {
	path = path[1:]
	pathArr := strings.Split(path, "/")
	if path == "" {
		pathArr = []string{}
	}
	ans := make([]string, 0)
	for _, el := range pathArr {
		if _, ok := this.children[el]; !ok {
			this.children[el] = &FileSystem{children: make(map[string]*FileSystem)}
		}
		this = this.children[el]
	}
	for k, _ := range this.children {
		ans = append(ans, k)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func (this *FileSystem) Mkdir(path string) {
	path = path[1:]
	pathArr := strings.Split(path, "/")
	if path == "" {
		pathArr = []string{}
	}
	for _, el := range pathArr {
		if _, ok := this.children[el]; !ok {
			this.children[el] = &FileSystem{children: make(map[string]*FileSystem)}
		}
		this = this.children[el]
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	filePath = filePath[1:]
	pathArr := strings.Split(filePath, "/")
	if filePath == "" {
		pathArr = []string{}
	}
	for _, el := range pathArr {
		if _, ok := this.children[el]; !ok {
			this.children[el] = &FileSystem{children: make(map[string]*FileSystem)}
		}
		this = this.children[el]
	}
	this.isFile = true
	this.content = this.content + content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	filePath = filePath[1:]
	pathArr := strings.Split(filePath, "/")
	if filePath == "" {
		pathArr = []string{}
	}
	for _, el := range pathArr {
		if _, ok := this.children[el]; !ok {
			this.children[el] = &FileSystem{children: make(map[string]*FileSystem)}
		}
		this = this.children[el]
	}
	return this.content
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */

// 42 / 63 testcases passed
func main() {
	fileSystem := Constructor()
	fmt.Println(fileSystem.Ls("/")) // return []
	fileSystem.Mkdir("/a/b/c")
	fileSystem.AddContentToFile("/a/b/c/d", "hello")
	fmt.Println(fileSystem.Ls("/"))                         // return ["a"]
	fmt.Println(fileSystem.ReadContentFromFile("/a/b/c/d")) // return "hello"
}
