#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QLineEdit>

class MainWindow : public QMainWindow
{
    Q_OBJECT
public:
    explicit MainWindow(QWidget *parent = 0);
    virtual ~MainWindow();
    void setFont();
signals:

public slots:
    void FontDlg();
    void postKeyReturn();

private:
    QLineEdit *line;
};

#endif // MAINWINDOW_H
