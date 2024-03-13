中文 | [English](./README.md) 

# 阅读之前

截止目前（2024.03），若想使用该工具激活录画功能，仍需要大量前置条件，请务必满足以下所有条件后再进行尝试：

1. Nvidia GPU，需要支持NVENC
2. LDJ-010的 bm2dx.dll
3. 合适的ASIO硬件/软件配置（我个人使用 XONAR AE，但有看到过声称可以录制到声音的 FlexASIO 配置）
4. 一块触摸屏，需要支持1280*720@60Hz，并符合框体的工作模式。确定触摸屏是否能用的步骤详见https://github.com/spice2x/spice2x.github.io/wiki/Configuring-touch-screens-as-subscreen。我买的是这个（不是广告）：【闲鱼】https://m.tb.cn/h.5EIeHuV?tk=HW5sWNI4zR2 HU7632 「快来捡漏【触摸 理想 L9 便携显示器便携屏 13.3 寸switch】」
5. 实现了必要响应字段的 xrpc 服务器（目前处于公开状态的只有氧无插件，本文会在后面内容介绍如何对插件进行修改）

# 010-record-api

lightning model video upload handler

## 安装

使用包中的010-record-api.exe文件

前往[release]()下载（目前不可用）

如果你有go，也可以使用命令安装（目前不可用）

```bash
go install github.com/bookqaq/010-record-api@latest
```

或者你也可以自己编译，这里不做介绍

## 配置本工具

1. 运行一次010-record-api.exe，运行完成后将在同级目录生成配置文件config.toml

   ![image-20240312162953082](https://github.com/bookqaq/010-record-api/blob/images/image-20240312162953082.png?raw=true)

2. 使用记事本编辑config.toml，设置listen_address与upload_service_address，如果仅本地运行，则保留默认配置即可。

   ![image-20240313170201467](https://github.com/bookqaq/010-record-api/blob/images/image-20240313170201467.png?raw=true)

3. 再次运行010-record-api.exe，此时工具应启动

   ![image-20240312163453651](https://github.com/bookqaq/010-record-api/blob/images/image-20240312163453651.png?raw=true)

## 配置本工具以外的必要项

以下步骤没有先后顺序

### 安装显卡驱动

影响未知，姑且写在这作为提醒

### 配置副屏触摸

请参考 https://github.com/spice2x/spice2x.github.io/wiki/Configuring-touch-screens-as-subscreen#step-by-step-instructions

### 配置spice2x（如果使用bemanitools，应忽视此步骤）

1. 下载 spice2x-24-02-13 及以后的版本，解压到游戏目录内

2. 打开spicecfg.exe，选择 options 标签栏，在Graphics (common) 中勾选 Disable D3D9 Device Hook

   ![image-20240312164557472](https://github.com/bookqaq/010-record-api/blob/images/image-20240312164557472.png?raw=true)

### 配置xrpc服务器

#### asphyxia

本文假设游戏，asphyxia与本程序均跑在<u>**同一台电脑**</u>上，且本程序生成的配置文件<u>**未被修改**</u>

##### 确认你的插件是否已经支持相关设置

启动氧无，在网页中选择左侧标签栏Plugins下的IIDX，如果右侧 Plugin Settings中有"Movie Upload URL"这一栏，则把值设置为图中所示即可，否则请看下方内容

![image-20240312165724390](https://github.com/bookqaq/010-record-api/blob/images/image-20240312165724390.png?raw=true)

##### 如何修改插件以支持

~~打开你的iidx@asphyxia插件所在目录，进行以下修改~~

摆了，直接使用我包中带的iidx@asphyxia，记得备份你使用的原插件

#### 其他服务器

本人只有氧无，其他服务器爱莫能助

### 修改bm2dx.dll

**仅 2023-09-05 LDJ-010，其他的我没有**

前往 http://localhost:4399/patcher/ ，像平常使用dll patcher一样，应用本patcher中唯一的选项，并替换你的 bm2dx.dll

如果你使用的服务器声明你的框体区域为日区，则不需要本步骤

## 启动后

在模式选择界面确认副屏是否出现了"動画"按钮，如果出现，那么你的环境配置成功了

# Contribution

Just submit your PRs, I'll check and reply.

# TODOs

In README.md