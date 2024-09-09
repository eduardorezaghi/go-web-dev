#include <stdio.h>

void doubleValue(int *x) {
    *x *= 2; // Dereference x, then multiply the value it points to by 2
}

int main() {
    int num = 5;
    printf("Initial value of num: %d\n", num); // Output: Initial value of num: 5
    doubleValue(&num); // Pass the address of num
    printf("Doubled value of num: %d\n", num); // Output: Doubled value of num: 10
    return 0;
}