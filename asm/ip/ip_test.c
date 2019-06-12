/* The following is a test of the above routine. */

int main ()
{
    int i;
    /* "ips" contains a selection of internet address strings. */
    const char * ips[] = {
        "127.0.0.1",
        "110.3.244.53",
        /* The following tests whether we can detect an invalid
           address. */
        "Bonzo Dog Doo-Dah Band",
        "182.118.20.178",
        /* The following tests whether it is OK to end the string with
           a non-NUL byte. */
        "74.125.16.64  ",
        "1234.567.89.122345",
    };
    int n_ips;

    /* Set "n_ips" to the number of elements in "ips". */

    n_ips = sizeof (ips) / sizeof (const char *);

    for (i = 0; i < n_ips; i++) {
        unsigned integer;
        const char * ip;

        ip = ips[i];
        integer = ip_to_int (ip);
        if (integer == INVALID) {
            printf ("'%s' is not a valid IP address.\n", ip);
        }
        else {
            printf ("'%s' is 0x%08x.\n", ip, integer);
        }
    }
    return 0;
}