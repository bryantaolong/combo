<a href="https://github.com/bryantaolong/combo">
  <img width="60px" height="60px" src="./assets/logo.png" align="right"  alt=""/>
</a>

# Combo CLI

**Combo** 是一个基于 **Go** 语言开发的命令行 To Do 应用，支持分组管理、任务管理、完成状态标记，并提供彩色输出，方便区分任务状态。

---

## ✨ 功能特性

- 创建/删除/列出分组
- 添加/删除/列出/清空任务
- 标记任务完成
- 默认分组与自定义分组支持
- 彩色输出区分任务状态：
  - 已完成任务：绿色
  - 未完成任务：默认颜色
  - 警告/错误信息：红色
  - 信息提示：黄色或蓝色

---

## 🚀 安装与构建

### 1. 安装 Go

确保本地安装了 Go 1.20+ 并配置好环境变量：

```bash
go version
````

### 2. 克隆项目

```bash
git clone https://github.com/bryantaolong/combo.git
cd combo
```

### 3. 构建可执行文件

#### Windows

```powershell
go build -o combo.exe
```

#### Linux / macOS

```bash
go build -o combo
```

> 可选：将可执行文件加入系统 PATH，方便全局使用。

---

## 🎮 使用说明

### 查看帮助

```bash
combo --help
combo todo --help
combo group --help
```

### 分组管理

```bash
# 创建分组
combo group add work

# 列出所有分组
combo group list

# 删除分组（默认分组不能删除）
combo group delete work
```

### 任务管理

```bash
# 添加任务
combo todo add "写日报" -g work
combo todo add "默认任务1"

# 列出任务
combo todo list -g work
combo todo list

# 标记任务完成
combo todo done 1 -g work

# 删除任务
combo todo delete 1 -g work

# 清空分组任务
combo todo clear -g work
```

---

## 示例

```powershell
# 创建分组
combo group add work

# 添加任务
combo todo add "写日报" -g work
combo todo add "开会准备" -g work

# 查看任务
combo todo list -g work

# 标记完成
combo todo done 1 -g work

# 删除任务
combo todo delete 2 -g work

# 删除分组
combo group delete work
```

---

## 文件结构

```
combo/
├─ cmd/           # cobra 命令模块
│  ├─ root.go
│  ├─ todo/
│  │  ├─ add.go
│  │  ├─ delete.go
│  │  ├─ list.go
│  │  ├─ done.go
│  │  └─ clear.go
│  └─ group/
│     ├─ group.go
│     ├─ add.go
│     ├─ list.go
│     └─ delete.go
├─ storage/       # 数据结构与存储逻辑
├─ main.go
└─ README.md
```

---

## 💾 数据存储

* 默认存储为 JSON 文件，位于用户目录下（可自定义）
* 每个分组对应一个任务列表
* 每个任务包含：

    * `ID`：唯一标识
    * `Content`：任务内容
    * `Done`：完成状态

---

## 📚 彩色输出

* 使用 [`fatih/color`](https://github.com/fatih/color) 实现
* 已完成任务：绿色
* 警告/错误信息：红色
* 信息提示：黄色或蓝色

---

## License

[MIT License](./LICENSE)
