// gcc main.c -o testGVfs `pkg-config --libs --cflags gio-unix-2.0`
#include <gio/gio.h>
#include <stdio.h>

int main(int argc, char const *argv[])
{
    int i;
    GVfs *vfs = g_vfs_get_default();
    const gchar * const * supported;
    supported = g_vfs_get_supported_uri_schemes (vfs);
    for (i = 0; supported[i] != NULL; i++) {
        g_print ("uri->%s is be supported!\n", supported[i]);
    }
    return 0;
}
