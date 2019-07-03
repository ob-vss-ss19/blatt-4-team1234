echo "starting all services"
start ./hallservice/main.exe
start ./movieservice/main.exe
start ./reservationservice/main.exe
start ./showservice/main.exe
start ./userservice/main.exe
echo "all services running"
echo "press any key to show all running services"
pause
micro list services
echo "press any key to show all available Halls"
pause
micro call go.micro.srv.hallservice HallService.GetAllHalls "{}"
echo "press any key to show all available Shows"
pause
micro call go.micro.srv.showservice ShowService.GetAllShows "{}"
echo "press any key to show all available reservations"
pause
micro call go.micro.srv.reservationservice ReservationService.GetAllReservations "{}"
echo "press any key to remove Hall with Id:1"
pause
micro call go.micro.srv.hallservice HallService.RemoveHall "{\"Id\":1}"
pause
micro call go.micro.srv.hallservice HallService.GetAllHalls "{}"
echo "press any key to show all available Shows"
pause
micro call go.micro.srv.showservice ShowService.GetAllShows "{}"
echo "press any key to show all available reservations"
pause
micro call go.micro.srv.reservationservice ReservationService.GetAllReservations "{}"
echo "press any key to remove Hall with Id:1"
exit 0