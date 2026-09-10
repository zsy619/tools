package xio

import (
	"os"
	"runtime"
	"strings"
)

// IsWin 判断当前系统是否为 Windows。
func IsWin() bool {
	return runtime.GOOS == "windows"
}

// IsMac 判断当前系统是否为 Darwin（macOS）。
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux 判断当前系统是否为 Linux。
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsSupportColor 判断当前控制台是否支持彩色输出。
// 支持：linux、mac，或 Windows 下的 ConEmu、Cmder、putty、git-bash.exe
// 不支持：Windows 自带的 cmd.exe、powerShell.exe
func IsSupportColor() bool {
	// Support color: "TERM=xterm" "TERM=xterm-vt220" "TERM=xterm-256color" "TERM=screen-256color"
	// Don't support color: "TERM=cygwin"
	envTerm := os.Getenv("TERM")
	if strings.Contains(envTerm, "xterm") || strings.Contains(envTerm, "screen") {
		return true
	}

	// like on ConEmu software, e.g "ConEmuANSI=ON"
	if os.Getenv("ConEmuANSI") == "ON" {
		return true
	}

	// like on ConEmu software, e.g "ANSICON=189x2000 (189x43)"
	if os.Getenv("ANSICON") != "" {
		return true
	}
	return false
}

// IsSupport256Color 判断当前控制台是否支持 256 色输出。
func IsSupport256Color() bool {
	// "TERM=xterm-256color" "TERM=screen-256color"
	return strings.Contains(os.Getenv("TERM"), "256color")
}

// IsSupportTrueColor 判断当前控制台是否支持真彩色输出。
func IsSupportTrueColor() bool {
	// "COLORTERM=truecolor"
	return strings.Contains(os.Getenv("COLORTERM"), "truecolor")
}
