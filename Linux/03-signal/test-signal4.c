#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <signal.h>

void printset(sigset_t *ped)
{
    // pending 打印函数
    int i;
    for (i = 1; i < 32; i++)
    {
        if (sigismember(ped, i) == 1)
        {
            putchar('1');
        }
        else
        {
            putchar('0');
        }
    }
    printf("\n");
}

int main()
{
    // 定义信号集
    sigset_t set, oldset, ped;
    // 初始化自定义信号集 set
    sigemptyset(&set);
    // 将 2 号信号 SIGINT 加入到 set
    sigaddset(&set, SIGINT);
    sigprocmask(SIG_BLOCK, &set, &oldset);
    while(1)
    {
        sigpending(&ped);
        printset(&ped);
        sleep(1);
    }
    return 0;
}
// 00000000000000000000000000000000
// Ctrl + C
// 01000000000000000000000000000000
