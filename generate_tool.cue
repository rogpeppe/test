@experiment(explicitopen)

package registry

import (
	"encoding/yaml"
	"tool/file"
)

issues!: "01-content-request.yml"!: _

command: generate: {
	for name, content in issues {
		file.Create
		filename: ".github/ISSUE_TEMPLATE/\(name)"
		contents: yaml.Marshal(content)
	}
}
