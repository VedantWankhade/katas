#include<stdio.h>

void reverse(char line[], char reversed[]) {
     
}

int main() {
    char line[] = "hello";

    for (int i = 0; line[i] != '\0'; i++) 
        printf("%c", line[i]);
    
    char reversed[10];

    reverse(line, reversed);

    printf("\n");
    for (int i = 0; reversed[i] != '\0'; i++)
        printf("%c",reversed[i]); 
    return 0;
}
