#include<stdio.h>

int main() {
    int c;
    while ((c = getchar()) != EOF) {
        if (c == '\t') {
            while ((c = getchar()) == '\t') {
                
            }
            putchar(' ');
        }
           putchar(c);
    }
}
