#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <stdlib.h>

void *tfn(void *arg)
{
    long int i;
    i = (long int)arg;
    if (i == 2)
    {
        pthread_exit(NULL);
    }

    sleep(i);
    printf("I'm %dth thread, Thread_ID = %lu\n", i+1, pthread_self());
    return NULL;
}

int main(int argc, char *argv[])
{
    long int n = 5, i;
    pthread_t tid;
    if (argc == 2)
    {
        n = atoi(argv[1]);
    }

    for (i = 0; i < n; i++)
    {
        pthread_create(&tid, NULL, tfn, (void *)i);
    }

    sleep(n);
    printf("I am main, I'm a thread!\n"
           "main_thread_ID = %lu\n", pthread_self());
    return 0;
}