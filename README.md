## 一个用于获取以太坊上开源智能合约的小脚本

### 数据源

https://etherscan.io/apis#contracts，etherscan提供的开源合约源代码的api，申请开发者key，提供合约的address即可获取源码等信息。address list亦来自于etherscan，https://etherscan.io/exportData?type=open-source-contract-codes，已提前下载了2021年3月21日的address list并放入文件中，需要更新list可以自行下载替换

### json解析

提供的json有多种不同格式，限于水平使用了较脏的解析代码

### 结果存储

为每个合约建立文件夹，还原各合约包含的多个源代码文件的目录结构，避免import出现问题

### 代理

由于wall的关系，我设置了代理，如果不需要请直接删掉

### Api key

访问api需要api key，申请开发者后在个人信息中获取

------

**English description is translated by Google Translator** 

## Contracts-source-code-collector

A small script to collect source codes of contracts from api of Etherscan

### Data source

https://etherscan.io/apis#contracts, the api of the open source contract source code provided by etherscan, apply for the developer key and provide the address of the contract to obtain the source code and other information. The address list is also from etherscan, https://etherscan.io/exportData?type=open-source-contract-codes. The address list on March 21, 2021 has been downloaded in advance and put into the file. If you need to update the list, download and replace by yourself.

### json parsing

The provided json has a variety of different formats, so I wrote a dirty parsing code.

### Result storage

create a folder for each contract, restore the directory structure of multiple source code files contained in each contract, and avoid import problems.

### Proxy

Due to the relationship of the wall, I set up a proxy, please delete it if you don’t need it.

### Api key

api key is required to access api, which can be obtained from personal information after applying for developer.


