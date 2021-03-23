package com.sammidev.customer11;

public class CellDriver {
    public static void main(String[] args) {
        DuaTipe<String,Integer> duaTipeA = new DuaTipe<String,Integer>();
        duaTipeA.setT1("ini string");
        duaTipeA.setT2(10);

        DuaTipe<Integer,Boolean> duaTipeB = new DuaTipe<Integer, Boolean>();
        duaTipeB.setT1(12);
        duaTipeB.setT2(true);
    }
}