#include <stdlib.h>
#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

void *tfn(void *arg)
{
    while(1)
    {
        printf("child thread...\n");
        // 设置取消点
        pthread_testcancel();
    }
}

int main(void)
{
    pthread_t tid;
    void *tret = NULL;
    pthread_create(&tid, NULL, tfn, NULL);
    sleep(1);
    pthread_cancel(tid);
    pthread_join(tid, &tret);
    printf("child thread exit code = %ld\n", (long int)tret);
    return 0;
}