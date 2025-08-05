#include<stdio.h>
#include<string.h>

int atoi(char text[]) {
    int n = strlen(text); 
    int ret = 0;
    int j = 0;
    int _n = n;
    for (int i = n - 1; i >= 0; i--) {
        int cn = text[i] - '0';
        for (int x = 0; x < j; x++) {
            cn *= 10;
        } 
        j++;
        ret += cn;
    } 
    return ret;
}

int main() {

    char text[] = "452";
    int n = atoi(text);
    printf("%d\n", n); 

    return 0;
}

