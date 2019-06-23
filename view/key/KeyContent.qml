import QtQuick 2.12
import QtGraphicalEffects 1.12
import Goxui 1.0

Item {

    id: root

    property color front: "#2a2c2f"
    property color background: "#e1e1e1"
    property string label: ""
    property string code: ""

    property int fontSize: 14
    property double score: Math.random()
    property alias align: label.horizontalAlignment

    anchors.fill: parent
    anchors.margins: 3

    Rectangle {
        id: bg
        color: root.background
        anchors.fill: parent
        radius: 3
    }

    Rectangle {
        id: tm
        color: "#ec6762"
        anchors.fill: parent
        opacity: 0
        radius: 3
    }

    Text {
        id: label
        text: root.label
        color: root.front
        font.pixelSize: root.fontSize
        anchors.fill: parent
        anchors.margins: 18
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
    }

    Component.onCompleted: {
        tm.opacity = KeyStat.GetScore(root.code)
    }

    Event {
        key: root.code
        enable: true
        onActive: function (data) {
            tm.opacity = data
        }
    }
}
