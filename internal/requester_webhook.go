package internal

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/validate-v1-namespace,mutating=false,failurePolicy=fail,sideEffects=None,groups="core",resources=namespaces,verbs=create;update;delete,versions=v1,name=vnamespace.kb.io,admissionReviewVersions=v1

// NamespaceLabelValidator is  used to handle the webhook in the main file
type RequesterValidator struct {
	Client  client.Client
	Decoder *admission.Decoder
}

// Handle checks if the namespacelabel name is the same as the namespace
func (hook *RequesterValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	log := ctrllog.Log.WithName("webhooks").WithName("InitInjector")
	name := req.UserInfo.Username
	log.Info("the namespace was changed", "username", name)

	return admission.Allowed("Namespace was modified")
}
