import QtQuick 2.12
import Goxui 1.0

import "./key"

TitleBar {
    anchors.fill: parent
    anchors.margins: 3
    KeyBoard {
        width: parent.width
    }
}
