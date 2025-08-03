#include<stdio.h>

int main() {
    int c, c_others = 0, c_wspaces = 0;
    int c_digits[10];
    for (int i = 0; i < 10; i++) {
        c_digits[i] = 0;
    }
    while ((c = getchar()) != EOF) {
        if (c <= '9' && c >='0') {
            c_digits[c - '0']++;  
        } else if(c =='\n' || c == '\t' || c==' '){
            c_wspaces++;
        } else {
            c_others++;
        }
    }

    long total_digits = 0;

    for (int i = 0; i < 10; i++) {
        total_digits += c_digits[i];
    }

    printf("\nFollowing is the number of different characters:\n");
    printf("Whitespaces: %d\n", c_wspaces);
    printf("Other characters: %d\n", c_others);
    printf("Total Digits: %ld\n", total_digits);
    printf("Individual digit count:\n");
    
    for (int i = 0; i < 10; i++) {
        printf("%d:%d\n", i, c_digits[i]);
    }
}
