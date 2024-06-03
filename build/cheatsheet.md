# Rego Cheat Sheet


## Rules - <sub><sup>The building blocks of Rego</sup></sub>


### Complete Rules 


Complete rules assign a single value. 
 ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie1xuICBcInVzZXJcIjoge1xuICAgIFwicm9sZVwiOiBcImFkbWluXCIsXG4gICAgXCJpbnRlcm5hbFwiOiB0cnVlXG4gIH1cbn0iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmRlZmF1bHQgYWxsb3cgOj0gZmFsc2VcblxuYWxsb3cgaWYge1xuXHRpbnB1dC51c2VyLnJvbGUgPT0gXCJhZG1pblwiXG5cdGlucHV0LnVzZXIuaW50ZXJuYWxcbn1cblxuZGVmYXVsdCByZXF1ZXN0X3F1b3RhIDo9IDEwMFxuXG5yZXF1ZXN0X3F1b3RhIDo9IDEwMDAgaWYgaW5wdXQudXNlci5pbnRlcm5hbFxucmVxdWVzdF9xdW90YSA6PSA1MCBpZiBpbnB1dC51c2VyLnBsYW4udHJpYWxcbiJ9))


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






### Partial Rules 


Partial rules generate and assign a set of values to a variable. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie1xuICBcInVzZXJcIjoge1xuICAgIFwidGVhbXNcIjogW1xuICAgICAgXCJvcHNcIixcbiAgICAgIFwiZW5nXCJcbiAgICBdXG4gIH1cbn1cbiIsInAiOiJwYWNrYWdlIGNoZWF0XG5cbmltcG9ydCBmdXR1cmUua2V5d29yZHNcblxucGF0aHMgY29udGFpbnMgXCIvaGFuZGJvb2svKlwiXG5cbnBhdGhzIGNvbnRhaW5zIHBhdGggaWYge1xuXHRzb21lIHRlYW0gaW4gaW5wdXQudXNlci50ZWFtc1xuXHRwYXRoIDo9IHNwcmludGYoXCIvdGVhbXMvJXYvKlwiLCBbdGVhbV0pXG59XG4ifQ%3D%3D))


