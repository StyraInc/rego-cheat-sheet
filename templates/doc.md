# Rego Cheat Sheet

{{ range .Sections }}
## {{ .Title }} - <sub><sup>{{ .Subtitle }}</sup></sub>

{{ if ne .Shared "" }}
_(Shared Code)_
```rego
{{ .Shared }}
```
{{ end }}

{{ range .Cheats }}
### {{ .Title }} 

{{ if ne .Text "" }}
{{ .Text }} ([Playground]({{ .PlaygroundLink }}))
{{ else }}
[Playground]({{ .PlaygroundLink }})
{{ end }}

```rego
{{ .CodeDisplay }}
```

{{ if ne .Output "" }}
```javascript
// Output
{{ .Output }}
```
{{ end }}



{{ end }}
{{ end }}

