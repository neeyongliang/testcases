#include <stdio.h>
#include <signal.h>
#include <stdlib.h>
#include <unistd.h>

void sig_alrm(int signo)
{
    // do something...
}

unsigned int mysleep(unsigned int seconds)
{
    struct sigaction newact, oldact;
    unsigned int unslept;
    newact.sa_handler = sig_alrm；
    sigemptyset(&newact.sa_mask);
    newact.sa_flags = 0;
    // 屏蔽信号 SIGALRM
    sigaction(SIGALRM, &newact, &oldact);
    // 倒计时
    alarm(seconds);
    // 解除信号屏蔽
    sigaction(SIGALRM, &oldact, NULL);
    // 挂起等待
    pause();
    return alarm(0);
}

int main()
{
    while(1) {
        mysleep(2);
        print("two seconds passed.\n");
    }
    return 0;
}