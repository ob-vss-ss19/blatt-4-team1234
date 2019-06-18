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
	- AddUserResponse: AddUserResponse{int64 Id}
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
	- AddMovieRequest{Movie Movie} **(Id will be ignored)**
	- AddMovieResponse{int64 Id}
	- Errors: InvalidArgument
	
4. RemoveMovie
	- RemoveMovieRequest{int64 Id}
	- RemoveMovieResponse{}
	- Micro-Calls: ShowService.RemoveShowsForMovie{int64 MovieId}
	- Errors: NotFound, Internal


## HallService

### DataStructure

Hall{
	int64 Id = 1;
	string Name = 2;
	int64 Rows = 3;
	int64 Columns = 4;
}

### Endpoints

1. GetAllHalls
	- GetAllHallsRequest{}
	- GetAllHallsResponse{Halls []Hall}

2. GetHall
	- GetHallRequest{int64 Id}
	- GetHallResponse{Hall Hall}
	- Errors: NotFound

3. AddHall
	- AddHallRequest{Hall Hall} **Id will be ignored**
	- AddHallResponse{int64 Id}
	- Errors: InvalidArgument
	
4. RemoveHall
	- RemoveHallRequest{int64 Id}
	- RemoveHallResponse{}
	- Micro-Calls: ShowService.RemoveShowsForHalls{int64 Id}
	- Errors: NotFound, Internal

## ShowService

### DataStructure

Show{
	int64 Id = 1;
	int64 MovieId = 2;
	int64 HallId = 3;
	string DateTime = 4;
}

### Endpoints

1. GetAllShows{}
	- GetAllShowsRequest{}
	- GetAllShowsResponse{Shows []Show}
	
2. GetShow
	- GetShowRequest{int64 Id}
	- GetShowResponse{Show Show}
	- Errors: NotFound

3. AddShow
	- AddShowRequest{Show Show} **Id will be ignored**
	- AddShowResponse{int64 Id}
	- Errors: InvalidArgument
	
4. RemoveShow
	- RemoveShowRequest{int64 Id}
	- RemoveShowResponse{}
	- Micro-Calls: ShowService.RemoveShowsForShows{int64 Id}
	- Errors: NotFound, Internal
	
5. RemoveShowsForMovie

6. RemoveShowsForHall

## ReservationService

### DataStructure

### Endpoints
