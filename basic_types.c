#include<stdio.h>
#include<limits.h>
#include<float.h>


void can_i_change_arr(const int arr[]) {
    // arr[0] = 99; // no i cannot
}
int main() {
    
    char c;
    int i;
    float f;
    double d = 22.222222222;
    printf("Size of char: %ld byte\n", sizeof c);
    printf("Size of int: %ld byte\n", sizeof i);
    printf("Size of float: %ld byte\n", sizeof f);
    printf("Size of double: %ld byte\n", sizeof d);

    // with qualifiers
    short int si;
    long int li;
    long double ld;
    printf("Size of short int: %ld byte\n", sizeof si);
    printf("Size of long int: %ld byte\n", sizeof li);
    printf("Size of long double: %ld byte\n", sizeof ld);

    //signed vs unsiged
    int i2 = -124;
    unsigned int ui2 = -231; // unacceptable value stored
    printf("%d\n", i2);
    printf("%u\n", i2); // wrong value displayed
    printf("%d\n", ui2); // wrong value displayed
    printf("%u\n", ui2); // correct value displayed - but thats actually the wrong value stored :(
    
    ui2 = 222; // acceptable value stored
    printf("%d\n", ui2);
    printf("%u\n", ui2); // correct value displayed
    
    // limits
    int li1 = 2147483647;
    printf("int max value: %d = %x\n", li1, li1);
    int li2 = 2147483648;
    printf("int max value: %d = %x\n", li2, li2);

    unsigned int li3 = 4294967295;
    printf("unsigned int max value: %u = %x\n", li3, li3);
    unsigned int li4 = 4294967296;
    printf("unsigned int max value: %u = %x\n", li4, li4);

    printf("Integer min: %d; max: %ld\n", INT_MIN, INT_MAX);
    printf("Character min: %d; max: %ld\n", CHAR_MIN, CHAR_MAX);
    printf("Float min: %f; max: %lf\n", FLT_MIN, FLT_MAX);



    return 0;
}
