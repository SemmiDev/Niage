package com.sammidev.customer7;

import java.util.Scanner;

/**
 * Created by sam on 16/03/21.
 */
public class StudiKasus3 {
    public static void main(String[] args) {

        boolean isNext = true;
        while (isNext) {

            Scanner scannerString = new Scanner(System.in);
            Scanner scannerInt = new Scanner(System.in);
            Scanner scannerDouble = new Scanner(System.in);

            String KodeBarang, NamaBarang, PotonganBiayaKirim;
            int Quantity, NoFakturBeli;
            double HargaBarang, BiayaKirim;

            System.out.print("Masukkan no faktur beli    : ");
            NoFakturBeli = scannerInt.nextInt();

            System.out.print("Masukkan kode barang       : ");
            KodeBarang = scannerString.nextLine();

            System.out.print("Masukkan nama barang       : ");
            NamaBarang = scannerString.nextLine();

            System.out.print("Masukkan Harga Barang      : ");
            HargaBarang = scannerDouble.nextDouble();

            System.out.print("Masukkan quantity          : ");
            Quantity = scannerInt.nextInt();

            System.out.print("Potongan biaya kirim (Y/N) : ");
            PotonganBiayaKirim = scannerString.nextLine();

            System.out.print("Masukkan Biaya Kirim       : ");
            BiayaKirim = scannerDouble.nextDouble();

            Double potonganBiayaKirim = 0.0;

            if (PotonganBiayaKirim.equalsIgnoreCase("Y")) {
                potonganBiayaKirim = BiayaKirim * 0.1;
                BiayaKirim = BiayaKirim - potonganBiayaKirim;
            }

            Distributor distributor = new Distributor(NamaBarang, potonganBiayaKirim);
            distributor.NoFakturBeli = NoFakturBeli;
            distributor.KodeBarang = KodeBarang;
            distributor.HargaBarang = HargaBarang;
            distributor.Quantity = Quantity;
            distributor.BiayaKirim = BiayaKirim;

            distributor.Biaya(HargaBarang, Quantity);
            distributor.Biaya(BiayaKirim);
            distributor.Biaya(HargaBarang, Quantity, BiayaKirim);

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

class Distributor {
    public int NoFakturBeli;
    public String KodeBarang;
    private String NamaBarang;
    public Double HargaBarang;
    public int Quantity;
    private Double PotonganBiayaKirim;
    public Double BiayaKirim;

    public Distributor(String namaBarang, Double PotonganBiayaKirim) {
        this.NamaBarang = namaBarang;
        this.PotonganBiayaKirim = PotonganBiayaKirim;
    }

    public void Biaya(Double hargaBarang, int quantity) {
        Double result = hargaBarang * quantity;
        System.out.println("Biaya dengan overloading 1 : " + result);
    }

    public void Biaya(Double biayaKirim) {
        Double result = (biayaKirim / 100) * 10;
        System.out.println("Biaya dengan overloading 2 : " +result);
    }

    public void Biaya(Double hargaBarang, int quantity, Double biayaKirim) {
        Double result = (hargaBarang * quantity) - (biayaKirim / 100) * 10;
        System.out.println("Biaya dengan overloading 3 : " +result);
    }
}