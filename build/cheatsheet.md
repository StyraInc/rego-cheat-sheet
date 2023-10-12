# Rego Cheat Sheet


## Rules - <sub><sup>The building blocks of Rego</sup></sub>




### Complete 


Complete rules assign a single value. 
 ([Playground](http://localhost:8181/?config=eyJpIjoie1xuICBcInVzZXJcIjoge1xuICAgIFwicm9sZVwiOiBcImFkbWluXCIsXG4gICAgXCJpbnRlcm5hbFwiOiB0cnVlXG4gIH1cbn0iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmRlZmF1bHQgYWxsb3cgOj0gZmFsc2VcblxuYWxsb3cgaWYge1xuXHRpbnB1dC51c2VyLnJvbGUgPT0gXCJhZG1pblwiXG5cdGlucHV0LnVzZXIuaW50ZXJuYWxcbn1cblxuZGVmYXVsdCByZXF1ZXN0X3F1b3RhIDo9IDEwMFxuXG5yZXF1ZXN0X3F1b3RhIDo9IDEwMDAgaWYgaW5wdXQudXNlci5pbnRlcm5hbFxucmVxdWVzdF9xdW90YSA6PSA1MCBpZiBpbnB1dC51c2VyLnBsYW4udHJpYWxcbiJ9))


```rego
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


Partial rules generate and assign a set of values to a variable. ([Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbnBhdGhzIGNvbnRhaW5zIHBhdGggaWYge1xuXHRwYXRoIDo9IFwiL2hhbmRib29rLypcIlxufVxuXG5wYXRocyBjb250YWlucyBwYXRoIGlmIHtcblx0c29tZSB0ZWFtIGluIGlucHV0LnVzZXIudGVhbXNcblx0cGF0aCA6PSBzcHJpbnRmKFwiL3RlYW1zLyV2LypcIiwgW3RlYW1dKVxufVxuIn0%3D))


```rego
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






## Iteration - <sub><sup>Make quick work of collections</sup></sub>




### Some 


Name local query variables. ([Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmFsbF9yZWdpb25zIDo9IHtcblx0XCJlbWVhXCI6IHtcIndlc3RcIiwgXCJlYXN0XCJ9LFxuXHRcIm5hXCI6IHtcIndlc3RcIiwgXCJlYXN0XCIsIFwiY2VudHJhbFwifSxcblx0XCJsYXRhbVwiOiB7XCJ3ZXN0XCIsIFwiZWFzdFwifSxcblx0XCJhcGFjXCI6IHtcIm5vcnRoXCIsIFwic291dGhcIn0sXG59XG5cbmFsbG93ZWRfcmVnaW9ucyBjb250YWlucyByZWdpb25faWQgaWYge1xuXHRzb21lIGFyZWEsIHJlZ2lvbnMgaW4gYWxsX3JlZ2lvbnNcblxuXHRzb21lIHJlZ2lvbiBpbiByZWdpb25zXG5cdHJlZ2lvbl9pZCA6PSBzcHJpbnRmKFwiJXNfJXNcIiwgW2FyZWEsIHJlZ2lvbl0pXG59XG4ifQ%3D%3D))


```rego
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


Apply conditions to many elements. ([Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmFsbG93IGlmIHtcblx0cHJlZml4IDo9IHNwcmludGYoXCIvZG9jcy8lcy9cIiwgW2lucHV0LnVzZXJJRF0pXG5cdGV2ZXJ5IHBhdGggaW4gaW5wdXQucGF0aHMge1xuXHRcdHN0YXJ0c3dpdGgocGF0aCwgcHJlZml4KVxuXHR9XG59XG4ifQ%3D%3D))


```rego
allow if {
	prefix := sprintf("/docs/%s/", [input.userID])
	every path in input.paths {
		startswith(path, prefix)
	}
}
```







## Control Flow - <sub><sup>Handle different conditions</sup></sub>




### Logical And 


Statements in rules are 'anded' together. ([Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbnZhbGlkX3N0YWZmX2VtYWlsIGlmIHtcblx0cmVnZXgubWF0Y2goYF5cXFMrQFxcUytcXC5cXFMrJGAsIGlucHV0LmVtYWlsKVxuXG5cdGVuZHN3aXRoKGlucHV0LmVtYWlsLCBcImV4YW1wbGUuY29tXCIpXG59XG4ifQ%3D%3D))


```rego
valid_staff_email if {
	regex.match(`^\S+@\S+\.\S+$`, input.email)

	endswith(input.email, "example.com")
}
```






### Logical Or 


Express OR with multiple rules, functions or the in keyword. ([Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbiMgdXNpbmcgbXVsdGlwbGUgcnVsZXNcbnZhbGlkX2VtYWlsIGlmIGVuZHN3aXRoKGlucHV0LmVtYWlsLCBcIkBleGFtcGxlLmNvbVwiKVxudmFsaWRfZW1haWwgaWYgZW5kc3dpdGgoaW5wdXQuZW1haWwsIFwiQGV4YW1wbGUub3JnXCIpXG52YWxpZF9lbWFpbCBpZiBlbmRzd2l0aChpbnB1dC5lbWFpbCwgXCJAZXhhbXBsZS5uZXRcIilcblxuIyB1c2luZyBmdW5jdGlvbnNcbmFsbG93ZWRfZmlyc3RuYW1lKG5hbWUpIGlmIG5hbWUgPT0gXCJqb2VcIlxuYWxsb3dlZF9maXJzdG5hbWUobmFtZSkgaWYgbmFtZSA9PSBcImphbmVcIlxuXG52YWxpZF9uYW1lIGlmIHtcblx0YWxsb3dlZF9maXJzdG5hbWUoaW5wdXQubmFtZSlcbn1cblxuIyB1c2luZyBgaW5gXG52YWxpZF9yZXF1ZXN0IGlmIHtcblx0aW5wdXQubWV0aG9kIGluIHtcIkdFVFwiLCBcIlBPU1RcIn1cbn1cbiJ9))


```rego
# using multiple rules
valid_email if endswith(input.email, "@example.com")
valid_email if endswith(input.email, "@example.org")
valid_email if endswith(input.email, "@example.net")

# using functions
allowed_firstname(name) if name == "joe"
allowed_firstname(name) if name == "jane"

valid_name if {
	allowed_firstname(input.name)
}

# using `in`
valid_request if {
	input.method in {"GET", "POST"}
}
```







## Testing - <sub><sup>Validate your policy's behavior</sup></sub>




### With 


Override input and data using the with keyword. ([Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmFsbG93IGlmIHtcblx0aW5wdXQuYWRtaW4gPT0gdHJ1ZVxufVxuXG50ZXN0X2FsbG93X3doZW5fYWRtaW4gaWYge1xuXHRhbGxvdyB3aXRoIGlucHV0IGFzIHtcImFkbWluXCI6IHRydWV9XG59XG4ifQ%3D%3D))


```rego
allow if {
	input.admin == true
}

test_allow_when_admin if {
	allow with input as {"admin": true}
}
```







## Debugging - <sub><sup>Find and fix problems</sup></sub>




### Print 


Use print in rules to inspect values at runtime. ([Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmFsbG93ZWRfdXNlcnMgOj0ge1wiYWxpY2VcIiwgXCJib2JcIiwgXCJjaGFybGllXCJ9XG5cbmFsbG93IGlmIHtcblx0c29tZSB1c2VyIGluIGFsbG93ZWRfdXNlcnNcblx0cHJpbnQodXNlcilcblx0aW5wdXQudXNlciA9PSB1c2VyXG59XG4ifQ%3D%3D))


```rego
allowed_users := {"alice", "bob", "charlie"}

allow if {
	some user in allowed_users
	print(user)
	input.user == user
}
```


```javascript
// Output
// alice
// bob
// charlie

```






## Builtins - <sub><sup>Handy functions for common tasks</sup></sub>




### Aggregates 


[Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbnZhbHMgOj0gWzUsIDEsIDQsIDIsIDNdXG5cbnZhbHNfY291bnQgOj0gY291bnQodmFscylcbnZhbHNfbWF4IDo9IG1heCh2YWxzKVxudmFsc19taW4gOj0gbWluKHZhbHMpXG52YWxzX3NvcnRlZCA6PSBzb3J0KHZhbHMpXG52YWxzX3N1bSA6PSBzdW0odmFscylcbiJ9)


```rego
vals := [5, 1, 4, 2, 3]

vals_count := count(vals)
vals_max := max(vals)
vals_min := min(vals)
vals_sorted := sort(vals)
vals_sum := sum(vals)
```


```javascript
// Output
{
  "vals_count": 5,
  "vals_max": 5,
  "vals_min": 1,
  "vals_sorted": [1, 2, 3, 4, 5],
  "vals_sum": 15
}
```





### Objects 


[Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbm9iaiA6PSB7XCJ1c2VyaWRcIjogXCIxODQ3MlwiLCBcInJvbGVzXCI6IFt7XCJuYW1lXCI6IFwiYWRtaW5cIn1dfVxuXG52YWwgOj0gb2JqZWN0LmdldChvYmosIFtcInJvbGVzXCIsIDAsIFwibmFtZVwiXSwgXCJtaXNzaW5nXCIpXG5cbmRlZmF1bHRlZF92YWwgOj0gb2JqZWN0LmdldChcblx0b2JqLFxuXHRbXCJyb2xlc1wiLCAwLCBcInBlcm1pc3Npb25zXCJdLFxuXHRcInVua25vd25cIixcbilcblxua2V5cyA6PSBvYmplY3Qua2V5cyhvYmopXG4ifQ%3D%3D)


```rego
obj := {"userid": "18472", "roles": [{"name": "admin"}]}

val := object.get(obj, ["roles", 0, "name"], "missing")

defaulted_val := object.get(
	obj,
	["roles", 0, "permissions"],
	"unknown",
)

keys := object.keys(obj)
```


```javascript
// Output
{
  "val": "admin",
  "defaulted_val": "unknown",

  "keys": [
    "roles",
    "userid"
  ]
}
```





### Strings 


[Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmV4YW1wbGVfc3RyaW5nIDo9IFwiQnVpbGQgUG9saWN5IGFzIENvZGUgd2l0aCBPUEEhXCJcblxuY2hlY2tfY29udGFpbnMgaWYgY29udGFpbnMoZXhhbXBsZV9zdHJpbmcsIFwiT1BBXCIpXG5jaGVja19zdGFydHN3aXRoIGlmIHN0YXJ0c3dpdGgoZXhhbXBsZV9zdHJpbmcsIFwiQnVpbGRcIilcbmNoZWNrX2VuZHN3aXRoIGlmIGVuZHN3aXRoKGV4YW1wbGVfc3RyaW5nLCBcIiFcIilcbmNoZWNrX3JlcGxhY2UgOj0gcmVwbGFjZShleGFtcGxlX3N0cmluZywgXCJPUEFcIiwgXCJTdHlyYVwiKVxuY2hlY2tfc3ByaW50ZiA6PSBzcHJpbnRmKFwiT1BBIGlzICVzIVwiLCBbXCJhd2Vzb21lXCJdKVxuIn0%3D)


```rego
example_string := "Build Policy as Code with OPA!"

check_contains if contains(example_string, "OPA")
check_startswith if startswith(example_string, "Build")
check_endswith if endswith(example_string, "!")
check_replace := replace(example_string, "OPA", "Styra")
check_sprintf := sprintf("OPA is %s!", ["awesome"])
```


```javascript
// Output
{
  "check_contains": true,
  "check_startswith": true,
  "check_endswith": true,
  "check_replace": "Build Policy as Code with Styra!",
  "check_sprintf": "OPA is awesome!"
}
```





### Regex 


[Playground](http://localhost:8181/?config=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmV4YW1wbGVfc3RyaW5nIDo9IFwiQnVpbGQgUG9saWN5IGFzIENvZGUgd2l0aCBPUEEhXCJcblxuY2hlY2tfbWF0Y2ggaWYgcmVnZXgubWF0Y2goYFxcdytgLCBleGFtcGxlX3N0cmluZylcblxuY2hlY2tfcmVwbGFjZSA6PSByZWdleC5yZXBsYWNlKGV4YW1wbGVfc3RyaW5nLCBgXFxzK2AsIFwiX1wiKVxuIn0%3D)


```rego
example_string := "Build Policy as Code with OPA!"

check_match if regex.match(`\w+`, example_string)

check_replace := regex.replace(example_string, `\s+`, "_")
```


```javascript
// Output
{
  "check_match": true,
  "check_replace": "Build_Policy_as_Code_with_OPA!"
}
```







