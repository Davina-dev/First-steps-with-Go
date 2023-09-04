type User struct {
	Name string
	Bio  string
    Age int
}


user := User{
  Name: "Davina",
  Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
  Age:  123,
}