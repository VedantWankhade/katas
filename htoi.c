#include<stdio.h>
#include<math.h>
#include<string.h>
#include<ctype.h>

/**
*   Convert a string of hexadecinal digits including an optional 0x or 0X into its 
*   equivalent integer value. The digits allowed are 0 thru 9 and a thru f or A thru F.
*/
int htoi(char hex[]) {
    int n = strlen(hex);
    long int ret = 0;
    for (int i = 0; i < n; i++) {
        if (!isdigit(hex[i]))
            hex[i] = tolower(hex[i]);
    }

    // we have hex strign in lowercase now
    int i = 0;
    if (hex[0] == '0' && hex[1] == 'x')
        i = 2;

    int x = 0;
    for (int j = n - 1; j >= i; j--) {
        int y;
        if (!isdigit(hex[j]))
            y = 10 + hex[j] - 'a';
        else y = hex[j] - '0';
        ret += y * pow(16, x++);
    }
    return ret;
}

int main() {
    // htoi("0xAA"); // will give error because we are passing string literal which are contstant
    // pass as arrays
    char a[] = "0XAA";
    long int ai = htoi(a);
    char b[] = "0x5a";
    long int bi = htoi(b);
    printf("%ld\n%ld\n", ai, bi);
    return 0;
}
