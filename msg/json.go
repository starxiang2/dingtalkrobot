package msg

import (
	"fmt"
)

type toJson struct {
}

func (t *toJson) ToJson() (string, error) {
	fmt.Println(t)
	return "", nil
	//marshal, err := json.Marshal(t)
	//if err != nil {
	//	return "", errors.New("json 转换失败")
	//}
	//return string(marshal), nil
}
