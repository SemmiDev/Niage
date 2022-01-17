class Material:
    id_material = 0
    nama_material = ""
    satuan = ""
    spesifikasi = ""
    harga = 0
    nama_supplier = ""

    def __init__(self, id_material, nama_material, satuan, spesifikasi, harga, nama_supplier):
        self.id_material = id_material
        self.nama_material = nama_material
        self.satuan = satuan
        self.spesifikasi = spesifikasi
        self.harga = harga
        self.nama_supplier = nama_supplier

    def tampilMaterial(self):
        print("ID Material : ", self.id_material)
        print("Nama Material : ", self.nama_material)
        print("Satuan : ", self.satuan)
        print("Spesifikasi : ", self.spesifikasi)
        print("Harga : ", self.harga)
        print("Nama Supplier : ", self.nama_supplier)

class Pemesanan:
    dataPemesanan = []
    dataMaterial = []

    no_invoice = 0
    id_material = 0
    jumlah = 0
    nama_gudang = ""
    tgl_pemesanan = ""

    def __init__(self):
        self.no_invoice = 0
        self.id_material = 0
        self.jumlah = 0
        self.nama_gudang = ""
        self.tgl_pemesanan = ""        

    def tampilPemesanan(self):
        totalHargaPemesanan = 0
        for pemesanan in self.dataPemesanan:
            print("No Invoice:", pemesanan.no_invoice)
            print("ID Material:", pemesanan.id_material)
            print("Jumlah:", pemesanan.jumlah)
            print("Nama Gudang:", pemesanan.nama_gudang)
            print("Tanggal Pemesanan:", pemesanan.tgl_pemesanan)
            
            harga = 0
            for material in self.dataMaterial:
                if material.id_material == pemesanan.id_material:
                    harga = material.harga
                    break
            totalHargaPemesanan += pemesanan.jumlah * harga
            print("Sub Total Harga:", pemesanan.jumlah * harga)
            print("--------------------------------------------")
        print("")
        print("Total Harga Pemesanan:", totalHargaPemesanan)
        print("")
    
    def hitungTotal(self):
        totalHargaPemesanan = 0
        for pemesanan in self.dataPemesanan:
            harga = 0
            for material in self.dataMaterial:
                if material.id_material == pemesanan.id_material:
                    harga = material.harga
                    break
            totalHargaPemesanan += pemesanan.jumlah * harga
        print("Total Harga Pemesanan:", totalHargaPemesanan)
        print("")

    def tambahItem(self, Material, no_invoice, jumlah, nama_gudang, tgl_pemesanan):
        pemesanan = Pemesanan()

        pemesanan.id_material = Material.id_material
        pemesanan.no_invoice = no_invoice
        pemesanan.jumlah = jumlah
        pemesanan.nama_gudang = nama_gudang
        pemesanan.tgl_pemesanan = tgl_pemesanan

        self.dataPemesanan.append(pemesanan)
        self.dataMaterial.append(Material)
    
    def hapusItem(self, id_material):
        for i in range(len(self.dataMaterial)):
            if self.dataMaterial[i].id_material == id_material:
                del self.dataMaterial[i]
                for j in range(len(self.dataPemesanan)):
                    if self.dataPemesanan[i].id_material == id_material:
                        del self.dataPemesanan[i]
                        break
                break

material1 = Material(1, "Buku",   "Pcs", "Buku Tulis",   10000, "Bambang")
material2 = Material(2, "Kertas", "Pcs", "Kertas HVS 1", 20000, "Andi")
material3 = Material(3, "Kertas", "Pcs", "Kertas HVS 2", 30000, "Dendi")
material4 = Material(4, "Kertas", "Pcs", "Kertas HVS 3", 40000, "Wisnu")
material5 = Material(5, "Kertas", "Pcs", "Kertas HVS 4", 50000, "Rahmah")

pemesanan1 = Pemesanan()
pemesanan1.tambahItem(material1, 1, 1, "Gudang 1", "12-12-2019")
pemesanan1.tambahItem(material2, 2, 1, "Gudang 2", "12-12-2020")
pemesanan1.tambahItem(material3, 3, 1, "Gudang 3", "12-12-2021")
pemesanan1.tambahItem(material4, 4, 1, "Gudang 4", "12-10-2019")
pemesanan1.tambahItem(material5, 5, 1, "Gudang 5", "12-12-2019")

pemesanan1.tampilPemesanan()
pemesanan1.hitungTotal()

pemesanan1.hapusItem(1)
pemesanan1.tampilPemesanan()
pemesanan1.hitungTotal()

pemesanan1.hapusItem(2)
pemesanan1.tampilPemesanan()
pemesanan1.hitungTotal()

pemesanan1.hapusItem(3)
pemesanan1.tampilPemesanan()
pemesanan1.hitungTotal()

pemesanan1.hapusItem(4)
pemesanan1.tampilPemesanan()
pemesanan1.hitungTotal()

pemesanan1.hapusItem(5)
pemesanan1.tampilPemesanan()
pemesanan1.hitungTotal()
