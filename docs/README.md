# Orch

Also factors out the DB
It just doesn't have to create a mock repository
It just uses a struct
Given the repository returns X, domainBooks should return Y
params, _ := getBookParams(req)
repoBooks, _ := &MockRepository{Db: Db}.GetBooks(params)
domainBooks, _ := domain.GetBooks(repoBooks)
t.Equals(t, expectedBooks, domainBooks)
