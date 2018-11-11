package main

import "os"
import "strings"
import "sort"
import "fmt"
import "log"
import "path/filepath"

type ByRank []string

const (
	SET_PATH_COMMAND = `PATH="%v"; export PATH;`
)

func get_path_rank(path string) int {
	prefix_ranks := map[string]int{
		"/home/":      -1,
		"/Users/":     -2,
		"/usr/local/": 1,
		"/usr/":       2,
		"/bin":        3,
		"/sbin":       4,
	}

	for k, v := range prefix_ranks {
		if strings.HasPrefix(path, k) {
			return v
		}
	}

	return 0
}

func remove_duplicates(paths []string) []string {
	new_list := make([]string, 0, len(paths))
	path_map := make(map[string]bool)

	for _, p := range paths {
		p, err := filepath.Abs(p)
		if err != nil {
			log.Fatalf("Could not get abs path for %v: %v", p, err)
		}
		if _, ok := path_map[p]; !ok {
			path_map[p] = true
			new_list = append(new_list, p)
		}
	}

	return new_list
}

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {
	rank_i := get_path_rank(a[i])
	rank_j := get_path_rank(a[j])
	return rank_i < rank_j
}

func main() {
	path := os.Getenv("PATH")
	search_dirs := strings.Split(path, ":")
	search_dirs = remove_duplicates(search_dirs)

	sort.Sort(ByRank(search_dirs))

	output := fmt.Sprintf(SET_PATH_COMMAND, strings.Join(search_dirs, ":"))
	fmt.Println(output)
}
