package sdlcanvas

import "github.com/veandco/go-sdl2/sdl"

var keyMap [262]string

func init() {
	keyMap[sdl.SCANCODE_ESCAPE] = "Escape"
	keyMap[sdl.SCANCODE_0] = "Digit0"
	keyMap[sdl.SCANCODE_1] = "Digit1"
	keyMap[sdl.SCANCODE_2] = "Digit2"
	keyMap[sdl.SCANCODE_3] = "Digit3"
	keyMap[sdl.SCANCODE_4] = "Digit4"
	keyMap[sdl.SCANCODE_5] = "Digit5"
	keyMap[sdl.SCANCODE_6] = "Digit6"
	keyMap[sdl.SCANCODE_7] = "Digit7"
	keyMap[sdl.SCANCODE_8] = "Digit8"
	keyMap[sdl.SCANCODE_9] = "Digit9"
	keyMap[sdl.SCANCODE_MINUS] = "Minus"
	keyMap[sdl.SCANCODE_EQUALS] = "Equal"
	keyMap[sdl.SCANCODE_BACKSPACE] = "Backspace"
	keyMap[sdl.SCANCODE_TAB] = "Tab"
	keyMap[sdl.SCANCODE_Q] = "KeyQ"
	keyMap[sdl.SCANCODE_W] = "KeyW"
	keyMap[sdl.SCANCODE_E] = "KeyE"
	keyMap[sdl.SCANCODE_R] = "KeyR"
	keyMap[sdl.SCANCODE_T] = "KeyT"
	keyMap[sdl.SCANCODE_Y] = "KeyY"
	keyMap[sdl.SCANCODE_U] = "KeyU"
	keyMap[sdl.SCANCODE_I] = "KeyI"
	keyMap[sdl.SCANCODE_O] = "KeyO"
	keyMap[sdl.SCANCODE_P] = "KeyP"
	keyMap[sdl.SCANCODE_LEFTBRACKET] = "BracketLeft"
	keyMap[sdl.SCANCODE_RIGHTBRACKET] = "BracketRight"
	keyMap[sdl.SCANCODE_RETURN] = "Enter"
	keyMap[sdl.SCANCODE_LCTRL] = "ControlLeft"
	keyMap[sdl.SCANCODE_A] = "KeyA"
	keyMap[sdl.SCANCODE_S] = "KeyS"
	keyMap[sdl.SCANCODE_D] = "KeyD"
	keyMap[sdl.SCANCODE_F] = "KeyF"
	keyMap[sdl.SCANCODE_G] = "KeyG"
	keyMap[sdl.SCANCODE_H] = "KeyH"
	keyMap[sdl.SCANCODE_J] = "KeyJ"
	keyMap[sdl.SCANCODE_K] = "KeyK"
	keyMap[sdl.SCANCODE_L] = "KeyL"
	keyMap[sdl.SCANCODE_SEMICOLON] = "Semicolon"
	keyMap[sdl.SCANCODE_APOSTROPHE] = "Quote"
	keyMap[sdl.SCANCODE_GRAVE] = "Backquote"
	keyMap[sdl.SCANCODE_LSHIFT] = "ShiftLeft"
	keyMap[sdl.SCANCODE_BACKSLASH] = "Backslash"
	keyMap[sdl.SCANCODE_Z] = "KeyZ"
	keyMap[sdl.SCANCODE_X] = "KeyX"
	keyMap[sdl.SCANCODE_C] = "KeyC"
	keyMap[sdl.SCANCODE_V] = "KeyV"
	keyMap[sdl.SCANCODE_B] = "KeyB"
	keyMap[sdl.SCANCODE_N] = "KeyN"
	keyMap[sdl.SCANCODE_M] = "KeyM"
	keyMap[sdl.SCANCODE_COMMA] = "Comma"
	keyMap[sdl.SCANCODE_PERIOD] = "Period"
	keyMap[sdl.SCANCODE_SLASH] = "Slash"
	keyMap[sdl.SCANCODE_RSHIFT] = "RightShift"
	keyMap[sdl.SCANCODE_KP_MULTIPLY] = "NumpadMultiply"
	keyMap[sdl.SCANCODE_LALT] = "AltLeft"
	keyMap[sdl.SCANCODE_SPACE] = "Space"
	keyMap[sdl.SCANCODE_CAPSLOCK] = "CapsLock"
	keyMap[sdl.SCANCODE_F1] = "F1"
	keyMap[sdl.SCANCODE_F2] = "F2"
	keyMap[sdl.SCANCODE_F3] = "F3"
	keyMap[sdl.SCANCODE_F4] = "F4"
	keyMap[sdl.SCANCODE_F5] = "F5"
	keyMap[sdl.SCANCODE_F6] = "F6"
	keyMap[sdl.SCANCODE_F7] = "F7"
	keyMap[sdl.SCANCODE_F8] = "F8"
	keyMap[sdl.SCANCODE_F9] = "F9"
	keyMap[sdl.SCANCODE_F10] = "F10"
	keyMap[sdl.SCANCODE_PAUSE] = "Pause"
	keyMap[sdl.SCANCODE_SCROLLLOCK] = "ScrollLock"
	keyMap[sdl.SCANCODE_KP_7] = "Numpad7"
	keyMap[sdl.SCANCODE_KP_8] = "Numpad8"
	keyMap[sdl.SCANCODE_KP_9] = "Numpad9"
	keyMap[sdl.SCANCODE_KP_MINUS] = "NumpadSubtract"
	keyMap[sdl.SCANCODE_KP_4] = "Numpad4"
	keyMap[sdl.SCANCODE_KP_5] = "Numpad5"
	keyMap[sdl.SCANCODE_KP_6] = "Numpad6"
	keyMap[sdl.SCANCODE_KP_PLUS] = "NumpadAdd"
	keyMap[sdl.SCANCODE_KP_1] = "Numpad1"
	keyMap[sdl.SCANCODE_KP_2] = "Numpad2"
	keyMap[sdl.SCANCODE_KP_3] = "Numpad3"
	keyMap[sdl.SCANCODE_KP_0] = "Numpad0"
	keyMap[sdl.SCANCODE_KP_DECIMAL] = "NumpadDecimal"
	keyMap[sdl.SCANCODE_PRINTSCREEN] = "PrintScreen"
	keyMap[sdl.SCANCODE_NONUSBACKSLASH] = "IntlBackslash"
	keyMap[sdl.SCANCODE_F11] = "F11"
	keyMap[sdl.SCANCODE_F12] = "F12"
	keyMap[sdl.SCANCODE_KP_EQUALS] = "NumpadEqual"
	keyMap[sdl.SCANCODE_F13] = "F13"
	keyMap[sdl.SCANCODE_F14] = "F14"
	keyMap[sdl.SCANCODE_F15] = "F15"
	keyMap[sdl.SCANCODE_F16] = "F16"
	keyMap[sdl.SCANCODE_F17] = "F17"
	keyMap[sdl.SCANCODE_F18] = "F18"
	keyMap[sdl.SCANCODE_F19] = "F19"
	keyMap[sdl.SCANCODE_UNDO] = "Undo"
	keyMap[sdl.SCANCODE_PASTE] = "Paste"
	keyMap[sdl.SCANCODE_AUDIOPREV] = "MediaTrackPrevious"
	keyMap[sdl.SCANCODE_CUT] = "Cut"
	keyMap[sdl.SCANCODE_COPY] = "Copy"
	keyMap[sdl.SCANCODE_AUDIONEXT] = "MediaTrackNext"
	keyMap[sdl.SCANCODE_KP_ENTER] = "NumpadEnter"
	keyMap[sdl.SCANCODE_RCTRL] = "ControlRight"
	keyMap[sdl.SCANCODE_MUTE] = "AudioVolumeMute"
	keyMap[sdl.SCANCODE_AUDIOPLAY] = "MediaPlayPause"
	keyMap[sdl.SCANCODE_AUDIOSTOP] = "MediaStop"
	keyMap[sdl.SCANCODE_VOLUMEDOWN] = "AudioVolumeDown"
	keyMap[sdl.SCANCODE_VOLUMEUP] = "AudioVolumeUp"
	keyMap[sdl.SCANCODE_KP_DIVIDE] = "NumpadDivide"
	keyMap[sdl.SCANCODE_RALT] = "AltRight"
	keyMap[sdl.SCANCODE_HELP] = "Help"
	keyMap[sdl.SCANCODE_HOME] = "Home"
	keyMap[sdl.SCANCODE_UP] = "ArrowUp"
	keyMap[sdl.SCANCODE_PAGEUP] = "PageUp"
	keyMap[sdl.SCANCODE_LEFT] = "ArrowLeft"
	keyMap[sdl.SCANCODE_RIGHT] = "ArrowRight"
	keyMap[sdl.SCANCODE_END] = "End"
	keyMap[sdl.SCANCODE_DOWN] = "ArrowDown"
	keyMap[sdl.SCANCODE_INSERT] = "Insert"
	keyMap[sdl.SCANCODE_DELETE] = "Delete"
	keyMap[sdl.SCANCODE_APPLICATION] = "ContextMenu"
}

func keyName(s sdl.Scancode) string {
	if int(s) >= len(keyMap) {
		return "Unidentified"
	}
	name := keyMap[s]
	if name == "" {
		return "Unidentified"
	}
	return name
}