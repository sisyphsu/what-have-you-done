import QtQuick 2.12
import QtQuick.Window 2.12
import Goxui 1.0

Window {
    visible: true
    width: 960
    height: 480
    title: qsTr("what have you done")
    Loader {
        id: loader
        anchors.fill: parent
        source: "Main.qml"
        function reload() {
            // system.clearComponentCache();
            source = "";
            source = "Main.qml";
        }
    }
    HotKey {
        sequence: "F9"
        onActivated: loader.reload()
    }
}
