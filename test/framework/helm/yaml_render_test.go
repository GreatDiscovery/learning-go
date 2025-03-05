package helm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
	"os"
	"testing"
	"text/template"
)

// Config 是用于存储要填充到模板中的数据结构
type Config struct {
	Name     string
	AppName  string
	LogLevel string
	DBHost   string
	DBPort   string
}

func TestYamlRender(t *testing.T) {
	// 定义模板数据
	config := Config{
		Name:     "my-config",
		AppName:  "MyApp",
		LogLevel: "INFO",
		DBHost:   "localhost",
		DBPort:   "5432",
	}

	// 读取模板文件
	tmpl, err := template.ParseFiles("config-template.yaml")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// 渲染模板到文件或标准输出
	file, err := os.Create("rendered-config.yaml")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// 将模板渲染到文件
	err = tmpl.Execute(file, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// 打印渲染的结果到标准输出
	fmt.Println("Rendered YAML Config:")
	err = tmpl.Execute(os.Stdout, config)
	if err != nil {
		log.Fatalf("Error executing template to stdout: %v", err)
	}

	// 1. 读取 YAML 文件
	yamlFile, err := ioutil.ReadFile("rendered-config.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	// 2. 解析 YAML 文件到 ConfigMap 对象
	var configMap v1.ConfigMap
	if err := yaml.Unmarshal(yamlFile, &configMap); err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// 3. 打印 ConfigMap 对象
	fmt.Printf("ConfigMap Name: %s\n", configMap.Name)
	fmt.Printf("Data: %v\n", configMap.Data)
}

type YamlConfig struct {
	Server struct {
		Listen     string `yaml:"listen"`
		ListenPeer string `yaml:"listenPeer"`
	} `yaml:"server"`
}

func TestYamlRenderYaml(t *testing.T) {
	// 1. 读取 YAML 文件
	yamlFile, err := ioutil.ReadFile("render-config-yaml.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	// 2. 解析 YAML 文件到 ConfigMap 对象
	var configMap v1.ConfigMap
	if err := yaml.Unmarshal(yamlFile, &configMap); err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// 3. 打印 ConfigMap 对象
	fmt.Printf("ConfigMap Name: %s\n", configMap.Name)
	fmt.Printf("Data[cluster.yaml]: %v\n", configMap.Data["cluster.yaml"])

	// 4. 读取 cluster.yaml 文件
	var yamlConfig YamlConfig
	yaml.Unmarshal([]byte(configMap.Data["cluster.yaml"]), &yamlConfig)
	assert.Equal(t, yamlConfig.Server.Listen, "127.0.0.1:18000")
	assert.Equal(t, yamlConfig.Server.ListenPeer, "127.0.0.1:18000")
}
