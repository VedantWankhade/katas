#include<stdio.h>
#include"add.h"

// compile both main.c and add.c
int main() {
    int a = add(2, 4);
    printf("%d\n", a);
    return 0;
}
