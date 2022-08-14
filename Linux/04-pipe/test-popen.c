#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

int main()
{
    FILE *r_fp, *w_fp;
    char buf[100];
    // 读取输出
    r_fp = popen("ls", "r");
    // 将结果写入管道
    w_fp = popen("wc -l", "w");
    while(fgets(buf, sizeof(buf), r_fp) != NULL)
    {
        fputs(buf, w_fp);
    }
    pclose(r_fp);
    pclose(w_fp);
    return 0;
}