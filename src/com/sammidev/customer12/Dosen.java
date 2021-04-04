package com.sammidev.customer12;

public class Dosen extends CivitasAkademik{
    private String nidn;

    public Dosen() {
    }


    public Dosen(String nidn, String nama, String prodi, String fakultas) {
        super(nama, prodi, fakultas);
        this.nidn = nidn;
    }

    public String getNidn() {
        return nidn;
    }

    public void setNidn(String nidn) {
        this.nidn = nidn;
    }
}
