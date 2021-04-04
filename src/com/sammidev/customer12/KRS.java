package com.sammidev.customer12;

import java.util.Scanner;

public class KRS {
    private Mahasiswa mahasiswa;
    private MataKuliah[] matkul = new MataKuliah[100];
    private int berisi = 0;
    private int totalSKS = 0;

    public KRS() {}

    public void tambahMatkul(MataKuliah m1) {
        if (this.totalSKS != 24) {
            this.totalSKS += m1.getSks();

            for (int i = 0; i < matkul.length; i++) {
                if (matkul[i] == null) {
                    this.matkul[i] = m1;
                    berisi++;
                    System.out.println("Info " + this.mahasiswa.getNim() + "\t: " + "Matkul Berhasil ditambahkan");
                    break;
                }
            }
        }else {
            System.out.println("Info " + this.mahasiswa.getNim() + "\t: " + "Anda mencapai batas SKS");
        }
    }

    public void tampilKRS() {
        System.out.println("" +
                "NIM         \t: " + this.mahasiswa.getNim() + "\n" +
                "NAMA        \t: " + this.mahasiswa.getNama() + "\n" +
                "PRODI       \t: " + this.mahasiswa.getProdi() + "\n" +
                "FAKULTAS    \t: " + this.mahasiswa.getFakultas() + "\n" +
                "MATA KULIAH \t: ");
        for (int i = 0; i < this.berisi; i++) {
            System.out.println("   * " +
                    this.matkul[i].getKodeMatkul() + "-" +
                    this.matkul[i].getNamaMatkul() + "-" +
                    this.matkul[i].getSks() + "sks-" +
                    this.matkul[i].getDosen().getNama()
            );
        }
    }

    public Mahasiswa getMahasiswa() {
        return mahasiswa;
    }

    public void setMahasiswa(Mahasiswa mahasiswa) {
        this.mahasiswa = mahasiswa;
    }

    public MataKuliah[] getMatkul() {
        return matkul;
    }

    public void setMatkul(MataKuliah[] matkul) {
        this.matkul = matkul;
    }

    public static void main(String[] args) {
//        Dosen d1 = new Dosen("123", "Aji Sulasti","TIF","FILKOM");
//        Dosen d2= new Dosen("124", "Suratno","TIF","FILKOM");
//        KRS krsKe1 = new KRS();
//
//        MataKuliah m1 = new MataKuliah("UBR001","PKN", 5, d1);
//        MataKuliah m2 = new MataKuliah("CIE006","PEMDAS", 4, d1);
//        MataKuliah m3 = new MataKuliah("CIE021","JARKOM", 5, d2);
//        MataKuliah m4 = new MataKuliah("CIE005","PEMWEB", 5, d2);
//        MataKuliah m5 = new MataKuliah("CIE005","Pancasila", 5, d1);
//        MataKuliah m6 = new MataKuliah("CIE005","STATISTIKA", 5, d1);
//
//        krsKe1.setMahasiswa(new Mahasiswa("1231321","Ahmad","SI","FILKOM"));
//        krsKe1.tambahMatkul(m1);
//        krsKe1.tambahMatkul(m2);
//        krsKe1.tambahMatkul(m3);
//        krsKe1.tambahMatkul(m4);
//        krsKe1.tambahMatkul(m5);
//        krsKe1.tambahMatkul(m6);
//        krsKe1.tampilKRS();


        // SCANNER APPROACH

        Scanner scanner = new Scanner(System.in);
        Scanner scanner2 = new Scanner(System.in);

        String nidnDosen,namaDosen,prodiDosen,fakultasDosen;
        String nimMhs,namaMhs,prodiMhs,fakultasMhs;

        String kodeMatkul, namaMatkul;
        int sks;

        Dosen d1 = new Dosen(), d2 = new Dosen();
        MataKuliah m1 = new MataKuliah(),m2= new MataKuliah(),m3= new MataKuliah(),m4= new MataKuliah(),m5= new MataKuliah(),m6= new MataKuliah();
        KRS KRS1 = new KRS();
        Mahasiswa mhs;

        System.out.println("---- DOSEN");
        for (int i = 1; i <= 2; i++) {
            System.out.print("Masukkan nidn dosen     " + i + " : ");
            nidnDosen = scanner.nextLine();
            System.out.print("Masukkan nama dosen     " +  + i + " : ");
            namaDosen = scanner.nextLine();
            System.out.print("Masukkan prodi dosen    "  + i + " : ");
            prodiDosen = scanner.nextLine();
            System.out.print("Masukkan fakultas dosen "   + i + " : ");
            fakultasDosen = scanner.nextLine();
            System.out.println("\n\n");

            if (i == 1) {
                d1.setNidn(nidnDosen);
                d1.setNama(namaDosen);
                d1.setProdi(prodiDosen);
                d1.setFakultas(fakultasDosen);
            }else {
                d2.setNidn(nidnDosen);
                d2.setNama(namaDosen);
                d2.setProdi(prodiDosen);
                d2.setFakultas(fakultasDosen);
            }
        }

        System.out.println("---- Mata Kuliah");
        for (int i = 1; i <= 6; i++) {
            System.out.print("Masukkan kode matkul        " + i + " : ");
            kodeMatkul = scanner.nextLine();
            System.out.print("Masukkan nama matkul        " + i + " : ");
            namaMatkul = scanner.nextLine();
            System.out.print("Masukkan jumlah sks         " + i + " : ");
            sks = scanner2.nextInt();
            System.out.println("\n\n");

            if (i == 1) {
                m1.setKodeMatkul(kodeMatkul);
                m1.setNamaMatkul(namaMatkul);
                m1.setSks(sks);
                m1.setDosen(d1);
            }else if(i == 2){
                m2.setKodeMatkul(kodeMatkul);
                m2.setNamaMatkul(namaMatkul);
                m2.setSks(sks);
                m2.setDosen(d1);
            }else if(i == 3){
                m3.setKodeMatkul(kodeMatkul);
                m3.setNamaMatkul(namaMatkul);
                m3.setSks(sks);
                m3.setDosen(d2);
            }else if(i == 4){
                m4.setKodeMatkul(kodeMatkul);
                m4.setNamaMatkul(namaMatkul);
                m4.setSks(sks);
                m4.setDosen(d2);
            }else if(i == 5){
                m5.setKodeMatkul(kodeMatkul);
                m5.setNamaMatkul(namaMatkul);
                m5.setSks(sks);
                m5.setDosen(d1);
            }else if(i == 6){
                m6.setKodeMatkul(kodeMatkul);
                m6.setNamaMatkul(namaMatkul);
                m6.setSks(sks);
                m6.setDosen(d1);
            }else {}
        }


        System.out.println("---- Mahasiswa");
        System.out.print("Masukkan NIM      mahasiswa : ");
        nimMhs = scanner.nextLine();
        System.out.print("Masukkan nama     mahasiswa : ");
        namaMhs = scanner.nextLine();
        System.out.print("Masukkan prodi    mahasiswa : ");
        prodiMhs= scanner.nextLine();
        System.out.print("Masukkan fakultas mahasiswa : ");
        fakultasMhs= scanner.nextLine();

        mhs = new Mahasiswa(nimMhs,namaMhs, prodiMhs, fakultasMhs);
        System.out.println("\n\n");

        KRS1.setMahasiswa(mhs);
        KRS1.tambahMatkul(m1);
        KRS1.tambahMatkul(m2);
        KRS1.tambahMatkul(m3);
        KRS1.tambahMatkul(m4);
        KRS1.tambahMatkul(m5);
        KRS1.tambahMatkul(m6);
        KRS1.tampilKRS();
    }
}