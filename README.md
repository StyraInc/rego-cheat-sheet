# rego-cheat-sheet

This project contains the code to generate the Rego Cheat Sheet,
in both MD and PDF formats.

Content is defined in YAML, Rego, and JSON in [cheats](./cheats),
and rendered to Markdown and PDF in [build](./build).

The cheats dir content is fairly self-explanatory.
Each top level dir is a section, each sub-dir is for a single 'cheat'.

## Development

To render the PDF after having made changes to the cheats, run `make build`.
This requires `tectonic` to be installed (`brew install tectonic`).