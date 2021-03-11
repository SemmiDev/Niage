#include <iostream>
using namespace std;
int main(){
    int s,luas,keliling;

    cout << "Masukan sisi persegi = ";
    cin >> s;
    
    luas = s*s;
    cout<<"Luas persegi adalah "<< luas << endl;
    
    keliling = 4 * s;
    cout << "Keliling persegi adalah "<< keliling;
    return 0;
}
