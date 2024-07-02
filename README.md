# tool
自定义可配置cli工具。

## 运行方式：

> 1.将工具路径配置在环境变量全局运行。(推荐)

> 2.将工具拷贝到需要执行的目录内运行。

## 首次使用

```shell
# 环境变量运行方式
tool
# 当前目录运行方式
./tool
```

> 输出:

```shell
请选择你要使用的功能：
0.运行调试服务

请输入你的选择:
0
开始运行(运行调试服务)服务...
```

## 配置命令（首次运行tool工具后，会在tool同路径生成tool.json文件）

> 编辑tool同级目录生成的```tool.json文件```

```json
[
  {
    "key": "设置npm镜像",
    "value": "npm config set registry https://registry.npmmirror.com"
  },
  {
    "key": "取消npm镜像",
    "value": "npm config set registry https://registry.npmjs.org"
  }
]
```

## 再次运行

```shell
./tool
```

> 输出:

```shell
请选择你要使用的功能：
0.设置npm镜像
1.取消npm镜像

请输入你的选择:
1
开始运行(取消npm镜像)命令...
```

## 个性化配置（在tool同路径生成tool.json文件会用作通用配置，意味着在任何位置运行tool工具，都将拥有这些配置，在特定位置创建tool.json可用于特殊场景）

> 例：

1.在当前项目`/work/app`内创建一个`tool.json`文件

```json
[
  {
    "key": "React Native 模拟器调试",
    "value": "adb reverse tcp:8081 tcp:8081"
  }
]
```

2.在`/work/app`位置执行`tool`命令

```shell
请选择你要使用的功能：
0.设置npm镜像
1.取消npm镜像
2.React Native 模拟器调试
```

3.在其他位置执行`tool`命令

```shell
请选择你要使用的功能：
0.设置npm镜像
1.取消npm镜像
```

