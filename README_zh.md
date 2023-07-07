# simple_encryption
[English](README.md) | 简体中文

在服务器端加密数据表ID

# 用法
## 密钥Json字符串
```
var contrast string = "{\"extraItem\":2,\"key\":[\"jb10=m/zkvpds=1/\",\"/*-+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz\"]}"
```
- extraItem: 额外项添加位置，暂时只支持1个字符
- key[0]: 加密用的key
- key[1]: 加密用的字符集，不在字符集内的字符将被忽略

## 初始化
- 函数：New
- 参数：
    - extraItem: 额外项添加位置，暂时只支持1个字符
    - key: 加密用的key
    - cryptKey: 加密用的字符集，不在字符集内的字符将被忽略
- 返回值：
    - se: SimpleEncryption对象
    - err: 错误信息
- 示例：
```
se, err := New(extraItem, key, cryptKey)
if err != nil {
    println("encrypt error:", err)
    return
}
```

## 使用Json字符串初始化
- 函数：NewJson
- 参数：
    - contrast: 密钥Json字符串
- 返回值：
    - se: SimpleEncryption对象
    - err: 错误信息
- 示例：
```
se, err := NewJson(contrast)
if err != nil {
    println("encrypt error:", err)
    return
}
```

## 加密
- 函数：encrypt
- 参数：
    - str: 需要加密的字符串
    - extra: 额外的加密项，暂时只支持1个字符
- 返回值：
    - enStr: 加密后的字符串
    - err: 错误信息
- 示例：
```
var (
    extra   string = "1"
    str     string = "1z3a+bc f"
    enStr   string = ""
    err     error  = nil
)
se, err := New(contrast)
if err != nil {
    println("encrypt error:", err)
    return
}
fmt.Println("[extra]:  ", extra, "[str]:", str)
enStr, err = se.encrypt(str, extra)
if err != nil {
    fmt.Println("encrypt error:", err)
    return
}
fmt.Println("[encrypt]:         ", enStr)
```
- 输出：
```
[extra]:      1 [str]:      1z3a+bc f
[encrypt]:                  ec/xLjLJ 1
```

## 解密
- 函数：decrypt
- 参数：
    - str: 需要解密的字符串
- 返回值：
    - reStr: 解密后的字符串
    - reExtra: 额外的加密项
    - err: 错误信息
- 示例：
```
var (
    str     string = "ec/xLjLJ 1"
    reStr   string = ""
    reExtra string = ""
    err     error  = nil
)
se, err := New(contrast)
if err != nil {
    println("encrypt error:", err)
    return
}
fmt.Println("[str]:               ", str)
reStr, reExtra, err = se.decrypt(str)
if err != nil {
    fmt.Println("decrypt error:", err)
    return
}
fmt.Println("[extra]:", reExtra, "[decrypt]:", reStr)
```
- 输出：
```
[str]:                ec/xLjLJ 1
[extra]: 1 [decrypt]: 1z3a+bc f
```
