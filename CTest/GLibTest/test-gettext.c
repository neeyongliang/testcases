#include <stdio.h>
#include <locale.h>
#include "config.h"
#include <glib.h>
#include <glib/gi18n-lib.h>

int main(int argc, char const *argv[])
{
    g_print ("\nPACKAGE: %s, LOCALDIR:%s\n", GETTEXT_PACKAGE, LOCALEDIR);
    bindtextdomain (GETTEXT_PACKAGE, LOCALEDIR);
    setlocale(LC_ALL,"");
    bind_textdomain_codeset (GETTEXT_PACKAGE, "UTF-8");
    textdomain (GETTEXT_PACKAGE);

    g_print ("\n-->%s<--\n", _("Unknown"));
    return 0;
}