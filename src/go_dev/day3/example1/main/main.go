package main

import (
	"fmt"
	"strings"
)

func main() {
	var url string
	var path string
	fmt.Scanf("%s%s", &url, &path)

	urlProcess(url)
	pathProcess(path)
}

func urlProcess(url string) string {

	result := strings.HasPrefix("http://", url)
	if !result {
		fmt.Printf("http://%s\n", url)
	}

	return url
}

func pathProcess(path string) string {

	result := strings.HasSuffix("/", path)
	if !result {
		fmt.Printf("%s/", path)
	}

	return path
}
