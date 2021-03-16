package com.sammidev.customer7;

import java.util.Scanner;

/**
 * Created by sam on 16/03/21.
 */
public class StudiKasus4 {
    public static void main(String[] args) {

        boolean isNext = true;
        while (isNext) {

            Scanner scannerString = new Scanner(System.in);
            Scanner scannerInt = new Scanner(System.in);
            Scanner scannerDouble = new Scanner(System.in);

            String NID, Nama, JenisKelamin;
            int Kepangkatan;
            double Gaji;

            System.out.print("Masukkan NID    : ");
            NID = scannerString.nextLine();

            System.out.print("Masukkan Nama   : ");
            Nama = scannerString.nextLine();

            System.out.print("Masukkan Jenis Kelamin    : ");
            JenisKelamin = scannerString.nextLine();

            System.out.print("Masukkan Kepangkatan      : \n");
            System.out.println("1 = AA\n2 = Lektor\n3 = Lektor Kepala\n4 = Guru Besar / Proferor\n");
            Kepangkatan = scannerInt.nextInt();

            System.out.print("Masukkan Gaji             : ");
            Gaji = scannerDouble.nextDouble();

            Dosen dosen = new Dosen();
            dosen.NID = NID;
            dosen.Nama = Nama;
            dosen.JenisKelamin = JenisKelamin;
            dosen.Kepangkatan = Kepangkatan;
            dosen.Gaji = Gaji;

            dosen.TotalGaji();
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

class DosenTetap {
    public String NID;
    public String Nama;
    public String JenisKelamin;
    public int Kepangkatan;
    public Double Gaji;
}

class Dosen extends DosenTetap {
    public double TunjanganFungsional() {
        if (this.Kepangkatan == 1) {
            return  1000000d;
        }else if (Kepangkatan == 2) {
            return  2000000d;
        }else if (Kepangkatan == 3) {
            return  3000000d;
        }else if (Kepangkatan == 4) {
            return  4000000d;
        }
        return 0d;
    }

    public double PajakPenghasilan() {
        return 0.015 * this.Gaji;
    }

    public void TotalGaji() {
        Double totalGaji = this.Gaji + TunjanganFungsional() - PajakPenghasilan();
        System.out.println("Total Gaji : " + totalGaji);
    }
}
