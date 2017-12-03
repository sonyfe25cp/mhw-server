package utils

import (
	"net/http"
	"encoding/json"
	"strconv"
	"com.github.sonyfe25cp.mhw-server/log"
	"net/url"
)

func WriteJson(w http.ResponseWriter, data interface{}) {
	json, err := json.Marshal(data)
	if err == nil {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(json))
	} else {
		w.Write([]byte(err.Error()))
	}
}

func StringtoInt(intString string) int {
	intValue, _ := strconv.Atoi(intString)
	return intValue
}

func StringtoIntWithDefault(intString string, defaultValue int) int {
	if len(intString) == 0 {
		return defaultValue
	} else {
		intValue, _ := strconv.Atoi(intString)
		return intValue
	}
}

func StringtoInt64(intString string) int64 {
	i, _ := strconv.ParseInt(intString, 10, 64)
	return i
}

func Int64ToString(intValue int64) string {
	return strconv.FormatInt(intValue, 10)
}

func IntToString(intValue int) string {
	return strconv.Itoa(intValue)
}

func DebugFormValues(form url.Values) {
	for k, v := range form {
		logs.Info(k, " : ", v)

	}
}
