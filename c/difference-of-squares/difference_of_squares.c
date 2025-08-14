#include "difference_of_squares.h"

unsigned int sum_of_squares(unsigned int number){
    unsigned long square_sum = 0L;
    for (unsigned int i = 1; i <= number; i++) {
        square_sum += i * i;
    }
    return square_sum;
}

unsigned int square_of_sum(unsigned int number) {
    unsigned long sum = 0L;
    for (unsigned int i = 1; i <= number; i++) {
        sum += i;
    }
    return sum * sum;
}
unsigned int difference_of_squares(unsigned int number){
    return (unsigned int) (square_of_sum(number) - sum_of_squares(number));
}
