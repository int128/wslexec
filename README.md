# wslexec [![CircleCI](https://circleci.com/gh/int128/wslexec.svg?style=shield)](https://circleci.com/gh/int128/wslexec)

A command line tool to run a command on WSL (Windows Subsystem for Linux).


## Getting Started

Download [the latest release](https://github.com/int128/wslexec/releases).

Rename the binary to command name with prefix `wsl`.
For example, rename the binary to `wslgit` to run the `git` command on WSL.


### Path translation

A Windows path (e.g. `C:\`) in command arguments are translated to a WSL path (e.g. `/mnt/c`).

For example,

```sh
wslgit.exe -C C:\Users\foo\example log
```

is translated to following:

```sh
git -C /mnt/c/Users/foo/example log
```


## Contributions

Feel free to open an issue or pull request.

### Release from CircleCI

Push a tag, then CircleCI will release a build to GitHub.

It requires following environment variable:

- `GITHUB_TOKEN`
