package main

type status struct {
	todo 		string
	inProgress 	string
	done 		string
}
type Tasks struct {
	id 			int
	description string
	status 		status // struct
	createdAt  	int
	updatedAt	int
}