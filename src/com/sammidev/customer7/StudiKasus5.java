package com.sammidev.customer7;

import java.util.Scanner;

/**
 * Created by sam on 16/03/21.
 */
public class StudiKasus5 {
    public static void main(String[] args) {
        boolean isNext = true;
        while (isNext) {

            Scanner scannerString = new Scanner(System.in);
            Scanner scannerInt = new Scanner(System.in);
            Scanner scannerDouble = new Scanner(System.in);

            String NamaUniv, Alamat, NoTelepon, PimpinanUniv, NIM, NamaMHS;
            int KodeFakultas, KodeProdi;
            double UTS, UAS, Tugas;

            System.out.print("Masukkan Nama Univ   : ");
            NamaUniv = scannerString.nextLine();

            System.out.print("Masukkan Alamat      : ");
            Alamat = scannerString.nextLine();

            System.out.print("Masukkan No Telp     : ");
            NoTelepon = scannerString.nextLine();

            System.out.print("Masukkan nama pimpinan univ   : ");
            PimpinanUniv = scannerString.nextLine();

            System.out.print("Masukkan kode fakultas: \n" +
                    "1. FTD\n" +
                    "2. HUmaniora\n");
            KodeFakultas = scannerInt.nextInt();

            System.out.print("Masukkan kode prodi: \n" +
                    "1. Informatika\n" +
                    "2. SI\n" +
                    "3. BD\n");
            KodeProdi = scannerInt.nextInt();

            System.out.print("Masukkan NIM: ");
            NIM = scannerString.nextLine();

            System.out.print("Masukkan nama mahasiswa: ");
            NamaMHS = scannerString.nextLine();

            System.out.print("Masukkan nilai UTS: ");
            UTS = scannerDouble.nextDouble();

            System.out.print("Masukkan nilai UAS: ");
            UAS = scannerDouble.nextDouble();

            System.out.print("Masukkan nilai Tugas: ");
            Tugas = scannerDouble.nextDouble();

            Prodi prodi = new Prodi();
            prodi.NamaUniv = NamaUniv;
            prodi.Alamat = Alamat;
            prodi.NoTelepon = NoTelepon;
            prodi.PimpinanUniv = PimpinanUniv;
            prodi.KodeFakultas = KodeFakultas;
            prodi.KodeProdi = KodeProdi;
            prodi.NIM = NIM;
            prodi.NamaMhs = NamaMHS;
            prodi.UTS = UTS;
            prodi.setUAS(UAS);
            prodi.setTugas(Tugas);
            prodi.GradeNilai();

            System.out.println("\n");

            System.out.println("Mau coba lagi? (Y/T)");
            String next = scannerString.next();
            if (next.equalsIgnoreCase("Y")) {
                isNext = true;
            }else if(next.equalsIgnoreCase("T")) {
                isNext = false;
            }else {
                isNext = false;
            }
        }
    }
}

class Universitas {
    public String NamaUniv;
    public String Alamat;
    public String NoTelepon;
    public String PimpinanUniv;
}

class Fakultas extends Universitas{
    public int KodeFakultas;
    public String namaFakultas() {
        if (this.KodeFakultas == 1) {
            return "FTD";
        }else if (this.KodeFakultas == 2) {
            return "Humaniora";
        }
        return "NOT FOUND";
    }
}

class Prodi extends Fakultas{
    public int KodeProdi;
    public String NIM;
    public String NamaMhs;
    public  Double UTS;
    private Double UAS;
    private Double Tugas;

    public void setUAS(Double UAS) {
        this.UAS = UAS;
    }

    public void setTugas(Double tugas) {
        Tugas = tugas;
    }

    public String NamaProdi() {
        if (this.KodeProdi == 1) {
            return "Informatika";
        }else if (this.KodeProdi == 2) {
            return "SI";
        }else if (this.KodeProdi == 3) {
            return "BD";
        }
        return "NOT FOUND";
    }

    public Double TotalNilai() {
        return (this.UTS * 0.3) + (this.UAS * 0.5) + (this.Tugas * 0.2);
    }

    public void GradeNilai() {
        String grade = "";
        if (TotalNilai() >= 80) {
            grade = "Grade = A";
        }else if (TotalNilai() >= 60) {
            grade = "Grade = B";
        }else if (TotalNilai() >= 40) {
            grade = "Grade = C";
        }else if (TotalNilai() >= 20) {
            grade = "Grade = D";
        }else if (TotalNilai() < 20) {
            grade = "Grade = E";
        }

        System.out.println("\n\n\nNama Univ    : " + this.NamaUniv);
        System.out.println("Alamat  : " + this.NamaUniv);
        System.out.println("No Telp : " + this.NamaUniv);
        System.out.println("Nama Pimpinan Univ : " + this.NamaUniv);
        System.out.println("Kode Fakultas : " + KodeFakultas);
        System.out.println("Nama fakultas : " + namaFakultas());
        System.out.println("Kode prodi    : " + KodeProdi);
        System.out.println("Nama Prodi    : " + NamaProdi());
        System.out.println("Nilai UTS     : " + this.UTS);
        System.out.println("Nilai UAS     : " + this.UAS);
        System.out.println("Nilai Tugas   : " + this.Tugas);
        System.out.println("Total nilai   : " + this.TotalNilai());
        System.out.println("Grade         : " + grade);
    }
}






