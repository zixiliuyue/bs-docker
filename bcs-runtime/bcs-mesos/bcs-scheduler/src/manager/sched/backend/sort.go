

package backend

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
	"strconv"
	"strings"
)

// TaskSorter bia name of []TaskGroup
type TaskSorter []*types.TaskGroup

// Len 用于排序
func (s TaskSorter) Len() int { return len(s) }

// Swap 用于排序
func (s TaskSorter) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less 用于排序
func (s TaskSorter) Less(i, j int) bool {

	// the time for building taskgroup
	a, _ := strconv.Atoi(strings.Split(s[i].ID, ".")[4])
	b, _ := strconv.Atoi(strings.Split(s[j].ID, ".")[4])

	return a < b
}
