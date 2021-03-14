package com.sammidev.customer6;

public class App {
    public static void main(String[] args) {
        Student student = new Student("08982983", "Ghiyatsi", 80.0, 90.0, 87.0);
        student.hitungNilai();
        student.cetakNilai();
    }
}

class Student {
    String nim;
    String nama;
    Double nilaiTugas;
    Double nilaiUts;
    Double nilaiUas;
    private Double pNilaiUts;
    private Double pNilaiTugas;
    private Double pNilaiUas;
    private Double nilaiAkhir;

    public Student(String nim, String nama, Double nilaiTugas, Double nilaiUts, Double nilaiUas) {
        this.nim = nim;
        this.nama = nama;
        this.nilaiUts = nilaiUts;
        this.nilaiTugas = nilaiTugas;
        this.nilaiUas = nilaiUas;
    }

    void hitungNilai() {
        this.pNilaiTugas = 0.2  * this.nilaiTugas;
        this.pNilaiUts   = 0.35 * this.nilaiUts;
        this.pNilaiUas   = 0.45 * this.nilaiUas;
        this.nilaiAkhir  = this.pNilaiUts + this.pNilaiTugas + this.pNilaiUas;
    }

    void cetakNilai() {
        System.out.println(
                "NIM         : " + this.nim + "\n" +
                "Nama        : " + this.nama + "\n" +
                "Nilai Tugas : " + this.nilaiTugas + " 20%" + "     : " + this.pNilaiTugas + "\n" +
                "Nilai UTS   : " + this.nilaiUts   + " 25%" + "     : " + this.pNilaiUts   + "\n" +
                "Nilai UAS   : " + this.nilaiUas   + " 45%" + "     : " + this.pNilaiUas   + "\n" +
                "Nilai Akhir : " + this.nilaiAkhir);
    }
}
