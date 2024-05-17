package _7_default

import (
	"fmt"
	"reflect"
	"strconv"
)

// Go 加载配置时，利用标签和反射设置默认值，适用于值不存在的但需要设置默认值的场景
// 定义新类型：输出具体的错误

// ErrorNotAStructPointer 错误：非结构体指针类型
type ErrorNotAStructPointer string

// 函数：格式化 v 的类型并强转为 ErrorNotAStructPointer
func newErrorNotAStructPointer(v interface{}) ErrorNotAStructPointer {
	return ErrorNotAStructPointer(fmt.Sprintf("%t", v))
}

// 实现Error接口
func (e ErrorNotAStructPointer) Error() string {
	return fmt.Sprintf("expected a struct, instead got a %T", string(e))
}

type ErrorUnsettable string

// Error implements the error interface.
func (e ErrorUnsettable) Error() string {
	return fmt.Sprintf("can't set field %s", string(e))
}

// ErrorUnSupportedType 定义错误类型结构体表示不支持的类型错误，持有反射类型
type ErrorUnSupportedType struct {
	t reflect.Type
}

func (e ErrorUnSupportedType) Error() string {
	return fmt.Sprintf("unsupported type %v", e.t)
}

// Apply 获取值，判断值是指针类型或者结构体类型 （涉及Kind枚举的比对）
func Apply(t interface{}) error {
	val := reflect.ValueOf(t)
	if val.Kind() != reflect.Ptr {
		return newErrorNotAStructPointer(t)
	}
	// 获取结构体指针指向的结构体的值
	ref := val.Elem()
	if ref.Kind() != reflect.Struct {
		return newErrorNotAStructPointer(t)
	}

	// 属性转换
	return parseFields(ref)
}

