#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <unistd.h>

int main()
{
    pid_t pid, w;
    pid = fork();
    if (pid == -1)
    {
        perror("fork error");
        exit(1);
    }
    else if (pid == 0) {
        // 子进程会在 3 秒后退出
        sleep(3);
        printf("child process: pid=%d\n", getpid());
        exit(0);
    }
    else if (pid > 0) {
        do {
            // 父进程接收到子进程退出，被 waitpid 捕获
            w = waitpid(pid, NULL, WNOHANG);
            if (w == 0) {
                printf("no child exited\n");
                sleep(1);
            }
        }while(w == 0);
        if (w == pid) {
            printf("catch a child process: pid=%d\n", w);
        } else {
            printf("waitpid error\n");
        }
    }
    return 0;
}