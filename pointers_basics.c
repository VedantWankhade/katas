#include<stdio.h>

int main() {
    
    char c = 'A';
    char *p = &c;
    printf("%x points to %c\n", p, *p);
    return 0;
}
