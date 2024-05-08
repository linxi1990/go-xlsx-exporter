// DO NOT EDIT!
// This code is auto generated by go-xlsx-exporter
// VERSION {{.Version}}
// go-protobuf {{.GoProtoVersion}}

{{- $G := .}}
{{- $NS := .Namespace}}

{{- range .Info.Imports}}
{{.}}
{{- end}}

import $protobuf from "protobufjs";

{{- $ENABLE_LONG:=false }}
{{- $ENABLE_READONLEY:=false }}

{{- $TRIM_ENCODE:=false }}
{{- $TRIM_DECODE:=false }}
{{- $TRIM_FROM_OBJECT:=true }}
{{- $TRIM_TO_OBJECT:=true }}
{{- $TRIM_VERIFY:=true }}
{{- $TRIM_TO_JSON:=true }}
{{- $TRIM_DELIMITED:=true }}

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

declare module $utils {
    interface Long {

    }
}

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {} as any);

export interface Long {
    /** Low bits */
    low: number;

    /** High bits */
    high: number;

    /** Whether unsigned or not */
    unsigned: boolean;
}

export class DataConverter {
    static convertHandler: (data: DataModel | object, fieldName:string, value: string, alias?: string)=>any;      
    static getConvertData(target: any, fieldName: string, value: any, alias?: string, cachable?: boolean) {
        target._converted = target._converted || {};
        if(target._converted[fieldName])
        {
            return target._converted[fieldName];
        }

        if(DataConverter.convertHandler == null)
        {
            throw `convert field ${fieldName} value need a convetor`;
        }

        var data = DataConverter.convertHandler(target, fieldName, value, alias);
        if(cachable) {
            target._converted[fieldName] = data;
        }
        return data;
    }
}

export class DataModel {
    private _converted = {};

    protected getConvertData(fieldName: string, value: any, alias?: string, cachable?: boolean): any
    {
        if(this._converted[fieldName])
        {
            return this._converted[fieldName];
        }

        if(DataConverter.convertHandler == null)
        {
            throw `convert field ${fieldName} value need a convetor`;
        }

        var data = DataConverter.convertHandler(this, fieldName, value, alias);
        if(cachable) {
            this._converted[fieldName] = data;
        }
        return data;
    } 
}

export namespace {{$NS}} {
    var ALLTYPES: {[key: string]: any} = {};

    {{/*生成枚举类型*/}}
    {{- range .Enums}}
    // Defined in table: {{.DefinedTable}}
    export enum {{.TypeName}} {
        {{- range .Items}}
        {{if ne .Desc ""}} /** {{.Desc}} */ {{end}}
        {{.FieldName}} = {{.Value}},
        {{end}}
    }
    {{end}}{{/* enum */}}

    {{- /*生成配置类类型*/}}
    {{- range .Consts}}
    {{$TypeName := .TypeName}}

    // Defined in table: {{.DefinedTable}}
    export const {{.TypeName}}: {
        {{- range .Items}}
        {{if ne .Desc ""}} /** {{.Desc}} */{{end}}  
            {{- if not .IsVoid }}                     
        {{.FieldName}}?: Readonly<{{type_format .StandardValueType .ValueType .IsArray}}>,
            {{end}}
            {{- if .Convertable}}
        get{{upperF .FieldName}}(): Readonly<{{get_alias .Alias}}>,
            {{- end}}
        {{end}} {{/*end .Items */}}
    } = {
        {{- range .Items}}
            {{- if not .IsVoid }}
        {{.FieldName}} : {{value_format .Value .}},
            {{end}}
            {{- if .Convertable}}
        get{{upperF .FieldName}}: function() {
            return DataConverter.getConvertData({{$TypeName}}, '{{.FieldName}}', {{$TypeName}}.{{.FieldName}}, '{{get_alias .Alias}}', {{.Cachable}});
        },
            {{- end}}
        {{end}} {{/*end .Items */}}
    }
    {{end}}{{/*end .Consts */}}

    {{- /*生成类类型*/}}
    {{- range .Tables}}