func parseFields(v reflect.Value) error {
	// 获取结构体中字段的总数，遍历每一个字段，获取字段值和类型
	for i := 0; i < v.NumField(); i++ {
		err := parseField(v.Field(i), v.Type().Field(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func parseField(value reflect.Value, field reflect.StructField) error {
	// 获取属性里面的标签 default
	defaultVal := field.Tag.Get("default")

	// 判断值是否为结构体类型 或 是否为指针类型
	isStruct := value.Kind() == reflect.Struct
	isStructPointer := value.Kind() == reflect.Ptr && value.Type().Elem().Kind() == reflect.Struct

	if (defaultVal == "" || defaultVal == "-") && !(isStruct || isStructPointer) {
		return nil
	}

	if !value.CanSet() {
		return nil
	}

	if !value.IsZero() {
		// A value is set on this field so there's no need to set a default
		// value.
		return nil
	}

	// 判断值的类型 非字符串可能涉及转换 字符串直接set
	switch value.Kind() {

	case reflect.String:
		value.SetString(defaultVal)
		return nil

	case reflect.Bool:
		// 判断是否可以转为bool
		b, err := strconv.ParseBool(defaultVal)
		if err != nil {
			return err
		}
		value.SetBool(b)
		return nil

	case reflect.Int:
		// 判断是否可以转为 4字节 int
		intVal, err := strconv.ParseInt(defaultVal, 10, 32)
		if err != nil {
			return err
		}
		value.SetInt(intVal)
		return nil

	case reflect.Int8:
		// 判断是否可以转为 1字节 int
		intVal, err := strconv.ParseInt(defaultVal, 10, 8)
		if err != nil {
			return err
		}
		value.SetInt(intVal)
		return nil

	case reflect.Int16:
		// 判断是否可以转为 2字节 int
		intVal, err := strconv.ParseInt(defaultVal, 10, 16)
		if err != nil {
			return err
		}
		value.SetInt(intVal)
		return nil

		// ！！记得特殊处理 int32 和 rune
	case reflect.Int32:
		// 判断是否可以转为 4字节 int
		intVal, err := parseInt32(defaultVal)
		if err != nil {
			return err
		}
		value.SetInt(int64(intVal))
		return nil
	case reflect.Int64:
		// 判断是否可以转为 4字节 int
		intVal, err := strconv.ParseInt(defaultVal, 10, 64)
		if err != nil {
			return err
		}
		value.SetInt(intVal)
		return nil

	case reflect.Uint:
		intVal, err := strconv.ParseInt(defaultVal, 10, 32)
		if err != nil {
			return err
		}
		value.SetUint(uint64(intVal))
		return nil

	case reflect.Uint8:
		intVal, err := strconv.ParseInt(defaultVal, 10, 8)
		if err != nil {
			return err
		}
		value.SetUint(uint64(intVal))
		return nil

	case reflect.Uint16:
		intVal, err := strconv.ParseInt(defaultVal, 10, 16)
		if err != nil {
			return err
		}
		value.SetUint(uint64(intVal))
		return nil

	case reflect.Uint32:
		intVal, err := strconv.ParseInt(defaultVal, 10, 32)
		if err != nil {
			return err
		}
		value.SetUint(uint64(intVal))
		return nil

	case reflect.Uint64:
		intVal, err := strconv.ParseInt(defaultVal, 10, 64)
		if err != nil {
			return err
		}
		value.SetUint(uint64(intVal))
		return nil

	case reflect.Float32:
		floatVal, err := strconv.ParseFloat(defaultVal, 32)
		if err != nil {
			return err
		}
		value.SetFloat(floatVal)
		return nil

	case reflect.Float64:
		floatVal, err := strconv.ParseFloat(defaultVal, 64)
		if err != nil {
			return err
		}
		value.SetFloat(floatVal)
		return nil

	case reflect.Slice:
		// 这里只处理 byte类型的切片（绝大多数情况下只有该类型），如：configItem []byte 默认值为："data"
		// 所以，获取切片类类型，判断是否Uint8
		switch value.Type().Elem().Kind() {
		// 仅处理 []byte
		case reflect.Uint8:
			value.SetBytes([]byte(defaultVal))
			return nil
		default:
			return ErrorUnSupportedType{value.Type()}
		}
	case reflect.Struct:
		// 处理结构体类型：注意不处理空结构体，处理逻辑递归处理结构体的每一个属性即可
		if value.NumField() == 0 {
			return nil
		}
		return parseFields(value)

	case reflect.Ptr:
		// 指针类型存储的是数据的地址
		// 处理各种指针类型：基本类型 和 复杂类型（切片和结构体）
		// 获取指针变量指向的元素的类型
		ref := value.Type().Elem()
		switch ref.Kind() {
		case reflect.String:
			// 设置地址
			value.Set(reflect.ValueOf(&defaultVal))
			return nil
		case reflect.Bool:
			// 判断默认值字符串是否可以转为Bool 可以便设置值地址作为默认值，不可以则返回错误
			boolVal, err := strconv.ParseBool(defaultVal)
			if err != nil {
				return err
			}
			value.Set(reflect.ValueOf(&boolVal))
			return nil
		case reflect.Int8:
			intVal, err := strconv.ParseInt(defaultVal, 10, 8)
			if err != nil {
				return err
			}
			// 64 转 8
			intSwitchResult := int8(intVal)
			// 设置
			value.Set(reflect.ValueOf(&intSwitchResult))
			return nil
		case reflect.Int:
			intVal, err := strconv.ParseInt(defaultVal, 10, 32)
			if err != nil {
				return err
			}
			// 64 转 32
			intSwitchResult := int(intVal)
			// 设置
			value.Set(reflect.ValueOf(&intSwitchResult))
			return nil
		case reflect.Int16:
			intVal, err := strconv.ParseInt(defaultVal, 10, 16)
			if err != nil {
				return err
			}
			// 64 转 16
			intSwitchResult := int16(intVal)
			// 设置
			value.Set(reflect.ValueOf(&intSwitchResult))
			return nil
		case reflect.Int32:
			intVal, err := parseInt32(defaultVal)
			if err != nil {
				return err
			}
			// 设置
			value.Set(reflect.ValueOf(&intVal))
			return nil
		case reflect.Int64:
			intVal, err := strconv.ParseInt(defaultVal, 10, 64)
			if err != nil {
				return err
			}
			// 设置
			value.Set(reflect.ValueOf(&intVal))
			return nil
		case reflect.Uint8:
			intVal, err := strconv.ParseInt(defaultVal, 10, 8)
			if err != nil {
				return err
			}
			// 64 转 8
			intSwitchResult := uint8(intVal)
			// 设置
			value.Set(reflect.ValueOf(&intSwitchResult))
			return nil
		case reflect.Uint:
			intVal, err := strconv.ParseInt(defaultVal, 10, 32)
			if err != nil {
				return err
			}
			// 64 转 32
			intSwitchResult := uint(intVal)
			// 设置
			value.Set(reflect.ValueOf(&intSwitchResult))
			return nil
		case reflect.Uint16:
			intVal, err := strconv.ParseInt(defaultVal, 10, 16)
			if err != nil {
				return err
			}
			// 64 转 16
			intSwitchResult := uint16(intVal)
			// 设置
			value.Set(reflect.ValueOf(&intSwitchResult))
			return nil
		case reflect.Uint32:
			intVal, err := strconv.ParseInt(defaultVal, 10, 32)
			if err != nil {
				return err
			}
			// 64 转 32
			intSwitchResult := uint32(intVal)
			// 设置
			value.Set(reflect.ValueOf(&intSwitchResult))
			return nil
		case reflect.Uint64:
			intVal, err := strconv.ParseInt(defaultVal, 10, 64)
			if err != nil {
				return err
			}
			// 64 转 8
			intSwitchResult := uint64(intVal)
			// 设置
			value.Set(reflect.ValueOf(&intSwitchResult))
			return nil
		case reflect.Float32:
			floatVal, err := strconv.ParseFloat(defaultVal, 32)
			if err != nil {
				return err
			}
			// 64 转 32
			floatSwitchResult := float32(floatVal)
			// 设置
			value.Set(reflect.ValueOf(&floatSwitchResult))
			return nil
		case reflect.Float64:
			floatVal, err := strconv.ParseFloat(defaultVal, 64)
			if err != nil {
				return err
			}
			// 设置
			value.Set(reflect.ValueOf(&floatVal))
			return nil

		case reflect.Slice:
			// 切片指针元素类型，同理仅处理[]byte
			refElemKind := ref.Elem().Kind()
			switch refElemKind {
			case reflect.Uint8:
				bytes := []byte(defaultVal)
				value.Set(reflect.ValueOf(&bytes))
				return nil
			default:
				return ErrorUnSupportedType{value.Type()}
			}

		case reflect.Struct:
			// 结构体指针类型：无属性不处理 和 值为nil新建对应结构体实例
			if ref.NumField() == 0 {
				return nil
			}
			if value.IsNil() {
				value.Set(reflect.New(ref))
			}
			// 执行递归
			return parseFields(value.Elem())
		default:
			return ErrorUnSupportedType{value.Type()}
		}
	default:
		return ErrorUnSupportedType{value.Type()}
	}

}

func parseInt32(str string) (int32, error) {
	// 处理 rune 和 int32 类型
	// 尝试转换为 int32
	intVal, err := strconv.ParseInt(str, 10, 32)
	if err == nil {
		// 是int32
		return int32(intVal), nil
	}

	// 是 rune 且 为单一字符才返回对应字符的unicode ，否则返回0
	// 放在一个[]rune
	runes := []rune(str)
	if len(runes) == 1 {
		return runes[0], nil
	} else {
		return 0, nil
	}
}
