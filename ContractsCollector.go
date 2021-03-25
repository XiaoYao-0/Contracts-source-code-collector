package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ApiJson struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Sourcecode           string `json:"SourceCode"`
		Abi                  string `json:"ABI"`
		ContractName         string `json:"ContractName"`
		CompilerVersion      string `json:"CompilerVersion"`
		OptimizationUsed     string `json:"OptimizationUsed"`
		Runs                 string `json:"Runs"`
		ConstructorArguments string `json:"ConstructorArguments"`
		EvmVersion           string `json:"EVMVersion"`
		Library              string `json:"Library"`
		LicenseType          string `json:"LicenseType"`
		Proxy                string `json:"Proxy"`
		Implementation       string `json:"Implementation"`
		SwarmSource          string `json:"SwarmSource"`
	} `json:"result"`
}

type Contract struct {
	Address string
	Name    string
	Json    string
	Sols    []Sol
}

type Sol struct {
	Name string
	Src  string
}

// Save the result to "contracts/{{index}}_{{contract.Name}}/{{sol.Name}}" e.g."contracts/56_ERCToken20/ownership.sol"
func (contract *Contract) save() error {
	contractDir := fmt.Sprintf("contracts/%s", contract.Name)
	err := os.Mkdir(contractDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error: Failed to save contract {%s}, address = \"%s\"\n", contract.Name, contract.Address)
		return err
	}
	for _, sol := range contract.Sols {
		solPath := ""
		if strings.Contains(sol.Name, "/") {
			end := strings.LastIndex(sol.Name, "/")
			solDir := fmt.Sprintf("%s/%s", contractDir, sol.Name[:end])
			err = os.MkdirAll(solDir, os.ModePerm)
			if err != nil {
				fmt.Printf("Error: Failed to create %s when saving cont ract {%s}, address = \"%s\"\n", solDir, contract.Name, contract.Address)
				return err
			}
			solPath = fmt.Sprintf("%s%s", solDir, sol.Name[end:])
		} else {
			solPath = fmt.Sprintf("%s/%s", contractDir, sol.Name)
		}
		f, err := os.Create(solPath)
		if err != nil {
			fmt.Printf("Error: Failed to create %s when saving contract {%s}, address = \"%s\"\n", solPath, contract.Name, contract.Address)
			return err
		}
		_, err = f.WriteString(sol.Src)
		if err != nil {
			fmt.Printf("Error: Failed to write %s when saving contract {%s}, address = \"%s\"\n", solPath, contract.Name, contract.Address)
			return err
		}
	}
	return nil
}

// Unmarshal the SourceCode json string of a contract
func (contract *Contract) unmarshal() error {
	var src interface{}
	if strings.HasPrefix(contract.Json, "{{") {
		contract.Json = contract.Json[1 : len(contract.Json)-1]
		err := json.Unmarshal([]byte(contract.Json), &src)
		if err != nil {
			fmt.Printf("Error: Failed to unmarshal {%s}, address = \"%s\"\n", contract.Name, contract.Address)
			return err
		}
		src0, ok := src.(map[string]interface{})
		if ok {
			for k, v := range src0 {
				if k == "sources" {
					src1, ok1 := v.(map[string]interface{})
					if ok1 {
						for k1, v1 := range src1 {
							src2, ok2 := v1.(map[string]interface{})
							if ok2 {
								for _, v2 := range src2 {
									v3, ok3 := v2.(string)
									if ok3 {
										sol := Sol{k1, v3}
										contract.Sols = append(contract.Sols, sol)
									} else {
										fmt.Printf("Error: Failed to unmarshal {%s}, address = \"%s\"\n", contract.Name, contract.Address)
										err = errors.New("unmarshal error")
										return err
									}
								}
							} else {
								fmt.Printf("Error: Failed to unmarshal {%s}, address = \"%s\"\n", contract.Name, contract.Address)
								err = errors.New("unmarshal error")
								return err
							}
						}
						return nil
					} else {
						fmt.Printf("Error: Failed to unmarshal {%s}, address = \"%s\"\n", contract.Name, contract.Address)
						err = errors.New("unmarshal error")
						return err
					}
				}
			}
		} else {
			fmt.Printf("Error: Failed to unmarshal {%s}, address = \"%s\"\n", contract.Name, contract.Address)
			err = errors.New("unmarshal error")
			return err
		}
	}
	err := json.Unmarshal([]byte(contract.Json), &src)
	if err != nil {
		solName := fmt.Sprintf("%s.sol", strings.Split(contract.Name, "_")[1])
		sol := Sol{solName, contract.Json}
		contract.Sols = append(contract.Sols, sol)
		return nil
	}
	src0, ok := src.(map[string]interface{})
	if ok {
		for k, v := range src0 {
			v0, ok0 := v.(map[string]interface{})
			if ok0 {
				for _, v1 := range v0 {
					if v2, ok := v1.(string); ok {
						sol := Sol{k, v2}
						contract.Sols = append(contract.Sols, sol)
					}
				}
			} else {
				fmt.Printf("Error: Failed to unmarshal {%s}, address = \"%s\"\n", contract.Name, contract.Address)
				err = errors.New("unmarshal error")
				return err
			}
		}
	} else {
		fmt.Printf("Error: Failed to unmarshal {%s}, address = \"%s\"\n", contract.Name, contract.Address)
		err = errors.New("unmarshal error")
		return err
	}
	return nil
}

