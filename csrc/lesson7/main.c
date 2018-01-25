#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
    int *a;
    a = (int *) malloc(sizeof(int)*12);
    for (int i = 0;i< 12;i++) {
        printf("please input the %dth number\n", i+1);
        scanf("%d", &a[i]);
    }
    for(int i=0; i<12; i++) 
        printf("%d", a[i]);
    free(a);
    return -1;
}