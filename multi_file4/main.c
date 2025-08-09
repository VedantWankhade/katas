#include<stdio.h>

int add(int, int);

// to dynamically link add.c
// create library file for add.c with
// gcc -fPIC -c add.c
// gcc -shared -o libadd.so add.o
// then compile main 
// gcc main.c -L. -ladd -o main
// to run 
// LD_LIBRARY_PATH=. ./main
int main() {
    int a = add(2, 2);
    printf("%d\n", a);
    return 0;
}
