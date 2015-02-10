package service

type MultiplicationRequest struct {
	Factor1, Factor2 int
}

type MultiplicationResponse struct {
	Product int
}

type MultiplicationService struct{}

func (t *MultiplicationService) Multiply(request *MultiplicationRequest, response *MultiplicationResponse) error {
	response.Product = request.Factor1 * request.Factor2
	return nil
}
