package leetcode

import (
	"regexp"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	var (
		res     = make([][]string, 0)
		filemap = make(map[string][]string)
	)

	for _, path := range paths {
		tempmap := getFileMaps(strings.Split(path, " "))

		for k, v := range tempmap {
			if _, ok := filemap[k]; ok {
				filemap[k] = append(filemap[k], v...)
			} else {
				filemap[k] = v
			}
		}
	}

	for _, v := range filemap {
		if len(v) > 1 {
			res = append(res, v)
		}
	}

	return res
}

func getFileMaps(paths []string) map[string][]string {
	var (
		res = make(map[string][]string)
		dir = paths[0]
	)

	r, _ := regexp.Compile(`(.+)\((.+)\)`)

	for _, path := range paths[1:] {
		match := r.FindStringSubmatch(path)

		res[match[2]] = append(res[match[2]], dir+"/"+match[1])
	}

	return res
}
