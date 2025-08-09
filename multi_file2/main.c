#include<stdio.h>
#include"add.c"

// just need to compile main.c; becauase add.c is included in main.c
// gcc main.c -o main
int main() {
    int a = add(2, 3);
    printf("%d\n", a);
    return 0;
}
