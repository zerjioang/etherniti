#include<stdbool.h>

# define level 32

bool isDigit(char c) {
  if ( c >= '0' && c <= '9' )   
      return true;

  return false;
}

void isNumericArray(char data[], int len, char *res){
    int idx;
    *res = 0x1;
    for (idx = 0; idx < len; idx++){
        if (!isDigit(data[idx])){
        	*res = 0x0;
        	break;
        }
    }
}

void lowerCase(char data[], int len, char *res){
    int idx;
	char c;
    for (idx = 0; idx < len; idx++){
    	c = data[idx];
        if (c >= 'A' && c <= 'Z'){
        	res[idx] = c | level;
        }
    }
}