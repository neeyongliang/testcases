#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/sem.h>

union semu
{
    int val;
    struct semid_ds* buf;
    unsigned short* array;
    struct seminfo* _buf;
};

static int sem_id;

// 设置信号量
static int set_semval()
{
    union semu semu_union;
    semu_union.val = 1;
    if (semctl(sem_id, 0, SETVAL, semu_union) == -1)
    {
        return 0;
    }
    return 1;

}

// P 操作，获取信号量
static int semaphore_p()
{
    struct sembuf sem_b;
    sem_b.sem_num = 0;
    sem_b.sem_op = -1;
    sem_b.sem_flg = SEM_UNDO;
    if (semop(sem_id, &sem_b, 1) == -1)
    {
        perror("sem_p err");
        return 0;
    }
    return 1;
}

// V 操作，释放信号量
static int semaphore_v()
{
    struct sembuf sem_b;
    sem_b.sem_num = 0;
    sem_b.sem_op = 1;
    sem_b.sem_flg = SEM_UNDO;
    if (semop(sem_id, &sem_b, 1) == -1)
    {
        perror("sem_v err");
        return 0;
    }
    return 1;
}

// 删除信号量
static void del_semvalue()
{
    union semu sem_union;
    if (semctl(sem_id, 0, IPC_RMID, sem_union) == -1)
    {
        perror("del err");
    }
}

int main()
{
    int i;
    pid_t pid;
    char ch='C';
    sem_id = semget((key_t)1000, 1, 0644 | IPC_CREAT);
    if (sem_id == -1)
    {
        perror("sem_c err");
        exit(-1);
    }

    if(!set_semval())
    {
        perror("init err");
        exit(-1);
    }

    pid = fork();
    // 创建子进程
    if (pid == -1)
    {
        // 创建失败，删除信号量
        del_semvalue();
    }
    else if (pid == 0)
    {
        ch = 'Z';
    }
    else
    {
        ch = 'C';
    }

    // 设置随机数种子
    srand((unsigned int)getpid());
    for (i = 0; i < 8; i++)
    {
        semaphore_p();
        printf("%c", ch);
        fflush(stdout);
        sleep(rand()%4);
        printf("%c", ch);
        fflush(stdout);
        sleep(1);
        // 释放信号量
        semaphore_v();
    }

    if (pid > 0)
    {
        wait(NULL);
        del_semvalue();
    }

    printf("\nprocess %d finished.\n", getpid());
    return 0;
}