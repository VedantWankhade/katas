#include<stdio.h>

extern int counter;

void inc();

int main() {
    printf("%d\n", counter);
    counter++;
    inc();
    printf("%d\n", counter);
}
