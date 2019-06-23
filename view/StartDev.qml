import QtQuick 2.12
import QtQuick.Window 2.12
import Goxui 1.0

Window {
    visible: true
    width: 980
    height: 297
    minimumWidth: width
    minimumHeight: height
    maximumWidth: width
    maximumHeight: height
    title: qsTr("what have you done")
    Loader {
        id: loader
        anchors.fill: parent
        source: "Main.qml"
        function reload() {
            console.info("reload Main.qml");
            loader.source = "";
            system.clearComponentCache();
            loader.source = "Main.qml";
        }
    }
    HotKey {
        sequence: "F9"
        onActivated: loader.reload()
    }
}
