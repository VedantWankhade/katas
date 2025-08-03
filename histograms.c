#include<stdio.h>

#define MAX_TEXT_LENGTH 100
#define POSSIBLE_CHAR_NUMBER 256
#define MAX_WORD_LENGTH 10

void char_frequencies_histogram(char[]);
void word_length_histogram(char[]);

int test() {
    printf("hello");
    // can skip returning a value - it will only give warning
}
int main() {
    int x = test();
    printf("X is %d", x); // garbage value in x
    char  text[MAX_TEXT_LENGTH];
    int c, i = 0;
    
    while ((c = getchar()) != EOF && i < MAX_TEXT_LENGTH)
        text[i++] = c;


    char_frequencies_histogram(text);
    word_length_histogram(text);
}

// print histogram of character frequencies
void char_frequencies_histogram(char text[]) {
    int c_chars[POSSIBLE_CHAR_NUMBER];
    for (int i = 0; i < POSSIBLE_CHAR_NUMBER; i++) {
        c_chars[i] = 0;
    }
    for (int i = 0; i < MAX_TEXT_LENGTH; i++) {
        c_chars[text[i] - '0']++; 
    } 

    for (int i = 0; i < POSSIBLE_CHAR_NUMBER; i++) {
        if (c_chars[i] > 0) {
            printf("%c: ", i + '0');
            for (int j = 0; j < c_chars[i]; j++) {
                printf("*");
            }
            printf("\n");
        } 
    }
}

void word_length_histogram(char text[]) {
    int words_length[MAX_WORD_LENGTH];
    
    for (int i = 0; i < MAX_WORD_LENGTH; i++) {
        words_length[i] = 0;
    }

    printf("-------------------\n");
    int count = 0;
    for (int i = 0; i < MAX_TEXT_LENGTH; i++) {
        if (text[i] == ' ' || text[i] == '\n' || text[i] == '\t') {
            words_length[count]++;
            count = 0;
        } else {
            count++;  
            // printf("%d, %c: %d\n", i, text[i], count);
        }
    }
    printf("--------------------\n");
    for (int i = 0; i < MAX_WORD_LENGTH; i++) {
        if (words_length[i] != 0)
        printf("%d: %d\n", i, words_length[i]);
    }
}


