#include<stdio.h>

int main() {

    int c;

    long words = 0, lines = 0;

    while ((c = getchar()) != EOF) {
        if (c == ' ' || c == '\t' || c == '\n') {

        if (c == '\n') {
                lines++;
            }
            c = getchar();
            while (c == ' ' || c == '\t' || c == '\n') {
               c = getchar() ;
                if (c == '\n')
                    lines ++;
            }         
            words++;
        }
    }

    printf("Lines: %ld, Words: %ld", lines, words);
    return 0;
}
