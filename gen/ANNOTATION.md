# Annotation

Format
```
@<annotation_name>(
    <annotation_body> // toml format
)
```

Example
```
@name(
    str1 = "val1"
    int1 = 123
    float2 = 123.123
    arrInt1 = [1,2,3]
    arrStr1 = ["a", "b", "c"]
)
```

## Notes

* false sensitive of the end of a annotation if in comments (toml) is the closing round brackets. Use toml without comments with round brackets
