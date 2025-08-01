#include<stdio.h>

#define MESSAGE "hello"
#define VAR int bro
#define VAR_NAME bro

int main() {
    printf(MESSAGE);
    printf("\n");

    VAR = 22;
    printf("%d\n", bro); // interesting that the lsp knows what bro is, even though it will be declared after the replacement!!
    printf("%d", VAR_NAME);

    VAR_NAME = 99;
    printf("\n%d", VAR_NAME);
    return 0;
}
