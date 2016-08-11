# packageUpdateTool
更新gopath下， src目录中所有的依赖包版本
==============================================

##使用说明
### 第一步：
  ``` go get -u github.com/zhangjunfang/packageUpdateTool```
```go

var bufpool *bpool.BufferPool

func main() {

    bufpool = bpool.NewBufferPool(48)

}

func someFunction() error {

     // Get a buffer from the pool
     buf := bufpool.Get()
     ...
     ...
     ...
     // Return the buffer to the pool
     bufpool.Put(buf)

     return nil
}
```
