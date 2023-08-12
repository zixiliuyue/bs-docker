

package versions

// all apiVersion supported list
var apiVersionMap = map[string][]string{
	"1.5":   apiSetV15,
	"1.6":   apiSetV16,
	"1.7":   apiSetV17,
	"1.8":   apiSetV18,
	"1.11":  apiSetV111,
	"1.12":  apiSetV112,
	"1.12+": apiSetV112,
	"1.13":  apiSetV112,
	"1.13+": apiSetV112,
	"1.14":  apiSetV112,
	"1.14+": apiSetV112,
	"1.15":  apiSetV112,
	"1.16":  apiSetV112,
	"1.17":  apiSetV112,
	"1.18":  apiSetV112,
	"1.19":  apiSetV112,
	"1.20":  apiSetV112,
}
