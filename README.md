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

1. GetAllUsers
	- GetAllUsersRequest{}
	- GetAllUsersResponse{users []User}

2. GetUser
	- GetUserRequest{int64 Id}
	- GetUserResponse: {}
	- Errors: NotFound

3. AddUser
	- AddUserRequest{User User} **(UserID will be ignored)**
	- AddUserResponse: AddUserResponse{}
	- Errors: InvalidArgument
	
4. RemoveUser
	- RemoveUserRequest{int64 UserID}
	- RemoveUserResponse{}
	- Micro-Calls: ReservationService.GetReservationsForUser{int64 UserID})	
	- Errors: NotFound, Internal, FailedPrecondition


## MovieService

### DataStructure

Movie{
	int64 Id = 1;
	string Title = 2;
	int64 Fsk = 3;
}

### Endpoints

1. GetAllMovies
	- GetAllMoviesRequest{}
	- GetAllMoviesResponse{Movies []Movie}
	
2. GetMovie
	- GetMovieRequest{int64 Id}
	- GetMovieResponse{Movie Movie}
	- Errors: NotFound
	
3. AddMovie
	- AddMovieRequest{Movie Movie} **(UserID will be ignored)**
	- AddMovieResponse{}
	- Errors: InvalidArgument
	
4. RemoveMovie
	- RemoveMovieRequest{int64 Id}
	- RemoveMovieResponse{}
	- Micro-Calls: ShowService.RemoveShowsForMovie{int64 MovieId}
	- Errors: NotFound, Internal


## HallService

### DataStructure

### Endpoints



## ShowService

### DataStructure

### Endpoints


## ReservationService

### DataStructure

### Endpoints
