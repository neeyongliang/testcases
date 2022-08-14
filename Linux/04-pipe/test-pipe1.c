#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/types.h>

int main()
{
    // 文件描述符数组
    int fd[2];
    // 创建管道
    int ret = pipe(fd);
    if (ret == -1)
    {
        perror("pipe");
        exit(1);
    }
    pid_t pid = fork();
    // 处理父进程，写
    if (pid > 0)
    {
        // 关闭读端
        close(fd[0]);
        char *p="hello, pipe\n";
        // 写数据
        write(fd[1], p, strlen(p) + 1);
        close(fd[1]);
        wait(NULL);
    }
    else if (pid == 0)
    {
        // 处理子进程，读
        // 关闭写端
        close(fd[1]);
        char buf[64] = {0};
        ret = read(fd[0], buf, sizeof(buf));
        close(fd[0]);
        // 将读到的数据写入标准输出
        write(STDOUT_FILENO, buf, ret);
    }
    return 0;
}