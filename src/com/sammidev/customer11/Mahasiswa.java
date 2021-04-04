package com.sammidev.customer11;


public class Mahasiswa {
    private MataKuliah[] mataKuliah = new MataKuliah[10];
    private String[] masalahIRS = new String[19];
    private int totalSKS;
    private String nama;
    private String jurusan;
    private long npm;

    public Mahasiswa(String nama, long npm){
        this.nama = nama;
        this.npm = npm;
    }

    public void setJurusan(String jurusan) {
        String at = String.valueOf(npm);
        char att = at.charAt(3);
        if (String.valueOf(att).equalsIgnoreCase("1")) {
            this.jurusan = "Ikmu Komputer";
        }else if (String.valueOf(att).equalsIgnoreCase("2")) {
            this.jurusan = "Sistem Informasi";
        }else {
            this.jurusan = "not found";
        }
    }

    public void addMatkul(MataKuliah mataKuliah){
        for (int i = 0; i < this.mataKuliah.length; i++) {
            if (this.mataKuliah[i] == null) {
                this.mataKuliah[i] = mataKuliah;
                break;
            }
        }
    }

    public void dropMatkul(MataKuliah mataKuliah){
        for (int i = 0; i < this.mataKuliah.length; i++) {
            if (this.mataKuliah[i] == mataKuliah) {
                this.mataKuliah[i] = null;
                break;
            }
        }
    }


    public void cekIRS(){
        /* TODO: implementasikan kode Anda di sini */
    }

    public String toString() {
        return this.nama;
    }

    public MataKuliah[] getMataKuliah() {
        return mataKuliah;
    }

    public void setMataKuliah(MataKuliah[] mataKuliah) {
        this.mataKuliah = mataKuliah;
    }

    public String[] getMasalahIRS() {
        return masalahIRS;
    }

    public void setMasalahIRS(String[] masalahIRS) {
        this.masalahIRS = masalahIRS;
    }

    public int getTotalSKS() {
        return totalSKS;
    }

    public void setTotalSKS(int totalSKS) {
        this.totalSKS = totalSKS;
    }

    public String getNama() {
        return nama;
    }

    public void setNama(String nama) {
        this.nama = nama;
    }

    public String getJurusan() {
        return jurusan;
    }

    public long getNpm() {
        return npm;
    }

    public void setNpm(long npm) {
        this.npm = npm;
    }
}