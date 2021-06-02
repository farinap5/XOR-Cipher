<h1 align="center">XOR Cipher</h1>
<h3 align="center">Simple encoding for files and messages</h3>

---
Do not use in environments that would require secure encryption!

```
Simple explanation of how exclusive or works.
LOGIC:
eXclusive OR          OR
INPUT|OUTPUT     INPUT|OUTPUT
-----|------     -----|------
A   B| XOR       A   B|  OR
0   0|  0        0   0|   0
0   1|  1        0   1|   1
1   0|  1        1   0|   1
1   1|  0        1   1|   1

abc - 01100001 01100010 01100011   %%% - 00001010 00000111 00011010
key - 01101011 01100101 01111001   key - 01101011 01100101 01111001
      -------- -------- --------         -------- -------- --------
%%% - 00001010 00000111 00011010   abc - 01100001 01100010 01100011

%%% = any sequence of bytes.
```
