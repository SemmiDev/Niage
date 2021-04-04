package com.sammidev;

import java.util.Scanner;

public class App {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        Scanner scanner2 = new Scanner(System.in);

        int size;
        System.out.print("");
        size = scanner.nextInt();

        int lengthSizeINT[] = new int[size];
        String[] lengthSizeSTR= new String[size];

        int tempINT;
        String tempSTR;

        for (int i = 0; i < size; i++) {
            System.out.print("");
            tempINT = scanner.nextInt();
            lengthSizeINT[i] = tempINT;

            System.out.print("");
            tempSTR = scanner2.nextLine();
            lengthSizeSTR[i] = tempSTR;
        }

        for (int i = 0; i < lengthSizeINT.length; i++) {
            System.out.println(lengthSizeINT[i]);
            System.out.println(lengthSizeSTR[i]);
        }

        int[] result = new int[size];
        for (int i = 0; i < result.length; i++) {
            String temp = lengthSizeSTR[i];
            String[] res = temp.split(" ");

            if (res.length != lengthSizeINT[i]) {
                result[i] = 0;
                break;
            }else {
                int[] tempor = new int[res.length];

                for (int j = 0; j < res.length; j++) {
                    int tempInt = Integer.parseInt(res[j]);
                    tempor[i] = tempInt;
                }
                int count = 0;
                for (int j = 0; j < tempor.length-1; j++) {
                    for (int k = j+1; k < res.length; k++) {
                        if (tempor[j] > tempor[k]) {
                            System.out.println("checked -> " + tempor[j] + " > " + tempor[k]);
                            count += 1;
                            continue;
                        }
                    }
                }
                result[i] = count;
            }
        }

        for (int r = 0; r < result.length; r++) {
            System.out.println(result[r]);
        }
    }
}
