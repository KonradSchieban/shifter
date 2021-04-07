/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package generator

import (
	//"encoding/json"
	"fmt"
	//runtime "k8s.io/apimachinery/pkg/runtime"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"os"
	//"log"
	"bufio"
	"log"
	"shifter/lib"
	"strconv"
)

func Yaml(path string, objects []lib.K8sobject) {

	// Create our output folder
	createFolder(path)

	// Iterate over our objects to write out
	for k, v := range objects {
		no := strconv.Itoa(k)

		kind := fmt.Sprintf("%v", v.Kind)

		f, err := os.Create(path + "/" + no + "-" + kind + ".yaml")
		if err != nil {
			fmt.Println(err)
		}

		log.Println("Creating file ", f.Name())

		defer f.Close()

		w := bufio.NewWriter(f)

		e := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, nil, nil)
		err = e.Encode(v.Object, w)
		if err != nil {
			fmt.Println(err)
		}
		w.Flush()
	}
}
