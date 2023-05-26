package responses

type createBookingResponseBooking struct {
	Id int `json:"id"`
}

type createBookingResponseData struct {
	Booking createBookingResponseBooking `json:"booking"`
}

func CreateBookingResponseFromStoreData(bookingId int) createBookingResponseData {
	res := createBookingResponseData{}

	res.Booking = createBookingResponseBooking{
		Id: bookingId,
	}

	return res
}
