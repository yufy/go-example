### 安装
安装go-swagger, 文档链接https://goswagger.io/

### 生成数据
* yml格式：swagger generate spec -o docs/swagger.yml
* json格式：swagger generate spec -o docs/swagger.json

### 启动服务
swagger serve ./docs/swagger.yml

### 注意
* swagger:meta和包声明之间不能有空行，否则将不会生成swagger:meta中定义的信息