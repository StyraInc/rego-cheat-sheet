# Rego Cheat Sheet

<!-- The source of truth for this file's contents is https://github.com/StyraInc/rego-cheat-sheet -->

<div style={{`{{display: "none"}}`}}>
```rego
package cheat
import rego.v1
```
<RunSnippet id="preamble.rego"/>
</div>

{{ range .Sections }}
## {{ .Title }} - <sub><sup>{{ .Subtitle }}</sup></sub>

{{ range .Cheats }}
### {{ .Title }} 

{{ if ne .Text "" }}
{{ .Text }} ([Try It]({{ .PlaygroundLink }}))
{{ else }}
[Try It]({{ .PlaygroundLink }})
{{ end }}

{{ if ne .Input "" }}
Input:
```json
{{ .Input }}
```
<RunSnippet id="input.{{.Title | urlquery}}.json"/>
{{ end }}

```rego
{{ .CodeDisplay }}
```

{{ if ne .Input "" }}
<RunSnippet command="data.cheat" files="#input.{{.Title | urlquery}}.json" depends="preamble.rego"/>
{{ else }}
<RunSnippet command="data.cheat" depends="preamble.rego"/>
{{ end }}
{{ end }}
{{ end }}

{{ .Footer }}
