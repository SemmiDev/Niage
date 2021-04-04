package com.sammidev.customer12;

/**
 * Created by sam on 04/04/21.
 */
public class Mahasiswa  extends CivitasAkademik{
    private String nim;

    public Mahasiswa() {}

    public Mahasiswa(String nim, String nama, String prodi, String fakultas) {
        super(nama,prodi,fakultas);
        this.nim = nim;
    }

    public String getNim() {
        return nim;
    }

    public void setNim(String nim) {
        this.nim = nim;
    }
}
