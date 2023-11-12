package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Docker", "Kubernetes", "Openshift", "Swarm"}
	courseName = append(courseName, "HTML", "CSS", "JS", "ReactJS", "VueJS")

	courseWeb := courseName[4:]

	fmt.Println(courseName)
	fmt.Println(courseWeb)
}
