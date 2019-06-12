#include <stdio.h>

/* '0.0.0.0' is not a valid IP address, so this uses the value 0 to
   indicate an invalid IP address. */

#define INVALID 0
#define byte unsigned char
#define uint8 unsigned char

#define asciiDot 46
#define asciiZero 48

/* Convert the character string in "ip" into an unsigned integer.

   This assumes that an unsigned integer contains at least 32 bits. */

unsigned int ip_to_int (const char * ip)
{
    /* The return value. */
    unsigned v = 0;
    /* The count of the number of bytes processed. */
    int i;
    /* A pointer to the next digit to process. */
    const char * start;

    start = ip;
    for (i = 0; i < 4; i++) {
        /* The digit being processed. */
        char c;
        /* The value of this byte. */
        int n = 0;
        while (1) {
            c = * start;
            start++;
            if (c >= '0' && c <= '9') {
                n *= 10;
                n += c - '0';
            }
            /* We insist on stopping at "." if we are still parsing
               the first, second, or third numbers. If we have reached
               the end of the numbers, we will allow any character. */
            else if ((i < 3 && c == '.') || i == 3) {
                break;
            }
            else {
                return INVALID;
            }
        }
        if (n >= 256) {
            return INVALID;
        }
        v *= 256;
        v += n;
    }
    return v;
}