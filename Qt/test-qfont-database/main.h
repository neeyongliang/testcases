#ifndef TST_QFONTDATABASE_H
#define TST_QFONTDATABASE_H

#include <QObject>

class tst_QFontDatabase : public QObject
{
    Q_OBJECT

public:
    tst_QFontDatabase();
    virtual ~tst_QFontDatabase();

public slots:
    void init();
    void cleanup();
    void printList_data();
    void printList(QStringList fonts);
private slots:
    void addAppFont();
};

#endif // TST_QFONTDATABASE_H
