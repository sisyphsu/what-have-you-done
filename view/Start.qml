import QtQuick 2.12
import QtQuick.Window 2.12
import Qt.labs.platform 1.1
import Goxui 1.0

Window {
    id: root

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

    Component.onCompleted: {
        x = Screen.width / 2 - width / 2
        y = Screen.height / 2 - height / 2
    }

    Main {
    }

    SystemTrayIcon {
        visible: true
        iconSource: "images/icon.png"
        menu: Menu {
            MenuItem {
                text: root.active ? "Hide" : "Show"
                shortcut: "Ctrl+Shift+V"
                onTriggered: root.triggle()
            }
            MenuSeparator {
            }
            MenuItem {
                text: qsTr("Quit")
                onTriggered: Qt.quit()
            }
        }
    }

    HotKey {
        sequence: "ctrl+shift+V"
        onActivated: root.triggle()
    }

    function triggle() {
        if (root.active) {
            root.hide()
        } else {
            root.show()
            root.raise()
        }
    }
}
