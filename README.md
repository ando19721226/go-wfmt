Adjust the display width of wide characters.

**Bad**

```go
fmt.Printf("%-10sA\n", "nihongo")
fmt.Printf("%-10sA\n", "日本語")
```

![1](1.png)

**Good!**

```go
wfmt.Printf("%-10sA\n", "nihongo")
wfmt.Printf("%-10sA\n", "日本語")
```

![2](2.png)

