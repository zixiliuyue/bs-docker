

package processor

import (
	"strconv"
	"strings"
)

func generateCloudListenerName(clbname, protocol string, port int) string {
	clbname = clbname + protocol + strconv.Itoa(port)
	clbname = strings.Replace(clbname, "_", "-", -1)
	clbname = strings.Replace(clbname, "@", "-", -1)
	clbname = strings.Replace(clbname, "*", "-", -1)
	clbname = strings.Replace(clbname, "$", "-", -1)
	return strings.ToLower(clbname)
}
