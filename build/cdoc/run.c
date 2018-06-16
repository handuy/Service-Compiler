#include <stdio.h>
#include <stdlib.h>

int main(int argc,char* argv[]){
    // Parent process will return a non-zero value from fork()
    printf("I'm the parent.\n");

    printf("This is my main program and it will continue running and doing anything i want to...\n");

    return 0;
}