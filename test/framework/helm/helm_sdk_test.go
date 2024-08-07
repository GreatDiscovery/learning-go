package helm

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"log"
	"os"
	"testing"
)

//使用helm sdk动态生成yaml文件，方便yaml文件查看和管理

func TestHelmRenderYaml(t *testing.T) {
	// 创建一个 helm 配置对象
	settings := cli.New()
	client := action.NewInstall(&action.Configuration{})
	// 加载 chart 文件，这里应该说是加载chart的目录代码，其中values写好了
	chartPath, err := client.ChartPathOptions.LocateChart("./demo", settings)
	if err != nil {
		log.Fatal(err)
	}
	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个 values 对象，用于存储 chart 的值
	vals := values.Options{}
	valueOpts := &vals
	p := getter.All(settings)
	val, err := valueOpts.MergeValues(p)
	if err != nil {
		log.Fatal(err)
	}
	// 选择只渲染某个模板
	val["renderService"] = false
	val["renderDeployment"] = false
	val["renderServiceAccount"] = false
	val["renderHpa"] = false
	val["renderIngress"] = false
	val["autoscaling"] = map[string]interface{}{
		"enabled": true,
	}

	// 创建一个 action 对象，用于执行 helm 操作
	client.DryRun = true              // 设置为 dry-run 模式，不实际安装 chart
	client.ReleaseName = "my-release" // 设置 release 名称
	client.Replace = true             // 设置为替换模式，如果 release 已存在则覆盖
	client.ClientOnly = true          // 设置为客户端模式，不与集群交互

	// 执行 helm template 操作，并获取结果
	rel, err := client.Run(chart, val)
	if err != nil {
		log.Fatal(err)
	}

	filePath := "./1.yaml"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	//及时关闭file句柄
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 将结果输出为 yaml 格式
	_, err = fmt.Fprintln(file, rel.Manifest)
	if err != nil {
		return
	}
}
