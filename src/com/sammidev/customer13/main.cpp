
#include <iostream>

using namespace std;

int main (){
	  int apel, melon, semangka, mangga, jeruk;

	  cout << "                   Selamat Datang di Toko XYZ\n" << endl;
	  cout << "Daftar Pemesanan, isi dengan [0] jika tidak ingin melakukan pemesanan" << endl;
	  cout << "Apel \tRp. 25.000 \t: ";
	  cin >> apel;
	  cout << "Melon \tRp. 15.000,- \t: ";
	  cin >> melon;
	  cout << "Semangka \tRp. 10.000,- \t: ";
	  cin >> semangka;
	  cout << "Mangga \tRp. 10.000,- \t: ";
	  cin >> mangga;
	  cout << "Jeruk \tRp. 8.000,- \t: ";
	  cin >> jeruk;
	  
	  cout << "                   Selamat Datang di Toko XYZ\n" << endl;
	  cout << "Daftar Pemesanan, isi dengan [0] jika tidak ingin melakukan pemesanan" << endl;
	  cout << "Apel	    Rp. 25.000	  : " << apel << " kg" << endl;
	  cout << "Melon    Rp. 15.000,-  : " << melon<< " kg"  << endl;
	  cout << "Semangka Rp. 10.000,-  : " << semangka << " kg" << endl;
	  cout << "Mangga   Rp. 10.000,-  : " << mangga << " kg" << endl;
	  cout << "Jeruk    Rp. 8.000,-   : " << jeruk << " kg" << endl;
	  
	  int bapel,bmelon,bsemangka,bmangga,bjeruk,total;
	  bapel = apel * 25000;
	  bmelon = melon * 15000;
	  bsemangka = semangka * 10000;
	  bmangga = mangga * 10000;
	  bjeruk = jeruk * 8000;
  
	cout << "\n\nRincian Pembelian" << endl;
	cout << "Apel	   : "  <<  apel	<<  " =  Rp " << bapel << endl;
	cout << "Melon	   : "  <<  melon	<<	" =  Rp " << bmelon << endl;
	cout << "Semangka  : "  <<	semangka<<	" =  Rp " << bsemangka << endl;
	cout << "Mangga	   : "	<<  mangga	<<	" =  Rp " << bmangga << endl;
	cout << "Jeruk	   : "	<<  jeruk	<<	" =  Rp " << bjeruk << endl;

	total = bapel + bmelon + bsemangka+bmangga+bjeruk; 
	int bersih = total;
	string diskon = "no";
	if (total > 100000)
	{ 
	    diskon = "yes \n 10%";
	    bersih = total - (total * 0.1);
	}

	cout << "Total harga    : Rp " << total << endl;
	cout << "Discount       : " << diskon << endl;
	cout << "Total bayar    : Rp " << bersih << endl;

  	return 0;
}
