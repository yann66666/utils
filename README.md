# utils
一些工具类
对一些标准库的封装,实现额外的功能

## sync_util
### Once 实现具有 error 返回值的sync.Once
```golang
if err := once.Do(func() error {
      //do something...
      return nil
}); err != nil {
      //do something
}
```
