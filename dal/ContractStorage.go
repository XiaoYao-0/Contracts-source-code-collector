package dal

import (
	"fmt"
	"github.com/XiaoYao-0/Contracts-source-code-collector/conf"
	"github.com/XiaoYao-0/Contracts-source-code-collector/domain"
	"os"
	"strings"
)

// Save the result
// filepath: "contracts/${{index}}_${{contract.Name}}/${{sol.Name}}" e.g."contracts/56_ERCToken20/ownership.sol"
func (contractRepo ContractRepo) Save(contract *domain.Contract) error {
	dir, err := conf.StorageDir()
	if err != nil {
		return err
	}
	contractDir := fmt.Sprintf("%s/%d_%s", dir, contract.Number, contract.Name)
	err = os.Mkdir(contractDir, os.ModePerm)
	if err != nil {
		return err
	}
	for _, sol := range contract.Sols {
		solPath := ""
		if strings.Contains(sol.Name, "/") {
			end := strings.LastIndex(sol.Name, "/")
			solDir := fmt.Sprintf("%s/%s", contractDir, sol.Name[:end])
			err = os.MkdirAll(solDir, os.ModePerm)
			if err != nil {
				return err
			}
			solPath = fmt.Sprintf("%s%s", solDir, sol.Name[end:])
		} else {
			solPath = fmt.Sprintf("%s/%s", contractDir, sol.Name)
		}
		f, err := os.Create(solPath)
		if err != nil {
			return err
		}
		_, err = f.WriteString(sol.Src)
		if err != nil {
			return err
		}
	}
	return nil
}
