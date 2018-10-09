var waitforme = new Array(-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1, -1, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, -1, -1, -1, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, -1, -1, -1, -1, -1, -1, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, -1, -1, -1, -1, -1);
function sd34d1dxwdpox(a) {
    var b, c2, c3, c4;
    var i, len, out; len = a.length;
    i = 0; out = "";
    while (i < len) {
        do { b = waitforme[a.charCodeAt(i++) & 0xff] } while (i < len && b == -1);
        if (b == -1) break;
        do { c2 = waitforme[a.charCodeAt(i++) & 0xff] } while (i < len && c2 == -1); if (c2 == -1) break;
        out += String.fromCharCode((b << 2) | ((c2 & 0x30) >> 4)); do {
            c3 = a.charCodeAt(i++) & 0xff; if (c3 == 61) return out;
            c3 = waitforme[c3]
        } while (i < len && c3 == -1);
        if (c3 == -1) break;
        out += String.fromCharCode(((c2 & 0XF) << 4) | ((c3 & 0x3C) >> 2));
        do {
            c4 = a.charCodeAt(i++) & 0xff; if (c4 == 61) return out; c4 = waitforme[c4]
        } while (i < len && c4 == -1);
        if (c4 == -1) break;
        out += String.fromCharCode(((c3 & 0x03) << 6) | c4)
    } return out
}
function xs2rsv345kl(a) {
    var b, i, len, c;
    var d, char3; b = "";
    len = a.length; i = 0;
    while (i < len) {
        c = a.charCodeAt(i++);
        switch (c >> 4) {
            case 0: case 1: case 2: case 3: case 4: case 5: case 6: case 7: b += a.charAt(i - 1);
                break;
            case 12: case 13: d = a.charCodeAt(i++); b += String.fromCharCode(((c & 0x1F) << 6) | (d & 0x3F));
                break;
            case 14: d = a.charCodeAt(i++); char3 = a.charCodeAt(i++); b += String.fromCharCode(((c & 0x0F) << 12) | ((d & 0x3F) << 6) | ((char3 & 0x3F) << 0));
                break
        }
    } return b
}
function htmlcode(a) {
    a = a.replace(/(\*)/g, 'A');
    a = a.replace(/(#)/g, 'b');
    a = a.replace(/(!)/g, 'c');
    a = a.replace(/(:)/g, 'M');
    a = a.replace(/(@)/g, 'N');
    return xs2rsv345kl(sd34d1dxwdpox(a))
}

process.stdout.write(htmlcode(process.argv[2]))