# packageUpdateTool
更新gopath下， src目录中所有的依赖包版本
==============================================

##使用说明
### 第一步：
``` go get -u github.com/zhangjunfang/packageUpdateTool```
```go
package main

import (
	"github.com/zhangjunfang/packageUpdateTool/update"
)

func main() {
	update.Update()
}
```
