#include <stdio.h>
#include <stdlib.h>

char * getstackmemory(void){
    char * str = "abcdefg";
    return str;
}

int main(int argc, char **argv) {
    char * ch;
    ch = (char *)malloc(sizeof(char)*12);
    printf("%x\n", ch[11]);
    free(ch);
    ch = "this is world";
    printf("%d\n",ch[11]);

    char *str = getstackmemory();
    printf("%d\n",str[2]);
    return -1;
}
