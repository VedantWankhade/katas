#include<stdio.h>

extern int add(int, int);

// have to compile add.c and main.c
// gcc add.c main.c -o main
int main() {
    int a = add(2, 4); 
    printf("%d\n", a);
    return 0;
}
