#include <stdio.h>
#include <stdlib.h>
#include <sys/msg.h>
#define MAX_TEXT 512

struct my_msg_st {
    long int my_msg_type;
    char anytext[MAX_TEXT];
};

int main()
{
    int idx = 1;
    int msgid;
    struct my_msg_st data;
    long int msg_to_rcv = 0;

    if (msgid == -1)
    {
        perror("msgrcv err");
        exit(-1);
    }

    while(idx < 5)
    {
        if (msgrcv(msgid, (void *)&data, BUFSIZ, msg_to_rcv, 0) == -1)
        {
            perror("msgrcv err");
            exit(-1);
        }

        printf("msg type:%ld\n", data.my_msg_type);
        printf("msg content is:%s", data.anytext);
        idx++;
    }

    //删除消息队列
    if (msgctl(msgid, IPC_RMID, 0) == -1)
    {
        perror("msgctl err");
        exit(-1);
    }

    exit(0);
}