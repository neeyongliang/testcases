#include "mainwindow.h"

#include <QApplication>
#include <QFile>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);    

    QFile qssFile(":/mystyle.qss");
    qssFile.open(QFile::ReadOnly);
    qApp->setStyleSheet(qssFile.readAll());
    qssFile.close();
    MainWindow w;
    w.show();
    return a.exec();
}
