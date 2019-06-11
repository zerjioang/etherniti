#include <stdio.h>

char* toUpper(char *s) {
   int c = 0;
   
   while (s[c] != '\0') {
      if (s[c] >= 'a' && s[c] <= 'z') {
         s[c] = s[c] - 32;
      }
      c++;
   }
   return s;
}

char* toLower(char *s) {
   int c = 0;
   
   while (s[c] != '\0') {
      if (s[c] >= 'A' && s[c] <= 'Z') {
         s[c] = s[c] + 32;
      }
      c++;
   }
   return s;
}

int main()
{
    char example[] = "Hello World";
    printf("The string is: %s\n", example);
    example = toUpper(example);
    printf("The string in upper case is: %s\n", example);
     example = toLower(example);
    printf("The string in lower case is: %s\n", example);
    return 0;
}