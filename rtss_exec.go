package main

import (
	"strings"
	"os/exec"
	"fmt"
)

const (
	RTSS_PARAM_TEXT = "--text"
	RTSS_PARAM_RESET = "--reset"
	RTSS_FMT_BR = "<Br>"
)

func execSakuRTSSCli_clean() {
    fmt.Printf("cmd path: %s\n", *rtss_cli_path)
	cmd := exec.Command(*rtss_cli_path, RTSS_PARAM_RESET)
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("cmd clean:\n%s\n", string(out))
        fmt.Printf("cmd err:\n%s\n", err)
    }
    fmt.Printf("cmd clean:\n%s\n", string(out))
}

func execSakuRTSSCli_text(list []Motar) {
    fmt.Printf("cmd path: %s\n", *rtss_cli_path)
	var builder strings.Builder
	for _, value := range list {
		builder.WriteString(fmt.Sprintf("%s, %s mil%s", value.Angle, value.Mil, RTSS_FMT_BR))
	}
	value := builder.String()

	cmd := exec.Command(*rtss_cli_path, RTSS_PARAM_TEXT, value)
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("cmd text:\n%s\n", string(out))
        fmt.Printf("cmd err:\n%s\n", err)
    }
    fmt.Printf("cmd text:\n%s\n", string(out))
}