    {{$TypeName := .TypeName}}
    // Defined in table: {{.DefinedTable}}
    /** Properties of a {{$TypeName}}. */
    export interface I{{$TypeName}} {
        {{range .Headers}}
        {{if ne .Desc ""}} /** {{.Desc}} */{{end}}
            {{- if not .IsVoid }}             
                {{- if $ENABLE_READONLEY}}                  
        {{.FieldName}}?: Readonly<{{type_format .StandardValueType .ValueType .IsArray}}>;
                {{- else}}
        {{.FieldName}}?: {{type_format .StandardValueType .ValueType .IsArray}};
                {{- end}}
            {{end}} {{/*end not Void*/}}
            {{- if .Convertable}}
        get{{upperF .FieldName}}(): Readonly<{{get_alias .Alias}}>;
            {{- end}}
        {{end}} {{/*end .Headers */}}
    }

     /** Represents a {{$TypeName}}. */
    export class {{$TypeName}} extends DataModel implements I{{$TypeName}} { 
        private static __id__ = {{.Id}};
        private static __type_name__ = "{{$TypeName}}";
        {{- if not .IsArray}}
        private static __array_type_name__ = "{{$TypeName}}_ARRAY";
        static getArrayType(): any {
            return ALLTYPES[{{$TypeName}}.__array_type_name__];
        }
        {{- end}}

        {{range .Headers}}
            {{- if ne .Desc ""}} /** {{.Desc}} */{{end}}
            {{- if not .IsVoid }} 
                {{- if $ENABLE_READONLEY}}   
                    {{- if .IsArray}}
            {{.FieldName}}?: Readonly<{{type_format .StandardValueType .ValueType .IsArray}}> = $util.emptyArray;
                    {{- else if eq .StandardValueType "bytes"}}
            {{.FieldName}}?: Readonly<{{type_format .StandardValueType .ValueType .IsArray}}> = $util.newBuffer([]);
                    {{else}}
            {{.FieldName}}?: Readonly<{{type_format .StandardValueType .ValueType .IsArray}}> = {{default .}};
                    {{end -}}   
                {{- else}}
                    {{- if .IsArray}}
                {{.FieldName}}?: {{type_format .StandardValueType .ValueType .IsArray}} = $util.emptyArray;
                    {{- else if eq .StandardValueType "bytes"}}
                {{.FieldName}}?: {{type_format .StandardValueType .ValueType .IsArray}} = $util.newBuffer([]);
                    {{else}}
                {{.FieldName}}?: {{type_format .StandardValueType .ValueType .IsArray}} = {{default .}};
                    {{end -}}  
                {{- end}}
            {{end -}} 
            {{- if .Convertable}}    
        get{{upperF .FieldName}}(): Readonly<{{get_alias .Alias}}> {
            return this.getConvertData("{{.FieldName}}", {{if .IsVoid}}null{{else}}this.{{.FieldName}}{{end}}, '{{get_alias .Alias}}', {{.Cachable}});
        };
            {{end}}
        {{end}}

        /**
         * Constructs a new {{$TypeName}}.
         * @param [properties] Properties to set
         */
        constructor(properties?: I{{$TypeName}}) {
            super();
            
            {{range .Headers}}
        {{- if not .IsVoid }}
            {{- if .IsArray}}
            this.{{.FieldName}} = [];
            {{end -}}
        {{end -}}     
        {{end}}        
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }{{/*end function */}}

        static create(properties?: {{$TypeName}}): {{$TypeName}} {
            return new {{$TypeName}}(properties);
        }

