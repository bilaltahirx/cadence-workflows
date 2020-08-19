package activites

import (
	"context"
	"fmt"
	"go.uber.org/cadence/activity"
)
func init()  {

	activity.Register(InsertUserToPg)
	return
}



func InsertUserToPg(ctx context.Context, podsRequest string) (*string, error) {


	fmt.Println(podsRequest)
	//logger := activity.GetLogger(ctx)
	//client := resty.New()
	//resp, err := client.R().
	//	SetBody(podsRequest).
	//	SetHeader("Content-Type", "application/json").
	//	Post("http://" + env.Env.GoWebServerIp + ":" + env.Env.GoWebServerPort + "/opr/insert_tenant_pods")
	//if err != nil {
	//	logger.Error(err.Error())
	//	return nil, err
	//}
	//if resp.IsError() {
	//	logger.Error(string(resp.Body()))
	//	return nil, errors.New(string(resp.Body()))
	//}
	//
	//var companyResponse model.CompanyResponse
	//json.Unmarshal(resp.Body(), &companyResponse)
	return &podsRequest, nil
}

func HelloWodld(ctx context.Context, podsRequest string) (*string, error) {


	fmt.Println(podsRequest)
	//logger := activity.GetLogger(ctx)
	//client := resty.New()
	//resp, err := client.R().
	//	SetBody(podsRequest).
	//	SetHeader("Content-Type", "application/json").
	//	Post("http://" + env.Env.GoWebServerIp + ":" + env.Env.GoWebServerPort + "/opr/insert_tenant_pods")
	//if err != nil {
	//	logger.Error(err.Error())
	//	return nil, err
	//}
	//if resp.IsError() {
	//	logger.Error(string(resp.Body()))
	//	return nil, errors.New(string(resp.Body()))
	//}
	//
	//var companyResponse model.CompanyResponse
	//json.Unmarshal(resp.Body(), &companyResponse)
	return &podsRequest, nil
}

