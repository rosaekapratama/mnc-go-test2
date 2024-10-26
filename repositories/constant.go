package repositories

const (
	// Query to calculate the outstanding balance
	queryFindAllTransactionByUserId = `
SELECT 
    t.id,
    u.id AS user_id,
    t.type AS transaction_type,
    t.status,
    t.cr,
    t.amount,
    t.remark AS remarks,
    t.balance_before,
    t.balance_after,
    t.created_dt
FROM 
    public.transactions t
JOIN 
    public.accounts a ON t.account_id = a.id
JOIN 
    public.users u ON a.user_id = u.id
WHERE
    u.id = ?
ORDER BY
    t.created_dt DESC;
`
)
