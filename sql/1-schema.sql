CREATE TABLE dbo.Movie(
	Id INT IDENTITY PRIMARY KEY,
	Title NVARCHAR(100) NOT NULL,
	Released DATE NULL,
	Rating DECIMAL(3,1),
	CONSTRAINT chk_movie_rating CHECK (Rating between 0 and 10),
)