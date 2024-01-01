# doit - Run shell commands in various directories using a yaml config

Write a yaml file and run commands in different directories.

# Install

## Build From Source

You will need `go`
```
# clone the repository, move to the root of the repository
git clone git@github.com:nickpalenchar/doit.git && cd doit

# when the program is ran from the repo, it builds and
# installs itself
go run cmd/main.go
```

# Getting Started

`doit` is single command cli that looks for a file `doit.yml`, and runs commands based on the specification.

## Step 1: Set up the config file.

Create a file called `doit.yml` and add a map with the key `__MAIN__`.

```yaml
__MAIN__:
  // todo
```

`doit` programs are made up of **directives** defined under `__MAIN__`. These directives have
**instructions** (usually shell commands), that are ran in accordance to the directive it's under.

That's a mouthful, let's start with the directive `IN <path>`

```yaml
__MAIN__:
  IN .:
```

Directives have a keyword followed by arguments (space separated). In this example, we're saying
"In the current directory" (Using the unix `.` that.)

`IN <path>` uses paths relative to where the `doit.yml` file is.

Now to add instructions, the simplest one is a string, which will run as a shell command.

```yaml
__MAIN__:
  IN .:
   - pwd
```

This program prints the current working directory (since it'll be ran in `.`, the directory we are also in).

Let's add a few more `IN` directives to see how the `pwd` instruction is affected:

```yaml
__MAIN__:
  IN .:
   - 'echo Current directory:'
   - pwd
  
  IN ..:
    - 'echo Parent directory:'
    - pwd
  
  IN /usr/bin:
    - 'echo a bin directory:'
    - pwd
```

Each of the commands is ran from the directory specified by `IN`. This is the premise of `doit`: run commands from different
directories without changing the one you're currently in.


```yaml
// doit.yml

// put everything under here:
__MAIN__:
  
  // Each IN statement precedes a directory which commands will
  // be defined and ran under
  IN ./firstDir:
    // list of commands. This one runs npm within `firstDir` directory
    - npm run build

  IN ./secondDir:
    - npm run build

  IN ./:
    // You can add as many commands as you want
    - git add .
    - git commit -m "$(date)"
    // commands can be written as a key/value pair, with special key
    // characters effecting behavior:

    // ? - the doit script will keep running even if this command
    // exits non-zero
    - ?: grep -i 'helloworld' .
```

Run `doit` in the same directory `doit.yml` is saved

# To do
- [x] Notating commands where non-0 should not exit the whole program
- [ ] running commands (directives?) in parallel
- [ ] `CMD` directory for defining subcommands (`doit mycmd`)
- [ ] glob support (`IN ./foo*`)