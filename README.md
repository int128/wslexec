# wslexec [![CircleCI](https://circleci.com/gh/int128/wslexec.svg?style=shield)](https://circleci.com/gh/int128/wslexec)

A command line tool to run a command on Bash on Windows (WSL; Windows Subsystem for Linux) without spawning a command prompt window.


## Why

Bash on Windows usually runs in a command prompt window.
It is inconvenience for X window applications such as Xterm or GNOME terminal.
`wslexec` allows running a command without spawning any command prompt window.

See also [Stack Overflow: WSL run linux from windows without spawning a cmd-window](https://stackoverflow.com/questions/41225711/wsl-run-linux-from-windows-without-spawning-a-cmd-window).


## Getting Started

Download [the latest release](https://github.com/int128/wslexec/releases).

Create a shortcut link with the following command line:

```sh
wslexec xterm
```

It will open a new terminal without any command prompt window.

If you are using zsh, run following one:

```sh
wslexec xterm -e zsh -l
```

`wslexec` accepts below options.

Name | Value | Defaults to
-----|-------|------------
`display` | `DISPLAY` environment variable | `:0`


## How it works

`wslexec` just runs `bash -c` in background.


## Contributions

Feel free to open an issue or pull request.

### Release from CircleCI

Push a tag, then CircleCI will release a build to GitHub.

It requires following environment variable:

- `GITHUB_TOKEN`
