#include <stdio.h>
#include <string.h>
#include "wrapper.h"

char * fort_(char *, int *);

char * RunModel(char inputStr[]) {
    int len = strlen(inputStr);
    return fort_(inputStr, &len);
}