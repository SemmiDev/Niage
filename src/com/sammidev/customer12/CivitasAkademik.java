package com.sammidev.customer12;

public class CivitasAkademik {

    private String nama;
    private String prodi;
    private String fakultas;

    public CivitasAkademik() {}

    public CivitasAkademik(String nama, String prodi, String fakultas) {
        this.nama = nama;
        this.prodi = prodi;
        this.fakultas = fakultas;
    }

    public String getNama() {
        return nama;
    }

    public void setNama(String nama) {
        this.nama = nama;
    }

    public String getProdi() {
        return prodi;
    }

    public void setProdi(String prodi) {
        this.prodi = prodi;
    }

    public String getFakultas() {
        return fakultas;
    }

    public void setFakultas(String fakultas) {
        this.fakultas = fakultas;
    }
}
