#include "hello.h"
#include <stdio.h>

int sayHello(const char *name, char *out) {
    int n;

    n = sprintf(out, "Hello, My name is %s!", name);

    return n;
}

void cHello() {
    printf("Hello from C!\n");
}

void printMessage(char* message) {
    printf("从 Go 语言接收的信息: %s\n", message);
}

int add(int a, int b) {
    return a + b;
}