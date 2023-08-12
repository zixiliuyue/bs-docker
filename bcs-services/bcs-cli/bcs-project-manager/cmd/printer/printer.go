

package printer

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/tidwall/pretty"
	"github.com/ugorji/go/codec"
)

// defaultTableWriter create the tablewriter instance
func defaultTableWriter() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderLine(false)
	table.SetRowLine(false)
	table.SetColWidth(150)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	return table
}

const (
	outputTypeJSON = "json"
	outputTypeYaml = "yaml"
	outputTypeWide = "wide"
)

func encodeJSON(v interface{}) error {
	var data []byte
	if err := encodeJSONWithIndent(4, v, &data); err != nil {
		return errors.Wrapf(err, "encode json failed")
	}
	fmt.Println(string(pretty.Color(pretty.Pretty(data), nil)))
	return nil
}

func encodeJSONWithIndent(indent int8, v interface{}, s *[]byte) error {
	enc := codec.NewEncoderBytes(s, &codec.JsonHandle{
		MapKeyAsString: true,
		Indent:         indent,
	})
	return enc.Encode(v)
}
