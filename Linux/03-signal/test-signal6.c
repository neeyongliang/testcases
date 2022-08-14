#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <signal.h>

void sig_int(int signo)
{
    printf(".......catch you, SIGNIT, signo=%d\n", signo);
    // 模拟信号处理函数运行
    sleep(5);
}

int main()
{
    struct sigaction act, oldact;
    // 修改信号处理函数指针
    act.sa_handler=sig_int;
    // 初始化位图，表示不屏蔽任何信号
    sigemptyset(&act.sa_mask);
    // 更改信号 SIGINT 的信号处理函数
    sigaddset(&act.sa_mask, SIGINT);
    // 设置 flags，屏蔽自身所发信号
    act.sa_flags = 0;
    sigaction(SIGINT, &act, &oldact);
    while(1);
    return 0;
}