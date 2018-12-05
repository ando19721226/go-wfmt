

**Bad**

```go
fmt.Printf("%-10sA\n", "nihongo")
fmt.Printf("%-10sA\n", "日本語")
//nihongo   A
//日本語       A
```

**Good!**

```go
wfmt.Printf("%-10sA\n", "nihongo")
wfmt.Printf("%-10sA\n", "日本語")
//nihongo   A
//日本語    A
```

