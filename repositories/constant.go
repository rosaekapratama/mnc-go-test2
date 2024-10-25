package repositories

const (
	// Query to calculate the outstanding balance
	queryGetOutstanding = `
		SELECT 
			COALESCE(
				SUM(installment_payment_amount) - 
				(SELECT COALESCE(SUM(payment_amount), 0) FROM payments WHERE loan_id = ?), 
				0
			) AS outstanding
		FROM 
			billings
		WHERE 
			loan_id = ?;
`

	// Query to check for delinquency
	queryIsDelinquent = `
		SELECT EXISTS (
			SELECT 1
			FROM billings b
			LEFT JOIN payments p ON b.loan_id = p.loan_id
			WHERE b.loan_id = ?
			  AND (
				  p.payment_dt IS NULL OR 
				  (p.payment_dt < (?::timestamptz - INTERVAL '2 weeks') AND b.due_dt < ?::timestamptz)
			  )
			  AND b.due_dt < (?::timestamptz - INTERVAL '1 week')  -- This checks if the due date is outside the first week
		) AS delinquent;
	`
)
