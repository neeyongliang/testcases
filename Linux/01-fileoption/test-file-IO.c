// gcc test-file-IO.c -o test-file-IO
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <unistd.h>

int main() {
    int fd = 0;
    char filename[20]="./test-file-IO.txt";
    fd = open(filename, O_RDWR|O_EXCL|O_TRUNC, S_IRWXG);;
    if (fd == -1) {
        perror("file open error.\n");
        exit(-1);
    }

    int i = 0, len = 0;
    char buf[100] = {0};
    while (i < 3)
    {
        scanf("%s", buf);
        len = strlen(buf);
        write(fd, buf, len);
        i++;
    }

    close(fd);
    printf("-----------------------------\n");

    fd = open(filename, O_RDONLY);
    if (fd == -1) {
        perror("file open again error.\n");
        exit(-1);
    }

    off_t f_size = 0;
    f_size = lseek(fd, 0, SEEK_END);         //获取文件长度
    lseek(fd, 0, SEEK_SET);                  //设置文件读写位置
    while (lseek(fd, 0, SEEK_CUR) != f_size) //读取文件
    {
        read(fd, buf, 1024);
        printf("%s\n", buf);
    }
    close(fd);
    return 0;
}