```rego
paths contains "/handbook/*"

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


Name local query variables. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmFsbF9yZWdpb25zIDo9IHtcblx0XCJlbWVhXCI6IHtcIndlc3RcIiwgXCJlYXN0XCJ9LFxuXHRcIm5hXCI6IHtcIndlc3RcIiwgXCJlYXN0XCIsIFwiY2VudHJhbFwifSxcblx0XCJsYXRhbVwiOiB7XCJ3ZXN0XCIsIFwiZWFzdFwifSxcblx0XCJhcGFjXCI6IHtcIm5vcnRoXCIsIFwic291dGhcIn0sXG59XG5cbmFsbG93ZWRfcmVnaW9ucyBjb250YWlucyByZWdpb25faWQgaWYge1xuXHRzb21lIGFyZWEsIHJlZ2lvbnMgaW4gYWxsX3JlZ2lvbnNcblxuXHRzb21lIHJlZ2lvbiBpbiByZWdpb25zXG5cdHJlZ2lvbl9pZCA6PSBzcHJpbnRmKFwiJXNfJXNcIiwgW2FyZWEsIHJlZ2lvbl0pXG59XG4ifQ%3D%3D))


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


Check conditions on many elements. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie1xuICBcInVzZXJJRFwiOiBcInUxMjNcIixcbiAgXCJwYXRoc1wiOiBbXG4gICAgXCIvZG9jcy91MTIzL25vdGVzLnR4dFwiLFxuICAgIFwiL2RvY3MvdTEyMy9xNC1yZXBvcnQuZG9jeFwiXG4gIF1cbn0iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmFsbG93IGlmIHtcblx0cHJlZml4IDo9IHNwcmludGYoXCIvZG9jcy8lcy9cIiwgW2lucHV0LnVzZXJJRF0pXG5cdGV2ZXJ5IHBhdGggaW4gaW5wdXQucGF0aHMge1xuXHRcdHN0YXJ0c3dpdGgocGF0aCwgcHJlZml4KVxuXHR9XG59XG4ifQ%3D%3D))


```rego
allow if {
	prefix := sprintf("/docs/%s/", [input.userID])
	every path in input.paths {
		startswith(path, prefix)
	}
}
```







## Control Flow - <sub><sup>Handle different conditions</sup></sub>


### Logical AND 


Statements in rules are joined with logical AND. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie1xuICBcImVtYWlsXCI6IFwiam9lQGV4YW1wbGUuY29tXCJcbn0iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbnZhbGlkX3N0YWZmX2VtYWlsIGlmIHtcblx0cmVnZXgubWF0Y2goYF5cXFMrQFxcUytcXC5cXFMrJGAsIGlucHV0LmVtYWlsKVxuXG5cdCMgYW5kXG5cdGVuZHN3aXRoKGlucHV0LmVtYWlsLCBcImV4YW1wbGUuY29tXCIpXG59XG4ifQ%3D%3D))


```rego
valid_staff_email if {
	regex.match(`^\S+@\S+\.\S+$`, input.email)

	# and
	endswith(input.email, "example.com")
}
```






### Logical OR 


Express OR with multiple rules, functions or the in keyword. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie1xuICBcImVtYWlsXCI6IFwib3BhQGV4YW1wbGUuY29tXCIsXG4gIFwibmFtZVwiOiBcImFubmFcIixcbiAgXCJtZXRob2RcIjogXCJHRVRcIlxufSIsInAiOiJwYWNrYWdlIGNoZWF0XG5cbmltcG9ydCBmdXR1cmUua2V5d29yZHNcblxuIyB1c2luZyBtdWx0aXBsZSBydWxlc1xudmFsaWRfZW1haWwgaWYgZW5kc3dpdGgoaW5wdXQuZW1haWwsIFwiQGV4YW1wbGUuY29tXCIpXG52YWxpZF9lbWFpbCBpZiBlbmRzd2l0aChpbnB1dC5lbWFpbCwgXCJAZXhhbXBsZS5vcmdcIilcbnZhbGlkX2VtYWlsIGlmIGVuZHN3aXRoKGlucHV0LmVtYWlsLCBcIkBleGFtcGxlLm5ldFwiKVxuXG4jIHVzaW5nIGZ1bmN0aW9uc1xuYWxsb3dlZF9maXJzdG5hbWUobmFtZSkgaWYge1xuICAgIHN0YXJ0c3dpdGgobmFtZSwgXCJhXCIpXG4gICAgY291bnQobmFtZSkgXHUwMDNlIDJcbn1cbmFsbG93ZWRfZmlyc3RuYW1lKFwiam9lXCIpICMgaWYgbmFtZSA9PSAnam9lJ1xuXG52YWxpZF9uYW1lIGlmIHtcblx0YWxsb3dlZF9maXJzdG5hbWUoaW5wdXQubmFtZSlcbn1cblxuIyB1c2luZyBgaW5gXG52YWxpZF9yZXF1ZXN0IGlmIHtcblx0aW5wdXQubWV0aG9kIGluIHtcIkdFVFwiLCBcIlBPU1RcIn1cbn1cbiJ9))


```rego
# using multiple rules
valid_email if endswith(input.email, "@example.com")
valid_email if endswith(input.email, "@example.org")
valid_email if endswith(input.email, "@example.net")

# using functions
allowed_firstname(name) if {
    startswith(name, "a")
    count(name) > 2
}
allowed_firstname("joe") # if name == 'joe'

valid_name if {
	allowed_firstname(input.name)
}

