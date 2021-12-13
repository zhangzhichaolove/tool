package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

/**
工具结构定义
*/
type Tool struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

/**
运行shell命令
*/
func execShell(strCommand string) string {
	cmd := exec.Command("/bin/bash", "-c", strCommand)
	out, err := cmd.Output()
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			fmt.Println(string(e.Stderr))
		} else {
			fmt.Println(err.Error())
		}
	}
	return string(out)
}

/**
检查文件是否存在
*/
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**+
主要逻辑
*/
func main() {
	dir, _ := os.Executable()
	currentPath := filepath.Dir(dir)
	configJsonPath := fmt.Sprintf("%s/tool.json", currentPath)
	//没有配置文件的情况下，生成默认配置
	if !checkFileIsExist(configJsonPath) {
		f, _ := os.Create(configJsonPath)
		jsonIndent, _ := json.MarshalIndent(&[]Tool{{Key: "示例命令", Value: "npm start"}}, "", "\t")
		f.WriteString(string(jsonIndent))
		f.Close()
	}
	filePtr, err := os.Open(configJsonPath)
	if err != nil {
		fmt.Println("读取配置失败-->", err.Error())
	}
	defer filePtr.Close()
	var tools []Tool
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&tools)
	if err != nil {
		fmt.Println("解析配置失败-->", err.Error())
	}
	tips := "请选择你要使用的功能：\n"
	for i, tool := range tools {
		tips += fmt.Sprintf("%d.%s\n", i, tool.Key)
	}
	fmt.Println(tips)
	fmt.Println("请输入你的选择:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	num, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("输入错误，请输入正确的序号!")
	}
	fmt.Println(fmt.Sprintf("开始运行(%s)服务...", tools[num].Key))
	fmt.Println(execShell(tools[num].Value))
}
