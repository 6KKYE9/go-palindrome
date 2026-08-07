# go-palindrome

判断回文，零依赖。忽略大小写、空白和标点，只比字母数字。

## 用法

```bash
go run . "A man a plan a canal Panama"   # 是回文
echo "hello" | go run .                  # 不是
```

空串当「不是」，避免把空输入误判成回文。
