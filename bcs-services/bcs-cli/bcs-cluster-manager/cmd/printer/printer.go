

package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/andreazorzetto/yh/highlight"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/tidwall/pretty"
	"github.com/ugorji/go/codec"
	"sigs.k8s.io/yaml"
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

	enc := codec.NewEncoderBytes(&data, &codec.JsonHandle{
		MapKeyAsString: true,
		Indent:         4,
	})

	err := enc.Encode(v)
	if err != nil {
		return errors.Wrapf(err, "encode json failed")
	}

	fmt.Println(string(pretty.Color(pretty.Pretty(data), nil)))

	return nil
}

func encodeYaml(v interface{}) error {
	marshal, err := json.Marshal(v)
	if err != nil {
		return err
	}

	original, err := yaml.JSONToYAML(marshal)
	if err != nil {
		return err
	}

	r := strings.NewReader(string(original))

	h, err := highlight.Highlight(r)
	if err != nil {
		return err
	}

	fmt.Print(h)

	return nil
}

func encodeWide(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.SetRowSeparator("")
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetTablePadding("    ")
	table.SetBorder(true)
	table.SetHeaderLine(false)
	table.SetNoWhiteSpace(true)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

// PrintList list命令格式化输出函数
func PrintList(flagOutput string, resp interface{}, header []string, data [][]string) error {
	if flagOutput == outputTypeJSON {
		return encodeJSON(resp)
	} else if flagOutput == outputTypeYaml {
		return encodeYaml(resp)
	}

	encodeWide(header, data)

	return nil
}
