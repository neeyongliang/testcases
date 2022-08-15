#include <stdio.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <sys/types.h>
#include <unistd.h>
#include <string.h>
#define SEGSIZE 4096

typedef struct {
    char name[8];
    int age;
} Stu;

int main()
{
    int shm_id, i;
    key_t key;
    char name[8];
    Stu *smap;
    // 获取关键字
    key = ftok("/", 0);
    if (key == -1)
    {
        perror("ftok error");
        return exit(-1);
    }
    printf("key=%d\n", key);
    shm_id = shmget(key, SEGSIZE, IPC_CREAT | IPC_EXCL | 0664);
    if (shm_id == -1)
    {
        perror("create shared memory error\n");
        return -1;
    }

    printf("shm_id=%d\n", shm_id);
    // 将线程与共享内存绑定
    smap = (Stu*)shmat(shm_id, NULL, 0);
    memset(name, 0x00, sizeof(name));
    strcpy(name, "Jorn");
    name[4] = '0';
    for (i = 0; i < 3; i++)
    {
        name[4] += 1;
        strncpy((smap + i)->name, name, 5);
        (smap + i)->age = 20 + i;
    }

    if (shmdt(smap) == -1)
    {
        perror("detach error");
        return -1;
    }

    return 0;
}