#include <stdio.h>
#include <stdlib.h>
#include <cs50.h>
#include <stdint.h>
#include <ctype.h>


bool checkjpg(uint8_t *buffer);


int main(int argc, char *argv[])
{
 if (argc != 2)
    {
        printf("Usage: ./recover FILE\n");
        return 1;
    }

    FILE *card = fopen(argv[1], "r");
    if (card == NULL)
    {
        fprintf(stderr, "Could not open %s.\n", argv[1]);
        return 2;
    }

    FILE *JPG;

    uint8_t buffer[512];

    bool endJPG = false;
    int readblock;
    int jpgnum = 0;

    while (fread(buffer, 1, 512, card) == 512)
    {
       if (checkjpg(buffer))
        {
            char filename[8];

            sprintf(filename, "%03.3i.jpg", jpgnum);
            printf("processing file: %s\n", filename);

            if (endJPG == true)
            {
                fclose(JPG);
            }
            else
            {
                endJPG = true;
            }

            JPG = fopen(filename, "w");
            fwrite(buffer, sizeof(buffer), 1, JPG);


            jpgnum++;
        }
          else{
            if (endJPG == true){
                fwrite(buffer, sizeof(buffer), 1, JPG);
            }
        }
    }
   fclose(card);
    fclose(JPG);
}

bool checkjpg(uint8_t *buffer){
    if (buffer[0] == 0xff && buffer[1] == 0xd8 && buffer[2] == 0xff && (buffer[3] & 0xf0) == 0xe0)
    {
        return true;
    }
    else
    {
        return false;
    }
}
