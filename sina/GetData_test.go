package sina

import "testing"

func TestSinaData_GetData(t *testing.T) {
	strings := []string{"sh513050", "sh600001"}
	(&SinaData{}).GetData(strings)
}
