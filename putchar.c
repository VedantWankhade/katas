#include<stdio.h>

int main() 
{
    putchar('h');
    putchar('e');
    printf("l"); // cursor is shared between putchar and printf
    
    putchar('\n');
    putchar(65);
    printf("\n"); 
    printf("%d = %c = %x = %b", 65, 65, 65, 65);

    return 0;
}
