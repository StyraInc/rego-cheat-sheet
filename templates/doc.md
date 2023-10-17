# Rego Cheat Sheet

{{ range .Sections }}
## {{ .Title }} - <sub><sup>{{ .Subtitle }}</sup></sub>

{{ range .Cheats }}
### {{ .Title }} 

{{ if ne .Text "" }}
{{ .Text }} ([Try It]({{ .PlaygroundLink }}))
{{ else }}
[Try It]({{ .PlaygroundLink }})
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

{{ .Footer }}
