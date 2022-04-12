package DeploymentConf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
type Deployment struct {
	IP string `yaml:"ip"`
	Port string `yaml:"port"`
}
type Config struct {
	Deployment Deployment `yaml:"deployment"`
}
var deployment *Deployment
func getDeployment()error{
	data,err:=ioutil.ReadFile("./Config/config.yaml")
	if err!=nil{
		fmt.Println("Error happened when reading deployment config in function DeploymentConf.getDeployment()")
		fmt.Println(err)
		return err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err!=nil{
		fmt.Println("Error happened when unmarshalling deployment config in function DeploymentConf")
		fmt.Println(err)
		return err
	}
	deployment=&config.Deployment
	return nil
}
func getPort()(string,error){
	err:=getDeployment()
	if err!=nil{
		return "",err
	}
	return fmt.Sprintf("%s",deployment.Port),nil
}