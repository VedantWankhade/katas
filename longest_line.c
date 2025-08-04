#include<stdio.h>

#define MAX_LINE_LENGTH 100

int get_line(char line[], int len) {
    int c;
    int n = 0;
    while ((c = getchar()) != '\n') {
        if (c == EOF)
            return 0;
        line[n] = c;
        n++;
    }
    return n;
}

int main() {
    int longest = 0;

    char longest_line[MAX_LINE_LENGTH], line[MAX_LINE_LENGTH];

    int n = 0;
    while ((n = get_line(line, MAX_LINE_LENGTH)) != 0) {
        if (n > longest) {
            longest = n;
            for (int i = 0; i < MAX_LINE_LENGTH; i++) {
                longest_line[i] = line[i];
            }
        }
    }  
    printf("Longest line: %s\nLength: %d\n",longest_line, longest);
    return 0;
}

