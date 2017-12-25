#include<stdio.h>

int main()
{
    int a = 1;
    int b = 0;
    int c = 0;
    int d = 0;
    int e = 0;
    int f = 0;
    int g = 0;
    int h = 0;
    label0: b = 79;
    label1: c = b;
    if (a !=0){
        goto label4;
    }
    if (1 !=0){
        goto label8;
    }
    label4: b *= 100;
    label5: b -= -100000;
    label6: c = b;
    label7: c -= -17000;
    label8: f = 1;
    label9: d = 2;
    label10: e = 2;
    if (b%d ==0 ){
        f=0;
    }
    e=b;
    label20: d -= -1;
    label21: g = d;
    label22: g -= b;
    if (g !=0){
        goto label10;
    }
    if (f !=0){
        goto label26;
    }
    label25: h -= -1;
    label26: g = b;
    label27: g -= c;
    if (g !=0){
        goto label30;
    }
    if (1 !=0){
        goto label2147483647;
    }
    label30: b -= -17;
    if (1 !=0){
        goto label8;
    }
    label2147483647: printf("ENDE h:  %i\n",h);
    return 0;

}