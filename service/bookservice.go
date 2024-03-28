package service
 import(
	"database/sql"
	"lms/models"
 )

func GettingBooks(db *sql.DB) ([]models.Book, error) {
	var books []models.Book
	rows, err := db.Query("SELECT * FROM  book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Bookid,&book.Genre,&book.Author,&book.Publisher,&book.Stockcount)
		if err!=nil{
			return nil, err
	}
		books = append(books, book)
	
		
	}
	return books,nil
}
func Addbooks(db *sql.DB,book models.Book)error{
	// Insert the new user into the database
	query:="INSERT INTO book(genre,author,publisher,stock_count) VALUES ($1,$2,$3,$4)"
	_, err := db.Exec(query,book.Genre,book.Author,book.Publisher,book.Stockcount)
	if err != nil {
		return err
	}
	return nil

}
func UpdateStock(db *sql.DB,book models.Book,id int)error{
	query:="UPDATE book set stock_count=$1 where id=$2"
	_,err:=db.Exec(query,book.Stockcount,id)
	if err!=nil{
		return err
	}
	return nil
}
	
func GetBook(db *sql.DB,id int)(* models.Book,error){
	var book models.Book
	query:="SELECT  * FROM book WHERE  id = $1"
	row:=db.QueryRow(query,id)
	
	err:=row.Scan(&book.Bookid,&book.Genre,&book.Author,&book.Publisher,&book.Stockcount)
	if err!=nil{
		return nil,err
	}
	return &book,nil

}
func Delete(book *models.Book,db *sql.DB)error{
	query:="DELETE FROM book WHERE id=$1"
	_,err:=db.Exec(query,book.Bookid)
	if err!=nil{
		return err
	}
	return nil



}

func Authenticate(userauth models.UserAuth,db *sql.DB )(*models.User,error){
	var user models.User
	
	err:=db.QueryRow("SELECT * from users where username=$1",userauth.Username).Scan(&user.Userid,&user.Username,&user.Age,&user.Email,&user.Password,&user.Isadmin,&user.Subid,&user.Subdate)
	if err!=nil{
		return nil,err
	}
	return &user ,nil
}


