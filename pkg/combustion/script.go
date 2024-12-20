package combustion

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/suse-edge/edge-image-builder/pkg/template"
)

//go:embed templates/script-base.sh.tpl
var combustionScriptBase string

func assembleScript(scripts []string, networkScript string, imageType string) (string, error) {
	slices.Sort(scripts)

	values := struct {
		NetworkScript string
		Scripts       []string
		ImageType     string
	}{
		NetworkScript: networkScript,
		Scripts:       scripts,
		ImageType:     imageType,
	}

	data, err := template.Parse("combustion-base", combustionScriptBase, values)
	if err != nil {
		return "", fmt.Errorf("parsing combustion base template: %w", err)
	}

	return data, nil
}
