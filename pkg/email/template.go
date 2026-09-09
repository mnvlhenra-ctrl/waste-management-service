package email

import "fmt"

// PaymentReminderEmail generates an email reminding the user
// to complete their payment.
func PaymentReminderEmail(
	username string,
	invoiceNumber string,
	amount float64,
	dueDate string,
	paymentURL string,
) string {

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Payment Reminder</title>
</head>

<body style="font-family: Arial, sans-serif; line-height: 1.6;">

	<h2>Payment Reminder</h2>

	<p>Hello <strong>%s</strong>,</p>

	<p>
		This is a reminder that your payment is still pending.
		Please complete your payment before the due date.
	</p>

	<table>
		<tr>
			<td><strong>Invoice Number</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Amount</strong></td>
			<td>: Rp %.2f</td>
		</tr>

		<tr>
			<td><strong>Due Date</strong></td>
			<td>: %s</td>
		</tr>
	</table>

	<br>

	<p>
		<a href="%s"
		   style="
		   display: inline-block;
		   padding: 10px 20px;
		   background-color: #2563eb;
		   color: white;
		   text-decoration: none;
		   border-radius: 5px;
		   ">
			Pay Now
		</a>
	</p>

	<p>
		If you have already completed the payment,
		please ignore this email.
	</p>

	<p>
		Thank you.
	</p>

</body>
</html>
`,
		username,
		invoiceNumber,
		amount,
		dueDate,
		paymentURL,
	)
}

// PaymentInvoiceEmail generates an invoice email
// after a payment has been successfully completed.
func PaymentInvoiceEmail(
	username string,
	invoiceNumber string,
	amount float64,
	paymentDate string,
	paymentMethod string,
) string {

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Payment Invoice</title>
</head>

<body style="font-family: Arial, sans-serif; line-height: 1.6;">

	<h2>Payment Successful</h2>

	<p>Hello <strong>%s</strong>,</p>

	<p>
		Your payment has been successfully completed.
		Thank you for your payment.
	</p>

	<hr>

	<table>
		<tr>
			<td><strong>Invoice Number</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Amount Paid</strong></td>
			<td>: Rp %.2f</td>
		</tr>

		<tr>
			<td><strong>Payment Date</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Payment Method</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Status</strong></td>
			<td>: PAID</td>
		</tr>
	</table>

	<hr>

	<p>
		This email serves as your payment confirmation.
		Please keep it for your records.
	</p>

	<p>
		Thank you for using our service.
	</p>

</body>
</html>
`,
		username,
		invoiceNumber,
		amount,
		paymentDate,
		paymentMethod,
	)
}
