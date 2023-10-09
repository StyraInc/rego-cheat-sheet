# Rego Style Guide


## Rules




### Complete


Complete rules assign a single value. 



```rego
import future.keywords

default allow := false
allow if {
	input.user.role == "admin"
	input.user.internal
}

default request_quota := 100
request_quota := 1000 if input.user.internal
request_quota := 50 if input.user.plan.trial
```




### Partial


Partial rules generate and assign a set of values to a variable.


```rego
import future.keywords

paths contains path if {
	path := "/handbook/*"
}

paths contains path if {
	some team in input.user.teams
	path := sprintf("/teams/%v/*", [team])
}
```


```javascript
// Output
{
  "paths": [
    "/handbook/*",
    "/teams/owl/*", "/teams/tiger/*"
  ]
}
```




## Iteration




### Some


Name local query variables.


```rego
import future.keywords

all_regions := {
	"emea": {"west", "east"},
	"na": {"west", "east", "central"},
	"latam": {"west", "east"},
	"apac": {"north", "south"},
}

allowed_regions contains region_id if {
	some area, regions in all_regions

	some region in regions
	region_id := sprintf("%s_%s", [area, region])
}
```


```javascript
// Output
{
  "allowed_regions": [
    "apac_north", "apac_south", "emea_east", ...
  ]
}
```



### Every


Apply conditions to many elements.


```rego
import future.keywords

allow if {
	required_prefix := sprintf("/docs/%s/", [input.userID])
	every path in input.paths {
		startswith(path, required_prefix)
	}
}
```





## Comprehensions


_(Shared Code)_
```rego
letters := ["q", "w", "i", "e", "r", "t", "y", "u", "i", " e", "y"]
vowels := ["a", "e", "i", "o", "u", "y"]
```



### Arrays



```rego
array_match_vowels := [match |
  some letter in letters
  some vowel in vowels
  letter == vowel
  match := letter
]
```


```javascript
// Output
{
  "array_match_vowels": [
    "i", "e", "y", "u", "i", "e", "y"
  ]
}

```



### Sets



```rego
set_match_vowels := {match |
  some letter in letters
  some vowel in vowels
  letter == vowel
  match := letter
}
```


```javascript
// Output
{
  "set_match_vowels": [
    "e", "i", "u", "y"
  ]
}

```



### Objects



```rego
object_check_vowels := {letter: is_vowel |
  some letter in letters
  is_vowel := letter in vowels
}
```


```javascript
// Output
{
  "object_check_vowels": {
    "e": true, "i": true, "q": false, "r": false, "t": false, "u": true, "w": false, "y": true
  }
}

```





e": true, "i": true, "q": false, "r": false, "t": false, "u": true, "w": false, "y": true
  }
}

```
















