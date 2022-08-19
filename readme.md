TABLE MIGRATION RULES:
1. Migrate following model first: Account, AccountDetail, PhotoProfile
2. Insert 2 dummy datas to mst_account table manually through psql terminal
3. Migrate ServiceDetail
4. Migrate Order and ServicePrice
5. Migrate Feedback, OrderRequest, OrderStatus, and VideoResult
6. Migrate Refund and VideoProfile

In order to prevent 'violate constraints key' kind of error in future development, please kindly insert at least one dummy data to each tables manually via psql terminal. My appologies for this inconvinient.