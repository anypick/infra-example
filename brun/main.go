package main

import (
	_ "example"
	_ "example/src"
	"flag"
	"fmt"
	"github.com/anypick/infra"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/utils/common"
)

func main() {
	// 生成网站：http://patorjk.com/software/taag
	banner := `
                    .__                                                  __  .__               
_______ __ __  _____|  |__           _____________  ____   _____   _____/  |_|__| ____   ____  
\_  __ \  |  \/  ___/  |  \   ______ \____ \_  __ \/  _ \ /     \ /  _ \   __\  |/  _ \ /    \ 
 |  | \/  |  /\___ \|   Y  \ /_____/ |  |_> >  | \(  <_> )  Y Y  (  <_> )  | |  (  <_> )   |  \
 |__|  |____//____  >___|  /         |   __/|__|   \____/|__|_|  /\____/|__| |__|\____/|___|  /
                  \/     \/          |__|                      \/                           \/
	`
	fmt.Println(banner)
	profile := flag.String("profile", "", "环境信息")
	flag.Parse()
	resource := ""
	if common.StrIsBlank(*profile) {
		resource = "resources/application.yml"
	} else {
		resource = fmt.Sprintf("resources/application-%s.yml", *profile)
	}
	yamlConf := props.NewYamlSource(resource)
	application := infra.New(*yamlConf)
	application.Start()
}
