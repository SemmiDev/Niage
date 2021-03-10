package com.sammidev.customer1;

public class App {
    public static void main(String[] args) {
        char[] A = {'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O'};
        char[][] B = new char[3][5];

        // 00 01 02 03  10 11 12 13  20 21 22 23
        // {A,D,G,J,M}  {B,E,H,K,N}  {C,F,I,L,O}

        int index = 0;
        for (int i = 0; i < B[0].length; i++) {
            for (int j = 0; j < B.length; j++) {
                B[j][i] = A[index];
                index++;
            }
        }
        for (int i = 0; i < B.length; i++) {
            for (int j = 0; j < B[0].length; j++) {
                System.out.print(B[i][j] + " ");
            }
            System.out.println("\n");
        }
    }
}