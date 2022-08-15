#include <stdio.h>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <sys/types.h>
#include <unistd.h>

typedef struct
{
    char name[8];
    int age;
} Stu;

int main()
{
    int shm_id, i;
    key_t key;
    Stu *smap;
    struct shmid_ds buf;
    key = ftok("/", 0);
    if (key == -1)
    {
        perror("ftok error");
        exit(-1);
    }

    printf("key=%d\n", shm_id);
    smap = (Stu*)shmat(shm_id, NULL, 0);
    for (i = 0; i < 3; i++)
    {
        printf("name: %s\n", (*(smap + i)).name);
        printf("age : %d\n", (*(smap + i)).age);
    }

    if (shmdt(smap) == -1)
    {
        perror("detach error");
        return -1;
    }

    shmctl(shm_id, IPC_RMID, &buf);
    return 0;
}
