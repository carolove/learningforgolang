#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
    char * ch;
    ch = (char *)malloc(sizeof(char)*12);
    printf("%x\n", ch[11]);
    free(ch);
    ch = "this is world";
    printf("%d\n",ch[11]);
    return -1;
}