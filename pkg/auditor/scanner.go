package auditor

import (
	"os/exec"
	"strings"
)

type Control struct {
	ID             string
	Title          string
	ComplianceCheck string
	Remediation    string
	Automated      bool
}

func ScanControl(control Control) (bool, error) {
	cmd := exec.Command("sh", "-c", control.ComplianceCheck)
	err := cmd.Run()
	
	if err == nil {
		return true, nil // Controle em conformidade
	}
	return false, nil // Falha na verificação
}
