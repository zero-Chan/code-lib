# 程序模板
- bin/ 存放二进制文件
- app/ 存放main.go文件
- executor/ 存放executor实例
- conf/ 存放conf.go文件
- conf/files/ 存放记录配置信息的文件
- models/ 存放通用模块
- prot/ 存放协议数据结构
- prot/exec-prot/ 存放exec类使用的请求体和响应体
- restmux/ 存放项目使用的多路复用器

# Start
```shell
cd ${go_path}/src
ln -s ${rest_path}/example rest_example
```

# 规范
## executor/　中，每一个大类作为目录，每一个大类的操作作为文件
e.g:
```
executor/client-exec		
executor/client-exec/client-add.go	# 编码client.add的`builder`和`executor`.  package client_exec
executor/client-exec/client-get.go	# 编码client.get的`builder`和`executor`.  package client_exec
```
- 命名规范
```
builder类:   XXXBuilder
executor类:  XXXExec
```
e.g: 
file: executor/client-exec/client-add.go
```golang
package client_exec

type ClientAddBuilder struct {
}

func (this *ClientAddBuilder) BuildHandlerFromHTTP(httpch *resthttp.HTTPChannel) (rest.Executor, *rest.RestResponse) {
	// 把　httpch 的数据来源转换为 ClientAddExec 需要的数据来源 requestBody
	// httpch -> ClientAddExec.RequestBody
	
	// 使用 requestBody 新建 ClientAddExec，　返回
	// return ClientAddExec
}

type ClientAddExec struct {
	restChannel  *rest.RestChannel
	requestBody  *prot.ClientAddRequest
}

func (this *ClientAddExec) Prepare() *rest.RestResponse {
}

func (this *ClientAddExec) Handle() *rest.RestResponse {
}

func (this *ClientAddExec) Finish() *rest.RestResponse {
}
```

## prot/exec-prot/　中，目录结构要和 executor/　里面的一样
e.g:
```
prot/exec-prot/client-exec-prot		
prot/exec-prot/client-exec/client-add.go	# 编码client.add的`Request`和`Response`.  package client_exec
prot/exec-prot/client-exec/client-get.go	# 编码client.get的`Request`和`Response`.  package client_exec
```