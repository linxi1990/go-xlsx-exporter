package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wingcd/go-xlsx-exporter/model"
	"github.com/wingcd/go-xlsx-exporter/settings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	pref "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	valueTypeRegx *regexp.Regexp
)

func init() {
	valueTypeRegx, _ = regexp.Compile(`^(?P<name>\w+)(?P<array>\[(?P<split>.?)\])?(?P<conv>\??)$`)
}

func CompileValueType(valueType string) (valiable bool, baseValue string, isArray bool, splitChar string, converable, isVoid bool) {
	valueType = strings.Replace(valueType, " ", "", -1)
	var match = valueTypeRegx.FindStringSubmatch(valueType)
	valiable = len(match) == 5
	if !valiable {
		return
	}
	baseValue = match[1]
	isArray = match[2] != ""
	splitChar = match[3]
	if splitChar == "" {
		splitChar = settings.ArraySplitChar
	}
	converable = match[4] != ""
	isVoid = IsVoid(baseValue)
	if isVoid {
		converable = true
	}
	return
}

func Split(s, sep string) []string {
	strs := strings.Split(s, sep)
	rstrs := make([]string, 0)
	needConn := false
	// 解决分隔符转义问题
	for si := 0; si < len(strs); si++ {
		s := strs[si]
		if s != "" {
			if needConn {
				last := rstrs[len(rstrs)-1] + s
				rstrs[len(rstrs)-1] = last
				needConn = false
			} else {
				rstrs = append(rstrs, s)
			}
		} else if si > 0 {
			needConn = true
		}
	}
	return rstrs
}

var standardTypes = map[string]string{
	"bool":    "bool",
	"int":     "int",
	"int32":   "int",
	"uint":    "uint",
	"uint32":  "uint",
	"int64":   "int64",
	"uint64":  "uint64",
	"float":   "float",
	"float32": "float",
	"double":  "double",
	"float64": "double",
	"string":  "string",
}

var supportProtoTypes = map[string]string{
	"bool":    "bool",
	"int":     "int32",
	"int32":   "int32",
	"uint":    "uint32",
	"uint32":  "uint32",
	"int64":   "int64",
	"uint64":  "uint64",
	"float":   "float",
	"float32": "float",
	"double":  "double",
	"float64": "double",
	"string":  "string",
}

func ConvertToStandardType(valueType string) string {
	if tp, ok := standardTypes[valueType]; ok {
		return tp
	}
	return valueType
}

func ToEnumString(valueType string, value int32) string {
	for _, e := range settings.ENUMS {
		if valueType == e.TypeName {
			for _, it := range e.Items {
				if it.Value == strconv.Itoa(int(value)) {
					return it.FieldName
				}
			}
		}
	}
	return ""
}

//
func ConvertEnumValue(info *model.DefineTableInfo, valueType, value string) (error, int32) {
	var ret int64 = -1
	var err error

	if value == "" {
		return err, 0
	}

	var findType = false
	var findField = false
	for _, item := range info.Items {
		if info.TypeName == valueType {
			findType = true
			if item.FieldName == value {
				ret, _ = strconv.ParseInt(item.Value, 10, 32)
				findField = true
				break
			}
		}
	}
	if !findType {
		err = errors.New("找不到类型：" + valueType)
	} else if !findField {
		err = errors.New(fmt.Sprintf("找不到类型：%s中的字段：%s", valueType, value))
	}

	return err, int32(ret)
}

// 获取枚举类型的值
func ParseEnumValue(info *model.DefineTableInfo, valueType, value string) (success bool, ret interface{}, isArray bool) {
	_, valueType, repeated, splitChar, _, _ := CompileValueType(valueType)

	var err error

	if !repeated {
		err, ret = ConvertEnumValue(info, valueType, value)
		if err != nil {
			fmt.Printf("[错误] 值类型转换失败:%v [类型:%s 值：%v] \n", err, valueType, value)
			return false, value, repeated
		}
	} else {
		ret = make([]interface{}, 0)

		rstrs := Split(value, splitChar)

		for _, vstr := range rstrs {
			err, rvalue := ConvertEnumValue(info, valueType, vstr)

			if err != nil {
				fmt.Printf("[错误] 数组类型转换失败:%s [类型:%s 值：%s 子项：%s] \n", err, valueType, value, vstr)
				return false, value, repeated
			}

			ret = append(ret.([]interface{}), rvalue)
		}
	}

	return true, ret, repeated
}

func ResolveEnumValue(valueType, cellValue string) (success bool, ret interface{}, isArray bool) {
	for _, item := range settings.DEFINES {
		if item.Category == model.DEFINE_TYPE_ENUM {
			if item.TypeName == valueType {
				success, ret, isArray = ParseEnumValue(item, valueType, cellValue)
				return
			}
		}
	}
	return false, nil, false
}

// 将表格中支持的类型转换为protobuf支持的类型
func ParseType(vtype string) (bool, string) {
	_, vtype, repeated, _, _, _ := CompileValueType(vtype)

	if tp, ok := supportProtoTypes[vtype]; !ok {
		return false, ""
	} else if repeated {
		return true, "repeated " + tp
	} else {
		return true, tp
	}
}

