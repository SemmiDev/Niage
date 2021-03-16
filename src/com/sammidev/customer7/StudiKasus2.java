package com.sammidev.customer7;

import java.util.Scanner;

/**
 * Created by sam on 16/03/21.
 */
public class StudiKasus2 {
    public static void main(String[] args) {
        boolean isNext = true;
        while (isNext) {

            Scanner scannerString = new Scanner(System.in);
            Scanner scannerInt = new Scanner(System.in);
            Scanner scannerDouble = new Scanner(System.in);

            String kodeAlatBerat,namaAlatBerat,namaPenyewa,noKTPPenyewa;
            int lamaSewaPerJam;
            double biayaSewaPerJam;

            System.out.print("Masukkan Kode alat berat      : ");
            kodeAlatBerat = scannerString.nextLine();
            System.out.print("Masukkan Nama alat berat      : ");
            namaAlatBerat = scannerString.nextLine();
            System.out.print("Masukkan Biaya sewa (per jam) : ");
            biayaSewaPerJam = scannerDouble.nextDouble();
            System.out.print("Masukkan Lama sewa (per jam)  : ");
            lamaSewaPerJam = scannerInt.nextInt();
            System.out.print("Masukkan Nama Penyewa         : ");
            namaPenyewa = scannerString.nextLine();
            System.out.print("Masukkan No KTP Penyewa       : ");
            noKTPPenyewa = scannerString.nextLine();

            SewaAlatBerat sewaAlatBerat = new SewaAlatBerat(kodeAlatBerat,namaAlatBerat,biayaSewaPerJam,lamaSewaPerJam,namaPenyewa,noKTPPenyewa);
            sewaAlatBerat.TotalBayar();

            System.out.println("\n\n");

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

class SewaAlatBerat {
    private String KodeAlatBerat;
    private String NamaAlatBerat;
    private Double BiayaSewaPerJam;
    private int    LamaSewaPerJam;
    private String NamaPenyewa;
    private String NoKTPPenyewa;

    public SewaAlatBerat(String kodeAlatBerat,
                         String namaAlatBerat,
                         Double biayaSewaPerJam,
                         int lamaSewaPerJam,
                         String namaPenyewa,
                         String noKTPPenyewa) {
        KodeAlatBerat = kodeAlatBerat;
        NamaAlatBerat = namaAlatBerat;
        BiayaSewaPerJam = biayaSewaPerJam;
        LamaSewaPerJam = lamaSewaPerJam;
        NamaPenyewa = namaPenyewa;
        NoKTPPenyewa = noKTPPenyewa;
    }

    public Double SubTotalSewa() {
        return this.BiayaSewaPerJam * this.LamaSewaPerJam;
    }

    public Double PotonganHarga() {
        if (this.LamaSewaPerJam < 12) {
            return 0d;
        }
        return SubTotalSewa() * 0.1;
    }

    public void TotalBayar() {
        System.out.println("Total Bayar : " + (SubTotalSewa() - PotonganHarga()));
    }
}
