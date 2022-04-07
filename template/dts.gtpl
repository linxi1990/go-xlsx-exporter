// DO NOT EDIT!
// This code is auto generated by go-xlsx-exporter
// VERSION {{.Version}}
// go-$protobuf {{.GoProtoVersion}}

{{- $G := .}}
{{- $NS := .Namespace}}

{{- range .Info.Imports}}
{{.}}
{{- end}}

export interface Long {
    /** Low bits */
    low: number;

    /** High bits */
    high: number;

    /** Whether unsigned or not */
    unsigned: boolean;
}

export class DataConverter {
    static convertHandler: (typeName: string, fieldName:string, value: string)=>any;

    static convertData(typeName: string, fieldName:string, value: string): any;
}

/** Namespace  */
export namespace {{$NS}} {
    {{/*生成枚举类型*/}}
    {{- range .Enums}}
    // Defined in table: {{.DefinedTable}}
    enum {{.TypeName}} {
        {{- range .Items}}
        {{- if ne .Desc ""}} //{{.Desc}} {{end}}
        {{.FieldName}} = {{.Value}},
        {{end}}
    }
    {{end}}{{/* enum */}}

    {{- /*生成配置类类型*/}}
    {{- range .Consts}}
    // Defined in table: {{.DefinedTable}}
    var {{.TypeName}}: {
        {{- range .Items}}
            {{- if not .IsVoid }}   
                {{- if ne .Desc ""}} //{{.Desc}} {{end}}                    
        {{.FieldName}}?: {{type_format .StandardValueType .ValueType .IsArray}},
            {{end}}
            {{- if .Convertable}}    
        get{{.FieldName}}(): {{get_alias .Alias}};
            {{end}}
        {{end}} {{/*end .Items */}}
    }
    {{end}}{{/*end .Consts */}}

    {{- /*生成类类型*/}}
    {{- range .Tables}}
    {{$TypeName := .TypeName}}
    // Defined in table: {{.DefinedTable}}
    /** Properties of a {{$TypeName}}. */
    interface I{{$TypeName}} {
        {{range .Headers}}
            {{- if not .IsVoid }}                
        {{.FieldName}}?: {{type_format .StandardValueType .ValueType .IsArray}};
            {{end}} {{/*end not Void*/}}
            {{- if .Convertable}}    
        get{{.FieldName}}(): {{get_alias .Alias}};
            {{end}}
        {{end}} {{/*end .Headers */}}
    }

    /** Represents a {{$TypeName}}. */
    class {{$TypeName}} implements I{{$TypeName}} {    
        /**
         * Constructs a new {{$TypeName}}.
         * @param [properties] Properties to set
         */
        constructor(properties?: I{{$TypeName}});

        {{range .Headers}}
            {{- if not .IsVoid }}
        public {{.FieldName}}?: {{type_format .StandardValueType .ValueType .IsArray}};
            {{- end}} {{/*end not Void*/}}
            {{- if .Convertable}}    
        get{{.FieldName}}(): {{get_alias .Alias}};
            {{end}}
        {{- end}} {{/*end .Headers */}}

        /**
         * Creates a new {{$TypeName}} instance using the specified properties.
         * @param [properties] Properties to set
         * @returns {{$TypeName}} instance
         */
        public static create(properties?: {{$TypeName}}): {{$TypeName}};

        /**
         * Encodes the specified {{$TypeName}} message. Does not implicitly {@link {{$TypeName}}.verify|verify} messages.
         * @param message {{$TypeName}} message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: I{{$TypeName}}, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified {{$TypeName}} message, length delimited. Does not implicitly {@link {{$TypeName}}.verify|verify} messages.
         * @param message {{$TypeName}} message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: I{{$TypeName}}, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a {{$TypeName}} message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns {{$TypeName}}
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): {{$TypeName}};

        /**
         * Decodes a {{$TypeName}} message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns {{$TypeName}}
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): {{$TypeName}};

        /**
         * Verifies a {{$TypeName}} message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a {{$TypeName}} message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns {{$TypeName}}
         */
        public static fromObject(object: { [k: string]: any }): {{$TypeName}};

        /**
         * Creates a plain object from a {{$TypeName}} message. Also converts values to other types if specified.
         * @param message {{$TypeName}}
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: {{$TypeName}}, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this {{$TypeName}} to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }
        
    {{- end}} {{/*end .Tables */}}
}{{/* namespace */}}