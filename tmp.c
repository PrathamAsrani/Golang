#include<stdio.h>
#include<stdlib.h>

int main(){
    printf("Hello world\n");
    int a = 32;
    char c = 'c';
    auto *ptr = &c;
    printf("%d", ptr);
}