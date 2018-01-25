#include <stdio.h>
#include <stdlib.h>

int main(int argc, char ** argv){
    char *str;
    str = (char *)malloc(sizeof(char)*6);
    printf("please input string\n");
    scanf("%s", str);
    printf("string is %s\n", str);
    free(str);
    return -1;
}