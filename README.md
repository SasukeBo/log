## 日志工具包

四种状态日志，format方式输出和line print输出，使用方法类似`fmt.Printf`和`fmt.Println`

### Examples

```go
log.Infoln("Starting server...")
// => 2019/11/08 09:49:21 INFO     ▶ Starting server...
log.Successln("Connect db Successfully!")
// => 2019/11/08 09:50:21 SUCCESS  ▶ Connect db Successfully!
log.Warn("TCP server already listening on %d", 8080)
// => 2019/11/08 09:51:21 WARN     ▶ TCP server already listening on 8080
log.Error("TCP dial falied: %v", err)
// => 2019/11/08 09:52:21 ERROR    ▶ TCP dial failed: timeout.
```
