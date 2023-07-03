package more_task

import (
	"fmt"
	"reflect"
	"strings"
)

type ReturnValue interface{}

type ConfigMap map[string]ReturnValue

func (m ConfigMap) SetKey(key string, value ReturnValue) error {
	m[key] = value

	return nil
}

func (m ConfigMap) Set(kv string) error {
	i := strings.Index(kv, "=")
	if i == -1 {
		return fmt.Errorf(" key value must seperate \"=\"")
	}

	k := kv[:i]
	v := kv[i+1:]

	return m.SetKey(k, v)
}
func (m ConfigMap) Get(key string, defval ReturnValue) (ReturnValue, error) {
	return m.get(key, defval)
}

func (m ConfigMap) GetStr(key string, defval ReturnValue) (string, error) {
	a, er := m.get(key, defval)
	if er != nil {
		return "", er
	}
	b, er := Value2string(a)
	return b, er
}
func (m ConfigMap) get(key string, defval ReturnValue) (ReturnValue, error) {
	v, ok := m[key]
	if !ok {
		return defval, nil
	}

	if defval != nil && reflect.TypeOf(defval) != reflect.TypeOf(v) {
		return nil, fmt.Errorf("%s expects type %T, not %T", key, defval, v)
	}

	return v, nil
}

func Value2string(v ReturnValue) (ret string, errstr error) {

	switch x := v.(type) {
	case bool:
		if x {
			ret = "true"
		} else {
			ret = "false"
		}
	case int:
		ret = fmt.Sprintf("%d", x)
	case string:
		ret = x
	case fmt.Stringer:
		ret = x.String()
	default:
		return "", fmt.Errorf("Invalid value type %T", v)
	}

	return ret, nil
}
