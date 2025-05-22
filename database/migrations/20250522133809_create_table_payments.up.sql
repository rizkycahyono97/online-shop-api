CREATE TABLE payments (
      id INT AUTO_INCREMENT PRIMARY KEY,
      order_id INT NOT NULL,
      user_id INT NOT NULL,
      payment_method ENUM('cod', 'va', 'credit', 'paylater') NOT NULL,
      amount DECIMAL(10, 2) NOT NULL,
      orders_status ENUM('pending', 'paid', 'failed', 'cancelled') DEFAULT 'pending',
      FOREIGN KEY (order_id) REFERENCES orders(id),
      FOREIGN KEY (user_id) REFERENCES users(id)
);