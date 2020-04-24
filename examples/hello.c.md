# Hello, World!

_Hello, World_ is a fun simple example of how to write program that prints
something out! This example is in C. In order to get anything done in C, you
need some help from the **standard libary**. To print something, we need this
specific incantation:
```c
#include <stdio.h>
```
The `std` part tells us that it's part of the standard library, and the `io`
part stands for **I**input/**O**output. Since we want to *output* text to the
screen, this is what we need.

The entry point of every C program is a function called `main`. Main is supposed
to take a specific form. Here's what it looks like.
```c
int main(void) {
```

Cool.

Now how do we print something? We have to use this little function calld
`printf` that **prints** and **formats** text. We don't need to format anything,
so we'll give it some normal looking text.
```c
printf("Hello, World!\n");
```

Radical! Now all we have to do is close that brace from earlier.
```c
}
```

And that's hello world in C! In order to run this, you need to first call
`literate hello.c.md`, which should produce `hello.c`. Then you can run `cc -o
hello hello.c` to get a program called `hello` that you can run with `./hello`.
