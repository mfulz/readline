
# Readline - ZSH-like console library in Go

Put demo GIF here

*This demo GIF has been made with a Sliver project client.*


## Introduction

This project started out of the wish to make an enhanced console for a security tool (Sliver, see below).
There are already several readline libraries available in Go ([github.com/chzyer/readline](https://github.com/chzyer/readline), 
and [github.com/lmorg/readline](https://github.com/lmorg/readline)), but being stricter readline implementations, their completion 
system is either limited (the former), or not as advanced as this one (the latter). The latter, however, is the basis I used for this 
library, and all credits due and diserved are in the **Warmests Thanks** section).

This project is not a full command-line application, which means it does not automatically understand and execute any commands.
However, having been developed in a project using the CLI [github.com/jessevdk/go-flags](https://github.com/jessevdk/go-flags) library,
it also includes some default utilities (completers) that are made to work with this library, which I humbly but highly recommend.
Please see the [Wiki](https://github.com/maxlandon/readline/docs/) (or the `examples/` directory) for information on how to use these utilities.

Additionally, this project is not POSIX-compliant in any way, nor it a complete and perfect reimplementation of Z-Shell parsing & completion engine.
Please see the `Project Status` for the list of things that I can support on my own (considering no other maintainers).


## Features Summary

A summarized of features supported by this library is following:

### Input & Editing 
- Vim / Emacs input and editing modes.
- Optional, live-refresh Vim status.
- As compared to [github.com/lmorg/readline](https://github.com/lmorg/readline), a few more patterns for Vim editing.

### Completion engine
- 3 types of completion categories (`Grid`, `List` and `Map`)
- Stackable, combinable completions (completion groups of any type & size can be proposed simultaneously).
- Controlable completion group sizes (if size is greater than completions, the completions will roll automatically)
- Virtual insertion of the current candidate, like in Z-shell.
- In `List` completion groups, ability to have alternative candidates (used for displaying `--long` and `-s` (short) options, with descriptions)
- Completions working anywhere in the input line (your cursor can be anywhere)
- Completions are searchable with *Ctrl-F*, like in lmorg's library.

### Prompt system & Colors
- 1-line and 2-line prompts, both being customizable.
- Function for refreshing the prompt, with optional behavior settings.
- Optional colors (can be disabled).

### Hints & Syntax highlighting
- As borrowed from [github.com/lmorg/readline](https://github.com/lmorg/readline), a hint system. See utilities for a default one.
- Also borrowed for lmorg, a syntax highlighting system. A default one is also available.
- The Hint system is now refreshed depending on the cursor position as well, like completions.

### Command history
- Borrowed from lmorg again, a history system.
- Added, the ability to have 2 different history sources (I used this for clients connected to a server, used by a single user).
- History is searchable like completions.
- Default history is an in-memory list.
- Quick history navigation with *Up*/*Down* arrow keys in Emacs mode, and *j*/*k* keys in Vim mode.

### Utilities
- Default Tab completer, Hint formatter and Syntax highlighter provided, using [github.com/jessevdk/go-flags](https://github.com/jessevdk/go-flags) 
command parser to build themselves. These are in the  `completers/` directory. Please look at the [Wiki page](https://github.com/maxlandon/readline/docs/) 
for how to use them. Also feel free to use them as an inspiration source to make your owns.
- Colors mercilessly copied from [github.com/evilsocket/islazy/](https://github.com/evilsocket/islazy) `tui/` package.
- Also in the `completers` directory, completion functions for environment variables (using Go's std lib for getting them), and dir/file path completions.


## Installation & Usage

As usual with Go, installation:
```
go get github.com/maxlandon/readline

```
Please see either the `examples` directory, or the Wiki for detailed instructions on how to use this library.


## Documentation

The complete documentation for this library can be found in the repo's [Wiki](https://github.com/maxlandon/readline/docs/). Below is the Table of Contents:

- **Input modes**
    - Vim
        - Setup
        - Usage
    - Emacs
        - Setup
        - Usage

- **Prompt system**
    - Vim mode
    - Multiline prompts
    - Custom prompts
    - Prompt refresh

- **Completion system**
    - Completion group types
    - Stacking up completions
    - Completion search
    - Other details & warnings
    - Completion utilities

- **Hint formatter & Syntax highlighter**
    - Live refresh demonstration

- **Command History**
    - Main / alternative sources
    - Navigation & search

- **Command & Completion utilities**
    - Using the library with [github.com/jessevdk/go-flags](https://github.com/jessevdk/go-flags)
    - Default completion engines with go-flags.
    - Colors usage and/or deactivation.


## Project Status & Support

Being alone working on this project and having only one lifetime (anyone able to solve this please call me), I can engage myself over the following:
- Support for any issue opened.
- Answering any questions related.
- Being available for any blame you'd like to make for my humble but passioned work. I don't mind, I need to go up.

Things I do intend to add a more or less foreseeable future:
- A better recursive command/subcommand default completer (see utilities), because the current one supports only `command subcommand --options` patterns, not `command subcommand subsubcommand`.
- A recursive option group completer: tools like `nmap` will use options like `-PA`, or `-sT`, etc. These are not supported.

Therefore, I do not intend to add any other features, as far as I can see. Of course, any good will submitting mockups and a big smile might be considered !


## Warmest Thanks

- First of all, the warmest thanks to Laurence Morgan, aka [lmorg](https://github.com/lmorg) for his [readline](https://github.com/lmorg/readline) library. 
This project could have never been done without the time he dedicated writing the shell in the first place. Please go check his [murex](https://github.com/lmorg/murex) shell as well !
- The [Sliver](https://github.com/BishopFox/sliver) implant framework project, which I used as a basis to make, test and refine this library. as well as all the GIFs and documentation pictures !


## Licences

This library is distributed under the Apache License (Version 2.0, January 2004) (http://www.apache.org/licenses/), similarly to lmorg's readline project.