func ParseBool(s string) (bool, error) {
	s = strings.ToUpper(s)
	switch s {
	case "是", "真", "yes", "YES", "1", "true", "TRUE", "True":
		return true, nil
	case "否", "假", "no", "NO", "0", "false", "FALSE", "False":
		return false, nil
	case "":
		return false, nil
	}

	return false, errors.New("invalid bool value")
}

// 值类型转换
func ConvertValue(vtype, value string) (error, interface{}) {
	var ret interface{}
	var err error
	if IsEnum(vtype) {
		_, ret, _ = ResolveEnumValue(vtype, value)
	} else {
		switch vtype {
		case "bool":
			ret, err = ParseBool(value)
		case "int", "int32":
			v, e := strconv.ParseInt(value, 10, 32)
			ret = int32(v)
			err = e
		case "uint", "uint32":
			v, e := strconv.ParseUint(value, 10, 32)
			ret = uint32(v)
			err = e
		case "int64":
			ret, err = strconv.ParseInt(value, 10, 64)
		case "uint64":
			ret, err = strconv.ParseUint(value, 10, 64)
		case "float":
			v, e := strconv.ParseFloat(value, 32)
			ret = float32(v)
			err = e
		case "double":
			ret, err = strconv.ParseFloat(value, 64)
		case "string":
			ret = value
		}
	}

	return err, ret
}

// 通过原始类型对值进行转换
func ParseValue(rawType, value string) (success bool, ret interface{}, isArray bool) {
	_, rawType, repeated, splitChar, _, _ := CompileValueType(rawType)

	var err error

	if !repeated {
		err, ret = ConvertValue(rawType, value)
		if err != nil {
			fmt.Printf("[错误] 值类型转换失败:%v [类型:%s 值：%v] \n", err, rawType, value)
			return false, value, repeated
		}
	} else {
		ret = make([]interface{}, 0)

		rstrs := Split(value, splitChar)

		for _, vstr := range rstrs {
			err, rvalue := ConvertValue(rawType, vstr)

			if err != nil {
				fmt.Printf("[错误] 数组类型转换失败:%v [类型:%s 值：%s 子项：%s] \n", err, rawType, value, vstr)
				return false, value, repeated
			}

			ret = append(ret.([]interface{}), rvalue)
		}
	}

	return true, ret, repeated
}

func TableType2PbType(pbType string, pbDesc *descriptorpb.FieldDescriptorProto, customType bool) {
	switch pbType {
	case "int", "int32":
		pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum()
	case "uint", "uint32":
		pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_UINT32.Enum()
	case "int64":
		pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum()
	case "uint64":
		pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_UINT64.Enum()
	case "float":
		pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_FLOAT.Enum()
	case "double":
		pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum()
	case "bool":
		pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum()
	case "string":
		pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	default:
		if IsEnum(pbType) {
			pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum()
			pbDesc.TypeName = proto.String(settings.PackageName + "." + pbType)
		} else if IsStruct(pbType) || customType {
			pbDesc.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
			pbDesc.TypeName = proto.String(settings.PackageName + "." + pbType)
		} else {
			panic("unknown pb type: " + pbType)
		}
	}
}

func Convert2PBValue(valueType string, value interface{}) (val pref.Value, err error) {
	switch valueType {
	case "int", "int32":
		val = pref.ValueOfInt32(value.(int32))
	case "uint", "uint32":
		val = pref.ValueOfUint32(value.(uint32))
	case "int64":
		val = pref.ValueOfInt64(value.(int64))
	case "uint64":
		val = pref.ValueOfUint64(value.(uint64))
	case "float":
		val = pref.ValueOfFloat32(value.(float32))
	case "double":
		val = pref.ValueOfFloat64(value.(float64))
	case "bool":
		val = pref.ValueOfBool(value.(bool))
	case "string":
		val = pref.ValueOfString(value.(string))
	default:
		if IsEnum(valueType) {
			val = pref.ValueOfEnum(pref.EnumNumber(value.(int32)))
		} else {
			err = errors.New("unknown pb type: " + valueType)
		}
	}

	return
}

