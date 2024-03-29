package control

var KeyCodeMap = map[uint16]string{
	0x0001: "VC_ESCAPE",
	0x003B: "VC_F1",
	0x003C: "VC_F2",
	0x003D: "VC_F3",
	0x003E: "VC_F4",
	0x003F: "VC_F5",
	0x0040: "VC_F6",
	0x0041: "VC_F7",
	0x0042: "VC_F8",
	0x0043: "VC_F9",
	0x0044: "VC_F10",
	0x0057: "VC_F11",
	0x0058: "VC_F12",
	0x005B: "VC_F13",
	0x005C: "VC_F14",
	0x005D: "VC_F15",
	0x0063: "VC_F16",
	0x0064: "VC_F17",
	0x0065: "VC_F18",
	0x0066: "VC_F19",
	0x0067: "VC_F20",
	0x0068: "VC_F21",
	0x0069: "VC_F22",
	0x006A: "VC_F23",
	0x006B: "VC_F24",
	0x0029: "VC_BACKQUOTE",
	0x0002: "VC_1",
	0x0003: "VC_2",
	0x0004: "VC_3",
	0x0005: "VC_4",
	0x0006: "VC_5",
	0x0007: "VC_6",
	0x0008: "VC_7",
	0x0009: "VC_8",
	0x000A: "VC_9",
	0x000B: "VC_0",
	0x000C: "VC_MINUS",
	0x000D: "VC_EQUALS",
	0x000E: "VC_BACKSPACE",
	0x000F: "VC_TAB",
	0x003A: "VC_CAPS_LOCK",
	0x001E: "VC_A",
	0x0030: "VC_B",
	0x002E: "VC_C",
	0x0020: "VC_D",
	0x0012: "VC_E",
	0x0021: "VC_F",
	0x0022: "VC_G",
	0x0023: "VC_H",
	0x0017: "VC_I",
	0x0024: "VC_J",
	0x0025: "VC_K",
	0x0026: "VC_L",
	0x0032: "VC_M",
	0x0031: "VC_N",
	0x0018: "VC_O",
	0x0019: "VC_P",
	0x0010: "VC_Q",
	0x0013: "VC_R",
	0x001F: "VC_S",
	0x0014: "VC_T",
	0x0016: "VC_U",
	0x002F: "VC_V",
	0x0011: "VC_W",
	0x002D: "VC_X",
	0x0015: "VC_Y",
	0x002C: "VC_Z",
	0x001A: "VC_OPEN_BRACKET",
	0x001B: "VC_CLOSE_BRACKET",
	0x002B: "VC_BACK_SLASH",
	0x0027: "VC_SEMICOLON",
	0x0028: "VC_QUOTE",
	0x001C: "VC_ENTER",
	0x0033: "VC_COMMA",
	0x0034: "VC_PERIOD",
	0x0035: "VC_SLASH",
	0x0039: "VC_SPACE",
	0x0E37: "VC_PRINTSCREEN",
	0x0046: "VC_SCROLL_LOCK",
	0x0E45: "VC_PAUSE",
	0x0E52: "VC_INSERT",
	0x0E53: "VC_DELETE",
	0x0E47: "VC_HOME",
	0x0E4F: "VC_END",
	0x0E49: "VC_PAGE_UP",
	0x0E51: "VC_PAGE_DOWN",
	0xE048: "VC_UP",
	0xE04B: "VC_LEFT",
	0xE04C: "VC_CLEAR",
	0xE04D: "VC_RIGHT",
	0xE050: "VC_DOWN",
	0x0045: "VC_NUM_LOCK",
	0x0E35: "VC_KP_DIVIDE",
	0x0037: "VC_KP_MULTIPLY",
	0x004A: "VC_KP_SUBTRACT",
	0x0E0D: "VC_KP_EQUALS",
	0x004E: "VC_KP_ADD",
	0x0E1C: "VC_KP_ENTER",
	0x0053: "VC_KP_SEPARATOR",
	0x004F: "VC_KP_1",
	0x0050: "VC_KP_2",
	0x0051: "VC_KP_3",
	0x004B: "VC_KP_4",
	0x004C: "VC_KP_5",
	0x004D: "VC_KP_6",
	0x0047: "VC_KP_7",
	0x0048: "VC_KP_8",
	0x0049: "VC_KP_9",
	0x0052: "VC_KP_0",
	0x002A: "VC_SHIFT_L",
	0x0036: "VC_SHIFT_R",
	0x001D: "VC_CONTROL_L",
	0x0E1D: "VC_CONTROL_R",
	0x0038: "VC_ALT_L",
	0x0E38: "VC_ALT_R",
	0x0E5B: "VC_META_L",
	0x0E5C: "VC_META_R",
	0x0E5D: "VC_CONTEXT_MENU",
	0xE05E: "VC_POWER",
	0xE05F: "VC_SLEEP",
	0xE063: "VC_WAKE",
	0xE022: "VC_MEDIA_PLAY",
	0xE024: "VC_MEDIA_STOP",
	0xE010: "VC_MEDIA_PREVIOUS",
	0xE019: "VC_MEDIA_NEXT",
	0xE06D: "VC_MEDIA_SELECT",
	0xE02C: "VC_MEDIA_EJECT",
	0xE020: "VC_VOLUME_MUTE",
	0xE030: "VC_VOLUME_UP",
	0xE02E: "VC_VOLUME_DOWN",
	0xE06C: "VC_APP_MAIL",
	0xE021: "VC_APP_CALCULATOR",
	0xE03C: "VC_APP_MUSIC",
	0xE064: "VC_APP_PICTURES",
	0xE065: "VC_BROWSER_SEARCH",
	0xE032: "VC_BROWSER_HOME",
	0xE06A: "VC_BROWSER_BACK",
	0xE069: "VC_BROWSER_FORWARD",
	0xE068: "VC_BROWSER_STOP",
	0xE067: "VC_BROWSER_REFRESH",
	0xE066: "VC_BROWSER_FAVORITES",
	0x0070: "VC_KATAKANA",
	0x0073: "VC_UNDERSCORE",
	0x0077: "VC_FURIGANA",
	0x0079: "VC_KANJI",
	0x007B: "VC_HIRAGANA",
	0x007D: "VC_YEN",
	0x007E: "VC_KP_COMMA",
	0xFF75: "VC_SUN_HELP",
	0xFF78: "VC_SUN_STOP",
	0xFF76: "VC_SUN_PROPS",
	0xFF77: "VC_SUN_FRONT",
	0xFF74: "VC_SUN_OPEN",
	0xFF7E: "VC_SUN_FIND",
	0xFF79: "VC_SUN_AGAIN",
	0xFF7A: "VC_SUN_UNDO",
	0xFF7C: "VC_SUN_COPY",
	0xFF7D: "VC_SUN_INSERT",
	0xFF7B: "VC_SUN_CUT",
}
