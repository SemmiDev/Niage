#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <iostream>
#include <string>
#include <algorithm>

using namespace std;

#define N 3 

typedef struct MHS{
	string nama; 
	string nim;  
} MHS;

typedef struct NILAI{
	double uas; 
	double uts;
	double tugas;
	double absen;
	double nAkhir; 
	char nHuruf; 
} NILAI;

 typedef struct DATAMHS{
	MHS mhs;
	NILAI nilai;
 } DATAMHS;

void judul();
void searchingByNIM(DATAMHS dataMhs[], string matchNIM);
void searchingByName(DATAMHS dataMhs[], string matchName);
void sortingMhs(DATAMHS dataMhs[]);
void bacaMhs();
void bacaNilai(int i);
void infoMhs(int i);
double hitungAkhir(double tugas, double absen, double uts, double uas);
char konversiHuruf(double na);

int main(int argc, char const *argv[])
{
    
    DATAMHS dataMhs[N];

    dataMhs[0].mhs.nama = "sammi 1";
    dataMhs[0].mhs.nim = "2003113948";
    dataMhs[0].nilai.uas = 100; 
    dataMhs[0].nilai.uts = 40; 
    dataMhs[0].nilai.tugas = 40; 
    dataMhs[0].nilai.absen = 20;


    dataMhs[1].mhs.nama = "faren 2";
    dataMhs[1].mhs.nim = "2003113949";
    dataMhs[1].nilai.uts = 100; 
    dataMhs[1].nilai.tugas = 70; 
    dataMhs[1].nilai.absen = 60;
    dataMhs[1].nilai.uas = 70; 

    dataMhs[2].mhs.nama = "devin 3";
    dataMhs[2].mhs.nim = "2003113950";
    dataMhs[2].nilai.uas = 20; 
    dataMhs[2].nilai.uts = 30; 
    dataMhs[2].nilai.tugas = 40; 
    dataMhs[2].nilai.absen = 40;

    double nAkhirnya1 = hitungAkhir(dataMhs[0].nilai.uas, dataMhs[0].nilai.uts, dataMhs[0].nilai.tugas, dataMhs[0].nilai.absen);
    char nHurufnya1 = konversiHuruf(nAkhirnya1);	

    double nAkhirnya2 = hitungAkhir(dataMhs[1].nilai.uas, dataMhs[1].nilai.uts, dataMhs[1].nilai.tugas, dataMhs[1].nilai.absen);
    char nHurufnya2 = konversiHuruf(nAkhirnya2);	
    
    double nAkhirnya3 = hitungAkhir(dataMhs[2].nilai.uas, dataMhs[2].nilai.uts, dataMhs[2].nilai.tugas, dataMhs[2].nilai.absen);
    char nHurufnya3 = konversiHuruf(nAkhirnya3);	
    
    dataMhs[0].nilai.nAkhir = nAkhirnya1;
    dataMhs[0].nilai.nHuruf = nHurufnya1;

    dataMhs[1].nilai.nAkhir = nAkhirnya2;
    dataMhs[1].nilai.nHuruf = nHurufnya2;

    dataMhs[2].nilai.nAkhir = nAkhirnya3;
    dataMhs[2].nilai.nHuruf = nHurufnya3;


    cout << "Nama Mahasiswa  : " <<  dataMhs[0].mhs.nama << endl;
    cout << "Nomor induk mahasiswa  : " <<  dataMhs[0].mhs.nim << endl;
    cout << "Nilai Akhir  : " <<  dataMhs[0].nilai.nAkhir << endl;
    cout << "Nilai Huruf  : " <<  dataMhs[0].nilai.nHuruf << endl << endl;
    

    cout << "Nama Mahasiswa  : " <<  dataMhs[1].mhs.nama << endl;
    cout << "Nomor induk mahasiswa  : " <<  dataMhs[1].mhs.nim << endl;
    cout << "Nilai Akhir  : " <<  dataMhs[1].nilai.nAkhir << endl;
    cout << "Nilai Huruf  : " <<  dataMhs[1].nilai.nHuruf << endl << endl;

    cout << "Nama Mahasiswa  : " <<  dataMhs[2].mhs.nama << endl;
    cout << "Nomor induk mahasiswa  : " <<  dataMhs[2].mhs.nim << endl;
    cout << "Nilai Akhir  : " <<  dataMhs[2].nilai.nAkhir << endl;
    cout << "Nilai Huruf  : " <<  dataMhs[2].nilai.nHuruf << endl;
    
    sortingMhs(dataMhs);

    searchingByName(dataMhs, "ren");
    searchingByNIM(dataMhs, "2003113949");

    return 0;
}

