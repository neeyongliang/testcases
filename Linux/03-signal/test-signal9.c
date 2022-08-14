#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>

void sys_error(char *str)
{
    perror(str);
    exit(1);
}

void do_sig_child(int signo)
{
    waitpid(0, NULL, WNOHANG);
}

int main(void)
{
    pid_t pid;
    int i;
    for (i = 0; i < 5; i++)
    {
        if ((pid = fork()) == 0)
        {
            break;
        }
        else if (pid < 0)
        {
            sys_error("fork");
        }
    }

    if (pid == 0)
    {
        // 子进程
        int n = 1;
        while (n--) {
            printf("child ID %d\n", getpid());
        }
        exit(i+1);
    }
    else if (pid > 0)
    {
        struct  sigaction act;
        act.sa_handler = do_sig_child;
        sigemptyset(&act.sa_mask);
        act.sa_flags = 0;
        // SIGCHILD 属于不可靠进程，没有队列
        sigaction(SIGCHLD, &act, NULL);
        while(1)
        {
            printf("parent ID %d\n", getpid());
            sleep(1);
        }
    }
    return 0;
}