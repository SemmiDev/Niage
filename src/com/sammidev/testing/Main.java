package com.sammidev.testing;

public class Main {
    public static void main(String[] args) {
        String[] a = new String[10];

        a[0] = "sammmidev 1";
        a[1] = "sammmidev 2";

        String b = "sammidev 3";
        for (int i = 0; i < a.length; i++) {
            if (a[i] == null) {
                a[i] = b;
                break;
            }
        }

        for (int i = 0; i < a.length; i++) {
            System.out.println(a[i]);
        }
    }
}