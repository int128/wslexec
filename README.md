# wslexec [![CircleCI](https://circleci.com/gh/int128/wslexec.svg?style=shield)](https://circleci.com/gh/int128/wslexec)

A wrapper to run a Linux command inside WSL (Windows Subsystem for Linux) from Windows native applications.

If you are looking for Git on WSL solution, [andy-5/wslgit](https://github.com/andy-5/wslgit) is the best.


## Getting Started

### Visual Studio Code and PHP

You can use PHP on WSL from Windows native Visual Studio Code.

Download [the latest release](https://github.com/int128/wslexec/releases) and save as `wslphp.exe`.

Open Visual Studio Code and configure path as follows:

```json
{
  "php.validate.executablePath": "C:\\Users\\USER\\Documents\\wslphp.exe"
}
```

### Visual Studio and Node.js

WIP


## How it works

Rename the binary to command name with prefix `wsl`.
For example, rename the binary to `wslgit` to run the `git` command inside WSL.

### Windows path translation

Windows paths (e.g. `C:\`) in command line are translated to WSL paths (e.g. `/mnt/c`).

For example,

```sh
wslgit.exe -C C:\Users\foo\example log
```

is translated to the following command:

```sh
git -C /mnt/c/Users/foo/example log
```

### WSL path translation

WSL paths (e.g. `/mnt/c`) in standard input are translated to Windows paths (e.g. `C:/`).

For example,

```sh
/mnt/c/Users/foo/example
```

is translated to the following command:

```sh
c:/Users/foo/example
```


## Contributions

Feel free to open issues and pull requests.
