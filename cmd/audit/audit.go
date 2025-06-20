package audit

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/meta"
	"sigs.k8s.io/yaml"

	"github.com/mohamedaminehamdi/kubepilot/cmd"
)

func New(opt cmd.Options) *cobra.Command {
	c := &cobra.Command{
		Use:          "audit TYPE NAME",
		Short:        "Audit a resource",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ns, _, _ := opt.ToRawKubeConfigLoader().Namespace()
			obj, err := opt.NewBuilder().
				NamespaceParam(ns).
				DefaultNamespace().
				Unstructured().
				ResourceNames(args[0], args[1]).
				Do().
				Object()
			if err != nil {
				return fmt.Errorf("get object: %w", err)
			}
			metaObj, err := meta.Accessor(obj)
			if err == nil {
				metaObj.SetManagedFields(nil)
			}
			data, err := yaml.Marshal(obj)
			if err != nil {
				return fmt.Errorf("marshal object: %w", err)
			}

			var buf bytes.Buffer
			_ = promptAudit.Execute(&buf, templateData{
				Data: string(data),
				Lang: opt.Lang,
			})

			return opt.NewChatGPTClient(" Auditing...").
				CreateCompletion(cmd.Context(), buf.String(), cmd.OutOrStdout())
		},
	}
	opt.AddFlags(c.Flags())
	return c
}
