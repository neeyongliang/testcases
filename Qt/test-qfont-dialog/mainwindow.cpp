#include <QtTest/QTest>

#include "mainwindow.h"
#include <QWidgetList>
#include <QApplication>
#include <QFontDialog>
#include <QTimer>
#include <QPushButton>
#include <iostream>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
{
    resize(600,600);
    QPushButton *btn = new QPushButton(tr("Change Font"), this);
    QPushButton *compare = new QPushButton(tr("Compare"), this);
    btn->move(200,200);
    btn->move(200,300);
    line = new QLineEdit(this);
    line->move(200,150);
    connect(btn, &QPushButton::clicked, this, &MainWindow::FontDlg);
    connect(compare, &QPushButton::clicked, this,&MainWindow::setFont);
}

MainWindow::~MainWindow()
{
}

void MainWindow::postKeyReturn()
{
}

void MainWindow::setFont()
{
    bool ok = false;
    QString fontName = "Arial";
    int fontSize = 24;

    QFont f1(fontName, fontSize);
    f1.setPixelSize(QFontInfo(f1).pixelSize());
    QFont f2 = QFontDialog::getFont(&ok, f1);
    std::cout << "QFontInfo:f2:pointSize->" << QFontInfo(f2).pointSize() << std::endl;
    std::cout << "QFontInfo:f1:pointSize->" << QFontInfo(f1).pointSize() << std::endl;
    QCOMPARE(QFontInfo(f2).pointSize(), QFontInfo(f1).pointSize());
}

void MainWindow::FontDlg()
{
    bool ok;
    QFont font = QFontDialog::getFont(&ok,this);
    if(ok)
    {
        line->setFont(font);
    }
    std::cout << "QFontInfo:font:pointSize->" << QFontInfo(font).pointSize() << std::endl;
}

QTEST_MAIN(MainWindow)
