package email

import "fmt"

// RegistrationEmail generates an email after successful registration.
func RegistrationEmail(
	username string,
	email string,
) string {

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Registration Successful</title>
</head>

<body style="font-family: Arial, sans-serif; line-height: 1.6;">

	<h2>Registration Successful</h2>

	<p>Hello <strong>%s</strong>,</p>

	<p>
		Your account has been successfully registered.
		Welcome to Waste Management Service.
	</p>

	<table>
		<tr>
			<td><strong>Email</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Status</strong></td>
			<td>: ACTIVE</td>
		</tr>
	</table>

	<br>

	<p>
		You can now use our waste management service.
	</p>

	<p>
		Thank you for joining us.
	</p>

</body>
</html>
`,
		username,
		email,
	)
}

// InvoiceEmail generates an email when a new invoice is created.
func InvoiceEmail(
	username string,
	invoiceNumber string,
	service string,
	amount float64,
	dueDate string,
) string {

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>New Invoice</title>
</head>

<body style="font-family: Arial, sans-serif; line-height: 1.6;">

	<h2>New Invoice</h2>

	<p>Hello <strong>%s</strong>,</p>

	<p>
		A new invoice has been created for your waste management service.
		Please complete the payment before the due date.
	</p>

	<table>
		<tr>
			<td><strong>Invoice Number</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Service</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Amount</strong></td>
			<td>: Rp %.0f</td>
		</tr>

		<tr>
			<td><strong>Due Date</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Status</strong></td>
			<td>: UNPAID</td>
		</tr>
	</table>

	<br>

	<p>
		Please complete your payment before the due date.
	</p>

	<p>
		Thank you.
	</p>

</body>
</html>
`,
		username,
		invoiceNumber,
		service,
		amount,
		dueDate,
	)
}

// CollectionReminderEmail generates an email reminder
// one day before the scheduled waste collection.
func CollectionReminderEmail(
	username string,
	service string,
	collectionDate string,
) string {

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Waste Collection Reminder</title>
</head>

<body style="font-family: Arial, sans-serif; line-height: 1.6;">

	<h2>Waste Collection Reminder</h2>

	<p>Hello <strong>%s</strong>,</p>

	<p>
		This is a reminder that your waste collection is scheduled for tomorrow.
		Please make sure your waste is ready for collection.
	</p>

	<table>
		<tr>
			<td><strong>Service</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Collection Date</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Reminder</strong></td>
			<td>: 1 day before collection</td>
		</tr>
	</table>

	<br>

	<p>
		Please prepare your waste before the scheduled collection time.
	</p>

	<p>
		Thank you for using our service.
	</p>

</body>
</html>
`,
		username,
		service,
		collectionDate,
	)
}

// PaymentInvoiceEmail generates an email
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
	<title>Payment Successful</title>
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
			<td>: Rp %.0f</td>
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
