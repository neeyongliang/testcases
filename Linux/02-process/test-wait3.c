#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

int main()
{
    pid_t pid, p, w;
    pid = fork();
    if (pid == -1)
    {
        perror("fork1 error");
        exit(1);
    }
    else if (pid == 0)
    {
        sleep(5);
        printf("first child process: pid=%d\n", getpid());
    }
    else if (pid > 0)
    {
        // 父进程继续创建子进程
        int i;
        p = pid;
        for (i = 0; i < 3; i++)
        {
            if ((pid = fork()) == 0)
                break;
        }
        if (pid == -1) {
            perror("fork error\n");
            exit(0);
        }
        else if (pid == 0)
        {
            printf("child process:pid=%d\n", getpid());
            exit(0);
        }
        else if (pid > 0)
        {
            w = waitpid(p, NULL, 0);
            if (w == p)
            {
                printf("catch child process: %d\n", w);
            }
            else
            {
                printf("waitpid error\n");
            }
        }
    }
    return 0;
}