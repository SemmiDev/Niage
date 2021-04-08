package com.sammidev.customer14;

public class Application {
    public static void main(String[] args) {
        // polimorfishm
        Budaya budayaSumbar1 = new BudayaSumbar("Rumah Gadang","Sumatera Barat", JenisBudaya.RUMAH_ADAT);
        Budaya budayaSumbar2 = new BudayaSumbar("Tari Piring","Sumatera Barat", JenisBudaya.TARIAN_ADAT);
        Budaya budayaSumbar3 = new BudayaSumbar("Pakaian Limpapeh","Sumatera Barat", JenisBudaya.BAJU_ADAT);

        // polimorfishm
        Budaya budayaAceh1 = new BudayaAceh("RUmah Krong Bade","Aceh", JenisBudaya.RUMAH_ADAT);
        Budaya budayaAceh2 = new BudayaAceh("Tari Saman.","Aceh", JenisBudaya.TARIAN_ADAT);
        Budaya budayaAceh3 = new BudayaAceh("Daro Buro","Aceh", JenisBudaya.BAJU_ADAT);

        // polimorfishm
        Budaya budayaJawa1 = new BudayaJawa("Rumah Joglo","Jawa", JenisBudaya.RUMAH_ADAT);
        Budaya budayaJawa2 = new BudayaJawa("Tari Bambangan Cakil","Jawa", JenisBudaya.TARIAN_ADAT);
        Budaya budayaJawa3 = new BudayaJawa("Kebaya","Jawa", JenisBudaya.BAJU_ADAT);

        System.out.println(budayaSumbar1.toString());
        System.out.println(budayaSumbar2.toString());
        System.out.println(budayaSumbar3.toString());

        System.out.println(budayaAceh1.toString());
        System.out.println(budayaAceh2.toString());
        System.out.println(budayaAceh3.toString());

        System.out.println(budayaJawa1.toString());
        System.out.println(budayaJawa2.toString());
        System.out.println(budayaJawa3.toString());
    }
}

class Budaya {

    // enkapsulasi
    private String namaBudaya;
    private String asalBudaya;
    private JenisBudaya jenisBudaya;

    // constructor
    public Budaya(String namaBudaya, String asalBudaya, JenisBudaya jenisBudaya) {
        this.namaBudaya = namaBudaya;
        this.asalBudaya = asalBudaya;
        this.jenisBudaya = jenisBudaya;
    }

    @Override
    public String toString() {
        return namaBudaya + " | " + asalBudaya + " | " + jenisBudaya.name();
    }
}

// inheritance
class BudayaSumbar extends Budaya{

    public BudayaSumbar(String nama, String asal, JenisBudaya jenisBudaya) {
        super(nama, asal, jenisBudaya);
    }
}

// inheritance
class BudayaAceh extends Budaya {

    public BudayaAceh(String namaBudaya, String asalBudaya, JenisBudaya jenisBudaya) {
        super(namaBudaya, asalBudaya, jenisBudaya);
    }
}

// inheritance
class BudayaJawa extends Budaya {

    public BudayaJawa(String namaBudaya, String asalBudaya, JenisBudaya jenisBudaya) {
        super(namaBudaya, asalBudaya, jenisBudaya);
    }
}

// enum
enum JenisBudaya {
    RUMAH_ADAT, BAJU_ADAT, TARIAN_ADAT, SENJATA_ADAT, SUKU, MAKANAN_DAERAH
}