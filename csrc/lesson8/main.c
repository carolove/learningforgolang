#include <stdio.h>
#include <stdlib.h>

// this is a error way, because return stack addr pointer
char * getstackmemory(void){
    char * str = "abcdefg";
    return str;
}

void dosomecalcatpurefunction( const char * from, char * to, int count) {
    for (int i = 0; i < count; i ++) {
        to[i] = from[i];
    }
    return;
}

void testconst(const char * str, char * const str1) {
    *str1 = '1';
    str++;
    *str1= *str;
    printf("%d\n", *str1);
    return;
}

int main(int argc, char **argv) {
    char * ch, *ch1;
    int count  = 12;
    ch = (char *)malloc(sizeof(char)*12);
    for(int i = 0; i < count ; i++) ch[i] = i;
    ch1 = (char *) malloc(sizeof(char)*12);
    dosomecalcatpurefunction(ch, ch1, count);
    for(int i = 0; i < count ; i++) printf("%d", ch1[i]);
    printf("\n");
    testconst(ch, ch1);
    free(ch);
    free(ch1);
    return -1;
}
