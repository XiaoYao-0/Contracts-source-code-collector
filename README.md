### Usage

Usage of collect.exe:
  -a string
        address: contract address
  -f string
        filepath: contract list filepath (default "contract-address-list.csv")
  -m string
        mode:
        'single': get one contract by name and address, require "-n"(name) and "-a"(address)
        'multi': get contracts by contracts list, require "-f"(filepath)
        example list file: contract-address-list.csv
        Note: For the actual contract source codes use the api endpoints at https://etherscan.io/apis#contracts
        "Txhash","ContractAddress","ContractName"

"0x2530fe6c225ba7394aba96ebc48ff2ae17949b155b256969282ed4431c599400","0xf73ee4f0b82c57ebede359dd5a98d368838b01ea","Valorant_Token"        "0x6b734835970ca79853a105a82ecfe607ed4a3330719201785847489438783036","0x0cd75d7b8fb785f186165bdc280489ea750bad17","Token"        "0x0190ac3ae873c41f7ef08a474789a01fa9947f5c5094668f674c5566868275c9","0xb9b5bea373074b869b721c1cb38cc837f1b061b5","LOCG"
"0x81c84e3904e1d324eb9c5110576515a1394a3e665ac1373741b2c17c459416cb","0xced0b3b9f30f12332c77ad937d6137ce94162d9d","VEAN"
        'example': get example contracts list from etherscan, require nothing
  -n string
        name: contract name

### 三种 mode

single：获取单个合约代码，需要 name 和 address

multi：通过规范的列表获取一堆合约的源代码，需要列表文件路径

example：通过etherscan获取一个包含10000个开源合约地址的列表

### 配置信息 Conf

conf.json 里面用 json 格式包含了配置信息

```json
{
    "contract_number":0,
    "proxy":"",
    "api_key":"",
    "storage_dir":"./contracts"
}
```

contract_number 这是合约的编号，编号只是用来防止重名的合约，用来给收集到的合约文件夹命名，比如收集A，B合约，命名就是0-A，1-B，然后这个值会变成2，下次获取就是3-C。通过列表获取一堆合约时，请**原子性**执行，只要这一堆没执行完，number 就不会自动变，会导致后面的命名编号重复。建议不要动这个玩意儿。

proxy 是代理 ip: port，因为 etherscan 需要翻墙

api_key 是 etherscan 的 api key，没有的话会 5s 才能访问一次，可以自己去官网申请一个

storage_dir 是下载到的合约代码存储路径

### 示例

![image-20210503225928991](README.assets\image-20210503225928991.png)