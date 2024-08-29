# Rego Cheat Sheet

<!-- The source of truth for this file's contents is https://github.com/StyraInc/rego-cheat-sheet -->

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
