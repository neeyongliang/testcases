// gcc testGregex.c -o testGregex `pkg-config --libs --cflags glib-2.0`
#include <stdio.h>
#include <glib.h>

static gboolean
is_gvfs_mount_error (gchar *target_path, gchar *error_namespace)
{
	gboolean is_gvfs_mount = FALSE;
	g_return_val_if_fail ((target_path != NULL), is_gvfs_mount);
	GRegex *regex = g_regex_new (error_namespace, 0, 0, NULL);
	GMatchInfo *match_info;
	// g_assert (regex != NULL);
	g_regex_match (regex, target_path, 0, &match_info);
	if (g_match_info_matches(match_info))
	{
		is_gvfs_mount = TRUE;
	}

	g_match_info_free (match_info);
	g_regex_unref (regex);
	return is_gvfs_mount;
}

int main(int argc, char *argv[]) {
	gboolean res = is_gvfs_mount_error ("ftp://sdfsdgsdg.com", "^(smb://|ftp://)");
	g_print ("File: %s, res is ---- %d\n", __FILE__, res);
	return 0;
}

