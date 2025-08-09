#include<stdio.h>
#include"alloc.h"

int main() {
    char *p = alloc(5);
    for (int i = 0; i < 5; i++) {
        printf("%c\t", p[i]);
    }
    printf("\n");
    for (int i = 0; i < 5; i++) {
        p[i] = i + 'A';
    }
    for (int i = 0; i < 5; i++) {
        printf("%c\t", p[i]);
    }
    printf("\n");
    afree(p);
    return 0;
}
