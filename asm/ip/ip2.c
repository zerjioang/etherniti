/* '0.0.0.0' is not a valid IP address, so this uses the value 0 to
   indicate an invalid IP address. */
#define byte unsigned char
#define uint8 unsigned char

#define asciiDot 46
#define asciiZero 48

unsigned int ip_to_int2 (const char * ip, const int len)
{
    /* The return value. */
    unsigned intIp = 0;

    byte octets[4][4];
    byte currentOctect = 0;
    byte currentOctectPos = 0;
    int i = 0;
    for(i = 0; i < len; i++){
        byte ipVal = ip[i];
        if(ipVal == asciiDot){
            octets[currentOctect][3] = currentOctectPos;
            //move to the next octect
            currentOctect++;
            currentOctectPos = 0;
        } else {
            // assign value to current octect
            octets[currentOctect][currentOctectPos] = ipVal - asciiZero;
            currentOctectPos++;
        }
    }
    // save last octet information
    octets[currentOctect][3] = currentOctectPos;

    // convert octects string bytes to decimal
    byte octectsDecimal[4];
    byte idx = 0;
    for(idx = 0; idx < 4; idx++){
        //process each octect data
        // convert octects to uint32
        // octets[0]*256³ + octets[1]*256² + octets[2]*256¹ + octets[1]*256⁰
        switch(octets[idx][3]){
        case 0:
            octectsDecimal[idx] = 0;
            break;
        case 1:
            octectsDecimal[idx] = octets[idx][0];
            break;
        case 2:
            octectsDecimal[idx] = octets[idx][0]*10 + octets[idx][1];
            break;
        case 3:
            octectsDecimal[idx] = octets[idx][0]*100 + octets[idx][1]*10 + octets[idx][2];
            break;
        }
    }
    // intIp = uint32(octectsDecimal[0])*16777216 + uint32(octectsDecimal[1])*65536 + uint32(octectsDecimal[2])*256 + uint32(octectsDecimal[3])
    intIp = octectsDecimal[3] | octectsDecimal[2]<<8 | octectsDecimal[1]<<16 | octectsDecimal[0]<<24;
    return intIp;
}