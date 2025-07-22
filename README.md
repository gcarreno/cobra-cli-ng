# Cobra-CLI Generator Next Generation
[![Build Status](https://github.com/gcarreno/cobra-cli-ng/actions/workflows/main.yaml/badge.svg?branch=main)](https://github.com/gcarreno/cobra-cli-ng/actions)
![Go](https://img.shields.io/github/languages/top/gcarreno/cobra-cli-ng?logo=go)
[![Last Commit](https://img.shields.io/github/last-commit/gcarreno/cobra-cli-ng?logo=github)](https://github.com/gcarreno/cobra-cli-ng)
![Commits Since Last Release](https://img.shields.io/github/commits-since/gcarreno/cobra-cli-ng/latest?logo=github)
![COntributors](https://img.shields.io/github/contributors/gcarreno/cobra-cli-ng?logo=github)
[![License](https://img.shields.io/github/license/gcarreno/cobra-cli-ng?logo=github)](https://github.com/gcarreno/cobra-cli-ng/blob/main/LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/gcarreno/cobra-cli-ng?label=latest%20release&logo=github)](https://github.com/gcarreno/cobra-cli-ng/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/gcarreno/cobra-cli-ng/total?logo=github)](https://github.com/gcarreno/cobra-cli-ng/releases)

> [!IMPORTANT]
> This iteration of the tool draws a lot of inspiration, and some snippets of code, from the original: [cobra-cli](https://github.com/spf13/cobra-cli).

> [!WARNING]
> This is very much a work in progress( WIP ).\
> Please allow me to flesh the whole thing out!

This tool provides a way to quick and easily create boilerplate code for all your commands and sub commands when using [cobra](https://github.com/spf13/cobra) and/or [viper](https://github.com/spf13/viper).

## To Do

1. Use emojis and colour on the output/usage.
2. Add the license features, always with a default of opt-out.
3. Have a rummage through the [original project's issues](https://github.com/spf13/cobra-cli/issues) and see what new features or fixes I can add.
4. Add more complex templates:
  - One that uses `PreRunE` to validate flags/args and `RunE`, in order to enjoy automatic usage output when we detect an error that is usage related.
  - One that exemplifies the use of an hierarchical config structure. Maybe with an `init` command that drops the config file with defaults.
5. ~~Tests~~ DONE.
6. _Moar_ tests!!

## Motivation

I've been using the original [cobra-cli](https://github.com/spf13/cobra-cli) for quite a while now. And I've enjoyed it immensely!!

Alas, the more I used it, but mostly, the more I investigated about less simple usage of [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper), the more I had to customize the current output of [cobra-cli](https://github.com/spf13/cobra-cli).

I recognize this is a **me** problem. But since I'm a programmer, I do have the means/tools to make it better. For me, and for anyone that has felt the same limitations and sees my work as an improvement.

Also, the original project has been inactive for about two years, as of July of 2025, hence my choice to make a new tool and not just contribute to the original.

## Differences/Updates

One of the things that frustrated me was the fact that with the original `init` command we could specify a path: `init [path]`.
This would create the appropriate folder, or folders, we specified. But then, it would make use change to that new path in order to add any new commands. It wasn't aware of the path we just asked it to start our project in.

With that in mind, I decided to have a file called `cobra-cli-ng.json`, track all the _projects_ we can ask the command to initialize. This, of course, implies we get one or more commands that do not exists in the original `cobra-cli` implementation. You can look at the `cobra-cli-ng.json` file as any other `JSON` file created by a package manager. I would strongly advise you to include it in your source versioning system.

---

The other thing that I noticed was the fact that the code that was created is a bit simplistic and narrow minded.

For example: It assumes we don't want to use `RunE` and `PreRunE`. These are a good choices if you want to validate your flags( and your arguments ), and if something is wrong, have `cobra` print the error and the usage.

To mitigate this, in the future, I'll have sub-commands of `projects` to add commands to a particular project.

This will allow something like this:
```
    ▾ app/
        ▾ commands/
            ▾ myserverapp/
                ▾ cmd/
                    commandone.go
                    commandtwo.go
                    root.go
                main.go
            ▾ myservercliapp/
                ▾ cmd
                    commandone.go
                    commandtwo.go                    
                    root.go
                main.go
```

Both those projects will be tracked in the `cobra-cli-ng.json` file for future command addition.

---

Another thing that has been very frustrating is the fact that `cobra-cli` will not ask to replace any file if it exists.

Also, the fact that it imposes the use of a `LICENSE` file and license headers on each source code file, is not kosher.

---

Finally, I'm thinking of providing more than a single option for the code generated.

Not sure **how** I'm going to express that in terms of the flags, but I'll cross that bridge when I get there.

## Install

You can, quickly, install this tool by using the following command:
```console
$ go install github.com/gcarreno/cobra-cli-ng@latest
```

Once installed, you should have `cobra-cli-ng` in your `$GOPATH/bin` folder. You can test if it's available by using this command:
```console
$ command -v cobra-cli-ng
```

## Commands

The available commands are:
- `init [path]`: This will create the main structure for `cobra` to start operating
- `add [command name]`: This will add a new command
- `projects`: This lists the projects( more on this later )

### `init` ( i, initialise, initialize, create )

Usage:
```console
$ cobra-cli-ng init --help

Initialize (cobra-cli-ng init) will create a new application, with the 
appropriate structure for a Cobra-based CLI application.

This init command must be run inside of a go module (please run "go mod init <MODNAME>" first)

USAGE
  cobra-cli-ng init [path] [flags]

ALIASES
  init, i, initialize, initialise, create

FLAGS
  -h, --help   help for init

GLOBAL FLAGS
  -c, --config string   config file (default is $HOME/.cobra-cli-ng.yaml)
```

### `add` ( a, command )

Usage:
```console
$ cobra-cli-ng add --help

Add (cobra-cli-ng add) will create a new command, with the 
appropriate structure for a Cobra-based CLI application, and 
register it to its parent (default rootCmd).

If you want your command to be public, pass in the command name
with an initial uppercase letter.

USAGE
  cobra-cli-ng add [command name] [flags]

ALIASES
  add, a, command

EXAMPLES
  # Adding a new command named server, resulting in a new cmd/server.go
  cobra-cli-ng add server

FLAGS
  -h, --help             help for add
  -p, --project string   a project name

GLOBAL FLAGS
  -c, --config string   config file (default is $HOME/.cobra-cli-ng.yaml)
```

### `projects` ( p ) [VERY WIP]

Usage:
```console
$ cobra-cli-ng projects --help

List (cobra-cli-ng projects) will list all the saved projects in "cobra-cli-ng.json".

USAGE
  cobra-cli-ng projects [flags]

ALIASES
  projects, p

FLAGS
  -h, --help   help for projects

GLOBAL FLAGS
  -c, --config string   config file (default is $HOME/.cobra-cli-ng.yaml)
```