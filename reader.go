package goconfig

import (
	"reflect"
	"strconv"
	"strings"
)

type Reader struct {
	cfg *Config
}

func (r *Reader) assign(properties interface{}) error {
	v := reflect.Indirect(reflect.ValueOf(properties))
	return r.parse(v)
}

func (r *Reader) parse(value reflect.Value) error {
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Kind() == reflect.Struct {
			r.parse(value.Field(i))
		} else {
			tag := value.Type().Field(i).Tag.Get("prop")
			v, e := r.getValue(tag)
			if e != nil {
				return e
			}
			switch value.Field(i).Kind() {
			case reflect.Int:
				pv, e := strconv.ParseInt(v, 10, 32)
				if e != nil {
					return e
				}
				value.Field(i).SetInt(pv)
			case reflect.Int64:
				pv, e := strconv.ParseInt(v, 10, 64)
				if e != nil {
					return e
				}
				value.Field(i).SetInt(pv)
			case reflect.Float32:
				pv, e := strconv.ParseFloat(v, 32)
				if e != nil {
					return e
				}
				value.Field(i).SetFloat(pv)
			case reflect.Float64:
				pv, e := strconv.ParseFloat(v, 64)
				if e != nil {
					return e
				}
				value.Field(i).SetFloat(pv)
			case reflect.Bool:
				pv, e := strconv.ParseBool(v)
				if e != nil {
					return e
				}
				value.Field(i).SetBool(pv)
			case reflect.String:
				value.Field(i).SetString(v)
			case reflect.Slice:
				separator, ok := value.Type().Field(i).Tag.Lookup("separator")
				if !ok {
					separator = ","
				}
				vs := strings.Split(v, separator)
				switch value.Field(i).Type().String() {
				case "[]int":
					ar := make([]int, len(vs))
					for k, _ := range vs {
						pv, e := strconv.ParseInt(vs[k], 10, 32)
						if e != nil {
							return e
						}
						ar[k] = int(pv)
					}
					value.Field(i).Set(reflect.ValueOf(ar))
				case "[]int64":
					ar := make([]int64, len(vs))
					for k, _ := range vs {
						pv, e := strconv.ParseInt(vs[k], 10, 64)
						if e != nil {
							return e
						}
						ar[k] = pv
					}
					value.Field(i).Set(reflect.ValueOf(ar))
				case "[]float32":
					ar := make([]float32, len(vs))
					for k, _ := range vs {
						pv, e := strconv.ParseFloat(vs[k], 32)
						if e != nil {
							return e
						}
						ar[k] = float32(pv)
					}
					value.Field(i).Set(reflect.ValueOf(ar))
				case "[]float64":
					ar := make([]float64, len(vs))
					for k, _ := range vs {
						pv, e := strconv.ParseFloat(vs[k], 64)
						if e != nil {
							return e
						}
						ar[k] = pv
					}
					value.Field(i).Set(reflect.ValueOf(ar))
				case "[]string":
					value.Field(i).Set(reflect.ValueOf(vs))
				case "[]bool":
					ar := make([]bool, len(vs))
					for k, _ := range vs {
						pv, e := strconv.ParseBool(vs[k])
						if e != nil {
							return e
						}
						ar[k] = pv
					}
					value.Field(i).Set(reflect.ValueOf(ar))
				}
			}
		}
	}
	return nil
}

func (r *Reader) getValue(key string) (string, error) {
	if v, b := getFromEnv(key); b {
		return v, nil
	}
	return r.cfg.getFromConfig(key)
}

func Init(fileName string, properties interface{}) error {
	config, err := load(fileName)
	if err != nil {
		return err
	}
	reader := &Reader{
		cfg: config,
	}
	return reader.assign(properties)
}
