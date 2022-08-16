#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <pthread.h>
#include <semaphore.h>

#define NUM 5

// 全局数组实现环形队列
int queue[NUM];

// 空格子信号量，产品信号量
sem_t blank_number, product_number;

void *producer(void *arg)
{
    int i = 0;
    while(1)
    {
        // 生产者将空格子--，为0则阻塞等待
        sem_wait(&blank_number);
        // 生产一个产品
        queue[i] = rand() % 1000 + 1;
        printf("---Produce---%d\n", queue[i]);
        sem_post(&product_number);
        i = (i + 1) % NUM;
        sleep(rand() % 1);
    }
}

void *consumer(void *arg)
{
    int i = 0;
    while(1)
    {
        // 消费者将产品--，为0则阻塞
        sem_wait(&product_number);
        printf("-Consume---%d     %lu\n", queue[i], pthread_self());
        // 消费一个产品
        queue[i] = 0;
        // 消费之后，将空格子数++
        sem_post(&blank_number);
        i = (i + 1) % NUM;
        sleep(rand() % 1);
    }
}

int main(int argc, char *argv[])
{
    pthread_t pid, cid;
    sem_init(&blank_number, 0, NUM);
    sem_init(&product_number, 0, 0);
    pthread_create(&pid, NULL, producer, NULL);
    pthread_create(&cid, NULL, consumer, NULL);
    pthread_create(&cid, NULL, consumer, NULL);
    pthread_join(pid, NULL);
    pthread_join(cid, NULL);
    sem_destroy(&blank_number);
    sem_destroy(&product_number);
    return 0;
}