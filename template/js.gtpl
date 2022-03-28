// DO NOT EDIT!
// This code is auto generated by go-xlsx-exporter
// VERSION {{.Version}}
// go-protobuf {{.GoProtoVersion}}

{{- $G := .}}
{{- $NS := .Namespace}}

/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = protobuf.Reader, $Writer = protobuf.Writer, $util = protobuf.util;

// Exported root namespace
var $root = protobuf.roots["default"] || (protobuf.roots["default"] = {});

$root.{{$NS}} = (function() {
    /**
     * Namespace {{$NS}}.
     * @exports {{$NS}}
     * @namespace
     */
    var {{$NS}} = {};

    {{/*生成枚举类型*/}}
    {{- range .Enums}}
    // Defined in table: {{.DefinedTable}}
    var {{.TypeName}} = {{$NS}}.{{.TypeName}} = (function() {
        var valuesById = {}, values = Object.create(valuesById);
    {{range .Items}}
        {{if ne .Desc ""}} //{{.Desc}} {{end}}
        values[valuesById[{{.Index}}] = "{{.FieldName}}"] = {{.Value}};
    {{end}}
        return values;
    })();
    {{end}}

    {{- /*生成配置类类型*/}}
    {{- range .Consts}}
    // Defined in table: {{.DefinedTable}}
    var {{.TypeName}} = {{$NS}}.{{.TypeName}} = (function() {
        var values = Object.create(valuesById);
    {{range .Items}}
    {{- if not .IsVoid }}
        {{if ne .Desc ""}} //{{.Desc}} {{end}}
        values.{{.FieldName}} = {{value_format .Value .}};
    {{- end}}
        {{- if .Convertable}}
        
        {{- end}}
    {{end}}
        return values;
    })();
    {{end}}

    {{- /*生成类类型*/}}

    {{- range .Tables}}
    {{$TypeName := .TypeName}}
    // Defined in table: {{.DefinedTable}}
     var {{$TypeName}} = {{$NS}}.{{$TypeName}} = (function(properties) {
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
    })();

    {{range .Headers}}
        {{- if not .IsVoid }}
            {{- if .IsArray}}
                {{- if ne .Desc ""}} //{{.Desc}} {{end}}
    {{$TypeName}}.prototype.{{.FieldName}} =  $util.emptyArray;
            {{else}}
    {{$TypeName}}.prototype.{{.FieldName}} =  {{default .}};
            {{end -}}   
        {{end -}} 
    {{end -}} 

    {{$TypeName}}.create = function create(properties) {
        return new {{$TypeName}}(properties);
    };

    {{$TypeName}}.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();

        {{range .Headers}}        
            {{- $wireType := wireType .}}
            {{- $count := calcOffset .}}

            {{- if not .IsVoid }}
                {{- if .IsArray}}
                    {{- if ne .Desc ""}} //{{.Desc}} {{end}}
        if (message.{{.FieldName}} != null && message.{{.FieldName}}.length) {
            writer.uint32(/* id {{.Index}}, wireType {{$wireType}} =*/{{$count}}).fork();
            for (var i = 0; i < message.TaskItem.length; ++i)
                writer.uint32(message.TaskItem[i]);
            writer.ldelim();
        }
                {{- else}}
                {{- $pbType := getPBType .ValueType}}
        if (message.{{.FieldName}} != null && Object.hasOwnProperty.call(message, "{{.FieldName}}"))
            writer.uint32(/* id {{.Index}}, wireType {{$wireType}} =*/{{$count}}).{{$pbType}}(message.{{.FieldName}});
                {{- end -}}   
            {{end -}} 
        {{end -}} 

        return writer;
    };

    {{$TypeName}}.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    {{$TypeName}}.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.{{$TypeName}}.Chain();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                {{- range .Headers}}                
                {{- $pbType := getPBType .ValueType}}
                case {{.Index}}:
                    {{- if .IsArray}}
                    if (!(message.{{.FieldName}} && message.{{.FieldName}}.length))
                        message.{{.FieldName}} = [];
                    if ((tag & 7) === 2) {
                        var end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2)
                            message.{{.FieldName}}.push(reader.{{$pbType}}());
                    } else
                        message.{{.FieldName}}.push(reader.{{$pbType}}());
                    {{- else}}
                    message.{{.FieldName}} = reader.{{$pbType}}();
                    {{- end}} 
                    break;
                {{- end}}
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

    {{end}}
})();

module.exports = $root;