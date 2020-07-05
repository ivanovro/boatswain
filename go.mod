module boatswain

go 1.14

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190515023456-b74e4c97951f

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190515063710-7b18d6600f6b

replace helm.sh/helm => github.com/alauda/helm v3.0.0-alpha.1.0.20190829021852-0235ba407f6d+incompatible

require (
	github.com/alauda/helm-crds v0.0.0-20200520071325-ff5c5e248d83
	github.com/imdario/mergo v0.3.9 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/apimachinery v0.0.0-20190624085041-961b39a1baa0
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/kube-openapi v0.0.0-20190603182131-db7b694dc208 // indirect
	k8s.io/utils v0.0.0-20190607212802-c55fbcfc754a // indirect
)
