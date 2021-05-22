package main

import (
	"flag"
	"fmt"
)

func cmdSingle(name, address string) {
	if name == "" || len(address) != 42 {
		fmt.Println("To use mode 'single', '-n'(name) and '-a'(address) are required, and address should be 42 bits long e.g.0xf73ee4f0b82c57ebede359dd5a98d368838b01ea")
		return
	}
	GetOneContract(name, address)
	return
}

func cmdMulti(filepath string) {
	if filepath == "" {
		fmt.Println("To use mode 'multi', '-f'(filepath) is required")
		return
	}
	GetContractsByList(filepath)
	return
}

func cmdExample() {
	GetSomeContracts()
	return
}

func cmdERC20Token() {
	GetTopERC20Tokens()
	return
}

func main() {
	mode := flag.String("m", "", "mode:"+
		"\n'single': get one contract by name and address, require \"-n\"(name) and \"-a\"(address)"+
		"\n'multi': get contracts by contracts list, require \"-f\"(filepath)"+
		"\nexample list file: contract-address-list.csv"+
		"\nNote: For the actual contract source codes use the api endpoints at https://etherscan.io/apis#contracts"+
		"\n\"Txhash\",\"ContractAddress\",\"ContractName\""+
		"\n\"0x2530fe6c225ba7394aba96ebc48ff2ae17949b155b256969282ed4431c599400\",\"0xf73ee4f0b82c57ebede359dd5a98d368838b01ea\",\"Valorant_Token\""+
		"\n\"0x6b734835970ca79853a105a82ecfe607ed4a3330719201785847489438783036\",\"0x0cd75d7b8fb785f186165bdc280489ea750bad17\",\"Token\""+
		"\n\"0x0190ac3ae873c41f7ef08a474789a01fa9947f5c5094668f674c5566868275c9\",\"0xb9b5bea373074b869b721c1cb38cc837f1b061b5\",\"LOCG\""+
		"\n\"0x81c84e3904e1d324eb9c5110576515a1394a3e665ac1373741b2c17c459416cb\",\"0xced0b3b9f30f12332c77ad937d6137ce94162d9d\",\"VEAN\""+
		"\n'example': get example contracts list from etherscan, require nothing"+
		"\n'ERC20': get top ERC20 Token address list, require nothing")
	name := flag.String("n", "", "name: contract name")
	address := flag.String("a", "", "address: contract address")
	filepath := flag.String("f", "contract-address-list.csv", "filepath: contract list filepath")
	flag.Parse()
	switch *mode {
	case "single":
		cmdSingle(*name, *address)
	case "multi":
		cmdMulti(*filepath)
	case "example":
		cmdExample()
	case "ERC20":
		cmdERC20Token()
	default:
		flag.Usage()
	}
}
