package UserType

type UserType int

const(
	Tourist UserType=1
	NormalUser UserType=2
	VerifiedUser UserType=3
	VerifiedTeacher UserType=4
	VerifiedSchool UserType=5
	Admin UserType=254
	Owner UserType=255
)