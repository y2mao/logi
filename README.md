# logi
A simple, high performance logger with delay writing, file rolling and level

#### Usage

```go
logi.Info("http", "server started")
logi.Infof("http", "listening port:%d", 18801)
logi.Error("api", "undefined API")
logi.Errorf("api", "%s failed. %v", "/user/profile", errors.New("Invalid ID"))

// Log will be flushed to disk automatically (depends on interval seconds or buffer size).
// If necessary, Call this method to flush data from buffer to disk directly. 
logi.FlushAll()
```

Generated files as following:
```
20160216.api.log
20160216.http.log
```

#### Options
logi enables user change options with command argument like  `./yourapp --logi-<opt>`. Following are valid options.

| Command | Default Value |  Description |
| --- | --- | --- |
| --logi-dir | ./logs | the directory where the log be written to |
| --logi-quiet | false | diable console output. Please turn it on in production env |
| --logi-rolling | 20060102 | rolling format. Please refer to Golang time format |
| --logi-bufsize | 4194304 | max buffer size for caching log. default size is 4Mbytes |
| --logi-interval | 15 | flush interval(sec) |