void searchingByName(DATAMHS dataMhs[], string matchName) {
    cout << "===================================================================" << endl;
    cout << "MATCH WITH " + matchName << endl;
    cout << "===================================================================" << endl;
    for(int x = 0; x < N; x++){
        if (dataMhs[x].mhs.nama.find(matchName, 0) != string::npos){
                cout << "Nama Mahasiswa  : " <<  dataMhs[x].mhs.nama << endl;
                cout << "Nomor induk mahasiswa  : " <<  dataMhs[x].mhs.nim << endl;
                cout << "Nilai Akhir  : " <<  dataMhs[x].nilai.nAkhir << endl;
                cout << "Nilai Huruf  : " <<  dataMhs[x].nilai.nHuruf << endl << endl; 
        }
    }
}

void searchingByNIM(DATAMHS dataMhs[], string matchNIM) {
    cout << "===================================================================" << endl;
    cout << "MATCH WITH " + matchNIM << endl;
    cout << "===================================================================" << endl;
    for(int x = 0; x < N; x++){
        if (dataMhs[x].mhs.nim.find(matchNIM, 0) != string::npos){
                cout << "Nama Mahasiswa  : " <<  dataMhs[x].mhs.nama << endl;
                cout << "Nomor induk mahasiswa  : " <<  dataMhs[x].mhs.nim << endl;
                cout << "Nilai Akhir  : " <<  dataMhs[x].nilai.nAkhir << endl;
                cout << "Nilai Huruf  : " <<  dataMhs[x].nilai.nHuruf << endl << endl; 
        }
    }
}

void sortingMhs(DATAMHS dataMhs[]) {
    cout << "===================================================================" << endl;
    cout << "SORTING DESC BY NILAI " << endl;
    cout << "===================================================================" << endl;
    double nilaiMhs[N];

    for(int i=0; i < N; i++){
        nilaiMhs[i] = dataMhs[i].nilai.nAkhir;
    }

    // Sort the array in descending order
    sort(nilaiMhs, nilaiMhs + N, greater<double>());
  
    // Print the desc sorted array
    cout << "\nDescending Sorted Array:\n";
    for (int i = 0; i < N; i++){
        for (int j = 0; j < N; j++){
            if (nilaiMhs[i] == dataMhs[j].nilai.nAkhir){
                cout << "Nama Mahasiswa  : " <<  dataMhs[j].mhs.nama << endl;
                cout << "Nomor induk mahasiswa  : " <<  dataMhs[j].mhs.nim << endl;
                cout << "Nilai Akhir  : " <<  dataMhs[j].nilai.nAkhir << endl;
                cout << "Nilai Huruf  : " <<  dataMhs[j].nilai.nHuruf << endl << endl; 
                break;
            }
        }
    }      
}

double hitungAkhir(double tugas, double absen, double uts, double uas){
	double nAkhirnya;
	
	nAkhirnya =tugas * 0.2 + absen * 0.1 + uts * 0.3 + uas * 0.4;
	return nAkhirnya;
}

char konversiHuruf(double na){
	char nHurufnya;

	if((na >= 85.0) && (na<= 100.0))
		nHurufnya = 'A';
	else if(na >= 70.0)
		nHurufnya = 'B';
	else if(na >= 60.0)
		nHurufnya = 'C';
	else if(na >= 50.0)
		nHurufnya = 'D';
	else nHurufnya = 'E';
	
	return nHurufnya;	
}
