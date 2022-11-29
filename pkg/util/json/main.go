package json

import (
	"encoding/json"
)

func StructToJSONStr(obj interface{}) (str string, err error) {
	b, marshalErr := json.Marshal(obj)
	if err != nil {
		err = marshalErr
		return
	}
	return string(b), nil
}
