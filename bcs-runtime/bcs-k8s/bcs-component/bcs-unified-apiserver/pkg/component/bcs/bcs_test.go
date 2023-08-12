

package bcs

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-unified-apiserver/pkg/config"
	"github.com/stretchr/testify/assert"
)

func initConf() {
	_, filename, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(path.Join(path.Dir(filename), "../../../"))
	if err != nil {
		panic(err)
	}
	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	// 初始化BCS配置
	bcsConfContentYaml, err := ioutil.ReadFile("./etc/config_dev.yaml")
	if err != nil {
		panic(err)
	}

	if err = config.G.ReadFrom(bcsConfContentYaml); err != nil {
		panic(err)
	}
}

func TestListResources(t *testing.T) {
	initConf()
	ctx := context.Background()

	resources, pag, err := ListPodResources(ctx, []string{"BCS-K8S-00000"}, "kube-system", 5, 0)
	assert.NoError(t, err)
	assert.Equal(t, len(resources), 0)
	assert.Equal(t, pag.Total, 0)
}
