#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

int main(int argc, char *argv[])
{
    if (argc < 2)
    {
        print("./a.out fifoname\n");
        exit(1);
    }
    // 判断 fifo 文件是否存在
    int ret = access(argv[1], F_OK);
    // 若 fifo 文件不存在就创建
    if (ret == -1)
    {
        int r = mkfifo(argv[1], 0664);
        if (r == -1)
        {
            perror("mkfifo");
            exit(1);
        }
        else
        {
            printf("fifo create success!\n");
        }
    }
    int fd = open(argv[1], O_WRONLY);

    while(1) {
        char *p = "hello, world!";
        write(fd, p, strlen(p) + 1);
        sleep(1);
    }
    close(fd);
    return 0;
}