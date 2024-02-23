package main

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

const (
	GET_ALL_ARTCILES_QUERY = "select * from articles"
	GET_ARTICLE_ID_QUERY   = "select * from articles where id=?"
	CREATE_ARTICLE_QUERY   = "insert into articles (title, content) values (?, ?)"
	DELETE_ARTICLE_QUERY   = "delete from articles where id=?"
)

func GetAllArticles() ([]Article, error) {
	rows, err := DB.Query(GET_ALL_ARTCILES_QUERY)
	if err != nil {
		return nil, err
	}
	var articles []Article
	for rows.Next() {
		article := new(Article)
		rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
		)

		articles = append(articles, *article)

	}
	return articles, nil
}

func GetArticleByID(id int) (*Article, error) {
	row := DB.QueryRow(GET_ARTICLE_ID_QUERY, id)

	article := new(Article)
	err := row.Scan(&article.ID, &article.Title, &article.Content)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func CreateArticle(article Article) error {
	query, err := DB.Prepare(CREATE_ARTICLE_QUERY)
	defer query.Close()
	if err != nil {
		return err
	}
	_, err = query.Exec(article.Title, article.Content)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id int) error {
	query, err := DB.Prepare(DELETE_ARTICLE_QUERY)
	defer query.Close()
	if err != nil {
		return err
	}
	_, err = query.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
