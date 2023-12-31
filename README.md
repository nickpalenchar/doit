# doit - Run shell commands in various directories using a yaml config

Write a yaml file and run commands in different directories.

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