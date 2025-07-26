package main

import "fmt"

func topologicalSort(input map[string][]string) []string {
	// start off our todo list with all courses without prerequisites
	todo := NewDeque()
	prereq_to_courses := make(map[string][]string)
	for course, prereqs := range input {
		if len(prereqs) == 0 {
			todo.append(course)
		}
		for _, prereq := range prereqs {
			courses, ok := prereq_to_courses[prereq]
			if !ok {
				courses = []string{}
			}
			prereq_to_courses[prereq] = append(courses, course)
		}
	}

	result := []string{}
	for !todo.isEmpty() {
		prereq := todo.popLeft()
		result = append(result, prereq)

		// remove this prereq from all successor courses
		// if any course doesn't have prerequisites add it to todo
		for _, course := range prereq_to_courses[prereq] {
			courses := input[course]
			for i := range courses {
				if courses[i] == prereq {
					input[course] = append(courses[0:i], courses[i+1:]...)
					break
				}
			}
			if len(input[course]) == 0 {
				todo.append(course)
			}
		}
	}

	// circular dependency
	if len(result) < len(input) {
		return nil
	}

	return result
}

func main() {
	input := make(map[string][]string)
	input["CSC300"] = []string{"CSC100", "CSC200"}
	input["CSC200"] = []string{"CSC100"}
	input["CSC100"] = []string{}

	fmt.Printf("%v\n", topologicalSort(input))
}