# using `in`
valid_request if {
	input.method in {"GET", "POST"}
}
```


```javascript
// Output
{
  "email": "opa@example.com",
  "name": "anna",
  "method": "GET"
}
```






## Testing - <sub><sup>Validate your policy's behavior</sup></sub>


### With 


Override input and data using the with keyword. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmFsbG93IGlmIHtcblx0aW5wdXQuYWRtaW4gPT0gdHJ1ZVxufVxuXG50ZXN0X2FsbG93X3doZW5fYWRtaW4gaWYge1xuXHRhbGxvdyB3aXRoIGlucHV0IGFzIHtcImFkbWluXCI6IHRydWV9XG59XG4ifQ%3D%3D))


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


Use print in rules to inspect values at runtime. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmFsbG93ZWRfdXNlcnMgOj0ge1wiYWxpY2VcIiwgXCJib2JcIiwgXCJjaGFybGllXCJ9XG5cbmFsbG93IGlmIHtcblx0c29tZSB1c2VyIGluIGFsbG93ZWRfdXNlcnNcblx0cHJpbnQodXNlcilcblx0aW5wdXQudXNlciA9PSB1c2VyXG59XG4ifQ%3D%3D))


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






## Comprehensions - <sub><sup>Rework and process collections</sup></sub>


### Arrays 


Produce ordered collections, maintaining duplicates.
 ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmRvdWJsZWQgOj0gW20gfFxuXHRzb21lIG4gaW4gWzEsIDIsIDMsIDNdXG5cdG0gOj0gbiAqIDJcbl1cbiJ9))


```rego
doubled := [m |
	some n in [1, 2, 3, 3]
	m := n * 2
]
```


```javascript
// Output
{
  "doubled": [2, 4, 6, 6]
}
```





### Sets 


Produce unordered collections without duplicates.
 ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbnVuaXF1ZV9kb3VibGVkIDo9IHttIHxcblx0c29tZSBuIGluIFsxMCwgMjAsIDMwLCAyMCwgMTBdXG5cdG0gOj0gbiAqIDJcbn1cbiJ9))


```rego
unique_doubled := {m |
	some n in [10, 20, 30, 20, 10]
	m := n * 2
}
```


```javascript
// Output
{
  "unique_doubled": [20, 40, 60]
}
```





### Objects 


Produce key:value data. Note, keys must be unique.
 ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmlzX2V2ZW4gOj0ge251bWJlcjogaXNfZXZlbiB8XG5cdHNvbWUgbnVtYmVyIGluIFsxLCAyLCAzLCA0XVxuXHRpc19ldmVuIDo9IChudW1iZXIgJSAyKSA9PSAwXG59XG4ifQ%3D%3D))


```rego
is_even := {number: is_even |
	some number in [1, 2, 3, 4]
	is_even := (number % 2) == 0
}
```


```javascript
// Output
{
  "is_even": {
    "1": false, "2": true, "3": false, "4": true
  }
}
```






## Builtins - <sub><sup>Handy functions for common tasks</sup></sub>


### Regex 


Pattern match and replace string data. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmV4YW1wbGVfc3RyaW5nIDo9IFwiQnVpbGQgUG9saWN5IGFzIENvZGUgd2l0aCBPUEEhXCJcblxuY2hlY2tfbWF0Y2ggaWYgcmVnZXgubWF0Y2goYFxcdytgLCBleGFtcGxlX3N0cmluZylcblxuY2hlY2tfcmVwbGFjZSA6PSByZWdleC5yZXBsYWNlKGV4YW1wbGVfc3RyaW5nLCBgXFxzK2AsIFwiX1wiKVxuIn0%3D))


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





### Strings 


Check and transform strings. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbmV4YW1wbGVfc3RyaW5nIDo9IFwiQnVpbGQgUG9saWN5IGFzIENvZGUgd2l0aCBPUEEhXCJcblxuY2hlY2tfY29udGFpbnMgaWYgY29udGFpbnMoZXhhbXBsZV9zdHJpbmcsIFwiT1BBXCIpXG5jaGVja19zdGFydHN3aXRoIGlmIHN0YXJ0c3dpdGgoZXhhbXBsZV9zdHJpbmcsIFwiQnVpbGRcIilcbmNoZWNrX2VuZHN3aXRoIGlmIGVuZHN3aXRoKGV4YW1wbGVfc3RyaW5nLCBcIiFcIilcbmNoZWNrX3JlcGxhY2UgOj0gcmVwbGFjZShleGFtcGxlX3N0cmluZywgXCJPUEFcIiwgXCJTdHlyYVwiKVxuY2hlY2tfc3ByaW50ZiA6PSBzcHJpbnRmKFwiT1BBIGlzICVzIVwiLCBbXCJhd2Vzb21lXCJdKVxuIn0%3D))


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





### Aggregates 


Summarize data. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbnZhbHMgOj0gWzUsIDEsIDQsIDIsIDNdXG5cbnZhbHNfY291bnQgOj0gY291bnQodmFscylcbnZhbHNfbWF4IDo9IG1heCh2YWxzKVxudmFsc19taW4gOj0gbWluKHZhbHMpXG52YWxzX3NvcnRlZCA6PSBzb3J0KHZhbHMpXG52YWxzX3N1bSA6PSBzdW0odmFscylcbiJ9))


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





### Objects: Extracting Data 


Work with key value and nested data. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbm9iaiA6PSB7XCJ1c2VyaWRcIjogXCIxODQ3MlwiLCBcInJvbGVzXCI6IFt7XCJuYW1lXCI6IFwiYWRtaW5cIn1dfVxuXG4jIHBhdGhzIGNhbiBjb250YWluIGFycmF5IGluZGV4ZXMgdG9vXG52YWwgOj0gb2JqZWN0LmdldChvYmosIFtcInJvbGVzXCIsIDAsIFwibmFtZVwiXSwgXCJtaXNzaW5nXCIpXG5cbmRlZmF1bHRlZF92YWwgOj0gb2JqZWN0LmdldChcblx0b2JqLFxuXHRbXCJyb2xlc1wiLCAwLCBcInBlcm1pc3Npb25zXCJdLCAjIHBhdGhcblx0XCJ1bmtub3duXCIsICMgZGVmYXVsdCBpZiBwYXRoIGlzIG1pc3Npbmdcbilcblxua2V5cyA6PSBvYmplY3Qua2V5cyhvYmopXG4ifQ%3D%3D))


```rego
obj := {"userid": "18472", "roles": [{"name": "admin"}]}

