# simple_encryption
English | [简体中文](README_zh.md)

Encrypt data table ID on server side

# Usage
## Comparison value string
```
var contrast string = "{\"extraItem\":2,\"key\":[\"jb10=m/zkvpds=1/\",\"/*-+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz\"]}"
```
- extraItem: Extra item add position, currently only supports 1 character
- key[0]: Encryption key
- key[1]: Character set used for encryption, characters not in the character set will be ignored

## Initialize
- Function: New
- Parameter:
    - extraItem: Extra item add position, currently only supports 1 character
    - key: Encryption key
    - cryptKey: Character set used for encryption, characters not in the character set will be ignored
- Return value:
    - se: SimpleEncryption object
    - err: Error message
- Example:
```
se, err := New(extraItem, key, cryptKey)
if err != nil {
    println("encrypt error:", err)
    return
}
```

## Initialize with Json string
- Function: NewJson
- Parameter:
    - contrast: Json string of key
- Return value:
    - se: SimpleEncryption object
    - err: Error message
- Example:
```
se, err := NewJson(contrast)
if err != nil {
    println("encrypt error:", err)
    return
}
```

## encrypt
- Function: encrypt
- Parameter:
    - str: String to be encrypted
    - extra: Extra encryption item, currently only supports 1 character
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
