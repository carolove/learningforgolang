#include <stdio.h>

int main(int argc, char **argv){
    int a, b,c;
    printf("please input a,b\n");
    scanf("%d, %d", &a, &b);
    c = a >= b?a:b;
    printf("a,b max is %d\n", c);
    return -1;
}