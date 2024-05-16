# go-learning
Go每日学习（版本 v1.21 见 go.mod）
- 基本知识
- 底层知识
- 网络编程
- 代码审计

项目统计： 统计一个项目中源码的行数、百分比等信息。可使用 [cloc](https://github.com/AlDanial/cloc/releases) ;下载对应可执行文件，重命名可执行文件名字，复制至目录 **C:\Windows\System32**。
> 按文件类型统计文件数量、空行数、注释行数、代码行数：`cloc .` 

> 文件维度下注释行占比、空行占比、代码行数：`cloc --by-file  --by-percent c  .`

- 辅助工具

uml结构生成：可以使用go-plantuml。

> 安装最新版本：go-plantuml generate -d . -o graph.puml.

> 生成UML图：go-plantuml generate -d . -o graph.puml.

- 标准结构

``` text
api：对外暴露的Api
cmd：项目的启动函数（项目的工具链）
conf：服务的配置文件（运行配置文件、docker文件、yaml、json、脚本文件）如：项目镜像打包、健康检查脚本、关闭脚本。
constant：按照业务层面、非业务层面区分的常量
docs：编码规范、todo、项目说明文档、业务说明文档、uml、项目做的好的总结。
internal：私有的业务相关代码如models（私有代码无法被外部导入）
pkg：可以公用的代码，来源于私有业务代码的抽取（项目用到的工具链的源代码）
version：定义项目的版本信息
vender: 项目自身对外部项目的依赖
build: 存在makefile（构建第三方工具所用到的脚本）
```