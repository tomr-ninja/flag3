# flag3 (flag-tree) — hierarchical commands for CLIs and more

For developing simple CLI tools Go standard library has the 'flag' package, and it's perfectly
fine as long as your tool, which is a single binary, has just one command.

If you want to have multiple commands and subcommands, like `git` or `docker`:

```shell
docker --log-level=debug run -p 8080:80 --name=web nginx
```

You still technically can use the `flag` package with some custom `os.Args` processing (as it's
described on [GoByExample](https://gobyexample.com/command-line-subcommands)), but it's not very
convenient, especially when you introduce not only level-0 and level-1, but level-2 and so on
subcommands. 

**flag3** brings hierarchy to your CLI commands, while doing very little magic on top of the
`flag` library and keeping your code imperative (as opposed to declarative style of 'frameworks'
like [cobra](https://github.com/spf13/cobra)).

## Usage

Imagine you have to implement the `docker .. run ..` thing that served as an example above.

### Step 1: declare a command-tree

```go
tree := flag3.NewCLI() // no need to specify the root 'docker' command itself because it's the name of the executable
tree.Subcommand("run")
tree.Subcommand("image").Subcommand("prune") // 'docker image prune' is an example of level-2 command
// ... some other commands as well ...
```

Using a tree you declare all the valid 'command paths'. For example, `docker image prune` is
a valid command, but `docker run prune` is not ('prune' would be considered an image name rather
than a subcommand).

### Step 2: parse your os.Args

```go
root, next, err := flag3.ParseCLI(tree)
```

Now, given that your executable is named `docker`, and you do exactly this:

```shell
docker --log-level=debug run -p 8080:80 --name=web nginx
```

your `root.Command` is gonna be 'docker', and `root.Args` is `['--log-level=debug']`.

After you call `next()`:

```go
command, ok := next()
```

your `command.Command` is 'run', and `command.Args` is `['-p', '8080:80', '--name=web', 'nginx']`.

So basically what flag3 does for you is just splitting os.Args meaningfully and defining what is
the command and what is it's arguments in each fragment. You still might want to use a
`flag.FlagSet` to actually parse the flags, of course.

For a more complete example, see `example/cli`.

## Is it bound to os.Args?

No.

`NewCLI()` and `ParseCLI` are just convenient functions for the most common scenario, which is
a CLI app.

By calling `New('mytree')` you can provide any name for the root node instead of os.Args[0], and
`Parse(args, trees...)` allows you to pass any strings array of arguments and multiple trees to match.
