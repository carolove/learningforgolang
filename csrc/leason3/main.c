#include <stdio.h>

int main(int argc, char **argv) {
    int x, y, m;
    printf("please input x, y\n");
    scanf("%d%d", &x, &y);
    m = x*y;
    printf("x:%d, y:%d, m:%d\n", x,y,m);
}