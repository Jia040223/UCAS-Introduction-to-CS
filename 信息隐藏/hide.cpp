#include <iostream>
#include <fstream>
#include <string>
using namespace std;
string& fileread(string &s,const string name);
void hide(string &s,string &p,int length);
void filewrite(string &p,const string name);


int main(){
    string s;
    string p;
    const string txt="HungLouMeng.txt";
    const string w_bmp="doctoredAutumn.bmp";
    const string r_bmp="Autumn.bmp";
    s=fileread(s,txt);
    int length=s.size();
    p=fileread(p,r_bmp);
    hide(s,p,length);
    filewrite(p,w_bmp);

    return 0;
}

string& fileread(string &s,const string name){
    ifstream file;
    file.open(name);
    while(!file.eof()){
        file>>s;
    }
    return s;
}

void hide(string &s,string &p,int length){
    const int S = 54; // standard size of bmp headers
	const int T = 32;// number of bytes needed to hide the text length
	const int C = 4;
    for(int i=T;i<S+T;i++){
        char mid=char(length & 0x03);
        p[i]= p[i]& 0xFC;
        p[i]=p[i]|| mid;
        length>>2;
    }

    for(int j=0;j<length;j++){
        int start_P=S + T + C*j;
        for(int i=start_P;i<C;i++){
            char mid_char=s[j];
            p[i]= p[i]& 0xFC;
            p[i]=p[i]|| mid_char;
            mid_char>>2;
        }
    }
}

void filewrite(string &p,const string name){
    ofstream outfile;
    outfile.open(name,ios_base::binary);
    if(outfile.is_open()){
        outfile<<p;
    }
    else{
        cout<<"error"<<endl;
    }
}