# shortpath

A Go utility that shortens directory paths for display in shell prompts or other contexts.

## Features

- Shortens long directory paths to at most 3 visible components, respecting home directory
- Finds git directories in the parent paths
- Can calculate execution time in microseconds precision
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
# assuming that you installed preexec from https://github.com/rcaloras/bash-preexec
[[ -f ~/.bash-preexec.sh ]] && source ~/.bash-preexec.sh
preexec() { export ___T0_=$(spwd -m) ; }
precmd() { export ELAPSED=$(spwd -d ${___T0_:-1}) ; }
```

See this [stackoverlow answer](https://stackoverflow.com/a/1703567) for more color examples.