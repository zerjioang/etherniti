int main(int argc, char const *argv[])
{
    /* code */
    unsigned integer;
    const char * ip = "110.3.244.53";
    integer = ip_to_int (ip);
    printf ("'%s' is 0x%08x.\n", ip, integer);
    integer = ip_to_int2 (ip, 12);
    printf ("'%s' is 0x%08x.\n", ip, integer);
    return 0;
}