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
echo "press any key to show all current reservations"
pause
micro call go.micro.srv.reservationservice ReservationService.GetAllReservations "{}"
echo "continue and send 2 reservation requests and 2 activation requests"
pause
micro call go.micro.srv.reservationservice ReservationService.RequestReservation "{\"ShowId\":1,\"Seats\":[{\"Row\":1,\"Column\":1}],\"UserId\":1}"
micro call go.micro.srv.reservationservice ReservationService.RequestReservation "{\"ShowId\":1,\"Seats\":[{\"Row\":1,\"Column\":1}],\"UserId\":2}"
micro call go.micro.srv.reservationservice ReservationService.ActivateReservation "{\"ReservationId\":5,\"UserId\":1}"
micro call go.micro.srv.reservationservice ReservationService.ActivateReservation "{\"ReservationId\":6,\"UserId\":2}"
pause
echo "press any key to show all reservations"
pause
micro call go.micro.srv.reservationservice ReservationService.GetAllReservations "{}"
echo "press any key to exit"
pause