# Rego Style Guide

{{ range .Sections }}
## {{ .Title }}

{{ if ne .Shared "" }}
_(Shared Code)_
```rego
{{ .Shared }}
```
{{ end }}

{{ range .Cheats }}
## {{ .Title }}

```rego
{{ .BodyTex }}
```

{{ if ne .Output "" }}
_(Output)_
```json
{{ .Output }}
```
{{ end }}

{{ end }}
{{ end }}

