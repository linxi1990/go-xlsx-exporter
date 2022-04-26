package generator

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/wingcd/go-xlsx-exporter/model"
	"github.com/wingcd/go-xlsx-exporter/settings"
	"github.com/wingcd/go-xlsx-exporter/utils"
)

var dtsTemplate = ""

var dtsGenetatorInited = false

func egistDTSFuncs() {
	if dtsGenetatorInited {
		return
	}
	dtsGenetatorInited = true

	funcs["value_format"] = jsValueFormat

	funcs["type_format"] = tsTypeFormat

	funcs["default"] = jsValueDefault

	funcs["get_alias"] = func(alias string) string {
		if alias == "" {
			return "any"
		}
		return alias
	}
}

var supportDTSTypes = map[string]string{
	"bool":   "boolean",
	"int":    "number",
	"uint":   "number",
	"int64":  "number",
	"uint64": "number",
	"float":  "number",
	"double": "number",
	"string": "string",
	"bytes":  "Uint8Array",
	"void":   "any",
}

var sdefaultDTSValue = map[string]string{
	"bool":   "false",
	"int":    "0",
	"uint":   "0",
	"int64":  "0",
	"uint64": "0",
	"float":  "0",
	"double": "0",
	"string": "\"\"",
	"bytes":  "new ArrayBuffer(0)",
	"void":   "null",
}

type dtsFileDesc struct {
	commonFileDesc

	Namespace string
	Info      *BuildInfo
	Enums     []*model.DefineTableInfo
	Consts    []*model.DefineTableInfo
	Tables    []*model.DataTable
}

type dtsGenerator struct {
}

func (g *dtsGenerator) Generate(info *BuildInfo) (save bool, data *bytes.Buffer) {
	egistDTSFuncs()

	if dtsTemplate == "" {
		data, err := ioutil.ReadFile("./template/dts.gtpl")
		if err != nil {
			log.Println(err)
			return false, nil
		}
		dtsTemplate = string(data)
	}

	tpl, err := template.New("dts").Funcs(funcs).Parse(dtsTemplate)
	if err != nil {
		log.Println(err.Error())
		return false, nil
	}

	var fd = dtsFileDesc{
		Namespace: settings.PackageName,
		Info:      info,
		Enums:     settings.ENUMS[:],
		Consts:    settings.CONSTS[:],
		Tables:    make([]*model.DataTable, 0),
	}
	fd.Version = settings.TOOL_VERSION
	fd.GoProtoVersion = settings.GO_PROTO_VERTION

	utils.PreProcessDefine(fd.Consts)
	for _, c := range fd.Consts {
		for _, it := range c.Items {
			if !it.IsEnum && !it.IsStruct {
				it.ValueType = supportDTSTypes[it.StandardValueType]
			}
		}
	}

	tables := settings.GetAllTables()
	utils.PreProcessTable(tables)
	for _, t := range tables {
		if t.TableType == model.ETableType_Message {
			fd.HasMessage = true
		}

		// 排除语言类型
		if t.TableType == model.ETableType_Language && !settings.GenLanguageType {
			continue
		}

		// 排除配置
		if t.TableType == model.ETableType_Define {
			continue
		}

		fd.Tables = append(fd.Tables, t)

		// 处理类型
		for _, h := range t.Headers {
			if !h.IsEnum && !h.IsStruct {
				if _, ok := supportDTSTypes[h.StandardValueType]; !ok {
					log.Printf("[错误] 不支持类型%s 表：%s 列：%s \n", h.RawValueType, t.DefinedTable, h.FieldName)
					return false, nil
				}
				h.ValueType = supportDTSTypes[h.StandardValueType]
			}
		}

		if t.NeedAddItems {
			// 添加数组类型
			table := model.DataTable{}
			table.DefinedTable = t.DefinedTable
			table.TypeName = t.TypeName + "_ARRAY"
			table.IsArray = true
			header := model.DataTableHeader{}
			header.Index = 1
			header.FieldName = "Items"
			header.TitleFieldName = header.FieldName
			header.IsArray = true
			header.ValueType = t.TypeName
			header.RawValueType = t.TypeName + "[]"
			header.IsMessage = true
			table.Headers = []*model.DataTableHeader{&header}

			fd.Tables = append(fd.Tables, &table)
		}
	}
	utils.PreProcessTable(fd.Tables)

	var buf = bytes.NewBufferString("")

	err = tpl.Execute(buf, &fd)
	if err != nil {
		log.Println(err)
		return false, nil
	}

	return true, buf
}

func init() {
	Regist("dts", &dtsGenerator{})
}
