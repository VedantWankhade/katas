#include<stdio.h>

int main() {
    char s[20];
    int i = 0;
    while (i < 20) {
        int c = getchar(); // better to store the returned "character" in int - because getchar can return EOF which is an int
        if (c != '\n' && c != '\r')
            s[i++] = c;
        else break;
    }

    for (int j = 0; j < i; j++) {
        putchar(s[j]);
    }

    printf("\n");

    int eof = EOF; // EOF is in stdio as #define EOF (-1)
    printf("%d\n", eof);

    int n;

    while ((n = getchar()) != EOF)
        putchar(n);
    
    printf("\n");

    return 0;

}
