package waf

import (
	"crypto/tls"
	"github.com/letheliu/hhjc-devops/entity/dto/waf"
)

var wafData WafData

type WafData struct {
	wafDto   waf.WafDto
	routes   []*waf.WafRouteDto
	wafCerts []WafCert
	rules    []*waf.WafRuleDataDto
}

type WafCert struct {
	Hostname string
	TlsCert  *tls.Certificate
}