        {{- if not $TRIM_ENCODE }}
        static encode(message: I{{$TypeName}}, writer?: $protobuf.Writer): $protobuf.Writer {
            if (!writer)
                writer = $Writer.create();
                
            {{range .Headers}} 
                {{- $wireType := get_wire_type .}}
                {{- $count := calc_wire_offset .}}           
                
            {{if ne .Desc ""}} /** {{.Desc}} */{{end}}  
                {{- if not .IsVoid }}
                    {{- if .IsArray}}                       
                            {{- if is_struct .ValueType}}
            if (message.{{.FieldName}} != null && message.{{.FieldName}}.length)
                for (var i = 0; i < message.{{.FieldName}}.length; ++i)
                    {{.ValueType}}.encode(message.{{.FieldName}}[i], writer.uint32(/* id {{.Index}}, wireType {{$wireType}} =*/{{$count}}).fork()).ldelim();
                            {{- else}}
            if (message.{{.FieldName}} != null && message.{{.FieldName}}.length) {
                {{- if out .StandardValueType "string" "bytes"}}
                writer.uint32(/* id {{.Index}}, wireType {{$wireType}} =*/{{$count}}).fork();
                {{- end}}
                for (var i = 0; i < message.{{.FieldName}}.length; ++i)
                    {{- if eq "string" .StandardValueType}}
                    writer.uint32(/* id {{.Index}}, wireType {{$wireType}} =*/{{$count}}).string(message.{{.FieldName}}[i]);
                    {{- else if eq "bytes" .StandardValueType}}
                    writer.uint32(/* id {{.Index}}, wireType {{$wireType}} =*/{{$count}}).bytes(message.{{.FieldName}}[i]);
                    {{- else}}
                    {{- $pbType := get_pb_type .StandardValueType}}
                    writer.{{$pbType}}(message.{{.FieldName}}[i]);
                    {{- end}}
                {{- if out .StandardValueType "string" "bytes"}}
                writer.ldelim();
                {{- end}}
            }
                            {{- end}}{{/*end message*/}}
                    {{- else}}
            if (message.{{.FieldName}} != null && Object.hasOwnProperty.call(message, "{{.FieldName}}"))
                            {{- if is_struct .ValueType}}
                {{.ValueType}}.encode(message.{{.FieldName}}, writer.uint32(/* id {{.Index}}, wireType {{$wireType}} =*/{{$count}}).fork()).ldelim();
                            {{- else}}
                    {{- $pbType := get_pb_type .StandardValueType}}
                writer.uint32(/* id {{.Index}}, wireType {{$wireType}} =*/{{$count}}).{{$pbType}}(message.{{.FieldName}});
                            {{- end}}{{/*end message*/}}
                    {{- end -}} {{/*end if*/}}   
                {{- end}} 
            {{end -}} {{/*end .Headers*/}}

            return writer;
        }

        {{- if not $TRIM_DELIMITED }}
        static encodeDelimited(message: I{{$TypeName}}, writer?: $protobuf.Writer): $protobuf.Writer {
            return this.encode(message, writer).ldelim();
        }
        {{- end}} {{/** end if trim delimited */}}

        {{- end}} {{/** end if trim encode */}}

