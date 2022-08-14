#include <stdio.h>
#include <signal.h>
#include <stdio.h>

void sig_alrm(int signo)
{
    // do something...
}

unsigned int mysleep(unsigned int seconds)
{
    struct sigaction newact, oldact;
    sigset_t newmask, oldmask, suspmask;
    unsigned int unslept;
    // 1. 为 SIGALRM 设置捕获函数
    newact.sa_handler = sig_alrm;
    sigemptyset(&newact.sa_mask);
    newact.sa_flags = 0;
    sigaction(SIGALRM, &newact, &oldact);
    // 2. 设置阻塞信号集，屏蔽 SIGALRM 信号
    sigemptyset(&newmask);
    sigaddset(&newmask, SIGALRM);
    sigprocmask(SIG_BLOCK, &newmask, &oldmask);
    // 3. 设置计时器
    alarm(seconds);
    // 4. 构造临时阻塞信号集
    suspmask = oldmask;
    sigdelset(&suspmask, SIGALRM);
    // 5. 采用临时阻塞信号集替换原有的阻塞信号集合（不包含 SIGALRM 信号）
    sigsuspend(&suspmask);
    unslept=alarm(0);
    // 6. 恢复 SIGALRM 原有动作
    sigaction(SIGALRM, &oldact, NULL);
    // 7. 解除对 SIGALRM 的屏蔽，呼应注释2
    sigprocmask(SIG_SETMASK, &oldmask, NULL);
    return unslept;
}

int main()
{
    while (1)
    {
        mysleep(2);
        printf("two seconds passed\n");
    }
    return 0;
}