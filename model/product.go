package model

import (
	"backend-shortcourse/go-webservice/database"
	"backend-shortcourse/go-webservice/form"
	"errors"
)

type ProductModel struct{}

func (pm ProductModel) ReadAll() ([]form.Product, error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return nil, err
	}
	defer database.CloseConnection(conn)

	var products []form.Product

	rows, err := conn.Query(`
		SELECT product_id, product_name, product_quantity,
			product_price
		FROM products;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p form.Product
		err = rows.Scan(&p.ProductID, &p.ProductName,
			&p.ProductQuantity, &p.ProductPrice)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (pm ProductModel) Read(id uint64) (form.Product, error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return form.Product{}, err
	}
	defer database.CloseConnection(conn)

	var product form.Product

	rows, err := conn.Query(`
		SELECT product_id, product_name, product_quantity,
			product_price
		FROM products
		WHERE product_id = ?;
	`, id)
	if err != nil {
		return form.Product{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&product.ProductID, &product.ProductName,
			&product.ProductQuantity, &product.ProductPrice)
		if err != nil {
			return form.Product{}, err
		}

		return product, nil
	}

	return form.Product{}, errors.New("no product belong to this id")
}

func (pm ProductModel) Add(id uint64,
	name string,
	quantity uint64,
	price float64,
	productType uint64) error {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	rows, err := conn.Query(`
		INSERT INTO products
			(product_id, product_name, product_quantity,
			product_price, product_product_type_id)
		VALUES
			(?, ?, ?, ?, ?)
	`, id, name, quantity, price, productType)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (pm ProductModel) Update(id uint64,
	name string,
	quantity uint64,
	price float64,
	productType uint64) error {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	if name != "" {
		rows, err := conn.Query(`
			UPDATE products
			SET product_name = ?
			WHERE product_id = ?;
		`, name, id)
		if err != nil {
			return err
		}
		defer rows.Close()
	}
	if quantity != 0 {
		rows, err := conn.Query(`
			UPDATE products
			SET product_quantity = ?
			WHERE product_id = ?;
		`, quantity, id)
		if err != nil {
			return err
		}
		defer rows.Close()
	}
	if price != 0 {
		rows, err := conn.Query(`
			UPDATE products
			SET product_price = ?
			WHERE product_id = ?;
		`, price, id)
		if err != nil {
			return err
		}
		defer rows.Close()
	}
	if productType != 0 {
		rows, err := conn.Query(`
			UPDATE products
			SET product_product_type_id = ?
			WHERE product_id = ?;
		`, productType, id)
		if err != nil {
			return err
		}
		defer rows.Close()
	}

	return nil
}

func (pm ProductModel) Delete(id uint64) error {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	rows, err := conn.Query(`
		DELETE FROM products
		WHERE product_id = ?;
	`, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
