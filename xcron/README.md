# 定时器

## 用法

```go
// Seconds field, required
xcron.New(xcron.WithSeconds())

// Seconds field, optional
xcron.New(xcron.WithParser(xcron.NewParser(
xcron.SecondOptional | xcron.Minute | xcron.Hour | xcron.Dom | xcron.Month | xcron.Dow | xcron.Descriptor,
)))
```
