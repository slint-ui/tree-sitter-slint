<!-- Copyright <C2><A9> SixtyFPS GmbH <info@slint.dev> ; SPDX-License-Identifier: GPL-3.0-only OR LicenseRef-Slint-Royalty-free-1.1 OR LicenseRef-Slint-commercial -->

[![Update from monorepo](https://github.com/slint-ui/tree-sitter-slint/actions/workflows/update_from_monorepo.yaml/badge.svg)](https://github.com/slint-ui/tree-sitter-slint/actions/workflows/update_from_monorepo.yaml)

# tree-sitter support for Slint

> Tree-sitter is a parser generator tool and an incremental parsing library. It
> can build a concrete syntax tree for a source file and efficiently update the
> syntax tree as the source file is edited.

(taken from tree-sitter page)

Use with vim/helix/... other editors.

Queries can be found in the [nvim-treesitter](https://github.com/nvim-treesitter/nvim-treesitter/pull/6027)
plugin to nvim as well as the [helix](https://github.com/helix-editor/helix/pull/9551) editor.

## DO NOT EDIT

This repository is for convenience of editor users only!

This repository regularly takes the grammar from there and generates the
language specific files from it.

Please file bugs in the `slint-ui/slint` repository.

Please go the the `slint-ui/slint` repository to edit this code: It is found
there in `editors/tree-sitter-slint`.

## History

This is the second version of the tree-sitter parser for Slint. The first can be found here:

https://github.com/jrmoulton/tree-sitter-slint

A big thank you to @jrmoulton for creating and maintaining it!