// see: https://farer.org/2020/04/17/go-protobuf-apiv2-reflect-dynamicpb/
// from: https://github.com/davyxu/tabtoy/blob/b1843f83b7314c66816a493e37cfbf11b9cdacc4/v3/gen/pbdata/dynamictype.go#L13
func BuildDynamicType(tables []*model.DataTable) (protoreflect.FileDescriptor, error) {
	var file descriptorpb.FileDescriptorProto
	file.Syntax = proto.String("proto3")
	file.Name = proto.String(settings.PackageName + ".proto")
	file.Package = proto.String(settings.PackageName)

	// 创建公共定义类型，枚举与结构
	for enumName, item := range settings.DEFINES {
		if item.Category == model.DEFINE_TYPE_ENUM {
			var ed descriptorpb.EnumDescriptorProto
			ed.Name = proto.String(enumName)

			for _, field := range item.Items {
				var vd descriptorpb.EnumValueDescriptorProto
				vd.Name = proto.String(field.FieldName)
				v, _ := strconv.Atoi(field.Value)
				vd.Number = proto.Int32(int32(v))
				ed.Value = append(ed.Value, &vd)
			}
			file.EnumType = append(file.EnumType, &ed)
		} else if item.Category == model.DEFINE_TYPE_STRUCT {
			var desc descriptorpb.DescriptorProto
			desc.Name = proto.String(item.TypeName)
			var idx = 0
			for _, field := range item.Items {
				if field.IsVoid {
					continue
				}
				idx++

				var fd descriptorpb.FieldDescriptorProto
				fd.Name = proto.String(field.FieldName)
				fd.JsonName = proto.String(field.FieldName)
				fd.Number = proto.Int32(int32(idx))
				TableType2PbType(field.ValueType, &fd, false)
				if field.IsArray {
					fd.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
				} else {
					fd.Label = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
				}
				desc.Field = append(desc.Field, &fd)
			}
			file.MessageType = append(file.MessageType, &desc)
		}
	}

	PreProcessTable(tables)
	// 创建表数据结构
	for _, tab := range tables {
		var desc descriptorpb.DescriptorProto
		desc.Name = proto.String(tab.TypeName)
		for _, field := range tab.Headers {
			if field.IsVoid {
				continue
			}

			var fd descriptorpb.FieldDescriptorProto
			fd.Name = proto.String(field.FieldName)
			fd.JsonName = proto.String(field.FieldName)
			fd.Number = proto.Int32(int32(field.Index))
			TableType2PbType(field.ValueType, &fd, false)
			if field.IsArray {
				fd.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
			} else {
				fd.Label = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
			}

			desc.Field = append(desc.Field, &fd)
		}
		file.MessageType = append(file.MessageType, &desc)

		if tab.IsDataTable {
			// 创建列表结构
			var itemsDesc descriptorpb.DescriptorProto
			itemsDesc.Name = proto.String(tab.TypeName + "_ARRAY")
			var itemsFD descriptorpb.FieldDescriptorProto
			itemsFD.Name = proto.String("Items")
			itemsFD.JsonName = proto.String("Items")
			itemsFD.Number = proto.Int32(int32(1))
			TableType2PbType(tab.TypeName, &itemsFD, true)
			itemsFD.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
			itemsDesc.Field = append(itemsDesc.Field, &itemsFD)

			file.MessageType = append(file.MessageType, &itemsDesc)
		}
	}

	return protodesc.NewFile(&file, nil)
}

// 生成proto文件描述，文件名可以任意设置
func BuildFileDesc(filename string, includeLanguage bool) (protoreflect.FileDescriptor, error) {
	var file descriptorpb.FileDescriptorProto
	file.Syntax = proto.String("proto3")
	file.Name = proto.String(filename + ".proto")
	file.Package = proto.String(settings.PackageName)

	// 创建公共定义类型，枚举与结构
	for _, item := range settings.ENUMS {
		if item.Category == model.DEFINE_TYPE_ENUM {
			var ed descriptorpb.EnumDescriptorProto
			ed.Name = proto.String(item.TypeName)

			for _, field := range item.Items {
				var vd descriptorpb.EnumValueDescriptorProto
				vd.Name = proto.String(field.FieldName)
				v, _ := strconv.Atoi(field.Value)
				vd.Number = proto.Int32(int32(v))
				ed.Value = append(ed.Value, &vd)
			}
			file.EnumType = append(file.EnumType, &ed)
		}
	}

	// 创建表数据结构
	var tables = settings.GetAllTables()
	PreProcessTable(tables)

	for _, tab := range tables {
		// 当不生成语言类型时，过滤语言类型
		if tab.IsLanguage && !includeLanguage {
			continue
		}

		var desc descriptorpb.DescriptorProto
		desc.Name = proto.String(tab.TypeName)
		for _, field := range tab.Headers {
			if field.IsVoid {
				continue
			}

			var fd descriptorpb.FieldDescriptorProto
			fd.Name = proto.String(field.FieldName)
			fd.JsonName = proto.String(field.FieldName)
			fd.Number = proto.Int32(int32(field.Index))
			TableType2PbType(field.ValueType, &fd, false)
			if field.IsArray {
				fd.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
			} else {
				fd.Label = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
			}

			desc.Field = append(desc.Field, &fd)
		}
		file.MessageType = append(file.MessageType, &desc)

		if tab.IsDataTable {
			// 创建列表结构
			var itemsDesc descriptorpb.DescriptorProto
			itemsDesc.Name = proto.String(tab.TypeName + "_ARRAY")
			var itemsFD descriptorpb.FieldDescriptorProto
			itemsFD.Name = proto.String("Items")
			itemsFD.JsonName = proto.String("Items")
			itemsFD.Number = proto.Int32(int32(1))
			TableType2PbType(tab.TypeName, &itemsFD, true)
			itemsFD.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
			itemsDesc.Field = append(itemsDesc.Field, &itemsFD)

			file.MessageType = append(file.MessageType, &itemsDesc)
		}
	}

	return protodesc.NewFile(&file, nil)
}
