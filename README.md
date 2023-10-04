# rego-cheat-sheet

Example project for generating a Rego cheatsheet.

Content is defined in YAML and Rego in [cheats](./cheats), and rendered to Markdown and PDF in [build](./build).

The cheats dir content is fairly self-explanatory. Each top level dir is a section, each sub-dir is a subsection for
a single 'cheat'. Cheats have an optional `output.json` file, sections have an optional `section.rego` file for
shared code.

```azure
cheats
├── comprehensions
│   ├── arrays
│   │   ├── cheat.rego
│   │   ├── cheat.yaml
│   │   └── output.json
│   ├── objects
│   │   ├── cheat.rego
│   │   ├── cheat.yaml
│   │   └── output.json
│   ├── section.rego
│   ├── section.yaml
│   └── sets
│       ├── cheat.rego
│       ├── cheat.yaml
│       └── output.json
└── types
    ├── arrays
    │   ├── cheat.rego
    │   └── cheat.yaml
    ├── numbers
    │   ├── cheat.rego
    │   └── cheat.yaml
    ├── objects
    │   ├── cheat.rego
    │   └── cheat.yaml
    ├── section.yaml
    ├── sets
    │   ├── cheat.rego
    │   └── cheat.yaml
    └── strings
        ├── cheat.rego
        └── cheat.yaml
```