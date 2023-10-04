# Rego Style Guide


## Types




## Strings

```rego
string := "hello rego!"
complicated_string := `new\nline\n...wow`
```




## Numbers

```rego
number := 1
also_number := 3.14
```




## Arrays

```rego
array := [1,2,"three",null]
```




## Sets

```rego
set := {1,2,1}

equal := {1,2,1} == {1,2}
```




## Objects

```rego
object := {
  "key": true,
  "foo": "bar",
}

nested_object := {
  "bird": {
    "nest": [ "egg" ]
  },
}
```





## Comprehensions


_(Shared Code)_
```rego
letters := ["q", "w", "i", "e", "r", "t", "y", "u", "i", " e", "y"]
vowels := ["a", "e", "i", "o", "u", "y"]
```



## Arrays

```rego
array_match_vowels := [match |
  some letter in letters
  some vowel in vowels
  letter == vowel
  match := letter
]
```


_(Output)_
```json
{
  "array_match_vowels": [
    "i", "e", "y", "u", "i", "e", "y"
  ]
}

```



## Sets

```rego
set_match_vowels := {match |
  some letter in letters
  some vowel in vowels
  letter == vowel
  match := letter
}
```


_(Output)_
```json
{
  "set_match_vowels": [
    "e", "i", "u", "y"
  ]
}

```



## Objects

```rego
object_check_vowels := {letter: is_vowel |
  some letter in letters
  is_vowel := letter in vowels
}
```


_(Output)_
```json
{
  "object_check_vowels": {
    "e": true, "i": true, "q": false, "r": false, "t": false, "u": true, "w": false, "y": true
  }
}

```





