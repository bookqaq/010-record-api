[中文](./README-zh.md) | English

# Before Reading

Until now（2024.03）, if you want to use the recording feature in game, a bunch of requirements are needed:

1. Nvidia GPU, support NVENC
2. LDJ-010 bm2dx.dll
3. Proper ASIO Hardware/Software configuration（I use XONAR AE, but have seen FlexASIO config that claims to work fine）
4. A screen with touch support, support 1280*720@60Hz, function the same as the one on cab. Whether the screen match requirements:  https://github.com/spice2x/spice2x.github.io/wiki/Configuring-touch-screens-as-subscreen
5. A proper implemented xrpc server（asphyxia for example）

# 010-record-api

lightning model video upload handler

## Installation

use 010-record-api.exe

download executable at [release](https://github.com/bookqaq/010-record-api/releases/)（unusable）

if you have go installed, you could install via command line（unusable）

```bash
go install github.com/bookqaq/010-record-api@latest
```

or compile yourself

## Config The Tool

1. open 010-record-api.exe once, a config file (config.toml) will be generated

   ![image-20240312162953082](https://github.com/bookqaq/010-record-api/blob/images/image-20240312162953082.png?raw=true)

2. edit config.toml, set listen_address and upload_service_address if necessary. If you want to run it locally, then you can skip this step.

   ![image-20240312170951686](https://github.com/bookqaq/010-record-api/blob/images/image-20240313170201467.png?raw=true)

3. open 010-record-api.exe again, service should start now

   ![image-20240312163453651](https://github.com/bookqaq/010-record-api/blob/images/image-20240312163453651.png?raw=true)

## Config Other Necessary Parts

No specfic order is required to finish these configurations.

### Install Driver for Graphics Card

this is just a hint

### Config Subscreen Touch

Refer https://github.com/spice2x/spice2x.github.io/wiki/Configuring-touch-screens-as-subscreen#step-by-step-instructions

### Config spice2x（Ignore this for bemanitools）

1. Download spice2x-24-02-13 or newer than this, extract to your folder

2. open spicecfg.exe, change toggle Disable D3D9 Device Hook (in Graphics (common), under options tab) to true  

   ![image-20240312164557472](https://github.com/bookqaq/010-record-api/blob/images/image-20240312164557472.png?raw=true)

### Config XRPC Server

#### Asphyxia

Assume game, asphyxia and 010-record-api.exe are **running on the same pc**, and config file of 010-record-api.exe **is not changed**

##### Check whether your plugin supports

open asphyxia, go to "IIDX" under Plugins. If "Movie Upload URL" appears in Plugin Settings, then your plugin support the feature.

![image-20240312165724390](https://github.com/bookqaq/010-record-api/blob/images/image-20240312165724390.png?raw=true)

##### If not support, How to edit your plugin

~~open your iidx@asphyxia folder, change below~~

find a zip named iidx-asphyxia-v1.4.4_a4.zip, remember to do a backup.

#### Other server

Ask your server's owner.

### Config bm2dx.dll

**only 2023-09-05 LDJ-010**

Open http://localhost:4399/patcher/ in your browser, apply changes to your LDJ-010 version of bm2dx.dll

Skip this step if your server region is already Japan

## After gamestart

Check "動画" button on sunscreen. If it appears,  configuration is completed.

# Contribution

Just submit your PRs, I'll check and reply.

# TODOs

- ~~Replace gin (would bump go to 1.22)~~
- ~~Basic http response process(we lost these abilities after we dont use gin)~~
- ~~Better video filename~~(sort of)
- ~~Better in-memory movie upload instance manage~~
- ~~Fix report as VIRUS on my pc~~(maybe fixed)
- Speed up the client upload speed by:
   - patch dll
   - implement direct file copy from RawPlayVideo
- finish the proxy package?