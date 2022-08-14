#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <signal.h>

// 自定义信号处理
void sig_int(int signo)
{
    printf(".....catch you, SIGNAL\n");
    // 信号处理函数执行
    signal(SIGINT, SIG_DFL);
}

int main()
{
    // 捕获信号，修改信号处理函数
    signal(SIGINT, sig_int);
    while(1);
    return 0;
}