# paths can contain array indexes too
val := object.get(obj, ["roles", 0, "name"], "missing")

defaulted_val := object.get(
	obj,
	["roles", 0, "permissions"], # path
	"unknown", # default if path is missing
)

keys := object.keys(obj)
```


```javascript
// Output
{
  "val": "admin",
  "defaulted_val": "unknown",

  "keys": ["roles", "userid"]
}
```





### Objects: Transforming Data 


Manipulate and make checks on objects. ([Try It](https://play.openpolicyagent.org/?state=eyJpIjoie30iLCJwIjoicGFja2FnZSBjaGVhdFxuXG5pbXBvcnQgZnV0dXJlLmtleXdvcmRzXG5cbnVuaW9uZWQgOj0gb2JqZWN0LnVuaW9uKHtcImZvb1wiOiB0cnVlfSwge1wiYmFyXCI6IGZhbHNlfSlcblxuc3Vic2V0IDo9IG9iamVjdC5zdWJzZXQoXG5cdHtcImZvb1wiOiB0cnVlLCBcImJhclwiOiBmYWxzZX0sXG5cdHtcImZvb1wiOiB0cnVlfSwgIyBzdWJzZXQgb2JqZWN0XG4pXG5cbnJlbW92ZWQgOj0gb2JqZWN0LnJlbW92ZShcblx0e1wiZm9vXCI6IHRydWUsIFwiYmFyXCI6IGZhbHNlfSxcblx0e1wiYmFyXCJ9LCAjIHJlbW92ZSBrZXlzXG4pXG4ifQ%3D%3D))


```rego
unioned := object.union({"foo": true}, {"bar": false})

subset := object.subset(
	{"foo": true, "bar": false},
	{"foo": true}, # subset object
)

removed := object.remove(
	{"foo": true, "bar": false},
	{"bar"}, # remove keys
)
```


```javascript
// Output
{
  "removed": { "foo": true },
  "subset": true,
  "unioned": { "bar": false, "foo": true }
}
```







---

The Rego Cheat Sheet is maintained by [Styra](http://styra.com), the
creators of OPA, and the Styra community. If you have any questions,
suggestions, or would like to get involved, please join us on our
[Slack](https://communityinviter.com/apps/styracommunity/signup).
