# simple_encryption
English | [简体中文](README_zh.md)
Encrypt database ID on the server side using

# Usage
## Comparison value string
```
var contrast string = "{\"extraItem\":2,\"key\":[\"jb10=m/zkvpds=1/\",\"/*-+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz\"]}"
```
- extraItem: Extra item add position, currently only supports 1 character
- key[0]: Encryption key
- key[1]: Character set used for encryption, characters not in the character set will be ignored

## encrypt
- Function: encrypt
- Parameter:
    - str: String to be encrypted
    - extra: Extra encryption item, currently only supports 1 character
    - contrast: Comparison value string
- Return value:
    - enStr: Encrypted string
    - err: Error message
- Example:
```
var (
    extra   string = "1"
    str     string = "1z3a+bc f"
    enStr   string = ""
    err     error  = nil
)
se, err := New(contrast)
if err != nil {
    t.Error("encrypt error:", err)
    return
}
println("[extra]:  ", extra, "[str]:", str)
enStr, err = se.encrypt(str, extra)
if err != nil {
    println("encrypt error:", err)
    return
}
println("[encrypt]:         ", enStr)
```
- Output:
```
[extra]:      1 [str]:      1z3a+bc f
[encrypt]:                  ec/xLjLJ 1
```

## decrypt
- Function: decrypt
- Parameter:
    - str: String to be decrypted
    - contrast: Comparison value string
- Return value:
    - reStr: Decrypted string
    - reExtra: Extra encryption item
    - err: Error message
- Example:
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
println("[str]:               ", str)
reStr, reExtra, err = se.decrypt(str)
if err != nil {
    println("decrypt error:", err)
    return
}
println("[extra]:", reExtra, "[decrypt]:", reStr)
```
- Output:
```
[str]:                ec/xLjLJ 1
[extra]: 1 [decrypt]: 1z3a+bc f
```
