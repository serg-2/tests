#include <string.h>
#include <stdio.h>
#include <openssl/aes.h>

char key[] = "thisisasecretkey";

int main(){
unsigned char text[]="hello world";
unsigned char enc_out[80];
unsigned char dec_out[80];

AES_KEY enc_key, dec_key;

AES_set_encrypt_key(key, 128, &enc_key);
AES_encrypt(text, enc_out, &enc_key);

//Ari
AES_set_decrypt_key(key, 128, &dec_key);
AES_decrypt(enc_out, dec_out, &dec_key);

int i;

printf("original:\t");
for(i=0;*(text+i)!=0x00;i++)
    printf("%02X ",*(text+i));
printf("\nencrypted:\t");
for(i=0;*(enc_out+i)!=0x00;i++)
    printf("%02X ",*(enc_out+i));
printf("\n");

printf("original:%s\t",text);
printf("\nencrypted:%s\t",enc_out);
printf("\ndecrypted:%s\t",dec_out);
printf("\n");

return 0;
}