// Collect sourceCode json of every contract address
func (contract *Contract) collect() error {
	fmt.Printf("Getting {%s}, address = \"%s\"...\n", contract.Name, contract.Address)
	// TODO：SET YOUR API-KEY HERE
	apiKey := ""
	url := fmt.Sprintf("https:// api.etherscan.io/api?module=contract&action=getsourcecode&address=%s&apikey=%s",
		contract.Address, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: Failed to get {%s}, address = \"%s\"\n", contract.Name, contract.Address)
		return err
	}
	apiJson := ApiJson{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to get {%s}, address = \"%s\"\n", contract.Name, contract.Address)
		return err
	}
	resp.Body.Close()
	err = json.Unmarshal([]byte(body), &apiJson)
	if err != nil {
		fmt.Printf("Error: Failed to get {%s}, address = \"%s\"\n", contract.Name, contract.Address)
		return err
	}
	if len(apiJson.Result) == 0 {
		fmt.Printf("Error: Result of {%s} is empty, address = \"%s\"\n", contract.Name, contract.Address)
		err = errors.New("collect result empty error")
		return nil
	}
	if len(apiJson.Result) > 1 {
		fmt.Printf("Warning: Result of {%s} has 2 or more srcs, address = \"%s\"\n", contract.Name, contract.Address)
	}
	contract.Json = apiJson.Result[0].Sourcecode
	err = contract.unmarshal()
	if err != nil {
		return err
	}
	err = contract.save()
	if err != nil {
		return err
	}
	return nil
}

// Read contracts-list to execute collecting method
func getContracts(filepath string) {
	var errList []string
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error: Can't open contracts list, filepath \"%s\"\n", filepath)
		return
	}
	defer f.Close()

	br := bufio.NewReader(f)
	lineNum := 2
	for i := 0; i < lineNum; i++ {
		_, _, err = br.ReadLine()
		if err != nil {
			fmt.Printf("Error: Can't read 1st of contracts list, filepath \"%s\"\n", filepath)
			return
		}
	}

	for i := lineNum - 1; ; i++ {

		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		split := strings.Split(string(a), ",")
		contract := &Contract{}
		contract.Address = split[1]
		contract.Name = fmt.Sprintf("%d_%s", i, split[2])
		contract.Json = ""
		contract.Sols = []Sol{}
		err = contract.collect()
		if err != nil {
			errList = append(errList, contract.Name+contract.Address)
		}
	}
	fmt.Println("Error List:", errList)
}

// Init the proxy and make root directory to save contracts
func init() {
	// Set up proxy for Chinese Wall
	// Replace it if you need proxy, otherwise delete it
	proxy := "http:// 127.0.0.1:7890" // replace it with your proxy address or use system proxy instead
	err := os.Setenv("HTTP_PROXY", proxy)
	if err != nil {
		fmt.Println("Error: Failed to set up http_proxy")
	}
	err = os.Setenv("HTTPS_PROXY", proxy)
	if err != nil {
		fmt.Println("Error: Failed to set up https_proxy")
	}

	// Make a new directory named {contracts} to save the result
	err = os.Mkdir("./contracts", os.ModePerm)
	if err != nil {
		fmt.Println("Error: Failed to make a new directory named {contracts} to save the result")
	}
}

func main() {
	getContracts("./contract-address-list.txt")
}
