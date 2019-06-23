import QtQuick 2.12

Column {

    id: root

    KeyRow {
        height: 40
        KeyButton {
            label: "Esc"
            code: "VC_ESCAPE"
        }
        KeyButton {
            label: "F1"
            code: "VC_F1"
        }
        KeyButton {
            label: "F2"
            code: "VC_F2"
        }
        KeyButton {
            label: "F3"
            code: "VC_F3"
        }
        KeyButton {
            label: "F4"
            code: "VC_F4"
        }
        KeyButton {
            label: "F5"
            code: "VC_F5"
        }
        KeyButton {
            label: "F6"
            code: "VC_F6"
        }
        KeyButton {
            label: "F7"
            code: "VC_F7"
        }
        KeyButton {
            label: "F8"
            code: "VC_F8"
        }
        KeyButton {
            label: "F9"
            code: "VC_F9"
        }
        KeyButton {
            label: "F10"
            code: "VC_F10"
        }
        KeyButton {
            label: "F11"
            code: "VC_F11"
        }
        KeyButton {
            label: "F12"
            code: "VC_F12"
        }
    }
    KeyRow {
        KeyButton {
            label: "~"
            code: "VC_BACKQUOTE"
        }
        KeyButton {
            label: "1"
            code: "VC_1"
        }
        KeyButton {
            label: "2"
            code: "VC_2"
        }
        KeyButton {
            label: "3"
            code: "VC_3"
        }
        KeyButton {
            label: "4"
            code: "VC_4"
        }
        KeyButton {
            label: "5"
            code: "VC_5"
        }
        KeyButton {
            label: "6"
            code: "VC_6"
        }
        KeyButton {
            label: "7"
            code: "VC_7"
        }
        KeyButton {
            label: "8"
            code: "VC_8"
        }
        KeyButton {
            label: "9"
            code: "VC_9"
        }
        KeyButton {
            label: "0"
            code: "VC_0"
        }
        KeyButton {
            label: "-"
            code: "VC_MINUS"
        }
        KeyButton {
            label: "+"
            code: "VC_EQUALS"
        }
        KeyButtonWide {
            label: "Delete"
            code: "VC_BACKSPACE"
            size: 1.5
            align: Text.AlignRight
        }
    }
    KeyRow {
        KeyButtonWide {
            label: "Tab"
            code: "VC_TAB"
            size: 1.4
            align: Text.AlignLeft
        }
        KeyButton {
            label: "Q"
            code: "VC_Q"
        }
        KeyButton {
            label: "W"
            code: "VC_W"
        }
        KeyButton {
            label: "E"
            code: "VC_E"
        }
        KeyButton {
            label: "R"
            code: "VC_R"
        }
        KeyButton {
            label: "T"
            code: "VC_T"
        }
        KeyButton {
            label: "Y"
            code: "VC_Y"
        }
        KeyButton {
            label: "U"
            code: "VC_U"
        }
        KeyButton {
            label: "I"
            code: "VC_I"
        }
        KeyButton {
            label: "O"
            code: "VC_O"
        }
        KeyButton {
            label: "P"
            code: "VC_P"
        }
        KeyButton {
            label: "["
            code: "VC_OPEN_BRACKET"
        }
        KeyButton {
            label: "]"
            code: "VC_CLOSE_BRACKET"
        }
        KeyButton {
            label: "\\"
            code: "VC_BACK_SLASH"
        }
    }
    KeyRow {
        KeyButtonWide {
            label: "Caps Lock"
            code: "VC_CAPS_LOCK"
            size: 1.6
            align: Text.AlignLeft
        }
        KeyButton {
            label: "A"
            code: "VC_A"
        }
        KeyButton {
            label: "S"
            code: "VC_S"
        }
        KeyButton {
            label: "D"
            code: "VC_D"
        }
        KeyButton {
            label: "F"
            code: "VC_F"
        }
        KeyButton {
            label: "G"
            code: "VC_G"
        }
        KeyButton {
            label: "H"
            code: "VC_H"
        }
        KeyButton {
            label: "J"
            code: "VC_J"
        }
        KeyButton {
            label: "K"
            code: "VC_K"
        }
        KeyButton {
            label: "L"
            code: "VC_L"
        }
        KeyButton {
            label: ";"
            code: "VC_SEMICOLON"
        }
        KeyButton {
            label: "'"
            code: "VC_QUOTE"
        }
        KeyButtonWide {
            label: "Enter"
            code: "VC_ENTER"
            size: 1.6
            align: Text.AlignRight
        }
    }
    KeyRow {
        KeyButtonWide {
            label: "Shift"
            code: "VC_SHIFT_L"
            size: 2
            align: Text.AlignLeft
        }
        KeyButton {
            label: "Z"
            code: "VC_Z"
        }
        KeyButton {
            label: "X"
            code: "VC_X"
        }
        KeyButton {
            label: "C"
            code: "VC_C"
        }
        KeyButton {
            label: "V"
            code: "VC_V"
        }
        KeyButton {
            label: "B"
            code: "VC_B"
        }
        KeyButton {
            label: "N"
            code: "VC_N"
        }
        KeyButton {
            label: "M"
            code: "VC_M"
        }
        KeyButton {
            label: ","
            code: "VC_COMMA"
        }
        KeyButton {
            label: "."
            code: "VC_PERIOD"
        }
        KeyButton {
            label: "/"
            code: "VC_SLASH"
        }
        KeyButtonWide {
            label: "Shift"
            code: "VC_SHIFT_R"
            size: 2
            align: Text.AlignRight
        }
    }
    KeyRow {
        KeyButton {
            label: "Control"
            code: "VC_CONTROL_L"
            size: 1.3
        }
        KeyButton {
            label: "Option"
            code: "VC_ALT_L"
            size: 1.3
        }
        KeyButton {
            label: "Command"
            code: "VC_META_L"
            size: 1.5
        }
        KeyButton {
            label: ""
            code: "VC_SPACE"
            size: 5
        }
        KeyButton {
            label: "Command"
            code: "VC_META_R"
            size: 1.5
        }
        KeyButton {
            label: "Option"
            code: "VC_ALT_R"
            size: 1.3
        }
        KeyButton {
            label: "Control"
            code: "VC_CONTROL_R"
            size: 1.3
        }
        KeyButton {
            label: "◀"
            code: "VC_LEFT"
            fontSize: 9
        }
        Column {

            property double size: 1

            width: parent.width * (size / parent.size)
            height: parent.height
            spacing: 3

            Item {
                width: parent.width
                height: (parent.height - parent.spacing) * 0.5
                KeyContent {
                    label: "▲"
                    code: "VC_UP"
                    fontSize: 9
                    anchors.bottomMargin: 0
                }
            }
            Item {
                width: parent.width
                height: (parent.height - parent.spacing) * 0.5
                KeyContent {
                    label: "▼"
                    code: "VC_DOWN"
                    fontSize: 9
                    anchors.topMargin: 0
                }
            }
        }
        KeyButton {
            label: "▶"
            code: "VC_RIGHT"
            fontSize: 9
        }
    }
}
