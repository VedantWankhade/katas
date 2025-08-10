#include<stdio.h>

int main() {
    char *filename = "/tmp/test.txt"; 
    FILE *fp = fopen(filename, "w"); 
    int wrote = fprintf(fp, "hello %s", filename);
    printf("wrote: %d\n", wrote);
    fclose(fp);
    return 0;
}
