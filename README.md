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
git clone git@github.com:bahrigencsoy/spwd.git
cd spwd

# Build and install
make build && sudo make install
```

### Sample shell profile

```sh
if [[ -n "$BASH_VERSION" ]]; then
LIME_YELLOW=$(tput setaf 190)
RED=$(tput setaf 1)
NORMAL=$(tput sgr0)
BRIGHT=$(tput bold)
WHITE=$(tput setaf 7)
export PS1="\[${LIME_YELLOW}\]\$(spwd -g)\[${RED}\]\$(spwd) \[${BRIGHT}\]\[${WHITE}\]$\[${NORMAL}\] "
# assuming that you installed preexec from https://github.com/rcaloras/bash-preexec
# the execution time of the last command will be stored in $ELAPSED variable
[[ -f ~/.bash-preexec.sh ]] && source ~/.bash-preexec.sh
preexec() { export ___T0_=$(spwd -m) ; }
precmd() { export ELAPSED=$(spwd -d ${___T0_:-1}) ; }
fi
if [[ -n "$ZSH_VERSION" ]]; then
autoload -U colors && colors
setopt PROMPT_SUBST
export PS1="%{$fg[yellow]%}\$(spwd -g)%{$fg[red]%}\$(spwd) %{$fg[white]%}%% %{$reset_color%}"
fi
alias cdr='cd "$(spwd -gf)"'
```

See this [stackoverlow answer](https://stackoverflow.com/a/1703567) for more bash color examples.