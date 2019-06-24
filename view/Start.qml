import QtQuick 2.12
import QtQuick.Window 2.12
import Qt.labs.platform 1.1

Window {
    visible: true
    width: 980
    height: 297
    minimumWidth: width
    minimumHeight: height
    maximumWidth: width
    maximumHeight: height
    color: "transparent"
    flags: Qt.FramelessWindowHint
    title: qsTr("what have you done")

    Main {
    }

    SystemTrayIcon {
        visible: true
        iconSource: "images/icon.png"
        menu: Menu {
            MenuItem {
                text: qsTr("Quit")
                onTriggered: Qt.quit()
            }
        }
    }
}
