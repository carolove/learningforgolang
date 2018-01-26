#include <stdio.h>
#include <stdlib.h>


#define ROW 3
#define COL 4

int main(int argc, char ** argv) {
    int **ma, **mb;
    int *ha, *hb;
    int i,j;
    ma = (int **) malloc(sizeof(int *)* ROW);

    if (ma == NULL ){
        printf("malloc failed!\n");
        return -1;
    }

    mb = (int **) malloc(sizeof(int *) * COL);
    if (mb == NULL) {
        free(ma);
        printf("malloc failed!\n");
        return -1;
    }

    for (i = 0; i < ROW; i ++ ){
        *(ma + i) = (int *) malloc (sizeof(int) * COL);
    }

    for (i = 0; i < COL; i ++ ) {
        *(mb + i) = (int *) malloc(sizeof (int) * ROW);
    }

    printf("please input matrix a!\n");
    for (i = 0; i < ROW; i++) {
        for (j = 0; j < COL ; j++) {
            scanf("%d", &ma[i][j]); 
        }
        printf("\n");
    }

    for (i = 0; i < ROW; i++) {
        for (j = 0; j < COL ; j++) {
            mb[j][i] = ma[i][j]; 
        }
    }

    printf("output matrix b: \n");

    for (i = 0; i < COL; i++) {
        for (j = 0; j < ROW ; j++) {
           printf("%8d ", mb[i][j]) ; 
        }
        printf("\n");
    }   

    for (i = 0; i < ROW; i ++ ){
       free(ma[i]);
    }

    for (i = 0; i < COL; i ++ ) {
       free(mb[i]);
    }

    free(ma);
    free(mb);

    return -1;
}