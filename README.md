# tool
自定义可配置cli工具。

## 首次使用

```shell
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

## 扩展功能

> 编辑tool同级目录生成的```tool.json文件```

```json
[
  {
    "key": "运行调试服务",
    "value": "npm start"
  },
  {
    "key": "app编译",
    "value": "cd android && ./gradlew assembleRelease && open app/build/outputs/apk/release"
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
0.运行调试服务
1.app编译

请输入你的选择:
1
开始运行(app编译)服务...
```