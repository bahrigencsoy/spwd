# shortpath

A Go utility that shortens directory paths for display in shell prompts or other contexts.

## Features

- Shortens long directory paths to at most 3 visible components
- Shows root directory, last 2 directories, and represents middle directories as dots
- Replaces home directory with `~`
- Respects symbolic links (like `pwd` command does)
- Fast and optimized for speed

## Installation

### Build from source

```bash
# Clone the repository
git clone <span><span style="color: rgb(150, 34, 73); font-weight: bold;">&lt;repository-url&gt;</span><span style="color: black; font-weight: normal;">
cd shortpath

# Build and install
make install
```

### Sample BASH prompt

```bash
LIME_YELLOW=$(tput setaf 190)
RED=$(tput setaf 1)
NORMAL=$(tput sgr0)
BRIGHT=$(tput bold)
WHITE=$(tput setaf 7)
export PS1="\[${LIME_YELLOW}\]\$(spwd -g)\[${RED}\]\$(spwd) \[${BRIGHT}\]\[${WHITE}\]$\[${NORMAL}\] "
```

See this [stackoverlow answer](https://stackoverflow.com/a/1703567) for more examples.