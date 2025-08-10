#include<stdio.h>
#include<stdarg.h>

void print_count_and_types(int count, ...) {
    printf("Count: %d\n", count);
    printf("Types:\n");

    va_list ap;
    va_start(ap, count);
    int arg1 = va_arg(ap, int);
    int arg2 = va_arg(ap, int);
    double arg3 = va_arg(ap, double);
    printf("%d\t%d\t%f\n", arg1, arg2, arg3);
    va_end(ap);
}

int main() {
    int a = 99;
    double b = 22.11;
    print_count_and_types(3, 22, a, b);
}

