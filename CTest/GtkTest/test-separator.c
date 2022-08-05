// gcc test-separator.c -o test-separator `pkg-config --cflags --libs gtk+-3.0`
#include <gtk/gtk.h>
#include <locale.h>
#include <stdio.h>

#define CSS_PATH "./test.css"

static void add_class()
{
    static GtkCssProvider *provider;
    GdkScreen *screen;
    GError *error = NULL;
    screen = gdk_screen_get_default();

    provider = gtk_css_provider_new ();
    gtk_css_provider_load_from_path (provider, CSS_PATH, &error);
    if (error) {
        g_print ("\nError Occur, becase:%s\n", error->message);
        return;
    }

    gtk_style_context_add_provider_for_screen (screen,
        GTK_STYLE_PROVIDER (provider),
        GTK_STYLE_PROVIDER_PRIORITY_APPLICATION);
}

int main(int argc, char *argv[])
{
    setlocale(LC_ALL, "zh_CN.UTF-8");

    GtkWidget *window; //定义一个构件指针
    GtkWidget *box;

    gtk_init(&argc, &argv); //初始化GTK环境

    add_class();

    window = gtk_window_new(GTK_WINDOW_TOPLEVEL); //新建一个标准的有框架窗口
    gtk_window_set_default_size (GTK_WINDOW (window), 400, 200);
    gtk_window_set_position (GTK_WINDOW (window), GTK_WIN_POS_CENTER);
    gtk_style_context_add_class (gtk_widget_get_style_context (window), "hehehe");

    box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 10);
    gtk_widget_set_margin_start (box, 10);
    gtk_widget_set_margin_top (box, 10);
    gtk_widget_set_margin_end (box, 10);
    gtk_widget_set_margin_bottom (box, 10);
    gtk_container_add(GTK_CONTAINER(window), box);
    gtk_widget_set_name (box, "hihihi");

    gtk_widget_show_all(window); //显示window

    gtk_main(); //启动GTK

    return 1;
}
