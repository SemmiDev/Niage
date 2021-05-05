#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <iostream>
#include <string>
#include <algorithm>
#include <sstream>

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
string underscoreToSpace(string text);
void searchingByNIM(DATAMHS dataMhs[], string matchNIM);
void searchingByName(DATAMHS dataMhs[], string matchName);
void sortingMhs(DATAMHS dataMhs[]);
void bacaMhs();
void bacaNilai(int i);
void infoMhs(int i);
double hitungAkhir(double tugas, double absen, double uts, double uas);
char konversiHuruf(double na);

DATAMHS dataMhs[N];

int main(int argc, char const *argv[]){

    judul();    
    bacaMhs();
	
    for(int i = 0; i < N; i++){
		infoMhs(i);
		printf("\n");
	}
    
    sortingMhs(dataMhs);

    string inputNama;
    string inputNIM;

    for (int i = 0; i < N; i++)
    {
        cout << "Pencarian bedasarkan nama : ";
        cin >> inputNama;
        searchingByName(dataMhs, inputNama);
        cout << endl << endl;
    }


    for (int i = 0; i < N; i++)
    {
        cout << "Pencarian bedasarkan NIM : ";
        cin >> inputNIM;
        searchingByNIM(dataMhs, inputNIM);
        cout << endl << endl;
    }
    
    return 0;
}

void judul(){
	printf("=================================================================\n");
	printf("\t\t\t SELAMAT DATANG\t\t\t \n");
	printf("\t\t PROGRAM MENGHITUNG NILAI AKHIR\t\t \n");
	printf("\t MAHASISWA PENGANTAR CODING UNIVERSITAS NEGERI PADANG \t \n");
	printf("=================================================================\n");
	printf("Dibuat oleh : ADELYA AMANDA \n");    
}


void bacaMhs(){
    char namanya[100];
    char nimnya[50];

    string nama;
    string namawithspace;
    string nim;

	printf("\nMembaca identitas sejumlah Mahasiswa\n");
	printf("========================================\n");
	for(int i = 0; i < N; i++){

        printf("Masukkan NAMA Mahasiswa: ");
		gets(namanya); fflush(stdin);
		
        printf("Masukan NIM : ");
		gets(nimnya); fflush(stdin); 

        // cout << "Masukkan NAMA Mahasiswa: ( pisahkan dengan _ ) : ";
        // cin >> nama;

		// cout << "Masukkan NIM : ";
		// cin >> nim;

        // namawithspace = underscoreToSpace(nama);

        nama = string(namanya);
        nim = string(nimnya);

		dataMhs[i].mhs.nama = nama; 
		dataMhs[i].mhs.nim = nim;
		
        bacaNilai(i);
	}
}

string underscoreToSpace(string text){
    for(int i = 0; i < text.length(); i++)
    {
        if(text[i] == '_')
            text[i] = ' ';
    }
    return text;
}

void bacaNilai(int i){
	double uasnya, utsnya, tugasnya, absennya, nAkhirnya;
	char nHurufnya;

	cout << "Masukkan nilai UAS : ";
	cin >> uasnya;

    cout << "Masukkan nilai UTS : ";
	cin >> utsnya;
    
	cout << "Masukkan nilai Tugas : ";
    cin >> tugasnya;

	cout << "Masukkan nilai Absen : ";
    cin >> absennya;

	dataMhs[i].nilai.uas = uasnya; 
	dataMhs[i].nilai.uts = utsnya; 
	dataMhs[i].nilai.tugas = tugasnya; 
	dataMhs[i].nilai.absen = absennya;

	nAkhirnya = hitungAkhir(uasnya, utsnya, tugasnya, absennya);
	nHurufnya = konversiHuruf(nAkhirnya);

	dataMhs[i].nilai.nAkhir = nAkhirnya;
	dataMhs[i].nilai.nHuruf = nHurufnya;
	
    cout << endl;
}

void infoMhs(int i){

	printf("\nInformasi Identitas Mahasiswa\n");
	printf("========================================\n");

    for (int i = 0; i < N; i++)
    {
        cout << "Nama Mahasiswa  : " << dataMhs[i].mhs.nama << endl;
        cout << "Nomor induk mahasiswa :" << dataMhs[i].mhs.nim << endl;
        cout << "Nilai Akhir   : " << dataMhs[i].nilai.nAkhir << endl;
        cout << "Nilai Huruf   : " << dataMhs[i].nilai.nHuruf << endl;
        cout << "========================================" << endl;
    }
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
    cout << "SORTING MAHASISWA DESC BY NILAI " << endl;
    cout << "===================================================================" << endl;
    double nilaiMhs[N];

    for(int i=0; i < N; i++){
        nilaiMhs[i] = dataMhs[i].nilai.nAkhir;
    }

    sort(nilaiMhs, nilaiMhs + N, greater<double>());
  
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
	return tugas * 0.2 + absen * 0.1 + uts * 0.3 + uas * 0.4;
}

char konversiHuruf(double na){
	char nHurufnya;

	if((na >= 85.0) && (na <= 100.0))
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
