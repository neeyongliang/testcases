#include <stdio.h>
#include <glib.h>
#include <locale.h>

gchar *get_samba_path(gchar *uri) {
	gchar **path_info, **samba_info, **share_info;
	gchar *cmd = NULL;
	int i=0;
	gchar *server = NULL, *share = NULL;
	path_info = g_strsplit_set(uri, ":", -1);
	g_print("path_info length: %d\n", g_strv_length(path_info));
	g_print("path_info_details: %s\n", path_info[1]);

	if (g_strv_length(path_info) <= 1) {
		return cmd;
	}

	samba_info = g_strsplit_set(path_info[1], ",=", -1);
	g_print("samba_info length: %d\n", g_strv_length(samba_info));
	while (i < g_strv_length(samba_info)) {
		g_print ("===> %s", samba_info[i]);
		if (g_strcmp0 (samba_info[i], "server") == 0) {
			server = samba_info[i+1];
			g_print (":%s\n", server);
			i = i + 2;
			continue;
		}

		if (g_strcmp0 (samba_info[i], "share") == 0) {
			share_info = g_strsplit_set(samba_info[i+1], "/", -1);
			share = share_info[0];
			g_print (":%s\n", share);
			i = i + 2;
			continue;
		}

		i = i + 2;
	}
	if (server && share) {
		cmd = g_strdup_printf("smb://%s/%s", server, share);
	}

	if (share_info) {
		g_strfreev(share_info);
	}

	if (samba_info) {
		g_strfreev(samba_info);
	}

	if (path_info) {
		g_strfreev(path_info);
	}

	return cmd;
}

int main(int argc, char const *argv[])
{
	setlocale(LC_ALL, "zh_CN");
	gchar *str_1 = "/run/user/1000/gvfs/smb-share:server=192.168.1.1,share=share_for_all/Test/阿里巴巴Java开发手册.pdf";
	gchar *str_2 = "/run/user/1000/gvfs/smb-share:domain=WORKGROUP,server=192.168.1.1,share=共享,user=niyl/C++简体中文版.PDF";

	gchar *str_3 = "/run/user/1000/gvfs/smb-server:server=192.168.1.1,share=share_for_all/Test/阿里巴巴Java开发手册.pdf";
	gchar *str_4 = "/run/user/1000/gvfs/smb-server:domain=WORKGROUP,server=192.168.1.1,share=共享,user=niyl/C++简体中文版.PDF";

	gchar *str_5 = "/run/user/1000/gvfs/smb-network:server=192.168.1.1,share=share_for_all/Test/阿里巴巴Java开发手册.pdf";
	gchar *str_6 = "/run/user/1000/gvfs/smb-network:domain=WORKGROUP,server=192.168.1.1,share=共享,user=niyl/C++简体中文版.PDF";

	gchar *cmd = get_samba_path(str_2);
	g_print ("cmd is %s", cmd);

	return 0;
}