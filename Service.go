package main

import (
	"errors"
	"fmt"
	"github.com/XiaoYao-0/Contracts-source-code-collector/conf"
	"github.com/XiaoYao-0/Contracts-source-code-collector/dal"
	"github.com/XiaoYao-0/Contracts-source-code-collector/domain"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Init() error {
	err := initProxy()
	if err != nil {
		return err
	}
	err = initTestNet()
	if err != nil {
		return err
	}
	err = initStorageDir()
	if err != nil {
		return err
	}
	return nil
}

func initStorageDir() error {
	dir, err := conf.StorageDir()
	if err != nil {
		return err
	}
	d, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}
	if !d.IsDir() {
		return errors.New("storage dir already exists but is not a directory")
	}
	return nil
}

func initProxy() error {
	proxy, err := conf.Proxy()
	if err != nil {
		return err
	}
	if proxy == "" {
		return nil
	}
	err = os.Setenv("HTTP_PROXY", proxy)
	if err != nil {
		return err
	}
	err = os.Setenv("HTTPS_PROXY", proxy)
	if err != nil {
		return err
	}
	return nil
}

func initTestNet() error {
	_, err := http.Get("https://api.etherscan.io/api")
	if err != nil {
		return errors.New("proxy may be incorrect, because 'https://api.etherscan.io/api' can't be accessed")
	}
	return nil
}

func commonGetContracts(filepath string) error {
	fmt.Println("Init...")
	err := Init()
	if err != nil {
		return err
	}
	number, err := conf.ContractNumber()
	if err != nil {
		return err
	}
	fmt.Println("Getting contracts list...")
	contracts, err := dal.GetContracts(filepath)
	if err != nil {
		return err
	}
	fmt.Println("Collecting contract...")
	for _, contract := range contracts {
		contract.Number = number
		fmt.Printf("[+]Collecting contract %v\n", contract)
		contract.Collect()
		contract.Save()
		if contract.Error != nil {
			fmt.Println("[!]Error:", contract.Error)
		}
		number++
	}
	err = conf.SetContractNumber(number)
	if err != nil {
		return err
	}
	return nil
}

// GetContractsByList Get your own list of contracts
func GetContractsByList(filepath string) {
	err := commonGetContracts(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}

// GetSomeContracts Get default list of contracts from etherscan
func GetSomeContracts() {
	fmt.Println("Please try to download contracts list from https://etherscan.io/exportData?type=open-source-contract-codes")
}

// GetOneContract Get one contract
func GetOneContract(name string, address string) {
	fmt.Println("Init...")
	err := Init()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	number, err := conf.ContractNumber()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var contract domain.Contract
	contract.Number = number
	contract.Name = name
	contract.Address = address
	contract.ContractRepo = dal.ContractRepo{}
	fmt.Printf("[+]Collecting contract %v\n", contract)
	contract.Collect()
	contract.Save()
	if contract.Error != nil {
		fmt.Println("[!]Error:", contract.Error)
	}
	number++
	err = conf.SetContractNumber(number)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}

func GetTopERC20Tokens() {
	fmt.Println("Init...")
	contractList, err := dal.CollectERC20TokenList()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully get ERC20Token-address-list, saving...")
	err = dal.SaveContractList(contractList, "ERC20Token.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully save ERC20Token-address-list, filename: ERC20Token.csv")
	return
}

func GetTopERC721Tokens() {
	fmt.Println("Init...")
	contractList, err := dal.CollectERC721TokenList()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully get ERC721Token-address-list, saving...")
	err = dal.SaveContractList(contractList, "ERC721Token.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully save ERC721Token-address-list, filename: ERC721Token.csv")
	return
}

func VerifyBalanceOfInERC721Tokens(dir string) {
	fmt.Println("Init...")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			files1, err := ioutil.ReadDir(fmt.Sprintf("%v/%v", dir, file.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if len(files1) != 1 {
				fmt.Printf("Warning: %v has multi sols\n", file.Name())
				continue
			}
			file1 := files1[0]
			bytes, _ := ioutil.ReadFile(fmt.Sprintf("%v/%v/%v", dir, file.Name(), file1.Name()))
			solCode := string(bytes)
			res := verifyBalanceOfInOneERC721Token(solCode)
			if res != "" {
				fmt.Printf("[-]ERC721Token %v not correct\n", file.Name())
			}
		}
	}
}

func verifyBalanceOfInOneERC721Token(src string) string {
	const StandardBalanceOfFunction = "function balanceOf(address owner) public view virtual override returns (uint256) {\n        require(owner != address(0), \"ERC721: balance query for the zero address\");\n        return _balances[owner];\n    }"
	if strings.Contains(src, StandardBalanceOfFunction) {
		return ""
	}
	return "1"
}
