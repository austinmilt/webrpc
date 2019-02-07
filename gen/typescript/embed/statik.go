// Code generated by statik. DO NOT EDIT.

// Package statik contains static assets.
package embed

import (
	"github.com/rakyll/statik/fs"
)

func init() {
	data := "PK\x03\x04\x14\x00\x08\x00\x00\x00\xf7$GN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00client.go.tmplUT\x05\x00\x01\x93\xb6[\\{{define \"client\"}}\n{{if .Services}}\n  // Client\n\n  {{range .Services}}\n  const {{.Name | constPathPrefix}} = \"/rpc/{{.Name}}/\"\n  {{end}}\n\n  {{range .Services}}\n\n    export class {{.Name}} implements {{.Name | serviceInterfaceName}} {\n      private hostname: string\n      private fetch: Fetch\n      private path = '/rpc/{{.Name}}'\n\n			constructor(hostname: string, fetch: Fetch) {\n				this.hostname = hostname\n				this.fetch = fetch\n			}\n\n			private url(name: string): string {\n				return this.hostname + this.path + name\n			}\n\n      {{range .Methods}}\n        {{.Name}}({{.Inputs | methodInputs}}) ({{.Outputs | methodOutputs}}) {\n					return this.fetch(\n						this.url('{{.Name}}'),\n						createHTTPRequest(params, headers)\n					).then((res) => {\n						if (!res.ok) {\n							return throwHTTPError(res)\n						}\n						{{range $output := .Outputs}}\n							return res.json().then((_data) => {return {{$output | newResponseConcreteType}}(_data)})\n						{{end}}\n					})\n				}\n      {{end}}\n    }\n\n  {{end}}\n{{end}}\n{{end}}\nPK\x07\x08\xbdZ\xd7\x00\xfb\x03\x00\x00\xfb\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00*tGN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00	\x00helpers.go.tmplUT\x05\x00\x01\xb1A\\\\{{define \"helpers\"}}\n\nexport interface WebRPCErrorJSON {\n  code: string\n  msg: string\n  meta: {\n    [index: string]: string\n  }\n}\n\nexport class WebRPCError extends Error {\n  code: string\n  meta: {\n    [index: string]: string\n  }\n\n  constructor(te: WebRPCErrorJSON) {\n    super(te.msg)\n\n    this.code = te.code\n    this.meta = te.meta\n  }\n}\n\nexport const throwHTTPError = (resp: Response) => {\n  return resp.json().then((err: WebRPCErrorJSON) => { throw new WebRPCError(err) })\n}\n\nexport const createHTTPRequest = (body: object = {}, headers: object = {}): object => {\n  return {\n    method: 'POST',\n    headers: { ...headers, 'Content-Type': 'application/json' },\n    body: JSON.stringify(body || {})\n  }\n}\n\nexport type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>\n{{end}}\nPK\x07\x08d\x1eg	\x1e\x03\x00\x00\x1e\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x07sGN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00proto.gen.go.tmplUT\x05\x00\x01\x8f?\\\\{{define \"proto\"}}\n\n/* tslint:disable */\n\n// This file has been generated by https://github.com/webrpc/webrpc\n// Do not edit.\n\n{{template \"types\" .}}\n\n{{template \"client\" .}}\n\n{{template \"server\" .}}\n\n{{template \"helpers\" .}}\n\n{{end}}\nPK\x07\x08\xb1:\x11\x13\xeb\x00\x00\x00\xeb\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xf7$GN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00server.go.tmplUT\x05\x00\x01\x93\xb6[\\{{define \"server\"}}\n{{if .Services}}\n  // TODO: Server\n{{end}}\n{{end}}\nPK\x07\x087\xb8\xf5\xcbG\x00\x00\x00G\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00<tGN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0d\x00	\x00types.go.tmplUT\x05\x00\x01\xd5A\\\\{{define \"types\"}}\n{{if .Messages}}\n  {{range .Messages}}\n    {{if .Type | isEnum}}\n      {{$enumName := .Name}}\n      {{range .Fields}}\n        // {{$enumName}}_{{.Name}} = {{.Type}} {{.Value}}\n      {{end}}\n    {{end}}\n    {{if .Type | isStruct  }}\n      export interface {{.Name | interfaceName}} {\n        {{range .Fields}}\n          {{.Name | exportedField}}{{.Optional | optional}}: {{.Type | fieldType}}\n        {{end}}\n      }\n\n      export class {{.Name}} implements {{.Name | interfaceName}} {\n        private _data: {{.Name | interfaceName}}\n        constructor(_data?: {{.Name | interfaceName}}) {\n          this._data = {}\n          if (_data) {\n            {{range .Fields}}\n              this._data['{{.Name | exportedField}}'] = _data['{{.Name | exportedField}}']\n            {{end}}\n          }\n        }\n        public toJSON(): object {\n          returns this._data\n        }\n      }\n    {{end}}\n  {{end}}\n{{end}}\n{{if .Services}}\n  {{range .Services}}\n    export interface {{.Name | serviceInterfaceName}} {\n      {{range .Methods}}\n        {{.Name}}: ({{.Inputs | methodInputs}}) {{.Outputs | methodOutputs}}\n      {{end}}\n    }\n  {{end}}\n{{end}}\n{{end}}\nPK\x07\x08\xfb\xd5bN\x98\x04\x00\x00\x98\x04\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xf7$GN\xbdZ\xd7\x00\xfb\x03\x00\x00\xfb\x03\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\x00\x00\x00\x00client.go.tmplUT\x05\x00\x01\x93\xb6[\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00*tGNd\x1eg	\x1e\x03\x00\x00\x1e\x03\x00\x00\x0f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81@\x04\x00\x00helpers.go.tmplUT\x05\x00\x01\xb1A\\\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x07sGN\xb1:\x11\x13\xeb\x00\x00\x00\xeb\x00\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\xa4\x07\x00\x00proto.gen.go.tmplUT\x05\x00\x01\x8f?\\\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xf7$GN7\xb8\xf5\xcbG\x00\x00\x00G\x00\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\xd7\x08\x00\x00server.go.tmplUT\x05\x00\x01\x93\xb6[\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00<tGN\xfb\xd5bN\x98\x04\x00\x00\x98\x04\x00\x00\x0d\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81c	\x00\x00types.go.tmplUT\x05\x00\x01\xd5A\\\\PK\x05\x06\x00\x00\x00\x00\x05\x00\x05\x00\\\x01\x00\x00?\x0e\x00\x00\x00\x00"
	fs.Register(data)
}