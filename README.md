# Edgeless 自动插件机器人 2
## 简介
该项目是为了使用 `Golang` 重新实现 [Edgeless 自动插件机器人](https://github.com/EdgelessPE/edgeless-bot)
## 特性 (WIP)
* 完全兼容 [Edgeless 自动插件机器人](https://github.com/EdgelessPE/edgeless-bot)，包括 Tasks，以实现无缝迁移
* 更快的构建速度
* 更好的代码结构
* 更高的拓展性

## 工作进度
截止至 `2021/11/28 0:30`
### 结构
- [x] `brand`
  品牌形象模块，放置版本、标题、作者等
- [ ] `config`
  配置解析模块，用于解析构建配置、任务配置
  - [x] `app`
    构建配置（仅数据结构）
  - [x] `task`
    任务配置（仅数据结构）
- [ ] `utils`
  实用工具，未定
- [ ] `provider`
  内容提供。例如爬虫
  - [x] `paspider`
    PortableApps 爬虫，可从 PA 爬取软件包并处理 (Powered By Chromium)
    - [ ] `autobuild`
      自动从 PA 软件包生成
  - [ ] `external`
    外置构建模块
- [ ] `publisher`
  用于打包并发布模块，可上传至云端。
  