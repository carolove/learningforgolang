#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
    char * ch;
    ch = (char *)malloc(sizeof(char)*12);
    printf("%x\n", &ch[0]);
    free(ch);
    ch = "this is world";
    printf("%x\n",&ch[0]);
    return -1;
}