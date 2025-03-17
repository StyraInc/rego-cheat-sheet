# Rego Cheat Sheet

This project contains the code to generate the
[Rego Cheat Sheet](https://docs.styra.com/opa/rego-cheat-sheet),
in both Markdown and [PDF](https://docs.styra.com/rego-cheat-sheet.pdf)
formats.

The cheat sheet is designed to be a quick reference for those,
learning and using Rego.

Rego is the declarative language used by
[Open Policy Agent](https://www.openpolicyagent.org/), the general-purpose
policy engine originally built by [Styra](https://www.styra.com/).

## Development

Content is defined in YAML, Rego, and JSON in [cheats](./cheats),
and rendered to Markdown and PDF in [build](./build) using a simple
Go program.

The cheats dir content is fairly self-explanatory.
Each top level dir is a section, each sub-dir is for a single 'cheat'.

To render the outputs, after having made changes to the cheats,
run `make build`.
This requires `tectonic` to be installed (`brew install tectonic`).

## Community

For questions, discussions and announcements related to Styra products, services and open source projects, please join
the Styra community on [Slack](https://inviter.co/styra)!
