package wallet



func TestService_Reject_success (t *testing.T) {
	s:=&service{}
}

//регистрируем пользователя
phone:=types.Phone("+992000000001")
account, err := s.RegisterAccount(phone)
if err != nil {
	t.ErrorF("Refect(): can't register account, error = %v", err)
	return
}

//пополняем его счет
err = s.Deposit(account.ID, 10_000_000)
if err != nil {
	t.Errorf("Reject(): can't deposit account, error = %v", err)
	return
}

//Осуществляем платеж на его счет
payment, err := s.Pay(account.ID, 1000_00, "auto")
if err != nil {
	t.Errorf("Reject(): can't deposit account, error = %v", err)
	return
}

//пробуем отменить платеж
err = s.Reject(payment.ID)
if err != nil {
	t.Errorf("Reject(): error = %v", err)
	return
}