        {{- if not $TRIM_DECODE }}
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): {{$TypeName}} {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length;
            var message: any = new {{$TypeName}}();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                {{- range .Headers}}  
                {{- if not .IsVoid}}        
                case {{.Index}}:
                    {{- if .IsArray}}                    
                    if (!(message.{{.FieldName}} && message.{{.FieldName}}.length))
                        message.{{.FieldName}} = [];

                        {{- if is_struct .ValueType}}
                    message.{{.FieldName}}.push({{.ValueType}}.decode(reader, reader.uint32()));                    
                        {{- else}}
                            {{- if out .ValueType "string" "bytes"}}
                    if ((tag & 7) === 2) {
                        var end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2)
                            message.{{.FieldName}}.push(reader.{{.PBValueType}}());
                    } else
                        message.{{.FieldName}}.push(reader.{{.PBValueType}}());
                            {{- else}}
                    message.{{.FieldName}}.push(reader.{{.PBValueType}}());
                            {{- end}} {{/*end if string */}}
                        {{- end}} {{/*end if message*/}}
                    {{- else}}
                        {{- if is_struct .ValueType}}
                    message.{{.FieldName}} = {{.ValueType}}.decode(reader, reader.uint32());
                        {{- else}}
                    message.{{.FieldName}} = reader.{{.PBValueType}}();
                        {{- end}} {{/*end message*/}}
                    {{- end}} 
                    break;
                {{- end}} {{/**end .IsVoid */}}
                {{- end}} {{/**end case */}}
                default:
                    reader.skipType(tag & 7);
                    break;
                } 
            }
            return message;
        }

        {{- if not $TRIM_DELIMITED }}
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): {{$TypeName}} {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        }        
        {{- end}} {{/** end if trim delimited */}}

        {{- end }} {{/** end if trim decode */}}

        {{- if not $TRIM_VERIFY }}
        static verify(message: { [k: string]: any }): (string|null) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            {{- range .Headers}}   
            {{- if not .IsVoid}}     
                if (message.{{.FieldName}} != null && message.hasOwnProperty("{{.FieldName}}")) {
                {{- if .IsArray}}
                    if (!Array.isArray(message.{{.FieldName}}))
                        return "{{.FieldName}}: array expected";
                    {{- if .IsEnum}}                
                    {{- $enums := get_enum_values .StandardValueType}}
                    for (var i = 0; i < message.{{.FieldName}}.length; ++i)
                        switch (message.{{.FieldName}}[i]) {
                            default:
                                return "{{.FieldName}}: enum value[] expected";
                            {{- range $enums}}
                            case {{.}}:
                            {{- end}}
                                break;
                        }
                    {{- else if is_interger .StandardValueType}}
                    for (var i = 0; i < message.{{.FieldName}}.length; ++i)
                        if (!$util.isInteger(message.{{.FieldName}}[i]))
                            return "{{.FieldName}}: integer[] expected";
                    {{- else if is_number .StandardValueType}}
                    for (var i = 0; i < message.{{.FieldName}}.length; ++i)
                        if (typeof message.{{.FieldName}}[i] !== "number")
                            return "{{.FieldName}}: number[] expected";
                    {{- else if is_bool .StandardValueType}}
                    for (var i = 0; i < message.{{.FieldName}}.length; ++i)
                        if (typeof message.{{.FieldName}}[i] !== "boolean")
                            return "{{.FieldName}}: boolean[] expected";
                    {{- else if is_string .StandardValueType}}
                    for (var i = 0; i < message.{{.FieldName}}.length; ++i)
                        if (!$util.isString(message.{{.FieldName}}[i]))
                            return "{{.FieldName}}: string[] expected";  
                    {{- else if is_bytes .StandardValueType}}
                    for (var i = 0; i < message.{{.FieldName}}.length; ++i)
                        if (!(message.{{.FieldName}}[i] && typeof message.{{.FieldName}}[i].length === "number" || $util.isString(message.{{.FieldName}}[i])))
                            return "{{.FieldName}}: buffer[] expected";  
                    {{- else if is_struct .ValueType}}     
                    for (var i = 0; i < message.{{.FieldName}}.length; ++i) {
                        var error = {{.ValueType}}.verify(message.{{.FieldName}}[i]);
                        if (error)
                            return "{{.FieldName}}." + error;
                    }       
                    {{- else}}
                        "error type {{.ValueType}} {{.FieldName}}";
                    {{- end}}
                {{- else}}
                    {{- if .IsEnum}}
                    {{- $enums := get_enum_values .StandardValueType}}
                    switch (message.{{.FieldName}}) {
                        default:
                            return "{{.FieldName}}: enum value expected";
                        {{- range $enums}}
                        case {{.}}:
                        {{- end}}
                            break;
                    }
                    {{- else if is_interger .StandardValueType}}
                    if (!$util.isInteger(message.{{.FieldName}}))
                        return "{{.FieldName}}: integer expected";
                    {{- else if is_number .StandardValueType}}
                    if (typeof message.{{.FieldName}} !== "number")
                        return "{{.FieldName}}: number expected";
                    {{- else if is_bool .StandardValueType}}
                    if (typeof message.{{.FieldName}} !== "boolean")
                        return "{{.FieldName}}: boolean expected";
                    {{- else if is_string .StandardValueType}}
                    if (!$util.isString(message.{{.FieldName}}))
                        return "{{.FieldName}}: string expected";
                    {{- else if is_bytes .StandardValueType}} 
                    if (!(message.{{.FieldName}} && typeof message.{{.FieldName}}.length === "number" || $util.isString(message.{{.FieldName}})))
                        return "{{.FieldName}}: buffer expected";
                    {{- else if is_struct .ValueType}}
                        if (message.{{.FieldName}} != null && message.hasOwnProperty("{{.FieldName}}")) {
                        var error = {{.ValueType}}.verify(message.{{.FieldName}});
                        if (error)
                            return "{{.FieldName}}." + error;
                        }
                    {{- else}}
                        "error type {{.ValueType}} {{.FieldName}}";
                    {{- end}}
                {{- end}} {{/* end if IsArray */}}
                }
            {{- end}} {{/* end if not .IsVoid */}}
            {{- end}} {{/**end range Headers */}}
            return null;
        }
        {{- end }} {{/** end if trim verify */}}

        {{- if not $TRIM_FROM_OBJECT }}
        static fromObject(object: { [k: string]: any }): {{$TypeName}} {
            if (object instanceof {{$TypeName}})
                return object;
            var message: any = new {{$TypeName}}();
            {{- range .Headers}}  
            {{- if not .IsVoid}}
                {{- if .IsArray}}
            if(object.{{.FieldName}}) {
                if (!Array.isArray(object.{{.FieldName}}))
                    throw TypeError("{{$TypeName}}.{{.FieldName}}: array expected");            
                message.{{.FieldName}} = [];
                for (var i = 0; i < object.{{.FieldName}}.length; ++i) {
                    {{- $Ofn := .FieldName}}
                    {{- if .IsEnum}} 
                        {{- $enum := get_enum .StandardValueType}}
                    switch (object.{{.FieldName}}[i]) {
                        default:
                            {{- range $enum.Items}}
                        case "{{.FieldName}}":
                        case {{.Value}}:
                            message.{{$Ofn}}[i] = {{.Value}};
                            break;
                        {{- end}}    
                    }                
                    {{- else if and $ENABLE_LONG (is_long .StandardValueType)}}
                    if ($util.Long)
                        (message.{{.FieldName}}[i] = $util.Long.fromValue(object.{{.FieldName}})).unsigned = {{if eq .StandardValueType "uint64"}}true{{else}}false{{end}};
                    else if (typeof object.{{.FieldName}} === "string")
                        message.{{.FieldName}}[i] = parseInt(object.{{.FieldName}}, 10);
                    else if (typeof object.{{.FieldName}} === "number")
                        message.{{.FieldName}}[i] = object.{{.FieldName}};
                    else if (typeof object.{{.FieldName}} === "object")
                        message.{{.FieldName}}[i] = new $util.LongBits(object.{{.FieldName}}[i].low >>> 0, object.{{.FieldName}}[i].high >>> 0).toNumber();
                    {{- else if eq .StandardValueType "int"}}
                    message.{{.FieldName}}[i] = object.{{.FieldName}}[i] | 0;
                    {{- else if eq .StandardValueType "uint"}}
                    message.{{.FieldName}}[i] = object.{{.FieldName}}[i] >>> 0;
                    {{- else if eq .StandardValueType "bool"}}
                    message.{{.FieldName}}[i] = Boolean(object.{{.FieldName}}[i]);
                    {{- else if is_number .StandardValueType}}
                    message.{{.FieldName}}[i] = Number(object.{{.FieldName}}[i]);                
                    {{- else if eq .StandardValueType "string"}}
                    message.{{.FieldName}}[i] = String(object.{{.FieldName}}[i]);
                    {{- else if eq .StandardValueType "bytes"}}
                    if (typeof object.{{.FieldName}}[i] === "string")
                        $util.base64.decode(object.{{.FieldName}}[i], message.{{.FieldName}}[i] = $util.newBuffer($util.base64.length(object.{{.FieldName}}[i])), 0);
                    else if (object.{{.FieldName}}[i].length)
                        message.{{.FieldName}}[i] = object.{{.FieldName}}[i];
                    {{- else if is_struct .ValueType}}
                    if (typeof object.{{.FieldName}}[i] !== "object")
                        throw TypeError("{{$TypeName}}.{{.FieldName}}: object expected");
                    message.{{.FieldName}}[i] = {{$TypeName}}.fromObject(object.{{.FieldName}}[i]);
                    {{- else}}
                        throw TypeError("{{$TypeName}}.{{.FieldName}}: object expected");
                    {{- end}}
                }
            }
                {{- else}}{{/**end if Array */}}    
                    {{- if not .IsEnum}}                
            if (object.{{.FieldName}} != null)
                    {{- end}}                
                    {{- if .IsEnum}} 
                    {{- $Ofn := .FieldName}}
                    {{- $enum := get_enum .StandardValueType}}
            switch (object.{{.FieldName}}) {
                default:
                    {{- range $enum.Items}}
                case "{{.FieldName}}":
                case {{.Value}}:
                    message.{{$Ofn}} = {{.Value}};
                    break;
                    {{- end}}    
                }               
                    {{- else if and $ENABLE_LONG (is_long .StandardValueType)}}
                if ($util.Long)
                    (message.{{.FieldName}} = $util.Long.fromValue(object.{{.FieldName}})).unsigned = {{if eq .StandardValueType "uint64"}}true{{else}}false{{end}};
                else if (typeof object.{{.FieldName}} === "string")
                    message.{{.FieldName}} = parseInt(object.{{.FieldName}}, 10);
                else if (typeof object.{{.FieldName}} === "number")
                    message.{{.FieldName}} = object.{{.FieldName}};
                else if (typeof object.{{.FieldName}} === "object")
                    message.{{.FieldName}} = new $util.LongBits(object.{{.FieldName}}.low >>> 0, object.{{.FieldName}}.high >>> 0).toNumber();
                    {{- else if eq .StandardValueType "int"}}
                message.{{.FieldName}} = object.{{.FieldName}} | 0;
                    {{- else if eq .StandardValueType "uint"}}
                message.{{.FieldName}} = object.{{.FieldName}} >>> 0;
                    {{- else if eq .StandardValueType "bool"}}
                message.{{.FieldName}} = Boolean(object.{{.FieldName}});
                    {{- else if is_number .StandardValueType}}
                message.{{.FieldName}} = Number(object.{{.FieldName}});
                    {{- else if eq .StandardValueType "string"}}
                message.{{.FieldName}} = String(object.{{.FieldName}});
                    {{- else if eq .StandardValueType "bytes"}}
                if (typeof object.{{.FieldName}} === "string")
                    $util.base64.decode(object.{{.FieldName}}, message.{{.FieldName}} = $util.newBuffer($util.base64.length(object.{{.FieldName}})), 0);
                else if (object.{{.FieldName}}.length)
                    message.{{.FieldName}} = object.{{.FieldName}};
                    {{- else if is_struct .ValueType}}
                if (typeof object.{{.FieldName}} !== "object")
                    throw TypeError("{{$TypeName}}.{{.FieldName}}: object expected");
                message.{{.FieldName}} = {{$TypeName}}.fromObject(object.{{.FieldName}});
                    {{- end}}
                {{- end}}
            {{- end}} {{/* end if not .IsVoid */}}
            {{- end}} {{/* end range Headers */}}
            return message;
        }
        {{- end }} {{/** end if trim from object */}}

        {{- if not $TRIM_TO_OBJECT }}
        static toObject(message: {{$TypeName}}, options?: $protobuf.IConversionOptions): { [k: string]: any } {
            if (!options)
                options = {};
            var object: any = {};
            if (options.arrays || options.defaults) {
            {{- range .Headers}}  
                {{- if not .IsVoid}}
                    {{- if .IsArray}}
                object.{{.FieldName}} = [];
                    {{- end}}
                {{- end}} {{/* end if not .IsVoid */}}
            {{- end}}
            }

            if (options.defaults) {
            {{- range .Headers}}  
            {{- if not .IsVoid}}
                {{- if and $ENABLE_LONG (is_long .StandardValueType)}}
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.{{.FieldName}} = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.{{.FieldName}} = options.longs === String ? "0" : 0;
                {{- else if .IsEnum}}
                {{$enumDefault := get_enum_default .StandardValueType}}
                object.{{.FieldName}} = options.enums === String ? "{{$enumDefault.FieldName}}" : {{$enumDefault.Value}};
                {{- else}}
                object.{{.FieldName}} = {{default .}};
                {{- end}}
            {{- end}} {{/* end if not .IsVoid */}}
            {{- end}} {{/* end range Headers */}}
            }

            {{- range .Headers}}  
            {{- if not .IsVoid}}
                {{- if .IsArray}}
                
            if (message.{{.FieldName}} && message.{{.FieldName}}.length) {
                object.{{.FieldName}} = [];
                for (var j = 0; j < message.{{.FieldName}}.length; ++j)
                    {{- if is_number .StandardValueType}}
                    object.{{.FieldName}}[j] = options.json && !isFinite(message.{{.FieldName}}[j]) ? String(message.{{.FieldName}}[j]) : message.{{.FieldName}}[j];
                    {{- else if and $ENABLE_LONG (is_long .StandardValueType)}}
                    if (typeof message.{{.FieldName}}[j] === "number")
                        object.{{.FieldName}}[j] = options.longs === String ? String(message.{{.FieldName}}[j]) : message.{{.FieldName}}[j];
                    else
                        object.{{.FieldName}}[j] = options.longs === String ? $util.Long.prototype.toString.call(message.{{.FieldName}}[j]) : options.longs === Number ? new $util.LongBits(message.{{.FieldName}}[j].low >>> 0, message.{{.FieldName}}[j].high >>> 0).toNumber() : message.{{.FieldName}}[j];
                    {{- else if .IsEnum}}
                    object.{{.FieldName}}[j] = options.enums === String ? {{.ValueType}}[message.{{.FieldName}}[j]] : message.{{.FieldName}}[j];
                    {{- else if eq .StandardValueType "bytes"}}
                    object.{{.FieldName}}[j] = options.bytes === String ? $util.base64.encode(message.{{.FieldName}}[j], 0, message.{{.FieldName}}[j].length) : options.bytes === Array ? Array.prototype.slice.call(message.{{.FieldName}}[j]) : message.{{.FieldName}}[j];
                    {{- else if  is_struct .ValueType}}
                    object.{{.FieldName}}[j] = {{.ValueType}}.toObject(message.{{.FieldName}}[j] as any, options);
                    {{- else}}
                    object.{{.FieldName}}[j] = message.{{.FieldName}}[j];
                    {{- end}}
            }
                {{- else}}
            if (message.{{.FieldName}} != null && message.hasOwnProperty("{{.FieldName}}"))
                    {{- if is_number .StandardValueType}}
                object.{{.FieldName}} = options.json && !isFinite(message.{{.FieldName}}) ? String(message.{{.FieldName}}) : message.{{.FieldName}};
                    {{- else if and $ENABLE_LONG (is_long .StandardValueType)}}
                if (typeof message.{{.FieldName}} === "number")
                    object.{{.FieldName}} = options.longs === String ? String(message.{{.FieldName}}) : message.{{.FieldName}};
                else
                    object.{{.FieldName}} = options.longs === String ? $util.Long.prototype.toString.call(message.{{.FieldName}}) : options.longs === Number ? new $util.LongBits(message.{{.FieldName}}.low >>> 0, message.{{.FieldName}}.high >>> 0).toNumber() : message.{{.FieldName}};
                    {{- else if .IsEnum}}
                    object.{{.FieldName}} = options.enums === String ? {{.ValueType}}[message.{{.FieldName}}] : message.{{.FieldName}};
                    {{- else if eq .StandardValueType "bytes"}}
                object.{{.FieldName}} = options.bytes === String ? $util.base64.encode(message.{{.FieldName}}, 0, message.{{.FieldName}}.length) : options.bytes === Array ? Array.prototype.slice.call(message.{{.FieldName}}) : message.{{.FieldName}};
                    {{- else if is_struct .ValueType}}
                object.{{.FieldName}} = {{.ValueType}}.toObject(message.{{.FieldName}} as any, options);
                    {{- else}}
                object.{{.FieldName}} = message.{{.FieldName}};
                    {{- end}}
                {{- end}}
            {{- end}} {{/** end not .IsVoid */}}
            {{- end}}{{/*end range headers */}}
            return object;
        }{{/*end toObject function*/}}        
        {{- end }} {{/** end if trim to object */}}

        {{- if not $TRIM_TO_JSON }}
        toJSON(): { [k: string]: any } {
            return {{$TypeName}}.toObject(this, $protobuf.util.toJSONOptions);
        }
        {{- end }} {{/** end if trim to json*/}}
    } {{/*end class */}}
    ALLTYPES["{{$TypeName}}"] = {{$TypeName}};
    
        {{- end}} {{/*end tables */}}

    {{- if .HasMessage}}
    // regist all messages
    export enum EMessageNames {
        {{- range .Tables}}
            {{- if is_message_table .TableType}}
                {{- if gt .Id 0}}
        E_MSG_{{.TypeName}} = {{.Id}},
                {{- end}}
            {{- end}}
        {{- end}} 
    }

    export class MessageFactory
    {    
        public static Types:{[id: number]: any} = {
        {{- range .Tables}}
            {{- if is_message_table .TableType}}
                {{- if gt .Id 0}}
            {{.Id}}: {{.TypeName}},
                {{- end}}
            {{- end}}
        {{- end}}  
        };

        public static CreateMessage<T extends DataModel>(id: number): T
        {
            if(!this.Types[id]) 
            {                
                return null as any;
            }

            return new this.Types[id]() as T;
        }

        public static LoadMessage<T extends DataModel>(id: number, bytes: Uint8Array): T
        {
            if(!this.Types[id]) 
            {                
                return null as any;
            }

            const cls = this.Types[id];
            try {
                let msg = cls.decode(bytes) as T;
                return msg;
            }catch(e) {
                console.error(e, "decode message error, id:" + id + " bytes:" + bytes.length);
                return null as any;
            }
        }
    }
    {{- end}}
}{{/*end namespace */}}