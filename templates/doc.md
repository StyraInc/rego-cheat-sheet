# Rego Cheat Sheet

{{ range .Sections }}
## {{ .Title }}

{{ if ne .Shared "" }}
_(Shared Code)_
```rego
{{ .Shared }}
```
{{ end }}

{{ range .Cheats }}
### {{ .Title }}

{{ if ne .Text "" }}
{{ .Text }}
{{ end }}

```rego
{{ .BodyTex }}
```

{{ if ne .Output "" }}
```javascript
// Output
{{ .Output }}
```
{{ end }}

{{ end }}
{{ end }}

