# Microservice Interface Contracts


## UserService

### DataStructure

User {
	int64 UserID
	string FirstName
	string LastName
	int64 Age
}
	
### Endpoints

A. GetAllUsers
	* GetAllUsersRequest{}
	* GetAllUsersResponse{[]User}
	* Errors: 'ADD POSSIBLE ERROR CODES'

B. GetUser
	*GetUserRequest{int64 UserID}
	*GetUserResponse: {}
	*Errors: 'ADD POSSIBLE ERROR CODES'

C. AddUser
	- AddUserRequest{
		int64 UserID
		string FirstName
		string LastName
		int64 Age
	}
	- AddUserResponse: AddUserResponse{}
	- Errors
	
D. RemoveUser
	- RemoveUserRequest{int64 UserID}
	- RemoveUserResponse{}
	- Errors:
	- Cross-Calls: ReservationService.GetReservationsForUser{})	

## MovieService

### DataStructure

### Endpoints


## HallService

### DataStructure

### Endpoints



## ShowService

### DataStructure

### Endpoints


## ReservationService

### DataStructure

### Endpoints
>>>>>>> started readme with endpoints
