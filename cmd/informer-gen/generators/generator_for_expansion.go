/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generators

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
)

// genExpansion produces a file for a group client, e.g. ExtensionsInfomers for the extension group.
type genExpansion struct {
	generator.GoGenerator
	groupPackagePath string
	// types in a group
	types   []*types.Type
	imports namer.ImportTracker
}

// We only want to call GenerateType() once per group.
func (g *genExpansion) Filter(c *generator.Context, t *types.Type) bool {
	return len(g.types) == 0 || t == g.types[0]
}

func (g *genExpansion) Imports(c *generator.Context) (imports []string) {
	imports = append(imports, g.imports.ImportLines()...)
	// KCP specific
	imports = append(imports, "github.com/kcp-dev/logicalcluster/v3")
	imports = append(imports, "kcpcache \"github.com/kcp-dev/apimachinery/v2/pkg/cache\"")
	return
}

func (g *genExpansion) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	for _, t := range g.types {
		if _, err := os.Stat(filepath.Join(g.groupPackagePath, strings.ToLower(t.Name.Name+"_expansion.go"))); os.IsNotExist(err) {
			sw.Do(expansionInterfaceTemplate, t)
		}
	}
	return sw.Error()
}

var expansionInterfaceTemplate = `
type $.|public$ClusterExpansion interface {
	Cluster(logicalcluster.Name) $.|public$Informer
}
`
