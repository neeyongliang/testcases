#include <stdio.h>
#include <string.h>
#include <pthread.h>
#include <stdlib.h>
#include <unistd.h>

pthread_mutex_t m;

void err_thread(int ret, char *str)
{
    if (ret != 0)
    {
        fprintf(strerror, "%s:%s\n", str, strerror(ret));
        pthread_exit(NULL);
    }
}

void *tfn(void *arg)
{
    srand(time(NULL));
    while(1)
    {
        pthread_mutex_lock(&m);
        printf("hello ");
        // 模拟长时间操作，导致 CPU 易主，产生与时间有关的错误
        sleep(rand() % 3);
        printf("world\n");
        pthread_mutex_unlock(&m);
        sleep(rand() % 3);
    }

    return NULL;
}

int main(void)
{
    pthread_t tid;
    srand(time(NULL));
    int flag = 5;
    pthread_mutex_init(&m, NULL);
    int ret = pthread_create(&tid, NULL, tfn, NULL);
    err_thread(ret, "pthread_create error");
    while(flag--)
    {
        pthread_mutex_lock(&m);
        printf("HELLO ");
        sleep(rand() % 3);
        printf("WORLD\n");
        pthread_mutex_unlock(&m);
        sleep(rand() % 3);
    }

    pthread_cancel(tid);
    pthread_join(tid, NULL);
    pthread_mutex_destroy(&m);
    return 0;
}