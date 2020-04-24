# Literate

A simple literate programming tool.

[From Wikipedia](https://en.wikipedia.org/wiki/Literate_programming):

> Literate programming is a programming paradigm introduced by Donald Knuth in
> which a computer program is given an explanation of its logic in a natural
> language, such as English, interspersed with snippets of macros and
> traditional source code, from which compilable source code can be generated.

This command line tool facilitates using markdown documents (like the one you're
reading!) as literate programs. It reads a specified markdown file, strips out
everything except the code, and writes it to a new file. I might write a
document on how to write "Hello, world" in C, name it `hello_world.c.md`, and
then use `literate hello_world.c.md` to produce the `hello_world.c` source file
which I can then compile. See [examples](/examples) for more.

## Installation

You must first install Go.
```
$ go install github.com/spencer-p/literate
```

## Usage

To produce code from a markdown file:
```
literate my_file.c.md
```
The suffix is not required. However, keep in mind that `literate` will use the
file name as-is with the `.md` trimmed.

## Roadmap

[ ] Ask for confirmation to overwrite files (if not already `literate` output)
[ ] Skip files that are already up-to-date
[ ] Add ability to traverse directories
  [ ] Add ability to parse globs (`*.c.md`)
[ ] Match filename suffix with code fence tags (`.c` files should only have code
blocks that were marked with `\`\`\`c`)
