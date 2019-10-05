package main

import (
	"path/filepath"
	"sort"
	"strings"
)

type ByRank []string

func (paths ByRank) Len() int      { return len(paths) }
func (paths ByRank) Swap(a, b int) { paths[a], paths[b] = paths[b], paths[a] }
func (paths ByRank) Less(a, b int) bool {
	rankA := getRankForPath(paths[a])
	rankB := getRankForPath(paths[b])
	return rankA < rankB
}

func GetOrderedPaths(pathList []string) ([]string, error) {
	uniquePathList, err := deduplicatePaths(pathList)
	if err != nil {
		return nil, err
	}

	sort.Strings(uniquePathList)
	sort.Sort(ByRank(uniquePathList))
	return uniquePathList, nil
}

func getRankForPath(path string) int {
	prefixRankings := map[string]int{
		"/home/":      -1,
		"/Users/":     -2,
		"/usr/local/": 1,
		"/usr/":       2,
		"/bin":        3,
		"/sbin":       4,
	}
	for k, v := range prefixRankings {
		if strings.HasPrefix(path, k) {
			return v
		}
	}

	return 0
}

func deduplicatePaths(paths []string) ([]string, error) {
	newList := make([]string, 0, len(paths))
	pathMap := make(map[string]bool)

	for _, p := range paths {
		p, err := filepath.Abs(p)
		if err != nil {
			return nil, err
		}
		if _, ok := pathMap[p]; !ok {
			pathMap[p] = true
			newList = append(newList, p)
		}
	}

	return newList, nil
}
