#include <stdio.h>
#include <gtk/gtk.h>

static void
test_prop_changed (GtkWidget *window,
                   GParamFlags *pspec,
                   gpointer user_data)
{
    g_print ("\n Hello, World! \n");
}

int main(int argc,char *argv[])
{
    GtkWidget *window;  //定义一个构件指针

    gtk_init(&argc,&argv);  //初始化GTK环境

    window = gtk_window_new(GTK_WINDOW_TOPLEVEL);//新建一个标准的有框架窗口

    gulong id = g_signal_connect (window,
                                  "notify::scale-factor",
                                  G_CALLBACK (test_prop_changed),
                                  NULL);
    gtk_widget_show(window); //显示window

    gtk_main();//启动GTK

    return 1;
}