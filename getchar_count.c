#include<stdio.h>

int main() {

    int count2 = 0;
    while (1) {
        int c = getchar();
        if (c != EOF && c != '\n') 
        count2++;
        else break;
    }
    printf("%d\n",count2); 

    // allow multiline
    int count = 0;
    while (getchar() != EOF)
        count++;
    printf("%d\n",count); 
}
