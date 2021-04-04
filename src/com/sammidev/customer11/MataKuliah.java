package com.sammidev.customer11;


public class MataKuliah {
    private final String[] kodeMataKuliah = {"IK","SI","CS"};
    private String kode;
    private String nama;
    private int sks;
    private int kapasitas;
    private Mahasiswa[] daftarMahasiswa;

    public MataKuliah(String kode, String nama, int sks, int kapasitas){
        this.kode = kode;
        this.nama = nama;
        this.sks = sks;
        this.kapasitas = kapasitas;

        for (int i = 0; i < kodeMataKuliah.length; i++) {
            if (kodeMataKuliah[i].equalsIgnoreCase(kode)) {
                this.kode = kode;
                break;
            }
        }
    }


    public void addMahasiswa(Mahasiswa mahasiswa) {
        for (int i = 0; i < daftarMahasiswa.length; i++) {
            if (daftarMahasiswa[i] == null) {
                daftarMahasiswa[i] = mahasiswa;
            }
        }
    }

    public void dropMahasiswa(Mahasiswa mahasiswa) {
        for (int i = 0; i < daftarMahasiswa.length; i++) {
            if (daftarMahasiswa[i] == mahasiswa) {
                daftarMahasiswa[i] = null;
            }
        }
    }

    public String toString() {
        /* TODO: implementasikan kode Anda di sini */
        return getNama();
    }

    public String getKode() {
        return kode;
    }

    public void setKode(String kode) {
        this.kode = kode;
    }

    public String getNama() {
        return nama;
    }

    public void setNama(String nama) {
        this.nama = nama;
    }

    public int getSks() {
        return sks;
    }

    public void setSks(int sks) {
        this.sks = sks;
    }

    public int getKapasitas() {
        return kapasitas;
    }

    public void setKapasitas(int kapasitas) {
        this.kapasitas = kapasitas;
    }

    public Mahasiswa[] getDaftarMahasiswa() {
        return daftarMahasiswa;
    }

    public void setDaftarMahasiswa(Mahasiswa[] daftarMahasiswa) {
        this.daftarMahasiswa = daftarMahasiswa;
    }
}
