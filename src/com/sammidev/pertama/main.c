#include <stdio.h>

int main(){
    
    int lengthBO = 3;
    int lengthB1 = 5;

    char A[15] = {'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O'};
    char B[3][5];

    int index = 0;

    for (int i = 0; i < lengthB1; i++) {
        for (int j = 0; j < lengthBO; j++) {
            B[j][i] = A[index];
            index++;
        }
    }

    for (int i = 0; i < lengthBO; i++) {
            for (int j = 0; j < lengthB1; j++) {
                printf("%c", B[i][j]);
            }
            printf("\n");
    }
    return 0;
}
