package com.sammidev.customer10;

import java.util.Scanner;

public class PinjamOnline {
    private double saldoSistem = 5000000;
    private double pinjaman = 0;
    public String nama;
    private int lamaPeminjaman = 0;

    public PinjamOnline(String nama) {
        this.nama = nama;
    }

    public double getSaldoSistem() {
        return saldoSistem;
    }

    public void setSaldoSistem(double saldoSistem) {
        this.saldoSistem = saldoSistem;
    }

    public void setPinjaman(double pinjaman) {
        this.pinjaman = pinjaman;
    }

    public double getPinjaman() {
        return pinjaman;
    }

    public void setLamaPeminjaman(int lamaPeminjaman) {
        this.lamaPeminjaman = lamaPeminjaman;
    }

    public int getLamaPeminjaman() {
        return lamaPeminjaman;
    }

    public void pinjam() {
        System.out.println("Selamat datang, " + this.nama);
        Scanner scanner = new Scanner(System.in);

        System.out.println("Silahkan masukkan nominal uang yang ingin dipinjam : ");
        this.pinjaman = scanner.nextDouble();

        if (this.pinjaman > this.saldoSistem) {
            System.out.println("Mohon maaf, anda tidak bisa meminjam melebihi kapasitas saldo sistem");
            return;
        }

        if (this.saldoSistem != 5000000 && this.pinjaman != 0) {
            System.out.println("Mohon maaf, anda harus melunasi peminjaman terlebih dahulu");
            return;
        }

        System.out.println("Silahkan masukkan lama peminjaman : ");
        this.lamaPeminjaman = scanner.nextInt();

        System.out.println("Pinjaman anda telah berhasil");
        System.out.println("============================");
    }

    public void kembalikan() {

    }
}