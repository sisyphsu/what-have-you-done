import QtQuick 2.12

Item {
    id: root

    property double size: 1

    property alias code: content.code
    property alias label: content.label
    property alias score: content.score
    property alias fontSize: content.fontSize

    width: parent.width * (root.size / parent.size)
    height: parent.height

    KeyContent {
        id: content
    }

}
