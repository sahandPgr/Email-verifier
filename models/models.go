package models

type DomainUrl struct {
	DomainUrl string `string:"domain-url"`
}

type Response struct {
	Domain      string `json:"domain"`
	HasMX       bool   `json:"has-mx"`
	HasSPF      bool   `json:"has-spf"`
	SPFRecord   string `json:"spf-record"`
	HasDMARC    bool   `json:"has-dmarc"`
	DMARCRecord string `json:"dmarc-record"`
}
