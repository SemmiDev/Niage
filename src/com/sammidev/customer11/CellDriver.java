// package com.sammidev.customer11;

public class CellDriver {
    public static void main(String[] args) {
        DuaTipe<String,Integer> duaTipeA = new DuaTipe<String,Integer>();
        duaTipeA.setT1("ini string");
        duaTipeA.setT2(10);

        DuaTipe<Integer,Boolean> duaTipeB = new DuaTipe<Integer, Boolean>();
        duaTipeB.setT1(12);
        duaTipeB.setT2(true);

        System.out.println(duaTipeA.getT1());
    }
}

class DuaTipe<T1,T2> {

    private T1 t1;
    private T2 t2;

    public T1 getT1() {
        return t1;
    }

    public void setT1(T1 t1) {
        this.t1 = t1;
    }

    public T2 getT2() {
        return t2;
    }

    public void setT2(T2 t2) {
        this.t2 = t2;
    }
}