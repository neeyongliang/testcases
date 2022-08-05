#include <QtTest/QTest>

#include <QFontDatabase>
#include <QStringList>
#include <QString>

#include <iostream>
#include "main.h"

tst_QFontDatabase::tst_QFontDatabase()
{
    std::cout << "tst_QFontDatabase" << std::endl;
}

tst_QFontDatabase::~tst_QFontDatabase()
{
    std::cout << "~tst_QFontDatabase" << std::endl;
}

void tst_QFontDatabase::init()
{
    std::cout << "init" << std::endl;
}

void tst_QFontDatabase::cleanup()
{
    std::cout << "cleanup" << std::endl;
}

void tst_QFontDatabase::printList(QStringList fonts)
{
    for (int i = 0; i < fonts.size(); ++i)
         std::cout << fonts.at(i).toLocal8Bit().constData() << std::endl;
}

void tst_QFontDatabase::printList_data()
{
    std::cout << "printList_data start" << std::endl;
    QTest::addColumn<QString>("list");
    std::cout << "printList_data ..." << std::endl;
    QTest::newRow("data") << NULL;
    std::cout << "printList_data end" << std::endl;
}

void tst_QFontDatabase::addAppFont()
{
    QFontDatabase db;

    const QStringList oldFamilies = db.families();
    QVERIFY(!oldFamilies.isEmpty());
    std::cout << "=====oldFamilies=====" << std::endl;
    printList(oldFamilies);

    QString FREE_MONO_FONT_PATH = "/usr/share/fonts/truetype/freefont/FreeMono.ttf";
    const int id = QFontDatabase::addApplicationFont(FREE_MONO_FONT_PATH);
    if (id == -1)
    {
        QSKIP("Skip...", SkipSingle);
        return;
    }

    std::cout << "===== const id: " << id << "=====" << std::endl;
    const QStringList addedFamilies = QFontDatabase::applicationFontFamilies(id);
    QVERIFY(!addedFamilies.isEmpty());
    std::cout << "=====addedFamilies=====" << std::endl;
    printList(addedFamilies);

    const QStringList newFamilies = db.families();
    QVERIFY(!newFamilies.isEmpty());
    QVERIFY(newFamilies.count() >= oldFamilies.count());
    std::cout << "=====newFamilies=====" << std::endl;
    printList(newFamilies);

    for (int i = 0; i < addedFamilies.count(); ++i)
    {
        std::cout << "i:" << i << std::endl;
        // addedFamilies.at(i))
        QVERIFY(newFamilies.contains("FreeMono [GNU ]"));
    }

    QVERIFY(QFontDatabase::removeApplicationFont(id));
    QVERIFY(db.families() == oldFamilies);
}

QTEST_MAIN(tst_QFontDatabase)
//#include "tst_QFontDatabase.moc"
