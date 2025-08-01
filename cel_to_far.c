#include<stdio.h>


// conver cel to far
int far(int cel) {
    return (9/5) * cel + 32;
}

float realFar(float cel) {
    return (9.0/5) * cel + 32;
}

void realNumber() {
    float start = 0;
    float step = 20;
    float max = 300;

    printf("Celsius\tFarenheit\n");
    for (float t = start; t < max; t += step) {
        printf("%6.2f\t%8.2f\n", t, realFar(t));
    }
}

int main() {
    int start = 0;
    int step = 20, max = 300;

    printf("Celsius\tFarenheit\n");

    for (int t = start; t < max; t += step) {
        printf("%4d\t%6d\n", t, far(t));
    } 

    realNumber();
}

