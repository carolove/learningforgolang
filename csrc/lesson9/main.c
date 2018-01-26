#include <stdio.h>
#include <stdlib.h>


#define ROW 3
#define COL 4

int main(int argc, char ** argv) {
    int **Ma, **Mb;
    int *Ha, *Hb;
    int i,j;
    Ma = (int **) malloc(sizeof(int *) * ROW);
    
    if (Ma == NULL) {
        printf("malloc failed~\n");
        return -1;
    }

    Mb = (int **) malloc(sizeof(int *) * COL);

    if (Mb == NULL) {
        free(Ma);
        printf("malloc failed~\n");
        return -1;
    }
    
    for (i = 0 ; i < ROW; i ++ ) {
        *(Ma + i) = (int *)malloc(sizeof(int)*COL);
    }

    for (i = 0 ; i < COL; i ++ ) {
        *(Mb + i) = (int *)malloc(sizeof(int)*ROW);
    }

    printf("input Matrix a's elements\n");
    for (i = 0; i < ROW ; i ++ ) {
        for (j = 0; j < COL; j ++) {
            scanf("%d", &Ma[i][j]);
        }
    }

    for (i = 0; i < ROW ; i ++ ) {
        for (j = 0; j < COL; j ++) {
            Mb[j][i] = Ma[i][j];
        }
    }

    for (i = 0; i < COL ; i ++ ) {
        for (j = 0; j < ROW; j ++) {
            printf("%8d ", Mb[i][j]);
        }
        printf("\n");
    }
    for (i = 0 ; i < ROW; i ++ ) {
        free(Ma[i]);
    }

    for (i = 0 ; i < COL; i ++ ) {
        free(Mb[i]);
    }
    free(Ma);
    free(Mb);
    return -1;
}