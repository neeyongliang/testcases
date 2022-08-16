#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

// 最大连接数
#define MAXLINE 80
// 服务端口
#define SERV_PORT 6666

int main()
{
    // 定义服务端与客户端结构体
    struct sockaddr_in servaddr, cliaddr;
    socklen_t cliaddr_len;
    int listenfd, connfd;
    char buf[MAXLINE];
    char str[INET_ADDRSTRLEN];
    int i, n;
    listenfd = socket(AF_INET, SOCK_STREAM, 0);
    // 服务器端口清理
    bzero(&servaddr, sizeof(servaddr));
    servaddr.sin_family = AF_INET;
    servaddr.sin_addr.s_addr = htonl(INADDR_ANY);
    servaddr.sin_port = htons(SERV_PORT);
    // 将套接字文件与服务器端口绑定
    bind(listenfd, (struct sockaddr *)&servaddr, sizeof(servaddr));
    listen(listenfd, 20);
    printf("Accepting connections...\n");
    while(1)
    {
        cliaddr_len = sizeof(cliaddr);
        connfd = accept(listenfd, (struct sockaddr *)&cliaddr, &cliaddr_len);
        n = recv(connfd, buf, MAXLINE, 0);
        printf("received from %s at PORT %d\n",
                  inet_ntop(AF_INET, &cliaddr.sin_addr, str, sizeof(str)),
                  ntohs(cliaddr.sin_port));
        for (i = 0; i < n; i++)
            buf[i] = toupper(buf[i]);
        send(connfd, buf, n, 0);
        //关闭连接
        close(connfd);
    }
    return 0;
}
