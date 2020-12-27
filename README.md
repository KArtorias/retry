# 重试函数

### 安装
```
go get github.com/KArtorias/retry
```
### 参数
```
retryTimes: 重试次数
retryTimeout: 重试超时时间，传0表示不设置重试超时
waitTime: 重试前等待时间，传0表示本次重试失败后立即开启下一次重试
funcName: 重试函数信息
f func()error: 函数主体，即重试执行内容
```

### 使用
参考retry_test.go
