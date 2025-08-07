#include<stdio.h>

int main() {
    int n = 9;
    printf("n\t=\t%08b\n", n);
    printf("n & 8\t=\t%08b\n", (n & 8));

    printf("n << 2\t=\t%08b\n", (n << 2));
    printf("~n\t=\t%08b\n", ~n);
    return 0;
}

