#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
    pid_t pid;
    pid = fork();
    if (pid == -1)
    {
        perror("fork error");
        exit(-1);
    }
    else if (pid > 0)
    {
        printf("parent process:pid=%d\n", getpid());
    }
    else if (pid == 0)
    {
        printf("child process: pid=%d\n", getpid());
        // execl("/bin/ls", "-a", "-l", "test-fork.c", NULL);
        // execlp("ls", "-a", "test-fork.c", NULL);
        char *arg[]={"-a", "-l", "test-fork.c", NULL};
        execvp("ls", arg);
        perror("error exec\n");
        printf("child process: pid=%d\n", getpid());
    }

    return 0;
}